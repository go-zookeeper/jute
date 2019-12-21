package jute

import (
	"bufio"
	"encoding/binary"
	"io"
	"math"
)

// BinaryEncoder writes binary encoded data to an underlying io.Writer.  Binary
// data is encoded in big endian with all bytes.
//   * `bool` will write out 1 bytes with `0` meaning `false` and `1` meaning
//      `true`.
//   * `byte`, `int`, and `long\ are written out in bigendian byte order
//      writing out 1, 4, and 8 bytes respectively.
//   * `float` and `double` are encoded as IEEE 754 as 4 or 8 bytes
//   * `ustring` and `buffer` write out the length encoded as an int (4 bytes)
//      followed by the string/buffer in bytes
//   * `vector` will write the number of items as an int (4 bytes) followed by
//      each item encoded in it's type.
//   * `map` will write the number of items as an int (4 bytes) followd by each
//      key and value encoded in it's own type.
//
// There is no header/tailer for the record itself.
type BinaryEncoder struct {
	w   *bufio.Writer
	err error
	buf [8]byte
}

// NewBinaryEncoder return a new BinaryEncoder wrapping an underlying io.Writer.
func NewBinaryEncoder(w io.Writer) *BinaryEncoder {
	bw, ok := w.(*bufio.Writer)
	if !ok {
		bw = bufio.NewWriter(w)
	}
	return &BinaryEncoder{w: bw}
}

func (s *BinaryEncoder) Encode(r RecordWriter) error {
	return r.Write(s)
}

// WriteStart marks the start of the encoded record/class.  In BinaryEncoder
// this performs no operation.
func (s *BinaryEncoder) WriteStart() error { return nil }

// WriteEnd marks the endof the encoded record/class and flush it to the writer.
func (s *BinaryEncoder) WriteEnd() error {
	return s.w.Flush()
}

// WriteBoolean will write a boolean value as a single byte: `1` for `true`,
// `0` for `false`.
func (s *BinaryEncoder) WriteBoolean(b bool) error {
	if b {
		return s.WriteByte(1)
	}
	return s.WriteByte(0)
}

// WriteByte will write a single byte.
func (s *BinaryEncoder) WriteByte(b byte) error {
	return s.w.WriteByte(b)
}

// WriteInt will write an int as 4 bytes in big endian byte order.
func (s *BinaryEncoder) WriteInt(i int32) error {
	v := s.buf[0:4]
	binary.BigEndian.PutUint32(v, uint32(i))
	_, err := s.w.Write(v)
	return err
}

// WriteLong will write a long as 8 bytes in big endian byte order.
func (s *BinaryEncoder) WriteLong(i int64) error {
	v := s.buf[0:8]
	binary.BigEndian.PutUint64(v, uint64(i))
	_, err := s.w.Write(v)
	return err
}

// WriteFloat will write a Float in IEEE 754 format as 4 bytes.
func (s *BinaryEncoder) WriteFloat(i float32) error {
	return s.WriteInt(int32(math.Float32bits(i)))
}

// WriteDouble will write a double value in IEEE 754 format as 8 bytes.
func (s *BinaryEncoder) WriteDouble(i float64) error {
	return s.WriteLong(int64(math.Float64bits(i)))
}

// WriteUstring will write a utf8 encoded string by first writing it's length as
// 4 bytes and then the byte of the string.
func (s *BinaryEncoder) WriteUstring(v string) error {
	if err := s.WriteInt(int32(len(v))); err != nil {
		return err
	}
	_, err := s.w.WriteString(v)
	return err
}

// WriteBuffer will write any byte slice by first writing it's length as 4
// bytes followed by the bytes in the slice.
func (s *BinaryEncoder) WriteBuffer(v []byte) error {
	if err := s.WriteInt(int32(len(v))); err != nil {
		return err
	}
	_, err := s.w.Write(v)
	return err
}

// WriteVectorStart will write out the number of items in the vector as 4
// bytes.  After calling WriteVectorStart the caller should write out each item.
func (s *BinaryEncoder) WriteVectorStart(l int) error {
	return s.WriteInt(int32(l))
} // WriteString will write a utf8 encoded string by first writing it's length as

// WriteVectorEnd will mark the end of a vector.  In BinaryEncoder this
// performs no operation.
func (s *BinaryEncoder) WriteVectorEnd() error { return nil }

// WriteMapStart will write out the number of items in the map as 4 bytes.
// After calling Write<apStart the caller should write out each key and value
// of each item.
func (s *BinaryEncoder) WriteMapStart(l int) error {
	return s.WriteInt(int32(l))
}

// WriteMapEnd will mark the end of a vector.  In BinaryEncoder this performs
// no operation.
func (s *BinaryEncoder) WriteMapEnd() error { return nil }

// WriteRecord will write a jute record/class.
func (s *BinaryEncoder) WriteRecord(r RecordWriter) error {
	return r.Write(s)
}
