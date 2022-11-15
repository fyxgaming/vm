package main

import "github.com/fyxgaming/vm/lib"

var this *lib.ExecContext

func main() {
	var err error
	this, err = lib.Initialize()
	if err != nil {
		this.Return(err)
		return
	}

	this.Spawn(this.Instance.Origin, "Init", this.CallData)
	this.Instance.Destroy()

	this.Return(nil)
}
