//go:build with_dhcp

package include

import (
	"proxy-core/dns"
	"proxy-core/dns/transport/dhcp"
)

func registerDHCPTransport(registry *dns.TransportRegistry) {
	dhcp.RegisterTransport(registry)
}
