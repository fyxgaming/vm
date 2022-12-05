package types

type Note struct {
	Data string `json:"data"`
	Lock []byte `json:"lock"`
}
