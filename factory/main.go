package main

import "github.com/fyxgaming/vm/lib"

func main() {}

//export Initialize
func Initialize() (retCode int) {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	this.Spawn(this.Instance.Origin, "Init", this.CallData)
	this.Destroy()

	return this.Return(nil)
}
