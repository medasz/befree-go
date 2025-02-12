package lib

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"strings"
)

type VMessNodeProxy struct {
	*VMessNode
	Type    string
	AlterId int
}

func NewVMessNodeProxy(n *VMessNode) *VMessNodeProxy {
	return &VMessNodeProxy{n, "vmess", 0}
}
func (n *VMessNodeProxy) GetType() string {
	return n.Type
}
func (n *VMessNodeProxy) GetServer() string {
	return n.Server
}
func (n *VMessNodeProxy) GetPort() int {
	return n.Port
}
func (n *VMessNodeProxy) GetPassword() any {
	return nil
}
func (n *VMessNodeProxy) GetCipher() any {
	return n.Cipher
}
func (n *VMessNodeProxy) GetUuid() any {
	return n.UUID
}
func (n *VMessNodeProxy) GetAlterId() any {
	return n.AlterId
}
func (n *VMessNodeProxy) GetSkipCertVerify() any {
	return nil
}

type VMessNode struct {
	Name   string
	Server string
	Port   int
	UUID   string
	Cipher string
}

func (v *VMessNode) ToClashProxy() NodeProxy {
	return NewVMessNodeProxy(v)
}

func (v *VMessNode) GetName() string {
	return v.Name
}

func (v *VMessNode) SetName(name string) {
	v.Name = name
}

type vmessNodeRaw struct {
	Add            string `json:"add"`
	Aid            any    `json:"aid"`
	Host           string `json:"host"`
	Id             string `json:"id"`
	Net            string `json:"net"`
	Path           string `json:"path"`
	Port           any    `json:"port"`
	Ps             string `json:"ps"`
	Tls            any    `json:"tls"`
	Type           string `json:"type"`
	Security       string `json:"security"`
	SkipCertVerify bool   `json:"skip-cert-verify"`
	Sni            string `json:"sni"`
	Cipher         string `json:"cipher"`
}

func NewVMessNode(rawData string) (Node, error) {
	rawData = cleanBase64String(rawData)
	rawDataBase, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return nil, err
	}
	vmessNode := &vmessNodeRaw{}
	err = json.Unmarshal(rawDataBase, &vmessNode)
	if err != nil {
		return nil, err
	}
	node := &VMessNode{
		Name:   strings.ReplaceAll(vmessNode.Ps, "#", ""),
		Server: vmessNode.Add,
		UUID:   vmessNode.Id,
		Cipher: "auto",
	}
	switch port := vmessNode.Port.(type) {
	case int:
		node.Port = port
	case string:
		node.Port, err = strconv.Atoi(port)
	case float64:
		node.Port = int(port)
	}
	if vmessNode.Cipher != "" {
		node.Cipher = vmessNode.Cipher
	}
	return node, nil
}
