package platform

import (
	"proxy-core/adapter"
	"proxy-core/common/process"
	"proxy-core/option"

	tun "github.com/sagernet/sing-tun"
	"github.com/sagernet/sing/common/logger"
)

type Interface interface {
	Initialize(networkManager adapter.NetworkManager) error
	UsePlatformAutoDetectInterfaceControl() bool
	AutoDetectInterfaceControl(fd int) error
	OpenTun(options *tun.Options, platformOptions option.TunPlatformOptions) (tun.Tun, error)
	CreateDefaultInterfaceMonitor(logger logger.Logger) tun.DefaultInterfaceMonitor
	Interfaces() ([]adapter.NetworkInterface, error)
	UnderNetworkExtension() bool
	IncludeAllNetworks() bool
	ClearDNSCache()
	ReadWIFIState() adapter.WIFIState
	SystemCertificates() []string
	process.Searcher
	SendNotification(notification *Notification) error
}

type Notification struct {
	Identifier string
	TypeName   string
	TypeID     int32
	Title      string
	Subtitle   string
	Body       string
	OpenURL    string
}
