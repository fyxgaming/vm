package types

type Note struct {
	Lock []byte `json:"lock"`
	Note string `json:"note"`
}
