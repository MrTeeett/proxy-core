//go:build with_wireguard

package include

import (
	"proxy-core/adapter/endpoint"
	"proxy-core/adapter/outbound"
	"proxy-core/protocol/wireguard"
)

func registerWireGuardOutbound(registry *outbound.Registry) {
	wireguard.RegisterOutbound(registry)
}

func registerWireGuardEndpoint(registry *endpoint.Registry) {
	wireguard.RegisterEndpoint(registry)
}
