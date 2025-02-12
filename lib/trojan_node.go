package lib

import (
	"net/url"
	"strconv"
	"strings"
)

type TrojanNodeProxy struct {
	*TrojanNode
	Type           string
	SkipCertVerify bool
}

func NewTrojanNodeProxy(n *TrojanNode) *TrojanNodeProxy {
	return &TrojanNodeProxy{n, "trojan", true}
}
func (n *TrojanNodeProxy) GetType() string {
	return n.Type
}
func (n *TrojanNodeProxy) GetServer() string {
	return n.Server
}
func (n *TrojanNodeProxy) GetPort() int {
	return n.Port
}
func (n *TrojanNodeProxy) GetPassword() any {
	return n.Password
}
func (n *TrojanNodeProxy) GetCipher() any {
	return nil
}
func (n *TrojanNodeProxy) GetUuid() any {
	return nil
}
func (n *TrojanNodeProxy) GetAlterId() any {
	return nil
}
func (n *TrojanNodeProxy) GetSkipCertVerify() any {
	return n.SkipCertVerify
}

type TrojanNode struct {
	Name     string
	Server   string
	Port     int
	Sni      string
	Password string
}

func (t *TrojanNode) ToClashProxy() NodeProxy {
	return NewTrojanNodeProxy(t)
}

func (t *TrojanNode) GetName() string {
	return t.Name
}

func (t *TrojanNode) SetName(name string) {
	t.Name = name
}

func NewTrojanNode(rawData string) (Node, error) {
	a := strings.Split(rawData, "@")
	portStr := strings.Split(strings.Split(a[1], ":")[1], "?")[0]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}
	name, err := url.QueryUnescape(strings.Split(strings.Split(strings.Split(a[1], ":")[1], "?")[1], "#")[1])
	if err != nil {
		name = "aasda"
	}
	name = strings.Trim(name, "\n")
	name = strings.Trim(name, "\r")
	name = strings.ReplaceAll(name, "#", "")
	node := &TrojanNode{
		Name:     name,
		Server:   strings.Split(a[1], ":")[0],
		Port:     port,
		Sni:      strings.Split(strings.Split(strings.Split(a[1], ":")[1], "?")[1], "#")[0],
		Password: a[0],
	}

	return node, nil
}
