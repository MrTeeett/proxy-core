package rule

import (
	"proxy-core/adapter"
)

var _ RuleItem = (*IPAcceptAnyItem)(nil)

type IPAcceptAnyItem struct{}

func NewIPAcceptAnyItem() *IPAcceptAnyItem {
	return &IPAcceptAnyItem{}
}

func (r *IPAcceptAnyItem) Match(metadata *adapter.InboundContext) bool {
	return len(metadata.DestinationAddresses) > 0
}

func (r *IPAcceptAnyItem) String() string {
	return "ip_accept_any=true"
}
