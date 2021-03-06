// Autogenerated jute compiler
// @generated from 'testdata/zookeeper.jute'

package proto // github.com/go-zookeeper/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type ErrorResponse struct {
	Err int32 // err
}

func (r *ErrorResponse) GetErr() int32 {
	if r != nil {
		return r.Err
	}
	return 0
}

func (r *ErrorResponse) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Err, err = dec.ReadInt()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *ErrorResponse) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteInt(r.Err); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *ErrorResponse) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ErrorResponse(%+v)", *r)
}
