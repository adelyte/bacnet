package encoding

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"controlenvy.com/bacnet"
)

const (
	flag16bits byte = 0xFE
	flag32bits byte = 0xFF
)

//Encoder is the struct used to turn bacnet types to byte arrays. All
//public methods of encoder can set the internal error value. If such
//error is set, all encoding methods will be no-ops. This allows to
//defer error checking after several encoding operations
type Encoder struct {
	buf *bytes.Buffer
	err error
}

func NewEncoder() Encoder {
	e := Encoder{
		buf: new(bytes.Buffer),
		err: nil,
	}
	return e
}

func (e *Encoder) Error() error {
	return e.err
}

func (e *Encoder) Bytes() []byte {
	return e.buf.Bytes()
}

//ContextUnsigned write a (context) tag / value pair where the value
//type is an unsigned int
func (e *Encoder) ContextUnsigned(tagNumber byte, value uint32) {
	if e.err != nil {
		return
	}
	length := valueLength(value)
	t := tag{
		ID:      tagNumber,
		Context: true,
		Value:   uint32(length),
		Opening: false,
		Closing: false,
	}
	encodeTag(e.buf, t)
	unsigned(e.buf, value)
}

//ContextObjectID write a (context)tag / value pair where the value
//type is an unsigned int
func (e *Encoder) ContextObjectID(tabNumber byte, objectID bacnet.ObjectID) {
	if e.err != nil {
		return
	}
	t := tag{
		ID:      tabNumber,
		Context: true,
		Value:   4, //length of objectID is 4
		Opening: false,
		Closing: false,
	}
	encodeTag(e.buf, t)
	v, err := objectID.Encode()
	if err != nil {
		e.err = err
		return
	}
	_ = binary.Write(e.buf, binary.BigEndian, v)
}

//AppData writes a tag and value of any standard bacnet application
//data type. Returns an error if v of an invalid type
func (e *Encoder) AppData(v interface{}) {
	if e.err != nil {
		return
	}
	if v == nil {
		t := tag{ID: applicationTagNull}
		encodeTag(e.buf, t)
		return
	}
	switch val := v.(type) {
	case bool:
		var t tag
		if val {
			t = tag{ID: applicationTagBoolean, Value: 0b001}
		} else {
			t = tag{ID: applicationTagBoolean, Value: 0b000}
		}
		fmt.Printf("AppData() encoding bool %v\n", val)
		encodeTag(e.buf, t)
		_ = binary.Write(e.buf, binary.BigEndian, val)
	case float64:
		e.err = fmt.Errorf("not implemented ")
	case float32:
		t := tag{ID: applicationTagReal, Value: 4}
		encodeTag(e.buf, t)
		_ = binary.Write(e.buf, binary.BigEndian, val)
	case string:
		//+1 because there will be one byte for the string encoding format
		t := tag{ID: applicationTagCharacterString, Value: uint32(len(val) + 1)}
		encodeTag(e.buf, t)
		_ = e.buf.WriteByte(utf8Encoding)
		_, _ = e.buf.Write([]byte(val))
	case uint32:
		length := valueLength(val)
		t := tag{ID: applicationTagUnsignedInt, Value: uint32(length)}
		encodeTag(e.buf, t)
		unsigned(e.buf, val)
	case bacnet.SegmentationSupport:
		v := uint32(val)
		length := valueLength(v)
		t := tag{ID: applicationTagEnumerated, Value: uint32(length)}
		encodeTag(e.buf, t)
		unsigned(e.buf, v)
	case bacnet.ObjectID:
		t := tag{ID: applicationTagObjectID, Value: 4}
		encodeTag(e.buf, t)
		v, err := val.Encode()
		if err != nil {
			e.err = err
			return
		}
		_ = binary.Write(e.buf, binary.BigEndian, v)
	default:
		e.err = fmt.Errorf("encodeAppdata: unknown type %T", v)
	}
}

func (e *Encoder) ContextAbstractType(tagNumber byte, v bacnet.PropertyValue) {
	encodeTag(e.buf, tag{ID: tagNumber, Context: true, Opening: true})
	length := valueLength(v.Value)

	// t := tag{ID: v.Type, Value: uint32(length)}
	if v.Type == 0x01 {
		// t := tagNumber<<4 | 0b0000_1001
		t := byte(0b1001_0001)
		e.buf.WriteByte(t)
	} else {
		t := tag{ID: v.Type, Value: uint32(length)}
		encodeTag(e.buf, t)
	}

	unsigned(e.buf, v.Value)
	encodeTag(e.buf, tag{ID: tagNumber, Context: true, Closing: true})
}

// valueLength calculates how large the necessary value needs to be to fit in the appropriate
// packet length
func valueLength(value uint32) int {
	/* length of enumerated is variable, as per 20.2.11 */
	if value < 0x100 {
		return 1
	} else if value < 0x10000 {
		return 2
	} else if value < 0x1000000 {
		return 3
	}
	return 4
}

//unsigned writes the value in the buffer using a variable-sized encoding
func unsigned(buf *bytes.Buffer, value uint32) int {
	switch {
	case value < 0x100:
		buf.WriteByte(uint8(value))
		return 1
	case value < 0x10000:
		_ = binary.Write(buf, binary.BigEndian, uint16(value))
		return 2
	case value < 0x1000000:
		// There is no default 24 bit integer in go, so we have to
		// write it manually (in big endian)
		buf.WriteByte(byte(value >> 16))
		_ = binary.Write(buf, binary.BigEndian, uint16(value))
		return 3
	default:
		_ = binary.Write(buf, binary.BigEndian, value)
		return 4
	}
}
