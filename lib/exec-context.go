package lib

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"github.com/fyxgaming/vm/util"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/bscript"
)

type Action byte

const (
	Auth   Action = 0
	Mint   Action = 1
	Call   Action = 2
	Spawn  Action = 3
	Deploy Action = 4
)

type ExecContext struct {
	Action   Action         `json:"action"`
	Contract *Outpoint      `json:"contract,omitempty"`
	Method   string         `json:"method,omitempty"`
	CallData []byte         `json:"callData,omitempty"`
	Stack    []*ExecContext `json:"stack,omitempty"`
	Parent   int32          `json:"parent,omitempty"`
	Instance *Instance      `json:"instance"`
	Events   []*Event       `json:"events,omitempty"`
	Children []*Child       `json:"children,omitempty"`
}

func (exec *ExecContext) Return(err error) {
	if err != nil {
		os.Stderr.WriteString(err.Error())
		panic("error")
	} else {
		retData, err := exec.MarshalJSON()
		if err != nil {
			exec.Return(err)
			return
		}
		_, err = os.Stdout.Write(retData)
		if err != nil {
			exec.Return(err)
			return
		}
	}
}

func (e *ExecContext) Script() (script *bscript.Script, err error) {
	ser, err := e.MarshalJSON()
	if err != nil {
		return
	}
	fmt.Printf("Exec.Script Data: %s\n", ser)
	script = bscript.NewFromBytes(e.Instance.Lock)
	if len(*script) == 0 {
		err = script.AppendOpcodes(bscript.OpFALSE)
		if err != nil {
			return
		}
	}
	err = script.AppendOpcodes(bscript.OpRETURN)
	if err != nil {
		return
	}
	err = script.AppendPushDataString("fyx")
	if err != nil {
		return
	}
	err = script.AppendPushData(binary.AppendVarint([]byte{}, int64(e.Action)))
	if err != nil {
		return
	}

	err = script.AppendPushData(binary.AppendVarint([]byte{}, int64(e.Parent)))
	if err != nil {
		return
	}
	if e.Contract != nil {
		err = script.AppendPushData(*e.Contract)
		if err != nil {
			return
		}
	} else {
		err = script.AppendOpcodes(bscript.OpFALSE)
		if err != nil {
			return
		}
	}
	err = script.AppendPushDataString(e.Method)
	if err != nil {
		return
	}
	if len(e.CallData) > 0 {
		err = script.AppendPushData(e.CallData)
		if err != nil {
			return
		}
	} else {
		err = script.AppendOpcodes(bscript.OpFALSE)
		if err != nil {
			return
		}
	}
	length := byte(len(e.Events))
	err = script.AppendPushData([]byte{length})
	if err != nil {
		return
	}
	for _, event := range e.Events {
		err = script.AppendPushDataString(event.Id)
		if err != nil {
			return
		}
		length := byte(len(event.Topics))
		err = script.AppendPushData([]byte{length})
		if err != nil {
			return
		}
		err = script.AppendPushDataArray(event.Topics)
		if err != nil {
			return
		}
	}

	fmt.Printf("Exec.Script Script: %x\n", script)
	return
}

func ParseScript(script []byte) (exec *ExecContext, err error) {
	fmt.Printf("Exec.ParseScript Script: %x\n", script)

	ops, err := bscript.DecodeParts(script)
	if err != nil {
		return
	}

	var op []byte
	var done bool
	lock := [][]byte{}
	for op, ops, done = util.Unshift(ops); !done; op, ops, done = util.Unshift(ops) {
		if len(op) == 1 && op[0] == bscript.OpRETURN {
			break
		}
		lock = append(lock, op)
	}
	if len(lock) == 0 {
		lock = append(lock, []byte{bscript.OpFALSE})
	}

	op, ops, done = util.Unshift(ops)
	if done || string(op) != "fyx" {
		return
	}

	lockScript, err := bscript.EncodeParts(lock)
	if err != nil {
		return
	}
	exec = &ExecContext{
		Instance: &Instance{
			Lock: lockScript,
		},
	}

	if op, ops, done = util.Unshift(ops); done {
		return
	}
	val, length := binary.Varint(op)
	if length == 0 {
		err = fmt.Errorf("invalid-action")
		return
	}
	exec.Action = Action(val)

	if op, ops, done = util.Unshift(ops); done {
		return
	}

	val, length = binary.Varint(op)
	if length == 0 || val < -1 || val > 2^16 {
		err = fmt.Errorf("invalid-parent")
		return
	}
	exec.Parent = int32(val)

	if op, ops, done = util.Unshift(ops); done {
		return
	}
	if len(op) == 36 {
		exec.Contract = NewOutpointFromBytes(op)
	}

	if op, ops, done = util.Unshift(ops); done {
		return
	}
	exec.Method = string(op)

	if op, ops, done = util.Unshift(ops); done {
		return
	}
	exec.CallData = op

	if op, ops, done = util.Unshift(ops); done {
		return
	}
	for i := uint8(0); i < op[0]; i++ {
		if op, ops, done = util.Unshift(ops); done {
			return
		}
		exec.Events = append(exec.Events, &Event{
			Id: string(op),
		})
		if op, ops, done = util.Unshift(ops); done {
			return
		}
		for j := uint8(0); j < op[0]; j++ {
			if op, ops, done = util.Unshift(ops); done {
				return
			}
			exec.Events[i].Topics = append(exec.Events[i].Topics, op)
		}
	}

	ser, err := exec.MarshalJSON()
	if err != nil {
		return
	}
	fmt.Printf("Exec.ParseScript Data: %s\n", ser)
	return
}

