package lib

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

var testData []string

func init() {
	// create an array of strings
	testData = []string{
		"76a91482ffceecbfde154800f5a039b99bc2ebdf44a88488ac",
		"76a91458f546cca253a9b2f099640892c334e5dbd250f388ac",
		"76a91439eeee255ae4e99502e31f60013701d4b9e7b6b788ac",
		"76a9145d57e162859bebfab270c58a41821685fc0e180f88ac",
		"76a91402d4c7b56182686b37312ea5529233254fc904a288ac",
		"76a914f7c44f6a0ae99b0cbc955c174f145e8d00ac7a5a88ac",
		"76a914443315bd2b1e48b4367968b1ad23f5c84dec57f388ac",
		"76a9141faee95d444b7f5b828baa2ab7ab9c399d415b3b88ac",
		"76a9140f1e81fc5e8f4b76feedb6016abe5fadfe8fa8e488ac",
		"76a914bc435ab0d1401625850ce5ca44bd8c19196eb28188ac",
		"76a9147d92755dffb01db79b1919f73fe7ca067042b08e88ac",
		"76a9143c125d98dd3fd48be5e9056fa98b55e2e25aae4d88ac",
		"76a91424dee3443421c58df3b7055550f1ea8fab2adfad88ac",
	}
}

// create a test for the Script() function
func TestScript(t *testing.T) {
	// create a new ExecContext
	exec := &ExecContext{}
	// create a new Instance
	exec.Instance = &Instance{}
	// set the lock to a random string
	for _, lock := range testData {
		exec.Instance.Lock = []byte(lock)
		// call the Script() function
		script, err := exec.Script()
		// check if the error is nil
		if err != nil {
			// if not, fail the test
			t.Fail()
		}
		// check if the script is not nil
		if script == nil {
			// if it is, fail the test
			t.Fail()
		}
	}
}

// create a test for the ParseScript() function
func TestParseScript(t *testing.T) {
	// create a new ExecContext
	exec := &ExecContext{}
	// create a new Instance
	exec.Instance = &Instance{}
	// set the lock to a random script
	for _, lock := range testData {
		decodedLock, err := hex.DecodeString(lock)
		if err != nil {
			log.Printf("Error from DecodeString: %s", err.Error())
			t.Fail()
		}
		exec.Instance.Lock = decodedLock
		// call the Script() function
		script, err := exec.Script()
		// check if the error is nil
		if err != nil {
			// if not, fail the test
			log.Printf("Error from Script(): %s", err.Error())
			t.Fail()
		}
		// check if the script is not nil
		if script == nil {
			// if it is, fail the test
			t.Fail()
			log.Printf("Script() returned a nil value")
		}
		// call the ParseScript() function
		psexec, err := ParseScript(*script)
		// check if the error is nil
		if err != nil {
			// if not, fail the test
			log.Printf("Error from ParseScript(): %s", err.Error())
			t.Fail()
		}
		// check if the exec is not nil
		if psexec == nil {
			// if it is, fail the test
			log.Printf("Error from ParseScript(): execution context returned by ParseScript for script %s is nil", lock)
			t.Fail()
		}
		fmt.Printf("Parsed context lock value: %v\n", hex.EncodeToString(psexec.Instance.Lock))
	}
}

// create a test for the Return() function
func TestReturn(t *testing.T) {
	// create a new ExecContext
	exec := &ExecContext{}
	// call the Return() function
	exec.Return(nil)
	// if it did not panic, pass the test
}

// create a test for the Build() function
func TestBuild(t *testing.T) {
	// create a new ExecContext
	exec := &ExecContext{}
	// call the Build() function
	btOutput, err := exec.Build()
	// if it did not panic, pass the test
	_ = btOutput
	if err != nil {
		t.Fail()
	}
}
