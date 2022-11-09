package main

import (
	"strconv"

	"github.com/fyxgaming/vm/lib"
	"github.com/fyxgaming/vm/receipt/types"
	"github.com/mailru/easyjson"
)

func main() {}

//export Grant
func Grant() (retCode int) {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	grant := types.BattleReceipt{}
	err = easyjson.Unmarshal(this.CallData, &grant)
	if err != nil {
		return this.Return(err)
	}
	this.Instance.Lock = grant.Lock
	this.Instance.Storage = this.CallData

	return this.Return(nil)
}

//export Claim
func Claim() (retCode int) {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	grant := types.BattleReceipt{}
	err = easyjson.Unmarshal(this.Instance.Storage, &grant)
	if err != nil {
		return this.Return(err)
	}

	this.Emit("claim-xp", []string{
		grant.Fighter.String(),
		strconv.FormatInt(int64(grant.Xp), 10),
		strconv.FormatBool(grant.Win),
		strconv.FormatBool(grant.PvP),
	})

	this.Destroy()
	return this.Return(nil)
}
