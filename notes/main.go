package main

import (
	"github.com/fyxgaming/vm/lib"
	"github.com/fyxgaming/vm/notes/types"
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

	req := &types.Note{}
	err = easyjson.Unmarshal(this.CallData, req)
	if err != nil {
		this.Return(err)
		return
	}
	this.Instance.Satoshis = 1
	this.Instance.Lock = req.Lock
	this.Instance.Storage = []byte(req.Data)
	this.Return(nil)
}
