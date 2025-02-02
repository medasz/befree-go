package lib

type Node interface {
	ToClashProxy()
	GetName() string
	SetName(string)
}
