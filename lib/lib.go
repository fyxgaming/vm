package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/bscript"
	"github.com/mailru/easyjson"
)

func Initialize() (exec *ExecContext, err error) {
	EXEC := os.Getenv("EXEC")
	exec = &ExecContext{}
	err = easyjson.Unmarshal([]byte(EXEC), exec)
	if err != nil {
		return
	}
	return
}

func (i *Instance) Destroy() {
	i.Satoshis = 0
	i.Lock = []byte{bscript.OpFALSE}
	// i.Storage = []byte{}
}

type Txo struct {
	Outpoint *Outpoint       `json:"outpoint,omitempty"`
	Satoshis uint64          `json:"sats"`
	Lock     []byte          `json:"lock"`
	Script   *bscript.Script `json:"script"`
}

type Instance struct {
	Txo      *Txo      `json:"txo,omitempty"`
	Outpoint *Outpoint `json:"outpoint,omitempty"`
	Origin   *Outpoint `json:"origin,omitempty"`
	Nonce    uint64    `json:"nonce,omitempty"`
	Kind     *Outpoint `json:"kind,omitempty"`
	Satoshis uint64    `json:"sats"`
	Lock     []byte    `json:"lock"`
	Storage  []byte    `json:"store,omitempty"`
	Code     *Outpoint `json:"code"`
}

type Child struct {
	Contract *Outpoint `json:"contract"`
	Method   string    `json:"method"`
	CallData []byte    `json:"callData"`
}

// type Parent struct {
// 	Idx      int       `json:"idx"`
// 	Outpoint *Outpoint `json:"outpoint"`
// 	Kind     *Outpoint `json:"kind"`
// 	Lock     []byte    `json:"lock"`
// 	Origin   *Outpoint `json:"origin"`
// 	Nonce    uint64    `json:"nonce"`
// }

type Event struct {
	Id     string   `json:"id"`
	Topics [][]byte `json:"topics"`
}

type Error struct {
	Code int
	Err  string
}

func (err *Error) Error() string {
	if err.Code > 0 {
		return fmt.Sprintf("%d-%s", err.Code, err.Err)
	} else {
		return err.Err
	}
}

type File struct {
	Outpoint *Outpoint `json:"outpoint,omitempty"`
	Data     []byte    `json:"data"`
	Type     string    `json:"type,omitempty"`
	Encoding string    `json:"enc,omitempty"`
	Name     string    `json:"name,omitempty"`
	Size     uint32    `json:"size,omitempty"`
	Hash     []byte    `json:"hash,omitempty"`
}

func (file *File) Build() (output *bt.Output, err error) {
	ops := [][]byte{
		{bscript.OpFALSE},
		{bscript.OpRETURN},
		[]byte("19HxigV4QyBv3tHpQVcUEQyq1pzZVdoAut"),
		file.Data,
	}

	ops = append(ops, []byte(file.Type))
	ops = append(ops, []byte(file.Encoding))
	ops = append(ops, []byte(file.Name))

	script, err := bscript.EncodeParts(ops)
	if err != nil {
		log.Println(err)
		return
	}

	output = &bt.Output{
		Satoshis:      0,
		LockingScript: bscript.NewFromBytes(script),
	}

	return
}
