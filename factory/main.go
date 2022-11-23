package main

import (
	"log"
	"os"

	"github.com/fyxgaming/vm/lib"
)

var this *lib.ExecContext

func main() {
	var err error
	log.Println("EXEC:", os.Getenv("EXEC"))
	this, err = lib.Initialize()
	if err != nil {
		log.Println("Initialize Err:", err)
		this.Return(err)
		return
	}

	this.Spawn(this.Instance.Origin, "Init", this.CallData)
	this.Instance.Destroy()

	this.Return(nil)
}
