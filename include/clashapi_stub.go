//go:build !with_clash_api

package include

import (
	"context"

	"proxy-core/adapter"
	"proxy-core/experimental"
	"proxy-core/log"
	"proxy-core/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func init() {
	experimental.RegisterClashServerConstructor(func(ctx context.Context, logFactory log.ObservableFactory, options option.ClashAPIOptions) (adapter.ClashServer, error) {
		return nil, E.New(`clash api is not included in this build, rebuild with -tags with_clash_api`)
	})
}
