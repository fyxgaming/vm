package main

import "github.com/fyxgaming/vm/lib"

func main() {}

//export Init
func Init() int {
	this, err := lib.Initialize()
	if err != nil {
		return this.Return(err)
	}

	this.Instance.Lock = []byte{}
	this.Instance.Satoshis = 1

}
