package lib

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/libsv/go-bt/v2/bscript"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Child) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Contract":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Contract")
					return
				}
				z.Contract = nil
			} else {
				if z.Contract == nil {
					z.Contract = new(Outpoint)
				}
				err = z.Contract.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Contract")
					return
				}
			}
		case "Method":
			z.Method, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Method")
				return
			}
		case "CallData":
			z.CallData, err = dc.ReadBytes(z.CallData)
			if err != nil {
				err = msgp.WrapError(err, "CallData")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Child) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Contract"
	err = en.Append(0x83, 0xa8, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74)
	if err != nil {
		return
	}
	if z.Contract == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Contract.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Contract")
			return
		}
	}
	// write "Method"
	err = en.Append(0xa6, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.Method)
	if err != nil {
		err = msgp.WrapError(err, "Method")
		return
	}
	// write "CallData"
	err = en.Append(0xa8, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.CallData)
	if err != nil {
		err = msgp.WrapError(err, "CallData")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Child) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Contract"
	o = append(o, 0x83, 0xa8, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74)
	if z.Contract == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Contract.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Contract")
			return
		}
	}
	// string "Method"
	o = append(o, 0xa6, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64)
	o = msgp.AppendString(o, z.Method)
	// string "CallData"
	o = append(o, 0xa8, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61)
	o = msgp.AppendBytes(o, z.CallData)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Child) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Contract":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Contract = nil
			} else {
				if z.Contract == nil {
					z.Contract = new(Outpoint)
				}
				bts, err = z.Contract.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Contract")
					return
				}
			}
		case "Method":
			z.Method, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Method")
				return
			}
		case "CallData":
			z.CallData, bts, err = msgp.ReadBytesBytes(bts, z.CallData)
			if err != nil {
				err = msgp.WrapError(err, "CallData")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Child) Msgsize() (s int) {
	s = 1 + 9
	if z.Contract == nil {
		s += msgp.NilSize
	} else {
		s += z.Contract.Msgsize()
	}
	s += 7 + msgp.StringPrefixSize + len(z.Method) + 9 + msgp.BytesPrefixSize + len(z.CallData)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Error) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Code":
			z.Code, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Code")
				return
			}
		case "Err":
			z.Err, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Err")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Error) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Code"
	err = en.Append(0x82, 0xa4, 0x43, 0x6f, 0x64, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Code)
	if err != nil {
		err = msgp.WrapError(err, "Code")
		return
	}
	// write "Err"
	err = en.Append(0xa3, 0x45, 0x72, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.Err)
	if err != nil {
		err = msgp.WrapError(err, "Err")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Error) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Code"
	o = append(o, 0x82, 0xa4, 0x43, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt(o, z.Code)
	// string "Err"
	o = append(o, 0xa3, 0x45, 0x72, 0x72)
	o = msgp.AppendString(o, z.Err)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Error) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Code":
			z.Code, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Code")
				return
			}
		case "Err":
			z.Err, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Err")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Error) Msgsize() (s int) {
	s = 1 + 5 + msgp.IntSize + 4 + msgp.StringPrefixSize + len(z.Err)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Event) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "Topics":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Topics")
				return
			}
			if cap(z.Topics) >= int(zb0002) {
				z.Topics = (z.Topics)[:zb0002]
			} else {
				z.Topics = make([][]byte, zb0002)
			}
			for za0001 := range z.Topics {
				z.Topics[za0001], err = dc.ReadBytes(z.Topics[za0001])
				if err != nil {
					err = msgp.WrapError(err, "Topics", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Event) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Id"
	err = en.Append(0x82, 0xa2, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.Id)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	// write "Topics"
	err = en.Append(0xa6, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Topics)))
	if err != nil {
		err = msgp.WrapError(err, "Topics")
		return
	}
	for za0001 := range z.Topics {
		err = en.WriteBytes(z.Topics[za0001])
		if err != nil {
			err = msgp.WrapError(err, "Topics", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Event) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Id"
	o = append(o, 0x82, 0xa2, 0x49, 0x64)
	o = msgp.AppendString(o, z.Id)
	// string "Topics"
	o = append(o, 0xa6, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Topics)))
	for za0001 := range z.Topics {
		o = msgp.AppendBytes(o, z.Topics[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Event) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "Topics":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Topics")
				return
			}
			if cap(z.Topics) >= int(zb0002) {
				z.Topics = (z.Topics)[:zb0002]
			} else {
				z.Topics = make([][]byte, zb0002)
			}
			for za0001 := range z.Topics {
				z.Topics[za0001], bts, err = msgp.ReadBytesBytes(bts, z.Topics[za0001])
				if err != nil {
					err = msgp.WrapError(err, "Topics", za0001)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Event) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.Id) + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Topics {
		s += msgp.BytesPrefixSize + len(z.Topics[za0001])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *File) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Outpoint":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Outpoint")
					return
				}
				z.Outpoint = nil
			} else {
				if z.Outpoint == nil {
					z.Outpoint = new(Outpoint)
				}
				err = z.Outpoint.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Outpoint")
					return
				}
			}
		case "Data":
			z.Data, err = dc.ReadBytes(z.Data)
			if err != nil {
				err = msgp.WrapError(err, "Data")
				return
			}
		case "Type":
			z.Type, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Type")
				return
			}
		case "Encoding":
			z.Encoding, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Encoding")
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Size":
			z.Size, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "Size")
				return
			}
		case "Hash":
			z.Hash, err = dc.ReadBytes(z.Hash)
			if err != nil {
				err = msgp.WrapError(err, "Hash")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *File) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "Outpoint"
	err = en.Append(0x87, 0xa8, 0x4f, 0x75, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	if err != nil {
		return
	}
	if z.Outpoint == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Outpoint.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Outpoint")
			return
		}
	}
	// write "Data"
	err = en.Append(0xa4, 0x44, 0x61, 0x74, 0x61)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	// write "Type"
	err = en.Append(0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Type)
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	// write "Encoding"
	err = en.Append(0xa8, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67)
	if err != nil {
		return
	}
	err = en.WriteString(z.Encoding)
	if err != nil {
		err = msgp.WrapError(err, "Encoding")
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "Size"
	err = en.Append(0xa4, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.Size)
	if err != nil {
		err = msgp.WrapError(err, "Size")
		return
	}
	// write "Hash"
	err = en.Append(0xa4, 0x48, 0x61, 0x73, 0x68)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Hash)
	if err != nil {
		err = msgp.WrapError(err, "Hash")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *File) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Outpoint"
	o = append(o, 0x87, 0xa8, 0x4f, 0x75, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	if z.Outpoint == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Outpoint.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Outpoint")
			return
		}
	}
	// string "Data"
	o = append(o, 0xa4, 0x44, 0x61, 0x74, 0x61)
	o = msgp.AppendBytes(o, z.Data)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "Encoding"
	o = append(o, 0xa8, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67)
	o = msgp.AppendString(o, z.Encoding)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Size"
	o = append(o, 0xa4, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendUint32(o, z.Size)
	// string "Hash"
	o = append(o, 0xa4, 0x48, 0x61, 0x73, 0x68)
	o = msgp.AppendBytes(o, z.Hash)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *File) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Outpoint":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Outpoint = nil
			} else {
				if z.Outpoint == nil {
					z.Outpoint = new(Outpoint)
				}
				bts, err = z.Outpoint.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Outpoint")
					return
				}
			}
		case "Data":
			z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)
			if err != nil {
				err = msgp.WrapError(err, "Data")
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Type")
				return
			}
		case "Encoding":
			z.Encoding, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Encoding")
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Size":
			z.Size, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Size")
				return
			}
		case "Hash":
			z.Hash, bts, err = msgp.ReadBytesBytes(bts, z.Hash)
			if err != nil {
				err = msgp.WrapError(err, "Hash")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *File) Msgsize() (s int) {
	s = 1 + 9
	if z.Outpoint == nil {
		s += msgp.NilSize
	} else {
		s += z.Outpoint.Msgsize()
	}
	s += 5 + msgp.BytesPrefixSize + len(z.Data) + 5 + msgp.StringPrefixSize + len(z.Type) + 9 + msgp.StringPrefixSize + len(z.Encoding) + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.Uint32Size + 5 + msgp.BytesPrefixSize + len(z.Hash)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Instance) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Txo":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Txo")
					return
				}
				z.Txo = nil
			} else {
				if z.Txo == nil {
					z.Txo = new(Txo)
				}
				err = z.Txo.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Txo")
					return
				}
			}
		case "Outpoint":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Outpoint")
					return
				}
				z.Outpoint = nil
			} else {
				if z.Outpoint == nil {
					z.Outpoint = new(Outpoint)
				}
				err = z.Outpoint.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Outpoint")
					return
				}
			}
		case "Origin":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Origin")
					return
				}
				z.Origin = nil
			} else {
				if z.Origin == nil {
					z.Origin = new(Outpoint)
				}
				err = z.Origin.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Origin")
					return
				}
			}
		case "Nonce":
			z.Nonce, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Nonce")
				return
			}
		case "Kind":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Kind")
					return
				}
				z.Kind = nil
			} else {
				if z.Kind == nil {
					z.Kind = new(Outpoint)
				}
				err = z.Kind.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Kind")
					return
				}
			}
		case "Satoshis":
			z.Satoshis, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Satoshis")
				return
			}
		case "Lock":
			z.Lock, err = dc.ReadBytes(z.Lock)
			if err != nil {
				err = msgp.WrapError(err, "Lock")
				return
			}
		case "Storage":
			z.Storage, err = dc.ReadBytes(z.Storage)
			if err != nil {
				err = msgp.WrapError(err, "Storage")
				return
			}
		case "Code":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Code")
					return
				}
				z.Code = nil
			} else {
				if z.Code == nil {
					z.Code = new(Outpoint)
				}
				err = z.Code.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Code")
					return
				}
			}
		case "Creator":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Creator")
					return
				}
				z.Creator = nil
			} else {
				if z.Creator == nil {
					z.Creator = new(Instance)
				}
				err = z.Creator.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Creator")
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Instance) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "Txo"
	err = en.Append(0x8a, 0xa3, 0x54, 0x78, 0x6f)
	if err != nil {
		return
	}
	if z.Txo == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Txo.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Txo")
			return
		}
	}
	// write "Outpoint"
	err = en.Append(0xa8, 0x4f, 0x75, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	if err != nil {
		return
	}
	if z.Outpoint == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Outpoint.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Outpoint")
			return
		}
	}
	// write "Origin"
	err = en.Append(0xa6, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e)
	if err != nil {
		return
	}
	if z.Origin == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Origin.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Origin")
			return
		}
	}
	// write "Nonce"
	err = en.Append(0xa5, 0x4e, 0x6f, 0x6e, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Nonce)
	if err != nil {
		err = msgp.WrapError(err, "Nonce")
		return
	}
	// write "Kind"
	err = en.Append(0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return
	}
	if z.Kind == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Kind.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Kind")
			return
		}
	}
	// write "Satoshis"
	err = en.Append(0xa8, 0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Satoshis)
	if err != nil {
		err = msgp.WrapError(err, "Satoshis")
		return
	}
	// write "Lock"
	err = en.Append(0xa4, 0x4c, 0x6f, 0x63, 0x6b)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Lock)
	if err != nil {
		err = msgp.WrapError(err, "Lock")
		return
	}
	// write "Storage"
	err = en.Append(0xa7, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Storage)
	if err != nil {
		err = msgp.WrapError(err, "Storage")
		return
	}
	// write "Code"
	err = en.Append(0xa4, 0x43, 0x6f, 0x64, 0x65)
	if err != nil {
		return
	}
	if z.Code == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Code.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Code")
			return
		}
	}
	// write "Creator"
	err = en.Append(0xa7, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72)
	if err != nil {
		return
	}
	if z.Creator == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Creator.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Creator")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Instance) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "Txo"
	o = append(o, 0x8a, 0xa3, 0x54, 0x78, 0x6f)
	if z.Txo == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Txo.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Txo")
			return
		}
	}
	// string "Outpoint"
	o = append(o, 0xa8, 0x4f, 0x75, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	if z.Outpoint == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Outpoint.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Outpoint")
			return
		}
	}
	// string "Origin"
	o = append(o, 0xa6, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e)
	if z.Origin == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Origin.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Origin")
			return
		}
	}
	// string "Nonce"
	o = append(o, 0xa5, 0x4e, 0x6f, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint64(o, z.Nonce)
	// string "Kind"
	o = append(o, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if z.Kind == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Kind.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Kind")
			return
		}
	}
	// string "Satoshis"
	o = append(o, 0xa8, 0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x73)
	o = msgp.AppendUint64(o, z.Satoshis)
	// string "Lock"
	o = append(o, 0xa4, 0x4c, 0x6f, 0x63, 0x6b)
	o = msgp.AppendBytes(o, z.Lock)
	// string "Storage"
	o = append(o, 0xa7, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65)
	o = msgp.AppendBytes(o, z.Storage)
	// string "Code"
	o = append(o, 0xa4, 0x43, 0x6f, 0x64, 0x65)
	if z.Code == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Code.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Code")
			return
		}
	}
	// string "Creator"
	o = append(o, 0xa7, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72)
	if z.Creator == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Creator.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Creator")
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Instance) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Txo":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Txo = nil
			} else {
				if z.Txo == nil {
					z.Txo = new(Txo)
				}
				bts, err = z.Txo.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Txo")
					return
				}
			}
		case "Outpoint":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Outpoint = nil
			} else {
				if z.Outpoint == nil {
					z.Outpoint = new(Outpoint)
				}
				bts, err = z.Outpoint.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Outpoint")
					return
				}
			}
		case "Origin":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Origin = nil
			} else {
				if z.Origin == nil {
					z.Origin = new(Outpoint)
				}
				bts, err = z.Origin.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Origin")
					return
				}
			}
		case "Nonce":
			z.Nonce, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Nonce")
				return
			}
		case "Kind":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Kind = nil
			} else {
				if z.Kind == nil {
					z.Kind = new(Outpoint)
				}
				bts, err = z.Kind.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Kind")
					return
				}
			}
		case "Satoshis":
			z.Satoshis, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Satoshis")
				return
			}
		case "Lock":
			z.Lock, bts, err = msgp.ReadBytesBytes(bts, z.Lock)
			if err != nil {
				err = msgp.WrapError(err, "Lock")
				return
			}
		case "Storage":
			z.Storage, bts, err = msgp.ReadBytesBytes(bts, z.Storage)
			if err != nil {
				err = msgp.WrapError(err, "Storage")
				return
			}
		case "Code":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Code = nil
			} else {
				if z.Code == nil {
					z.Code = new(Outpoint)
				}
				bts, err = z.Code.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Code")
					return
				}
			}
		case "Creator":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Creator = nil
			} else {
				if z.Creator == nil {
					z.Creator = new(Instance)
				}
				bts, err = z.Creator.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Creator")
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Instance) Msgsize() (s int) {
	s = 1 + 4
	if z.Txo == nil {
		s += msgp.NilSize
	} else {
		s += z.Txo.Msgsize()
	}
	s += 9
	if z.Outpoint == nil {
		s += msgp.NilSize
	} else {
		s += z.Outpoint.Msgsize()
	}
	s += 7
	if z.Origin == nil {
		s += msgp.NilSize
	} else {
		s += z.Origin.Msgsize()
	}
	s += 6 + msgp.Uint64Size + 5
	if z.Kind == nil {
		s += msgp.NilSize
	} else {
		s += z.Kind.Msgsize()
	}
	s += 9 + msgp.Uint64Size + 5 + msgp.BytesPrefixSize + len(z.Lock) + 8 + msgp.BytesPrefixSize + len(z.Storage) + 5
	if z.Code == nil {
		s += msgp.NilSize
	} else {
		s += z.Code.Msgsize()
	}
	s += 8
	if z.Creator == nil {
		s += msgp.NilSize
	} else {
		s += z.Creator.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Txo) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Outpoint":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Outpoint")
					return
				}
				z.Outpoint = nil
			} else {
				if z.Outpoint == nil {
					z.Outpoint = new(Outpoint)
				}
				err = z.Outpoint.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Outpoint")
					return
				}
			}
		case "Satoshis":
			z.Satoshis, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Satoshis")
				return
			}
		case "Lock":
			z.Lock, err = dc.ReadBytes(z.Lock)
			if err != nil {
				err = msgp.WrapError(err, "Lock")
				return
			}
		case "Script":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Script")
					return
				}
				z.Script = nil
			} else {
				if z.Script == nil {
					z.Script = new(bscript.Script)
				}
				err = z.Script.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Script")
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Txo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Outpoint"
	err = en.Append(0x84, 0xa8, 0x4f, 0x75, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	if err != nil {
		return
	}
	if z.Outpoint == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Outpoint.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Outpoint")
			return
		}
	}
	// write "Satoshis"
	err = en.Append(0xa8, 0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Satoshis)
	if err != nil {
		err = msgp.WrapError(err, "Satoshis")
		return
	}
	// write "Lock"
	err = en.Append(0xa4, 0x4c, 0x6f, 0x63, 0x6b)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Lock)
	if err != nil {
		err = msgp.WrapError(err, "Lock")
		return
	}
	// write "Script"
	err = en.Append(0xa6, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74)
	if err != nil {
		return
	}
	if z.Script == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Script.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Script")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Txo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Outpoint"
	o = append(o, 0x84, 0xa8, 0x4f, 0x75, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74)
	if z.Outpoint == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Outpoint.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Outpoint")
			return
		}
	}
	// string "Satoshis"
	o = append(o, 0xa8, 0x53, 0x61, 0x74, 0x6f, 0x73, 0x68, 0x69, 0x73)
	o = msgp.AppendUint64(o, z.Satoshis)
	// string "Lock"
	o = append(o, 0xa4, 0x4c, 0x6f, 0x63, 0x6b)
	o = msgp.AppendBytes(o, z.Lock)
	// string "Script"
	o = append(o, 0xa6, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74)
	if z.Script == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Script.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Script")
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Txo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Outpoint":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Outpoint = nil
			} else {
				if z.Outpoint == nil {
					z.Outpoint = new(Outpoint)
				}
				bts, err = z.Outpoint.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Outpoint")
					return
				}
			}
		case "Satoshis":
			z.Satoshis, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Satoshis")
				return
			}
		case "Lock":
			z.Lock, bts, err = msgp.ReadBytesBytes(bts, z.Lock)
			if err != nil {
				err = msgp.WrapError(err, "Lock")
				return
			}
		case "Script":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Script = nil
			} else {
				if z.Script == nil {
					z.Script = new(bscript.Script)
				}
				bts, err = z.Script.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Script")
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Txo) Msgsize() (s int) {
	s = 1 + 9
	if z.Outpoint == nil {
		s += msgp.NilSize
	} else {
		s += z.Outpoint.Msgsize()
	}
	s += 9 + msgp.Uint64Size + 5 + msgp.BytesPrefixSize + len(z.Lock) + 7
	if z.Script == nil {
		s += msgp.NilSize
	} else {
		s += z.Script.Msgsize()
	}
	return
}
