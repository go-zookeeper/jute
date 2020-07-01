// Autogenerated jute compiler
// @generated from 'testdata/test.jute'

package test // com/github/gozookeeper/jute/test

import (
	"fmt"

	"com/github/gozookeeper/jute/test2"
	"github.com/go-zookeeper/jute/lib/go/jute"
)

type External struct {
	Shared    *Shared                 // shared
	SharedMap map[int32]*test2.Shared // sharedMap
}

func (r *External) GetShared() *Shared {
	if r != nil && r.Shared != nil {
		return r.Shared
	}
	return nil
}

func (r *External) GetSharedMap() map[int32]*test2.Shared {
	if r != nil && r.SharedMap != nil {
		return r.SharedMap
	}
	return nil
}

func (r *External) Read(dec jute.Decoder) (err error) {
	var size int
	if err = dec.ReadStart(); err != nil {
		return err
	}
	if err = dec.ReadRecord(r.Shared); err != nil {
		return err
	}
	size, err = dec.ReadMapStart()
	if err != nil {
		return err
	}
	r.SharedMap = make(map[int32]*test2.Shared)
	var k1 int32
	var v1 *test2.Shared
	for i := 0; i < size; i++ {
		k1, err = dec.ReadInt()
		if err != nil {
			return err
		}
		if err = dec.ReadRecord(v1); err != nil {
			return err
		}
		r.SharedMap[k1] = v1
	}
	if err = dec.ReadMapEnd(); err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *External) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteRecord(r.Shared); err != nil {
		return err
	}
	if err := enc.WriteMapStart(len(r.SharedMap)); err != nil {
		return err
	}
	for k, v := range r.SharedMap {
		if err := enc.WriteInt(k); err != nil {
			return err
		}
		if err := enc.WriteRecord(v); err != nil {
			return err
		}
	}
	if err := enc.WriteMapEnd(); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *External) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("External(%+v)", *r)
}
