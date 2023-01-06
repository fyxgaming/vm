// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package lib

import (
	json "encoding/json"
	bscript "github.com/libsv/go-bt/v2/bscript"
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

func easyjsonEc607727DecodeGithubComFyxgamingVmLib(in *jlexer.Lexer, out *Txo) {
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
		case "sats":
			out.Satoshis = uint64(in.Uint64())
		case "lock":
			if in.IsNull() {
				in.Skip()
				out.Lock = nil
			} else {
				out.Lock = in.Bytes()
			}
		case "script":
			if in.IsNull() {
				in.Skip()
				out.Script = nil
			} else {
				if out.Script == nil {
					out.Script = new(bscript.Script)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Script).UnmarshalJSON(data))
				}
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
func easyjsonEc607727EncodeGithubComFyxgamingVmLib(out *jwriter.Writer, in Txo) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Outpoint != nil {
		const prefix string = ",\"outpoint\":"
		first = false
		out.RawString(prefix[1:])
		out.Base64Bytes(*in.Outpoint)
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
	{
		const prefix string = ",\"script\":"
		out.RawString(prefix)
		if in.Script == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Script).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Txo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEc607727EncodeGithubComFyxgamingVmLib(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Txo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEc607727EncodeGithubComFyxgamingVmLib(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Txo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEc607727DecodeGithubComFyxgamingVmLib(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Txo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEc607727DecodeGithubComFyxgamingVmLib(l, v)
}
func easyjsonEc607727DecodeGithubComFyxgamingVmLib1(in *jlexer.Lexer, out *Instance) {
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
		case "txo":
			if in.IsNull() {
				in.Skip()
				out.Txo = nil
			} else {
				if out.Txo == nil {
					out.Txo = new(Txo)
				}
				(*out.Txo).UnmarshalEasyJSON(in)
			}
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
		case "code":
			if in.IsNull() {
				in.Skip()
				out.Code = nil
			} else {
				if out.Code == nil {
					out.Code = new(Outpoint)
				}
				if in.IsNull() {
					in.Skip()
					*out.Code = nil
				} else {
					*out.Code = in.Bytes()
				}
			}
		case "creator":
			if in.IsNull() {
				in.Skip()
				out.Creator = nil
			} else {
				if out.Creator == nil {
					out.Creator = new(Instance)
				}
				(*out.Creator).UnmarshalEasyJSON(in)
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
func easyjsonEc607727EncodeGithubComFyxgamingVmLib1(out *jwriter.Writer, in Instance) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Txo != nil {
		const prefix string = ",\"txo\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Txo).MarshalEasyJSON(out)
	}
	if in.Outpoint != nil {
		const prefix string = ",\"outpoint\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
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
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		if in.Code == nil {
			out.RawString("null")
		} else {
			out.Base64Bytes(*in.Code)
		}
	}
	{
		const prefix string = ",\"creator\":"
		out.RawString(prefix)
		if in.Creator == nil {
			out.RawString("null")
		} else {
			(*in.Creator).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Instance) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEc607727EncodeGithubComFyxgamingVmLib1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Instance) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEc607727EncodeGithubComFyxgamingVmLib1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Instance) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEc607727DecodeGithubComFyxgamingVmLib1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Instance) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEc607727DecodeGithubComFyxgamingVmLib1(l, v)
}
func easyjsonEc607727DecodeGithubComFyxgamingVmLib2(in *jlexer.Lexer, out *File) {
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
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				out.Data = in.Bytes()
			}
		case "type":
			out.Type = string(in.String())
		case "enc":
			out.Encoding = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "size":
			out.Size = uint32(in.Uint32())
		case "hash":
			if in.IsNull() {
				in.Skip()
				out.Hash = nil
			} else {
				out.Hash = in.Bytes()
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
func easyjsonEc607727EncodeGithubComFyxgamingVmLib2(out *jwriter.Writer, in File) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Outpoint != nil {
		const prefix string = ",\"outpoint\":"
		first = false
		out.RawString(prefix[1:])
		out.Base64Bytes(*in.Outpoint)
	}
	{
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.Data)
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	if in.Encoding != "" {
		const prefix string = ",\"enc\":"
		out.RawString(prefix)
		out.String(string(in.Encoding))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	if in.Size != 0 {
		const prefix string = ",\"size\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Size))
	}
	if len(in.Hash) != 0 {
		const prefix string = ",\"hash\":"
		out.RawString(prefix)
		out.Base64Bytes(in.Hash)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v File) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEc607727EncodeGithubComFyxgamingVmLib2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v File) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEc607727EncodeGithubComFyxgamingVmLib2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *File) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEc607727DecodeGithubComFyxgamingVmLib2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *File) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEc607727DecodeGithubComFyxgamingVmLib2(l, v)
}
func easyjsonEc607727DecodeGithubComFyxgamingVmLib3(in *jlexer.Lexer, out *Event) {
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
					var v34 []uint8
					if in.IsNull() {
						in.Skip()
						v34 = nil
					} else {
						v34 = in.Bytes()
					}
					out.Topics = append(out.Topics, v34)
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
func easyjsonEc607727EncodeGithubComFyxgamingVmLib3(out *jwriter.Writer, in Event) {
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
			for v36, v37 := range in.Topics {
				if v36 > 0 {
					out.RawByte(',')
				}
				out.Base64Bytes(v37)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEc607727EncodeGithubComFyxgamingVmLib3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEc607727EncodeGithubComFyxgamingVmLib3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEc607727DecodeGithubComFyxgamingVmLib3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEc607727DecodeGithubComFyxgamingVmLib3(l, v)
}
func easyjsonEc607727DecodeGithubComFyxgamingVmLib4(in *jlexer.Lexer, out *Error) {
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
		case "Code":
			out.Code = int(in.Int())
		case "Err":
			out.Err = string(in.String())
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
func easyjsonEc607727EncodeGithubComFyxgamingVmLib4(out *jwriter.Writer, in Error) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Code\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Code))
	}
	{
		const prefix string = ",\"Err\":"
		out.RawString(prefix)
		out.String(string(in.Err))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Error) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEc607727EncodeGithubComFyxgamingVmLib4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Error) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEc607727EncodeGithubComFyxgamingVmLib4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Error) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEc607727DecodeGithubComFyxgamingVmLib4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Error) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEc607727DecodeGithubComFyxgamingVmLib4(l, v)
}
func easyjsonEc607727DecodeGithubComFyxgamingVmLib5(in *jlexer.Lexer, out *Child) {
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
func easyjsonEc607727EncodeGithubComFyxgamingVmLib5(out *jwriter.Writer, in Child) {
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

// MarshalJSON supports json.Marshaler interface
func (v Child) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEc607727EncodeGithubComFyxgamingVmLib5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Child) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEc607727EncodeGithubComFyxgamingVmLib5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Child) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEc607727DecodeGithubComFyxgamingVmLib5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Child) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEc607727DecodeGithubComFyxgamingVmLib5(l, v)
}