func (e *ExecContext) Build() (output *bt.Output, err error) {
	output = &bt.Output{}
	output.Satoshis = e.Instance.Satoshis
	output.LockingScript, err = e.Script()
	return
}

func (e *ExecContext) Spawn(contract *Outpoint, method string, callData []byte) {
	e.Children = append(e.Children, &Child{
		Contract: contract,
		Method:   method,
		CallData: callData,
	})
}

func (e *ExecContext) Emit(event string, topics [][]byte) {
	e.Events = append(e.Events, &Event{
		Id:     event,
		Topics: topics,
	})
}

func ParseScriptASM(script []byte) (exec *ExecContext, err error) {
	fmt.Printf("Exec.ParseScriptASM Script: %x\n", script)

	fromBytesScript := bscript.NewFromBytes(script)
	strScript, err := fromBytesScript.ToASM()
	if err != nil {
		return
	}
	fmt.Printf("Exec.ParseScript strscript: %s\n", strScript)
	mystrary := strings.Split(strScript, " ")

	var lock, op string
	var ops []string
	var done bool
	for op, ops, done = util.Unshift(mystrary); !done; op, ops, done = util.Unshift(ops) {
		if op == "OP_RETURN" {
			lock = strings.TrimRight(lock, " ")
			break
		}
		lock += op + " "
	}
	if len(lock) == 0 {
		lock = string(bscript.OpFALSE)
	}

	lockScript, err := bscript.NewFromASM(lock)
	if err != nil {
		return
	}

	fmt.Printf("Exec.ParseScript Script: %s\n", lockScript)

	exec = &ExecContext{
		Instance: &Instance{
			Lock: *lockScript,
		},
	}

	op, ops, done = util.Unshift(ops)
	fyxScript, err := bscript.NewFromHexString(op)
	if err != nil {
		return
	}
	if done || string(*fyxScript) != "fyx" {
		return
	}

	if op, ops, done = util.Unshift(ops); done {
		return
	}
	fyxScript, err = bscript.NewFromASM(op)
	if err != nil {
		return
	}
	val, length := binary.Varint(*fyxScript)
	if length == 0 {
		err = fmt.Errorf("invalid-action")
		return
	}
	exec.Action = Action(val)

	if op, ops, done = util.Unshift(ops); done {
		return
	}
	fyxScript, err = bscript.NewFromASM(op)
	if err != nil {
		return
	}
	val, length = binary.Varint(*fyxScript)
	if length == 0 || val < -1 || val > 2^16 {
		err = fmt.Errorf("invalid-parent")
		return
	}
	exec.Parent = int32(val)

	if op, ops, done = util.Unshift(ops); done {
		return
	}
	fyxScript, err = bscript.NewFromASM(op)
	if err != nil {
		return
	}
	if len(op) == 36 {
		exec.Contract = NewOutpointFromBytes(*fyxScript)
	}

	if op, ops, done = util.Unshift(ops); done {
		return
	}
	fyxScript, err = bscript.NewFromHexString(op)
	if err != nil {
		return
	}
	exec.Method = string(*fyxScript)

	if op, ops, done = util.Unshift(ops); done {
		return
	}
	fyxScript, err = bscript.NewFromHexString(op)
	if err != nil {
		return
	}
	val, _ = binary.Varint(*fyxScript)
	if val != int64(bscript.OpFALSE) {
		exec.CallData = []byte(*fyxScript)
	}

	if op, ops, done = util.Unshift(ops); done {
		return
	}
	fyxScript, err = bscript.NewFromASM(op)
	if err != nil {
		return
	}
	for i := uint8(0); i < []byte(*fyxScript)[0]; i++ {
		if op, ops, done = util.Unshift(ops); done {
			return
		}
		fyxScript, err = bscript.NewFromASM(op)
		if err != nil {
			return
		}
		exec.Events = append(exec.Events, &Event{
			Id: string(*fyxScript),
		})
		if op, ops, done = util.Unshift(ops); done {
			return
		}
		var j_fyxScript *bscript.Script
		j_fyxScript, err = bscript.NewFromASM(op)
		if err != nil {
			return
		}
		for j := uint8(0); j < []byte(*j_fyxScript)[0]; j++ {
			if op, ops, done = util.Unshift(ops); done {
				return
			}
			j_fyxScript, err = bscript.NewFromASM(op)
			if err != nil {
				return
			}
			exec.Events[i].Topics = append(exec.Events[i].Topics, []byte(*j_fyxScript))
		}
	}

	ser, err := exec.MarshalJSON()
	if err != nil {
		return
	}
	fmt.Printf("Exec.ParseScript Data: %s\n", ser)
	return
}
