package include

import (
	"context"

	"proxy-core/adapter"
	"proxy-core/adapter/endpoint"
	"proxy-core/adapter/inbound"
	"proxy-core/adapter/outbound"
	C "proxy-core/constant"
	"proxy-core/dns"
	"proxy-core/dns/transport"
	"proxy-core/dns/transport/fakeip"
	"proxy-core/dns/transport/hosts"
	"proxy-core/dns/transport/local"
	"proxy-core/log"
	"proxy-core/option"
	"proxy-core/protocol/block"
	"proxy-core/protocol/direct"
	protocolDNS "proxy-core/protocol/dns"
	"proxy-core/protocol/group"
	"proxy-core/protocol/http"
	"proxy-core/protocol/mixed"
	"proxy-core/protocol/naive"
	"proxy-core/protocol/redirect"
	"proxy-core/protocol/shadowsocks"
	"proxy-core/protocol/shadowtls"
	"proxy-core/protocol/socks"
	"proxy-core/protocol/ssh"
	"proxy-core/protocol/tor"
	"proxy-core/protocol/trojan"
	"proxy-core/protocol/tun"
	"proxy-core/protocol/vless"
	"proxy-core/protocol/vmess"

	E "github.com/sagernet/sing/common/exceptions"
)

func InboundRegistry() *inbound.Registry {
	registry := inbound.NewRegistry()

	tun.RegisterInbound(registry)
	redirect.RegisterRedirect(registry)
	redirect.RegisterTProxy(registry)
	direct.RegisterInbound(registry)

	socks.RegisterInbound(registry)
	http.RegisterInbound(registry)
	mixed.RegisterInbound(registry)

	shadowsocks.RegisterInbound(registry)
	vmess.RegisterInbound(registry)
	trojan.RegisterInbound(registry)
	naive.RegisterInbound(registry)
	shadowtls.RegisterInbound(registry)
	vless.RegisterInbound(registry)

	registerQUICInbounds(registry)
	registerStubForRemovedInbounds(registry)

	return registry
}

func OutboundRegistry() *outbound.Registry {
	registry := outbound.NewRegistry()

	direct.RegisterOutbound(registry)

	block.RegisterOutbound(registry)
	protocolDNS.RegisterOutbound(registry)

	group.RegisterSelector(registry)
	group.RegisterURLTest(registry)

	socks.RegisterOutbound(registry)
	http.RegisterOutbound(registry)
	shadowsocks.RegisterOutbound(registry)
	vmess.RegisterOutbound(registry)
	trojan.RegisterOutbound(registry)
	tor.RegisterOutbound(registry)
	ssh.RegisterOutbound(registry)
	shadowtls.RegisterOutbound(registry)
	vless.RegisterOutbound(registry)

	registerQUICOutbounds(registry)
	registerWireGuardOutbound(registry)
	registerStubForRemovedOutbounds(registry)

	return registry
}

func EndpointRegistry() *endpoint.Registry {
	registry := endpoint.NewRegistry()

	registerWireGuardEndpoint(registry)

	return registry
}

func DNSTransportRegistry() *dns.TransportRegistry {
	registry := dns.NewTransportRegistry()

	transport.RegisterTCP(registry)
	transport.RegisterUDP(registry)
	transport.RegisterTLS(registry)
	transport.RegisterHTTPS(registry)
	transport.RegisterPredefined(registry)
	hosts.RegisterTransport(registry)
	local.RegisterTransport(registry)
	fakeip.RegisterTransport(registry)

	registerQUICTransports(registry)
	registerDHCPTransport(registry)

	return registry
}

func registerStubForRemovedInbounds(registry *inbound.Registry) {
	inbound.Register[option.ShadowsocksInboundOptions](registry, C.TypeShadowsocksR, func(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.ShadowsocksInboundOptions) (adapter.Inbound, error) {
		return nil, E.New("ShadowsocksR is deprecated and removed in proxy-core 1.6.0")
	})
}

func registerStubForRemovedOutbounds(registry *outbound.Registry) {
	outbound.Register[option.ShadowsocksROutboundOptions](registry, C.TypeShadowsocksR, func(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.ShadowsocksROutboundOptions) (adapter.Outbound, error) {
		return nil, E.New("ShadowsocksR is deprecated and removed in proxy-core 1.6.0")
	})
}
