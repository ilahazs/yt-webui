package urlutil

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

// lookupIPFunc is a function type matching net.LookupIP, used for dependency injection in tests.
type lookupIPFunc func(host string) ([]net.IP, error)

var lookupIP lookupIPFunc = net.LookupIP

// SetLookupIP overrides the default DNS lookup function. Primarily used in testing.
func SetLookupIP(fn func(host string) ([]net.IP, error)) {
	lookupIP = fn
}

// ValidateURL validates that a raw URL is valid, uses http or https, and does not resolve to a local or private address.
func ValidateURL(rawURL string) error {
	if strings.TrimSpace(rawURL) == "" {
		return fmt.Errorf("URL cannot be empty")
	}

	parsed, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("malformed URL: %w", err)
	}

	scheme := strings.ToLower(parsed.Scheme)
	if scheme != "http" && scheme != "https" {
		return fmt.Errorf("invalid URL scheme %q, only http and https are allowed", parsed.Scheme)
	}

	host := parsed.Hostname()
	if host == "" {
		return fmt.Errorf("URL host is missing")
	}

	hostLower := strings.ToLower(host)
	// Reject explicit local hostnames
	if hostLower == "localhost" || strings.HasSuffix(hostLower, ".local") || strings.HasSuffix(hostLower, ".localhost") {
		return fmt.Errorf("local hostnames are not allowed")
	}

	// If the host is a raw IP address, validate it directly
	if ip := net.ParseIP(host); ip != nil {
		if err := validateIP(ip); err != nil {
			return err
		}
		return nil
	}

	// Resolve the hostname and check all resolved IP addresses
	ips, err := lookupIP(host)
	if err != nil {
		return fmt.Errorf("failed to resolve host %q: %w", host, err)
	}

	if len(ips) == 0 {
		return fmt.Errorf("no IP addresses found for host %q", host)
	}

	for _, ip := range ips {
		if err := validateIP(ip); err != nil {
			return fmt.Errorf("host %q resolves to blocked IP: %w", host, err)
		}
	}

	return nil
}

func validateIP(ip net.IP) error {
	// If IPv4-mapped IPv6 address, convert it to IPv4 for proper evaluation
	if ipv4 := ip.To4(); ipv4 != nil {
		ip = ipv4
	}

	if ip.IsLoopback() {
		return fmt.Errorf("loopback IP address %s is blocked", ip)
	}
	if ip.IsPrivate() {
		return fmt.Errorf("private IP address %s is blocked", ip)
	}
	if ip.IsLinkLocalUnicast() {
		return fmt.Errorf("link-local unicast IP address %s is blocked", ip)
	}
	if ip.IsLinkLocalMulticast() {
		return fmt.Errorf("link-local multicast IP address %s is blocked", ip)
	}
	if ip.IsMulticast() {
		return fmt.Errorf("multicast IP address %s is blocked", ip)
	}
	if ip.IsUnspecified() {
		return fmt.Errorf("unspecified IP address %s is blocked", ip)
	}

	return nil
}
