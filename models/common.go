package models

const (
	ErrSystem   = -1
	ErrNotFound = 1
)

type CodeMsg struct {
	Code int    `json:"code"`
	Msg string `json:"msg"`
}

func NewErrorMsg(msg string) *CodeMsg {
	return &CodeMsg{1, msg}
}

func NewNormalMsg(info string) *CodeMsg {
	return &CodeMsg{0, msg}
}