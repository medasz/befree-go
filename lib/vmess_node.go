package lib

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
)

type VMessNode struct {
	Name   string
	Server string
	Port   int
	UUID   string
	Cipher string
}

func (v *VMessNode) ToClashProxy() {

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
	Tls            string `json:"tls"`
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
		Name:   vmessNode.Ps,
		Server: vmessNode.Add,
		UUID:   vmessNode.Id,
		Cipher: "auto",
	}
	switch port := vmessNode.Port.(type) {
	case int:
		node.Port = port
	case string:
		node.Port, err = strconv.Atoi(port)
	}
	if vmessNode.Cipher != "" {
		node.Cipher = vmessNode.Cipher
	}
	return node, nil
}
