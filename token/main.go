package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"strconv"

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
	if this.Parent == nil ||
		!this.Parent.Origin.Equal(*this.Contract) {
		return this.Return(fmt.Errorf("invalid-parent"))
	}

	this.Instance.Storage = strconv.FormatUint(mintReq.Supply, 10)
	this.Instance.Lock = mintReq.Lock
	this.Instance.Satoshis = 1

	log.Println("Emit")
	this.Emit("transfer", []string{"", base64.StdEncoding.EncodeToString(mintReq.Lock)})

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

	balance, err := strconv.ParseUint(this.Instance.Storage, 10, 64)
	if err != nil {
		return this.Return(err)
	}

	for _, exec := range this.Stack {
		if !exec.Instance.Kind.Equal(*this.Contract) {
			continue
		}
		for _, e := range exec.Events {
			if e.Id != "combine" || len(e.Topics) < 2 {
				continue
			}

			if this.Instance.Outpoint.String() == e.Topics[0] {
				log.Println("COMBINE", e.Topics)
				amount, err := strconv.ParseUint(e.Topics[1], 10, 64)
				if err != nil {
					return this.Return(err)
				}
				balance += amount
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

			this.Emit("transfer", []string{
				base64.StdEncoding.EncodeToString(send.To),
				base64.StdEncoding.EncodeToString(this.Instance.Lock),
				strconv.FormatUint(send.Amount, 10),
			})

			sendData, err := send.MarshalJSON()
			if err != nil {
				return this.Return(err)
			}
			this.Spawn(this.Contract, "recv", string(sendData))

		}
	}

	this.Instance.Storage = strconv.FormatUint(balance, 10)
	if balance == 0 {
		this.Destroy()
	}

	return this.Return(nil)
}

//export recv
func recv() int {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	if this.Parent == nil ||
		!this.Parent.Kind.Equal(*this.Contract) {
		return this.Return(fmt.Errorf("invalid-recv"))
	}

	var send types.Send
	err = easyjson.Unmarshal([]byte(this.CallData), &send)
	if err != nil {
		return this.Return(err)
	}

	this.Instance.Satoshis = 1
	this.Instance.Storage = strconv.FormatUint(send.Amount, 10)
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

	balance, err := strconv.ParseUint(this.Instance.Storage, 10, 64)
	if err != nil {
		return this.Return(err)
	}

	this.Emit("combine", []string{
		this.CallData,
		strconv.FormatUint(balance, 10),
	})

	this.Destroy()
	return this.Return(nil)
}
