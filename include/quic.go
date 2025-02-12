//go:build with_quic

package include

import (
	"proxy-core/adapter/inbound"
	"proxy-core/adapter/outbound"
	"proxy-core/dns"
	"proxy-core/dns/transport/quic"
	"proxy-core/protocol/hysteria"
	"proxy-core/protocol/hysteria2"
	_ "proxy-core/protocol/naive/quic"
	"proxy-core/protocol/tuic"
	_ "proxy-core/transport/v2rayquic"
)

func registerQUICInbounds(registry *inbound.Registry) {
	hysteria.RegisterInbound(registry)
	tuic.RegisterInbound(registry)
	hysteria2.RegisterInbound(registry)
}

func registerQUICOutbounds(registry *outbound.Registry) {
	hysteria.RegisterOutbound(registry)
	tuic.RegisterOutbound(registry)
	hysteria2.RegisterOutbound(registry)
}

func registerQUICTransports(registry *dns.TransportRegistry) {
	quic.RegisterTransport(registry)
	quic.RegisterHTTP3Transport(registry)
}
