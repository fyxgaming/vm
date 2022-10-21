package types

type MintReq struct {
	Supply uint64 `json:"supply"`
	Lock   []byte `json:"lock"`
}
type Send struct {
	To     []byte `json:"to"`
	Amount uint64 `json:"amount"`
}
type SendReq struct {
	Sends []*Send `json:"sends"`
}
