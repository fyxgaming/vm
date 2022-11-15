package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/fyxgaming/vm/lib"
	"github.com/fyxgaming/vm/token/types"
	"github.com/mailru/easyjson"
)

var this *lib.ExecContext

func main() {
	var err error
	this, err = lib.Initialize()
	if err != nil {
		this.Return(err)
		return
	}

	switch this.Method {
	case "Init":
		Init()
	case "Send":
		Send()
	case "recv":
		recv()
	case "Combine":
		Combine()
	}

}

//export Init
func Init() {
	mintReq := types.MintReq{}
	err := easyjson.Unmarshal([]byte(this.CallData), &mintReq)
	if err != nil {
		this.Return(err)
		return
	}

	// Ensure Init is being called by the contract itself
	if this.Parent < 0 ||
		!this.Stack[this.Parent].Instance.Origin.Equal(*this.Contract) {
		this.Return(fmt.Errorf("invalid-parent"))
		return
	}

	this.Instance.Storage = binary.BigEndian.AppendUint64([]byte{}, mintReq.Supply)
	// binary.BigEndian.AppendUint64([]byte{}, )
	this.Instance.Lock = mintReq.Lock
	this.Instance.Satoshis = 1

	log.Println("Emit")
	this.Emit("transfer", [][]byte{{}, mintReq.Lock})

	log.Println("Done")
	this.Return(nil)
}

//export Send
func Send() {
	this, err := lib.Initialize()
	if err != nil {
		this.Return(err)
		return
	}

	if !this.Instance.Kind.Equal(*this.Contract) {
		this.Return(fmt.Errorf("invalid-send"))
		return
	}

	balance := binary.BigEndian.Uint64(this.Instance.Storage)

	for _, exec := range this.Stack {
		if !exec.Instance.Kind.Equal(*this.Contract) {
			continue
		}
		for _, e := range exec.Events {
			if e.Id != "combine" || len(e.Topics) < 2 {
				continue
			}

			if this.Instance.Outpoint.Equal(e.Topics[0]) {
				log.Println("COMBINE", e.Topics)
				balance += binary.BigEndian.Uint64(e.Topics[1])
			}
		}
	}

	var sendReq types.SendReq
	if len(this.CallData) > 0 {
		err = easyjson.Unmarshal([]byte(this.CallData), &sendReq)
		if err != nil {
			this.Return(err)
			return
		}

		for _, send := range sendReq.Sends {
			if send.Amount > balance {
				this.Return(&lib.Error{
					Code: 402,
					Err:  "insufficient-balance",
				})
				return
			}
			balance -= send.Amount

			this.Emit("transfer", [][]byte{
				send.To,
				this.Instance.Lock,
				binary.BigEndian.AppendUint64([]byte{}, send.Amount),
			})

			sendData, err := send.MarshalJSON()
			if err != nil {
				this.Return(err)
				return
			}
			this.Spawn(this.Contract, "recv", sendData)

		}
	}

	this.Instance.Storage = binary.BigEndian.AppendUint64([]byte{}, balance)
	if balance == 0 {
		this.Instance.Destroy()
	}

	this.Return(nil)
}

//export recv
func recv() {
	this, err := lib.Initialize()
	if err != nil {
		this.Return(err)
		return
	}

	if this.Parent < 0 ||
		!this.Stack[this.Parent].Instance.Kind.Equal(*this.Contract) {
		this.Return(fmt.Errorf("invalid-recv"))
		return
	}

	var send types.Send
	err = easyjson.Unmarshal([]byte(this.CallData), &send)
	if err != nil {
		this.Return(err)
		return
	}

	this.Instance.Satoshis = 1
	this.Instance.Storage = binary.BigEndian.AppendUint64([]byte{}, send.Amount)
	this.Instance.Lock = send.To

	this.Return(nil)
}

//export Combine
func Combine() {
	this, err := lib.Initialize()
	if err != nil {
		this.Return(err)
		return
	}

	if !this.Instance.Kind.Equal(*this.Contract) {
		this.Return(fmt.Errorf("invalid-combine"))
		return
	}

	balance := binary.BigEndian.Uint64(this.Instance.Storage)
	this.Emit("combine", [][]byte{
		this.CallData,
		binary.BigEndian.AppendUint64([]byte{}, balance),
	})

	this.Instance.Storage = binary.BigEndian.AppendUint64([]byte{}, 0)
	this.Instance.Destroy()
	this.Return(nil)
}
