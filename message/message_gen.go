package message

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Message) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Message, zb0002)
	}
	for zb0001 := range *z {
		(*z)[zb0001], err = dc.ReadIntf()
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Message) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0003 := range z {
		err = en.WriteIntf(z[zb0003])
		if err != nil {
			err = msgp.WrapError(err, zb0003)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Message) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		o, err = msgp.AppendIntf(o, z[zb0003])
		if err != nil {
			err = msgp.WrapError(err, zb0003)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Message) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Message, zb0002)
	}
	for zb0001 := range *z {
		(*z)[zb0001], bts, err = msgp.ReadIntfBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Message) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		s += msgp.GuessSize(z[zb0003])
	}
	return
}
