package urlutil

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

func TestValidateURL(t *testing.T) {
	// Set up mock DNS lookup
	origLookup := lookupIP
	defer SetLookupIP(origLookup)

	SetLookupIP(func(host string) ([]net.IP, error) {
		switch host {
		case "example.com", "youtube.com", "www.youtube.com":
			return []net.IP{net.ParseIP("93.184.216.34")}, nil // Public IP
		case "localhost.localdomain":
			return []net.IP{net.ParseIP("127.0.0.1")}, nil // Resolves to loopback
		case "private.host":
			return []net.IP{net.ParseIP("192.168.1.50")}, nil // Resolves to private IP
		case "multi.host":
			return []net.IP{
				net.ParseIP("93.184.216.34"), // Public
				net.ParseIP("10.0.0.5"),      // Private
			}, nil
		case "unresolvable.host":
			return nil, fmt.Errorf("no such host")
		default:
			return nil, fmt.Errorf("unknown host in test mock")
		}
	})

	tests := []struct {
		name        string
		urlStr      string
		expectError bool
		errContains string
	}{
		// Valid cases
		{
			name:        "Valid HTTP",
			urlStr:      "http://example.com",
			expectError: false,
		},
		{
			name:        "Valid HTTPS with path",
			urlStr:      "https://youtube.com/watch?v=dQw4w9WgXcQ",
			expectError: false,
		},
		{
			name:        "Valid HTTPS with port",
			urlStr:      "https://example.com:443/test",
			expectError: false,
		},

		// Empty/Malformed cases
		{
			name:        "Empty URL",
			urlStr:      "",
			expectError: true,
			errContains: "URL cannot be empty",
		},
		{
			name:        "Whitespace URL",
			urlStr:      "   ",
			expectError: true,
			errContains: "URL cannot be empty",
		},
		{
			name:        "Malformed URL parsing",
			urlStr:      "http://%4Gexample.com", // %4G is invalid percent encoding
			expectError: true,
			errContains: "malformed URL",
		},

		// Scheme cases
		{
			name:        "Invalid scheme ftp",
			urlStr:      "ftp://example.com",
			expectError: true,
			errContains: "only http and https are allowed",
		},
		{
			name:        "Invalid scheme file",
			urlStr:      "file:///etc/passwd",
			expectError: true,
			errContains: "only http and https are allowed",
		},
		{
			name:        "No scheme",
			urlStr:      "example.com",
			expectError: true,
			errContains: "only http and https are allowed",
		},

		// Host/Hostname cases
		{
			name:        "No host",
			urlStr:      "https:///path",
			expectError: true,
			errContains: "URL host is missing",
		},
		{
			name:        "Localhost name",
			urlStr:      "http://localhost",
			expectError: true,
			errContains: "local hostnames are not allowed",
		},
		{
			name:        "Localhost sub-suffix",
			urlStr:      "http://sub.localhost",
			expectError: true,
			errContains: "local hostnames are not allowed",
		},
		{
			name:        "Local suffix",
			urlStr:      "http://device.local/status",
			expectError: true,
			errContains: "local hostnames are not allowed",
		},

		// Direct IP cases
		{
			name:        "Loopback IPv4",
			urlStr:      "http://127.0.0.1",
			expectError: true,
			errContains: "loopback IP address 127.0.0.1 is blocked",
		},
		{
			name:        "Loopback IPv6",
			urlStr:      "http://[::1]",
			expectError: true,
			errContains: "loopback IP address ::1 is blocked",
		},
		{
			name:        "Private IPv4 Class A",
			urlStr:      "http://10.0.0.1",
			expectError: true,
			errContains: "private IP address 10.0.0.1 is blocked",
		},
		{
			name:        "Private IPv4 Class B",
			urlStr:      "http://172.16.0.1",
			expectError: true,
			errContains: "private IP address 172.16.0.1 is blocked",
		},
		{
			name:        "Private IPv4 Class C",
			urlStr:      "http://192.168.1.1",
			expectError: true,
			errContains: "private IP address 192.168.1.1 is blocked",
		},
		{
			name:        "IPv4-mapped IPv6 Loopback",
			urlStr:      "http://[::ffff:127.0.0.1]",
			expectError: true,
			errContains: "loopback IP address 127.0.0.1 is blocked",
		},
		{
			name:        "IPv4-mapped IPv6 Private",
			urlStr:      "http://[::ffff:192.168.1.1]",
			expectError: true,
			errContains: "private IP address 192.168.1.1 is blocked",
		},
		{
			name:        "Link-local Unicast IPv4",
			urlStr:      "http://169.254.1.1",
			expectError: true,
			errContains: "link-local unicast IP address 169.254.1.1 is blocked",
		},
		{
			name:        "Unspecified IPv4",
			urlStr:      "http://0.0.0.0",
			expectError: true,
			errContains: "unspecified IP address 0.0.0.0 is blocked",
		},

		// DNS resolution cases
		{
			name:        "Hostname resolves to loopback",
			urlStr:      "http://localhost.localdomain",
			expectError: true,
			errContains: "resolves to blocked IP: loopback IP address 127.0.0.1 is blocked",
		},
		{
			name:        "Hostname resolves to private IP",
			urlStr:      "http://private.host",
			expectError: true,
			errContains: "resolves to blocked IP: private IP address 192.168.1.50 is blocked",
		},
		{
			name:        "Hostname resolves to mixed public and private IP",
			urlStr:      "http://multi.host",
			expectError: true,
			errContains: "resolves to blocked IP: private IP address 10.0.0.5 is blocked",
		},
		{
			name:        "Hostname cannot be resolved",
			urlStr:      "http://unresolvable.host",
			expectError: true,
			errContains: "failed to resolve host",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateURL(tt.urlStr)
			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.errContains)
				}
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("expected error containing %q, got %v", tt.errContains, err)
				}
			} else {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
			}
		})
	}
}
