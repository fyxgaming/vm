package main

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/fyxgaming/vm/lib"
	"github.com/fyxgaming/vm/token/types"
	"github.com/mailru/easyjson"
)

func main() {}

//export Init
func Init() (retCode int) {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	mintReq := types.MintReq{}
	err = easyjson.Unmarshal([]byte(this.CallData), &mintReq)
	if err != nil {
		return this.Return(err)
	}

	// Ensure Init is being called by the contract itself
	if this.Parent < 0 ||
		!this.Stack[this.Parent].Instance.Origin.Equal(*this.Contract) {
		return this.Return(fmt.Errorf("invalid-parent"))
	}

	this.Instance.Storage = binary.BigEndian.AppendUint64([]byte{}, mintReq.Supply)
	// binary.BigEndian.AppendUint64([]byte{}, )
	this.Instance.Lock = mintReq.Lock
	this.Instance.Satoshis = 1

	log.Println("Emit")
	this.Emit("transfer", [][]byte{{}, mintReq.Lock})

	log.Println("Done")
	return this.Return(nil)
}

//export Send
func Send() int {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	if !this.Instance.Kind.Equal(*this.Contract) {
		return this.Return(fmt.Errorf("invalid-send"))
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
			return this.Return(err)
		}

		for _, send := range sendReq.Sends {
			if send.Amount > balance {
				return this.Return(&lib.Error{
					Code: 402,
					Err:  "insufficient-balance",
				})
			}
			balance -= send.Amount

			this.Emit("transfer", [][]byte{
				send.To,
				this.Instance.Lock,
				binary.BigEndian.AppendUint64([]byte{}, send.Amount),
			})

			sendData, err := send.MarshalJSON()
			if err != nil {
				return this.Return(err)
			}
			this.Spawn(this.Contract, "recv", sendData)

		}
	}

	this.Instance.Storage = binary.BigEndian.AppendUint64([]byte{}, balance)
	if balance == 0 {
		this.Instance.Destroy()
	}

	return this.Return(nil)
}

//export recv
func recv() int {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	if this.Parent < 0 ||
		!this.Stack[this.Parent].Instance.Kind.Equal(*this.Contract) {
		return this.Return(fmt.Errorf("invalid-recv"))
	}

	var send types.Send
	err = easyjson.Unmarshal([]byte(this.CallData), &send)
	if err != nil {
		return this.Return(err)
	}

	this.Instance.Satoshis = 1
	this.Instance.Storage = binary.BigEndian.AppendUint64([]byte{}, send.Amount)
	this.Instance.Lock = send.To

	return this.Return(nil)
}

//export Combine
func Combine() int {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	if !this.Instance.Kind.Equal(*this.Contract) {
		return this.Return(fmt.Errorf("invalid-combine"))
	}

	balance := binary.BigEndian.Uint64(this.Instance.Storage)
	this.Emit("combine", [][]byte{
		this.CallData,
		binary.BigEndian.AppendUint64([]byte{}, balance),
	})

	this.Instance.Storage = binary.BigEndian.AppendUint64([]byte{}, 0)
	this.Instance.Destroy()
	return this.Return(nil)
}
