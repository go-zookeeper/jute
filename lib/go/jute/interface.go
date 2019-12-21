package jute

// RecordWriter defines how a jute record (class) will write to an encoder
// protocol.
type RecordWriter interface {
	Write(Encoder) error
}

// RecordReader defines how a jute record (class) will be read from an decoder
// protocol
type RecordReader interface {
	Read(Decoder) error
}

// Encoder defines how to encode a record to a destination like a network socket.
type Encoder interface {
	WriteStart() error
	WriteEnd() error

	WriteByte(byte) error
	WriteBoolean(bool) error
	WriteInt(int32) error
	WriteLong(int64) error
	WriteFloat(float32) error
	WriteDouble(float64) error
	WriteUstring(string) error
	WriteBuffer([]byte) error

	WriteVectorStart(len int) error
	WriteVectorEnd() error

	WriteMapStart(len int) error
	WriteMapEnd() error

	WriteRecord(RecordWriter) error
}

// Decoder defines how to dencode a record from a source like a network socket.
type Decoder interface {
	ReadStart() error
	ReadEnd() error

	ReadByte() (byte, error)
	ReadBoolean() (bool, error)
	ReadInt() (int32, error)
	ReadLong() (int64, error)
	ReadFloat() (float32, error)
	ReadDouble() (float64, error)
	ReadUstring() (string, error)
	ReadBuffer() ([]byte, error)

	ReadVectorStart() (int, error)
	ReadVectorEnd() error

	ReadMapStart() (int, error)
	ReadMapEnd() error

	ReadRecord(RecordReader) error
}
