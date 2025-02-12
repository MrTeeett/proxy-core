//go:build with_quic

package v2rayquic

import "proxy-core/transport/v2ray"

func init() {
	v2ray.RegisterQUICConstructor(NewServer, NewClient)
}
