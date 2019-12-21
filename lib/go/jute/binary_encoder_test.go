package jute

import (
	"bytes"
	"encoding/hex"
	"testing"
)

var _ Encoder = &BinaryEncoder{}

func TestBinaryEncoderBase(t *testing.T) {
	w := &bytes.Buffer{}
	enc := NewBinaryEncoder(w)
	if err := enc.WriteStart(); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteBoolean(true); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteBoolean(false); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteByte('f'); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteByte('b'); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteInt(19406); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteInt(2147483647); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteInt(-420); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteLong(19406); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteLong(9223372036854775807); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteFloat(3.14159265); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteDouble(3.14159265); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteUstring("hello"); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteBuffer([]byte{0x01, 0x02, 0x03, 0x04}); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteEnd(); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	want := []byte{
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

	if !bytes.Equal(want, w.Bytes()) {
		t.Errorf("unexpected results\n\twant: %s\n\tgot : %s", hex.Dump(want), hex.Dump(w.Bytes()))
	}
}

func TestBinaryEncoderVector(t *testing.T) {
	slice := []int32{5, 4, 3, 2, 1}

	w := &bytes.Buffer{}
	enc := NewBinaryEncoder(w)
	if err := enc.WriteStart(); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteVectorStart(len(slice)); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	for _, i := range slice {
		if err := enc.WriteInt(i); err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
	}

	if err := enc.WriteVectorEnd(); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteEnd(); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	want := []byte{
		0x00, 0x00, 0x00, 0x05, // length of vector
		0x00, 0x00, 0x00, 0x05, // 5
		0x00, 0x00, 0x00, 0x04, // 4
		0x00, 0x00, 0x00, 0x03, // 3
		0x00, 0x00, 0x00, 0x02, // 2
		0x00, 0x00, 0x00, 0x01, // 1
	}

	if !bytes.Equal(want, w.Bytes()) {
		t.Errorf("unexpected results\n\twant: %s\n\tgot : %s", hex.Dump(want), hex.Dump(w.Bytes()))
	}
}

func TestBinaryEncoderMap(t *testing.T) {
	m1 := map[string]int32{
		"one": 1,
		"two": 2,
	}

	w := &bytes.Buffer{}
	enc := NewBinaryEncoder(w)
	if err := enc.WriteStart(); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteMapStart(len(m1)); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	for k, v := range m1 {
		if err := enc.WriteUstring(k); err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if err := enc.WriteInt(v); err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
	}

	if err := enc.WriteMapEnd(); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if err := enc.WriteEnd(); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	want := []byte{
		0x00, 0x00, 0x00, 0x02, // length of map
		0x00, 0x00, 0x00, 0x03, // string length of "one"
		0x6f, 0x6e, 0x65, // 'o', 'n', 'e'
		0x00, 0x00, 0x00, 0x01, // int: 1
		0x00, 0x00, 0x00, 0x03, // lenght of "two"
		0x74, 0x77, 0x6f, // 't', 'w', 'o'
		0x00, 0x00, 0x00, 0x02, // int: 2
	}

	if !bytes.Equal(want, w.Bytes()) {
		t.Errorf("unexpected results\n\twant: %s\n\tgot : %s", hex.Dump(want), hex.Dump(w.Bytes()))
	}
}
