//go:build !with_acme

package tls

import (
	"context"
	"crypto/tls"

	"proxy-core/adapter"
	"proxy-core/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func startACME(ctx context.Context, options option.InboundACMEOptions) (*tls.Config, adapter.Service, error) {
	return nil, nil, E.New(`ACME is not included in this build, rebuild with -tags with_acme`)
}
