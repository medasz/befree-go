package lib

type Node interface {
	ToClashProxy() NodeProxy
	GetName() string
	SetName(string)
}

type NodeProxy interface {
	GetName() string
	GetType() string
	GetServer() string
	GetPort() int
	GetPassword() any
	GetCipher() any
	GetUuid() any
	GetAlterId() any
	GetSkipCertVerify() any
}
