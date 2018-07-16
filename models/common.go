package models

const (
	ErrSystem   = -1
	ErrNotFound = 1
)

type CodeMsg struct {
	Code int    `json:"code"`
	Msg string `json:"msg"`
}

func NewErrorMsg(Msg string) *CodeMsg {
	return &CodeMsg{1, Msg}
}

func NewNormalMsg(Msg string) *CodeMsg {
	return &CodeMsg{0, Msg}
}