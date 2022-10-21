// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package lib

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE910b2f5DecodeFyxVmLib(in *jlexer.Lexer, out *ExecContext) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "action":
			out.Action = string(in.String())
		case "contract":
			if in.IsNull() {
				in.Skip()
				out.Contract = nil
			} else {
				if out.Contract == nil {
					out.Contract = new(Outpoint)
				}
				if in.IsNull() {
					in.Skip()
					*out.Contract = nil
				} else {
					*out.Contract = in.Bytes()
				}
			}
		case "method":
			out.Method = string(in.String())
		case "callData":
			if in.IsNull() {
				in.Skip()
				out.CallData = nil
			} else {
				out.CallData = in.Bytes()
			}
		case "stack":
			if in.IsNull() {
				in.Skip()
				out.Stack = nil
			} else {
				in.Delim('[')
				if out.Stack == nil {
					if !in.IsDelim(']') {
						out.Stack = make([]*ExecContext, 0, 8)
					} else {
						out.Stack = []*ExecContext{}
					}
				} else {
					out.Stack = (out.Stack)[:0]
				}
				for !in.IsDelim(']') {
					var v3 *ExecContext
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(ExecContext)
						}
						(*v3).UnmarshalEasyJSON(in)
					}
					out.Stack = append(out.Stack, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "parent":
			if in.IsNull() {
				in.Skip()
				out.Parent = nil
			} else {
				if out.Parent == nil {
					out.Parent = new(Parent)
				}
				easyjsonE910b2f5DecodeFyxVmLib1(in, out.Parent)
			}
		case "instance":
			if in.IsNull() {
				in.Skip()
				out.Instance = nil
			} else {
				if out.Instance == nil {
					out.Instance = new(Instance)
				}
				easyjsonE910b2f5DecodeFyxVmLib2(in, out.Instance)
			}
		case "events":
			if in.IsNull() {
				in.Skip()
				out.Events = nil
			} else {
				in.Delim('[')
				if out.Events == nil {
					if !in.IsDelim(']') {
						out.Events = make([]*Event, 0, 8)
					} else {
						out.Events = []*Event{}
					}
				} else {
					out.Events = (out.Events)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *Event
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Event)
						}
						easyjsonE910b2f5DecodeFyxVmLib3(in, v4)
					}
					out.Events = append(out.Events, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "spawn":
			if in.IsNull() {
				in.Skip()
				out.Spawn = nil
			} else {
				in.Delim('[')
				if out.Spawn == nil {
					if !in.IsDelim(']') {
						out.Spawn = make([]*Spawn, 0, 8)
					} else {
						out.Spawn = []*Spawn{}
					}
				} else {
					out.Spawn = (out.Spawn)[:0]
				}
				for !in.IsDelim(']') {
					var v5 *Spawn
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						if v5 == nil {
							v5 = new(Spawn)
						}
						easyjsonE910b2f5DecodeFyxVmLib4(in, v5)
					}
					out.Spawn = append(out.Spawn, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE910b2f5EncodeFyxVmLib(out *jwriter.Writer, in ExecContext) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"action\":"
		out.RawString(prefix[1:])
		out.String(string(in.Action))
	}
	{
		const prefix string = ",\"contract\":"
		out.RawString(prefix)
		if in.Contract == nil {
			out.RawString("null")
		} else {
			out.Base64Bytes(*in.Contract)
		}
	}
	{
		const prefix string = ",\"method\":"
		out.RawString(prefix)
		out.String(string(in.Method))
	}
	if len(in.CallData) != 0 {
		const prefix string = ",\"callData\":"
		out.RawString(prefix)
		out.Base64Bytes(in.CallData)
	}
	if len(in.Stack) != 0 {
		const prefix string = ",\"stack\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v10, v11 := range in.Stack {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					(*v11).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Parent != nil {
		const prefix string = ",\"parent\":"
		out.RawString(prefix)
		easyjsonE910b2f5EncodeFyxVmLib1(out, *in.Parent)
	}
	{
		const prefix string = ",\"instance\":"
		out.RawString(prefix)
		if in.Instance == nil {
			out.RawString("null")
		} else {
			easyjsonE910b2f5EncodeFyxVmLib2(out, *in.Instance)
		}
	}
	if len(in.Events) != 0 {
		const prefix string = ",\"events\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v12, v13 := range in.Events {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					easyjsonE910b2f5EncodeFyxVmLib3(out, *v13)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Spawn) != 0 {
		const prefix string = ",\"spawn\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v14, v15 := range in.Spawn {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					easyjsonE910b2f5EncodeFyxVmLib4(out, *v15)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ExecContext) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE910b2f5EncodeFyxVmLib(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ExecContext) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE910b2f5EncodeFyxVmLib(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ExecContext) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE910b2f5DecodeFyxVmLib(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ExecContext) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE910b2f5DecodeFyxVmLib(l, v)
}
func easyjsonE910b2f5DecodeFyxVmLib4(in *jlexer.Lexer, out *Spawn) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "contract":
			if in.IsNull() {
				in.Skip()
				out.Contract = nil
			} else {
				if out.Contract == nil {
					out.Contract = new(Outpoint)
				}
				if in.IsNull() {
					in.Skip()
					*out.Contract = nil
				} else {
					*out.Contract = in.Bytes()
				}
			}
		case "method":
			out.Method = string(in.String())
		case "callData":
			if in.IsNull() {
				in.Skip()
				out.CallData = nil
			} else {
				out.CallData = in.Bytes()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE910b2f5EncodeFyxVmLib4(out *jwriter.Writer, in Spawn) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"contract\":"
		out.RawString(prefix[1:])
		if in.Contract == nil {
			out.RawString("null")
		} else {
			out.Base64Bytes(*in.Contract)
		}
	}
	{
		const prefix string = ",\"method\":"
		out.RawString(prefix)
		out.String(string(in.Method))
	}
	{
		const prefix string = ",\"callData\":"
		out.RawString(prefix)
		out.Base64Bytes(in.CallData)
	}
	out.RawByte('}')
}
func easyjsonE910b2f5DecodeFyxVmLib3(in *jlexer.Lexer, out *Event) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = string(in.String())
		case "topics":
			if in.IsNull() {
				in.Skip()
				out.Topics = nil
			} else {
				in.Delim('[')
				if out.Topics == nil {
					if !in.IsDelim(']') {
						out.Topics = make([][]uint8, 0, 2)
					} else {
						out.Topics = [][]uint8{}
					}
				} else {
					out.Topics = (out.Topics)[:0]
				}
				for !in.IsDelim(']') {
					var v22 []uint8
					if in.IsNull() {
						in.Skip()
						v22 = nil
					} else {
						v22 = in.Bytes()
					}
					out.Topics = append(out.Topics, v22)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE910b2f5EncodeFyxVmLib3(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"topics\":"
		out.RawString(prefix)
		if in.Topics == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v24, v25 := range in.Topics {
				if v24 > 0 {
					out.RawByte(',')
				}
				out.Base64Bytes(v25)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonE910b2f5DecodeFyxVmLib2(in *jlexer.Lexer, out *Instance) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "outpoint":
			if in.IsNull() {
				in.Skip()
				out.Outpoint = nil
			} else {
				if out.Outpoint == nil {
					out.Outpoint = new(Outpoint)
				}
				if in.IsNull() {
					in.Skip()
					*out.Outpoint = nil
				} else {
					*out.Outpoint = in.Bytes()
				}
			}
		case "origin":
			if in.IsNull() {
				in.Skip()
				out.Origin = nil
			} else {
				if out.Origin == nil {
					out.Origin = new(Outpoint)
				}
				if in.IsNull() {
					in.Skip()
					*out.Origin = nil
				} else {
					*out.Origin = in.Bytes()
				}
			}
		case "nonce":
			out.Nonce = uint64(in.Uint64())
		case "kind":
			if in.IsNull() {
				in.Skip()
				out.Kind = nil
			} else {
				if out.Kind == nil {
					out.Kind = new(Outpoint)
				}
				if in.IsNull() {
					in.Skip()
					*out.Kind = nil
				} else {
					*out.Kind = in.Bytes()
				}
			}
		case "sats":
			out.Satoshis = uint64(in.Uint64())
		case "lock":
			if in.IsNull() {
				in.Skip()
				out.Lock = nil
			} else {
				out.Lock = in.Bytes()
			}
		case "store":
			if in.IsNull() {
				in.Skip()
				out.Storage = nil
			} else {
				out.Storage = in.Bytes()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE910b2f5EncodeFyxVmLib2(out *jwriter.Writer, in Instance) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Outpoint != nil {
		const prefix string = ",\"outpoint\":"
		first = false
		out.RawString(prefix[1:])
		out.Base64Bytes(*in.Outpoint)
	}
	if in.Origin != nil {
		const prefix string = ",\"origin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(*in.Origin)
	}
	if in.Nonce != 0 {
		const prefix string = ",\"nonce\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Nonce))
	}
	if in.Kind != nil {
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(*in.Kind)
	}
	{
		const prefix string = ",\"sats\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Satoshis))
	}
	{
		const prefix string = ",\"lock\":"
		out.RawString(prefix)
		out.Base64Bytes(in.Lock)
	}
	if len(in.Storage) != 0 {
		const prefix string = ",\"store\":"
		out.RawString(prefix)
		out.Base64Bytes(in.Storage)
	}
	out.RawByte('}')
}
func easyjsonE910b2f5DecodeFyxVmLib1(in *jlexer.Lexer, out *Parent) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "idx":
			out.Idx = int(in.Int())
		case "outpoint":
			if in.IsNull() {
				in.Skip()
				out.Outpoint = nil
			} else {
				if out.Outpoint == nil {
					out.Outpoint = new(Outpoint)
				}
				if in.IsNull() {
					in.Skip()
					*out.Outpoint = nil
				} else {
					*out.Outpoint = in.Bytes()
				}
			}
		case "kind":
			if in.IsNull() {
				in.Skip()
				out.Kind = nil
			} else {
				if out.Kind == nil {
					out.Kind = new(Outpoint)
				}
				if in.IsNull() {
					in.Skip()
					*out.Kind = nil
				} else {
					*out.Kind = in.Bytes()
				}
			}
		case "lock":
			if in.IsNull() {
				in.Skip()
				out.Lock = nil
			} else {
				out.Lock = in.Bytes()
			}
		case "origin":
			if in.IsNull() {
				in.Skip()
				out.Origin = nil
			} else {
				if out.Origin == nil {
					out.Origin = new(Outpoint)
				}
				if in.IsNull() {
					in.Skip()
					*out.Origin = nil
				} else {
					*out.Origin = in.Bytes()
				}
			}
		case "nonce":
			out.Nonce = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE910b2f5EncodeFyxVmLib1(out *jwriter.Writer, in Parent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"idx\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Idx))
	}
	{
		const prefix string = ",\"outpoint\":"
		out.RawString(prefix)
		if in.Outpoint == nil {
			out.RawString("null")
		} else {
			out.Base64Bytes(*in.Outpoint)
		}
	}
	{
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		if in.Kind == nil {
			out.RawString("null")
		} else {
			out.Base64Bytes(*in.Kind)
		}
	}
	{
		const prefix string = ",\"lock\":"
		out.RawString(prefix)
		out.Base64Bytes(in.Lock)
	}
	{
		const prefix string = ",\"origin\":"
		out.RawString(prefix)
		if in.Origin == nil {
			out.RawString("null")
		} else {
			out.Base64Bytes(*in.Origin)
		}
	}
	{
		const prefix string = ",\"nonce\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Nonce))
	}
	out.RawByte('}')
}
