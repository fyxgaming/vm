package main

import (
	"encoding/hex"
	"fmt"

	"github.com/fyxgaming/vm/item/types"
	"github.com/fyxgaming/vm/lib"
	"github.com/mailru/easyjson"
)

const VALIDATOR_LOCK = "" //FYX:PUBKEY:

func main() {}

//export New
func New() int {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	if this.Parent == nil || hex.EncodeToString(this.Parent.Lock) != VALIDATOR_LOCK {
		return this.Return(fmt.Errorf("invalid-minter"))
	}

	var itemReq *types.ItemReq
	err = easyjson.Unmarshal([]byte(this.CallData), itemReq)
	if err != nil {
		return this.Return(err)
	}

	this.Emit("ItemMint", itemReq.Topics)
	this.Instance.Storage = itemReq.Storage
	return this.Return(nil)
}
