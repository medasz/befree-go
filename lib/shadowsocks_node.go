package lib

type ShadowsocksNode struct {
	Name     string
	Server   string
	Port     int
	Password string
	Cipher   string
}

func ParseShadowsocksNode() *ShadowsocksNode {
	return &ShadowsocksNode{}
}

func (n *ShadowsocksNode) ToClashProxy() {

}
