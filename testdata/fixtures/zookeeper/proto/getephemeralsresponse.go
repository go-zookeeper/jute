// Autogenerated jute compiler
// @generated from 'testdata/zookeeper.jute'

package proto // github.com/go-zookeeper/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type GetEphemeralsResponse struct {
	Ephemerals []string // ephemerals
}

func (r *GetEphemeralsResponse) GetEphemerals() []string {
	if r != nil && r.Ephemerals != nil {
		return r.Ephemerals
	}
	return nil
}

func (r *GetEphemeralsResponse) Read(dec jute.Decoder) (err error) {
	var size int
	if err = dec.ReadStart(); err != nil {
		return err
	}
	size, err = dec.ReadVectorStart()
	if err != nil {
		return err
	}
	if size < 0 {
		r.Ephemerals = nil
	} else {
		r.Ephemerals = make([]string, size)
		for i := 0; i < size; i++ {
			r.Ephemerals[i], err = dec.ReadString()
			if err != nil {
				return err
			}
		}
	}
	if err = dec.ReadVectorEnd(); err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *GetEphemeralsResponse) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.Ephemerals), r.Ephemerals == nil); err != nil {
		return err
	}
	for _, v := range r.Ephemerals {
		if err := enc.WriteString(v); err != nil {
			return err
		}
	}
	if err := enc.WriteVectorEnd(); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *GetEphemeralsResponse) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetEphemeralsResponse(%+v)", *r)
}
