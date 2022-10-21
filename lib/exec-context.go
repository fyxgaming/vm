package lib

import (
	"os"

	"github.com/fyxgaming/vm/util"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/bscript"
)

type ExecContext struct {
	Action   string         `json:"action"`
	Contract *Outpoint      `json:"contract"`
	Method   string         `json:"method"`
	CallData []byte         `json:"callData,omitempty"`
	Stack    []*ExecContext `json:"stack,omitempty"`
	Parent   *Parent        `json:"parent,omitempty"`
	Instance *Instance      `json:"instance"`
	Events   []*Event       `json:"events,omitempty"`
	Spawn    []*Spawn       `json:"spawn,omitempty"`
}

func (exec *ExecContext) Return(err error) (retCode int) {
	if err != nil {
		os.Stderr.WriteString(err.Error())
		if err, ok := err.(*Error); ok {
			retCode = err.Code
		} else {
			retCode = 500
		}
	} else {
		retData, err := exec.MarshalJSON()
		if err != nil {
			return exec.Return(err)
		}
		_, err = os.Stdout.Write(retData)
		if err != nil {
			return exec.Return(err)
		}
	}
	return
}

func (e *ExecContext) Destroy() {
	e.Instance.Satoshis = 0
	e.Instance.Lock = []byte{}
	e.Instance.Storage = []byte{}
}

func (e *ExecContext) Script() (script *bscript.Script, err error) {
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
	err = script.AppendPushDataString(e.Action)
	if err != nil {
		return
	}
	if e.Instance.Origin != nil {
		err = script.AppendPushData(*e.Instance.Origin)
		if err != nil {
			return
		}
	} else {
		err = script.AppendOpcodes(bscript.OpFALSE)
		if err != nil {
			return
		}
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
	return
}

func ParseScript(script []byte) (exec *ExecContext, err error) {
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
	exec.Action = string(op)

	if op, ops, done = util.Unshift(ops); done {
		return
	}
	if len(op) == 36 {
		exec.Instance.Origin = NewOutpointFromBytes(op)
	}

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

	return
}

func (e *ExecContext) Build() (output *bt.Output, err error) {
	output = &bt.Output{}
	output.Satoshis = e.Instance.Satoshis
	output.LockingScript, err = e.Script()
	return
}

func (e *ExecContext) Mint(contract *Outpoint, method string, callData []byte) {
	e.Spawn = append(e.Spawn, &Spawn{
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
