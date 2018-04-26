package adapter

import (
	"github.com/containernetworking/plugins/pkg/ns"
)

type NS struct {
}

func NewNS() NS {
	return NS{}
}

func (NS) GetNS(nspath string) (ns.NetNS, error) {
	return ns.GetNS(nspath)
}
