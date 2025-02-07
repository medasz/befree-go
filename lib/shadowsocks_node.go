package lib

import (
	"encoding/base64"
	"net/url"
	"strconv"
	"strings"
)

type ShadowsocksNodeProxy struct {
	*ShadowsocksNode
	Type string
}

func NewShadowsocksNodeProxy(n *ShadowsocksNode) *ShadowsocksNodeProxy {
	return &ShadowsocksNodeProxy{n, "ss"}
}
func (n *ShadowsocksNodeProxy) GetType() string {
	return n.Type
}
func (n *ShadowsocksNodeProxy) GetServer() string {
	return n.Server
}
func (n *ShadowsocksNodeProxy) GetPort() int {
	return n.Port
}
func (n *ShadowsocksNodeProxy) GetPassword() any {
	return n.Password
}
func (n *ShadowsocksNodeProxy) GetCipher() any {
	return n.Cipher
}
func (n *ShadowsocksNodeProxy) GetUuid() any {
	return nil
}
func (n *ShadowsocksNodeProxy) GetAlterId() any {
	return nil
}
func (n *ShadowsocksNodeProxy) GetSkipCertVerify() any {
	return nil
}

type ShadowsocksNode struct {
	Name     string
	Server   string
	Port     int
	Cipher   string
	Password string
}

func (t *ShadowsocksNode) ToClashProxy() NodeProxy {
	return NewShadowsocksNodeProxy(t)
}

func (t *ShadowsocksNode) GetName() string {
	return t.Name
}

func (t *ShadowsocksNode) SetName(name string) {
	t.Name = name
}

func NewShadowsocksNode(rawData string) (Node, error) {
	parts := strings.Split(rawData, "#")
	name, err := url.QueryUnescape(parts[1])
	if err != nil {
		name = "xxxx"
	}
	name = strings.TrimSpace(name)
	var cipher, password, server, port string
	if strings.Contains(parts[0], "@") {
		parts2 := strings.Split(parts[0], "@")
		cipherPassword, err := base64.StdEncoding.DecodeString(cleanBase64String(parts2[0]))
		if err != nil {
			return nil, err
		}
		cipherPasswords := strings.Split(string(cipherPassword), ":")
		if cipherPasswords[0] == "ss" {
			cipherPassword2 := strings.Trim(cipherPasswords[1], "//")
			cipherPassword, err = base64.StdEncoding.DecodeString(cleanBase64String(cipherPassword2))
			if err != nil {
				return nil, err
			}
		}
		cipherPasswords = strings.Split(string(cipherPassword), ":")
		cipher = cipherPasswords[0]
		password = cipherPasswords[1]
		serverPort := strings.Split(parts2[1], ":")
		server = strings.TrimSpace(serverPort[0])
		port = strings.TrimSpace(serverPort[1])
	} else {
		parts2, err := base64.StdEncoding.DecodeString(cleanBase64String(parts[0]))
		if err != nil {
			return nil, err
		}
		parts3 := strings.Split(string(parts2), "@")
		cipherPassword := strings.Split(cleanBase64String(parts3[0]), ":")
		cipher = cipherPassword[0]
		password = cipherPassword[1]

		serverPort := strings.Split(parts3[1], ":")
		server = strings.TrimSpace(serverPort[0])
		port = strings.TrimSpace(serverPort[1])
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	return &ShadowsocksNode{
		Name:     name,
		Server:   server,
		Port:     portInt,
		Cipher:   cipher,
		Password: password,
	}, nil
}
