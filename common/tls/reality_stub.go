//go:build !with_reality_server

package tls

import (
	"context"

	"proxy-core/log"
	"proxy-core/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func NewRealityServer(ctx context.Context, logger log.Logger, options option.InboundTLSOptions) (ServerConfig, error) {
	return nil, E.New(`reality server is not included in this build, rebuild with -tags with_reality_server`)
}
