package jute

import (
	"bufio"
	"encoding/binary"
	"io"
	"math"
)

// BinaryDecoder reads binary encoded data from an underlying io.Reader.  Binary
// data is read exactly the same way it is written.  See BinaryEncoder for the
// format.
type BinaryDecoder struct {
	r   *bufio.Reader
	buf [8]byte
}

// NewBinaryDecoder return a new BinaryDecoder wrapping an underlying io.Reader.
func NewBinaryDecoder(r io.Reader) *BinaryDecoder {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}
	return &BinaryDecoder{r: br}
}

// ReadStart performs no operation in BinaryDecoder.
func (d *BinaryDecoder) ReadStart() error { return nil }

// ReadEnd performs no operation in BinaryDecoder.
func (d *BinaryDecoder) ReadEnd() error { return nil }

// ReadBoolean will read a byte and return `false` if the value is zero and
// `true` if the value is non-zero.
func (d *BinaryDecoder) ReadBoolean() (bool, error) {
	v, err := d.ReadByte()
	if err != nil {
		return false, err
	}
	if v > 0 {
		return true, nil
	}
	return false, nil
}

// ReadByte will read a single byte.
func (d *BinaryDecoder) ReadByte() (byte, error) {
	return d.r.ReadByte()
}

// ReadInt will read 4 bytes in big endian byte order.
func (d *BinaryDecoder) ReadInt() (int32, error) {
	buf := d.buf[0:4]
	if _, err := io.ReadFull(d.r, buf); err != nil {
		return 0, err
	}
	return int32(binary.BigEndian.Uint32(buf)), nil
}

// ReadLong will read 8 bytes in a big endian byte order.
func (d *BinaryDecoder) ReadLong() (int64, error) {
	buf := d.buf[0:8]
	if _, err := io.ReadFull(d.r, buf); err != nil {
		return 0, err
	}
	return int64(binary.BigEndian.Uint64(buf)), nil
}

// ReadFloat will read a 32-bit float in IEEE 754 binary representation.
func (d *BinaryDecoder) ReadFloat() (float32, error) {
	v, err := d.ReadInt()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(uint32(v)), nil

}

// ReadDouble will read a 64-bit float in IEEE 754 binary representation.
func (d *BinaryDecoder) ReadDouble() (float64, error) {
	v, err := d.ReadLong()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(uint64(v)), nil

}

// ReadString will read a utf-8 encoded string first by reading the length
// encoded as an int (4 bytes) and then reading that number of bytes.
func (d *BinaryDecoder) ReadString() (*string, error) {
	// TODO: optimize for small reads
	p, err := d.ReadBuffer()
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, err
	}
	return String(string(p)), nil
}

// ReadBuffer will read a byte slice first by reading the length encoded as an
// int (4 bytes) and then reading that number of bytes.
func (d *BinaryDecoder) ReadBuffer() ([]byte, error) {
	size, err := d.ReadInt()
	if err != nil {
		return nil, err
	}

	if size < 0 {
		return nil, nil
	}

	buf := make([]byte, int(size))
	_, err = io.ReadFull(d.r, buf)
	return buf, err
}

// ReadVectorStart will read for the length of the vector as an int (4 bytes)
// and return the vector size.  The caller should then decode each item for
// that vector.
func (d *BinaryDecoder) ReadVectorStart() (int, error) {
	i, err := d.ReadInt()
	return int(i), err
}

// ReadVectorEnd performs no operation for BinaryDecoder.
func (d *BinaryDecoder) ReadVectorEnd() error { return nil }

// ReadMapStart will read for the length of the map as an int (4 bytes)
// and return the map size.  The caller should then decode the key and value
// for each item in the map.
func (d *BinaryDecoder) ReadMapStart() (int, error) {
	i, err := d.ReadInt()
	return int(i), err
}

// ReadMapEnd performs no operation for BinaryDecoder.
func (d *BinaryDecoder) ReadMapEnd() error { return nil }

// ReadRecord will read a jute record/class.
func (d *BinaryDecoder) ReadRecord(r RecordReader) error {
	return r.Read(d)
}
