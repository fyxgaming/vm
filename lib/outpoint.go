package lib

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
)

type Outpoint []byte

func NewOutpoint(txid []byte, vout uint32) *Outpoint {
	if len(txid) == 32 {
		o := make([]byte, 0, 36)
		o = append(o, txid...)
		o = binary.BigEndian.AppendUint32(o, vout)
		return (*Outpoint)(&o)
	}
	return nil
}

func NewOutpointFromBytes(b []byte) *Outpoint {
	o := make([]byte, 0, len(b))
	o = append(o, b...)
	return (*Outpoint)(&o)
}

func NewOutpointFromString(s string) (*Outpoint, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return (*Outpoint)(&b), nil
}

func (o Outpoint) Txid() []byte {
	return o[:32]
}

func (o Outpoint) Vout() uint32 {
	return binary.BigEndian.Uint32(o[32:])
}

func (o Outpoint) String() string {
	return base64.StdEncoding.EncodeToString(o)
}

func (o *Outpoint) Equal(c []byte) bool {
	return bytes.Equal(*o, c)
}

// func (o *Outpoint) MarshalJSON() ([]byte, error) {
// 	return []byte(o.String()), nil
// }

// // UnmarshalJSON deserializes ByteArray to hex
// func (o *Outpoint) UnmarshalJSON(data []byte) error {
// 	buf, err := hex.DecodeString(string(data))

// 	*o = buf
// 	// var x string
// 	// err := json.Unmarshal(data, &x)
// 	// if err == nil {
// 	// 	str, e := hex.DecodeString(x)
// 	// 	*s = ByteString([]byte(str))
// 	// 	err = e
// 	// }

// 	return err
// }
