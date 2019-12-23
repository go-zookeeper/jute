package jute

import (
	"bytes"
	"reflect"
	"testing"
)

var _ Decoder = &BinaryDecoder{}

func TestBinaryDecoderBase(t *testing.T) {
	input := []byte{
		0x01,                   // boolean: true
		0x00,                   // boolean: false
		0x66,                   // byte: 'f'
		0x62,                   // byte: 'b'
		0x00, 0x00, 0x4b, 0xce, // int:19406
		0x7f, 0xff, 0xff, 0xff, // int: 2147483647
		0xff, 0xff, 0xfe, 0x5c, // int: -420
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4b, 0xce, // long: 19406
		0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // long: 9,223,372,036,854,775,807
		0x40, 0x49, 0x0f, 0xdb, // float: 3.141592564
		0x40, 0x09, 0x21, 0xfb, 0x53, 0xc8, 0xd4, 0xf1, // double: 3.141592564
		// string: hello
		0x00, 0x00, 0x00, 0x05, // string len
		0x68, 0x65, 0x6c, 0x6c, 0x6f, // 'h', 'e', 'l', 'l', 'o'
		// buffer: 0x01, 0x02, 0x03, 0x04
		0x00, 0x00, 0x00, 0x04, // buffer len
		0x01, 0x02, 0x03, 0x04, // buffer contents
	}

	r := bytes.NewReader(input)
	dec := NewBinaryDecoder(r)
	if err := dec.ReadStart(); err != nil {
		t.Errorf("unexpected err: %v", err)
	}

	t1, err := dec.ReadBoolean()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if t1 != true {
		t.Errorf("ReadBoolean: expected true, got %t", t1)
	}

	t2, err := dec.ReadBoolean()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if t2 != false {
		t.Errorf("ReadBoolean: expected false, got %t", t2)
	}

	b1, err := dec.ReadByte()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if b1 != 'f' {
		t.Errorf("ReadByte: expected 'f', got '%c'", b1)
	}

	b2, err := dec.ReadByte()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if b2 != 'b' {
		t.Errorf("ReadByte: expected 'b', got '%c'", b2)
	}

	i1, err := dec.ReadInt()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if i1 != 19406 {
		t.Errorf("ReadInt: expected 19406, got '%d'", i1)
	}

	i2, err := dec.ReadInt()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if i2 != 2147483647 {
		t.Errorf("ReadInt: expected 2147483647, got '%d'", i2)
	}

	i3, err := dec.ReadInt()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if i3 != -420 {
		t.Errorf("ReadInt: expected -420, got '%d'", i3)
	}

	l1, err := dec.ReadLong()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if l1 != 19406 {
		t.Errorf("ReadLong: expected 19406, got '%d'", l1)
	}

	l2, err := dec.ReadLong()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if l2 != 9223372036854775807 {
		t.Errorf("ReadLong: expected 9223372036854775807, got '%d'", l1)
	}

	f1, err := dec.ReadFloat()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if f1 != 3.1415927 {
		t.Errorf("ReadFloat: expected 3.1415927, got %g", f1)
	}

	d1, err := dec.ReadDouble()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if d1 != 3.14159265 {
		t.Errorf("ReadDouble: expected 3.14159265, got %g", d1)
	}

	s1, err := dec.ReadUstring()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	if *s1 != "hello" {
		t.Errorf("ReadString: expected 'hello' got %q", *s1)
	}

	buf1, err := dec.ReadBuffer()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
	want := []byte{0x01, 0x02, 0x03, 0x04}
	if !bytes.Equal(buf1, want) {
		t.Errorf("ReadString: expected '%v' got '%v'", want, buf1)
	}

	if err := dec.ReadEnd(); err != nil {
		t.Errorf("unexpected err: %v", err)
	}

	unread := dec.r.Buffered() + r.Len()
	if unread != 0 {
		t.Errorf("did not fully read the input (%d bytes left)", unread)
	}
}

func TestBinaryDecoderVector(t *testing.T) {
	input := []byte{
		0x00, 0x00, 0x00, 0x05, // length of vector
		0x00, 0x00, 0x00, 0x05, // 5
		0x00, 0x00, 0x00, 0x04, // 4
		0x00, 0x00, 0x00, 0x03, // 3
		0x00, 0x00, 0x00, 0x02, // 2
		0x00, 0x00, 0x00, 0x01, // 1
	}
	r := bytes.NewReader(input)
	dec := NewBinaryDecoder(r)
	if err := dec.ReadStart(); err != nil {
		t.Errorf("unexpected err: %v", err)
	}

	size, err := dec.ReadVectorStart()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}

	slice := make([]int32, size)
	for i := 0; i < size; i++ {
		v, err := dec.ReadInt()
		if err != nil {
			t.Errorf("unexpected err: %v", err)
		}
		slice[i] = v
	}

	if err := dec.ReadVectorEnd(); err != nil {
		t.Errorf("unexpected err: %v", err)
	}

	if err := dec.ReadEnd(); err != nil {
		t.Errorf("unexpected err: %v", err)
	}

	unread := dec.r.Buffered() + r.Len()
	if unread != 0 {
		t.Errorf("did not fully read the input (%d bytes left)", unread)
	}

	want := []int32{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(want, slice) {
		t.Errorf("vector/slice does not match (want %v, got %v)", want, slice)
	}
}

func TestBinaryDecoderMap(t *testing.T) {
	input := []byte{
		0x00, 0x00, 0x00, 0x02, // length of map
		0x00, 0x00, 0x00, 0x03, // string length of "one"
		0x6f, 0x6e, 0x65, // 'o', 'n', 'e'
		0x00, 0x00, 0x00, 0x01, // int: 1
		0x00, 0x00, 0x00, 0x03, // lenght of "two"
		0x74, 0x77, 0x6f, // 't', 'w', 'o'
		0x00, 0x00, 0x00, 0x02, // int: 2
	}
	r := bytes.NewReader(input)
	dec := NewBinaryDecoder(r)
	if err := dec.ReadStart(); err != nil {
		t.Errorf("unexpected err: %v", err)
	}

	size, err := dec.ReadMapStart()
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}

	m1 := make(map[string]int32, size)
	for i := 0; i < size; i++ {
		k, err := dec.ReadUstring()
		if err != nil {
			t.Errorf("unexpected err: %v", err)
		}

		v, err := dec.ReadInt()
		if err != nil {
			t.Errorf("unexpected err: %v", err)
		}
		m1[*k] = v
	}

	if err := dec.ReadMapEnd(); err != nil {
		t.Errorf("unexpected err: %v", err)
	}

	if err := dec.ReadEnd(); err != nil {
		t.Errorf("unexpected err: %v", err)
	}

	unread := dec.r.Buffered() + r.Len()
	if unread != 0 {
		t.Errorf("did not fully read the input (%d bytes left)", unread)
	}

	want := map[string]int32{
		"one": 1,
		"two": 2,
	}
	if !reflect.DeepEqual(want, m1) {
		t.Errorf("vector/slice does not match (want %v, got %v)", want, m1)
	}

}
