package lib

import (
	"log"
	"reflect"
	"testing"

	"github.com/libsv/go-bt/v2/bscript"
)

// create a test for the Initialize() function
func TestInitialize(t *testing.T) {
	// call the Initialize() function
	exec, err := Initialize()
	// check if the error is nil
	if err != nil {
		// if not, fail the test
		t.Fail()
	}
	// check if the exec is not nil
	if exec == nil {
		// if it is, fail the test
		t.Fail()
	}
}

// create a test for the Destroy() function
func TestDestroy(t *testing.T) {
	// create a new ExecContext
	exec, err := Initialize()
	if err != nil {
		log.Println("Could not run Initialize() function")
		t.Fail()
	}
	// call the Destroy() function
	exec.Instance.Destroy()
	// check if the error is nil
	if exec.Instance.Satoshis != 0 || !reflect.DeepEqual(exec.Instance.Lock, []byte{bscript.OpFALSE}) {
		// if not, fail the test
		t.Fail()
	}
}
