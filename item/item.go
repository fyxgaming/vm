package main

import "github.com/fyxgaming/vm/lib"

func main() {}

//export Issue
func Issue() int {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	this.Instance.Storage = this.CallData

}
