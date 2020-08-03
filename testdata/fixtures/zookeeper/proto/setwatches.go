// Autogenerated jute compiler
// @generated from 'testdata/zookeeper.jute'

package proto // github.com/go-zookeeper/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type SetWatches struct {
	RelativeZxid int64    // relativeZxid
	DataWatches  []string // dataWatches
	ExistWatches []string // existWatches
	ChildWatches []string // childWatches
}

func (r *SetWatches) GetRelativeZxid() int64 {
	if r != nil {
		return r.RelativeZxid
	}
	return 0
}

func (r *SetWatches) GetDataWatches() []string {
	if r != nil && r.DataWatches != nil {
		return r.DataWatches
	}
	return nil
}

func (r *SetWatches) GetExistWatches() []string {
	if r != nil && r.ExistWatches != nil {
		return r.ExistWatches
	}
	return nil
}

func (r *SetWatches) GetChildWatches() []string {
	if r != nil && r.ChildWatches != nil {
		return r.ChildWatches
	}
	return nil
}

func (r *SetWatches) Read(dec jute.Decoder) (err error) {
	var size int
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.RelativeZxid, err = dec.ReadLong()
	if err != nil {
		return err
	}
	size, err = dec.ReadVectorStart()
	if err != nil {
		return err
	}
	if size < 0 {
		r.DataWatches = nil
	} else {
		r.DataWatches = make([]string, size)
		for i := 0; i < size; i++ {
			r.DataWatches[i], err = dec.ReadString()
			if err != nil {
				return err
			}
		}
	}
	if err = dec.ReadVectorEnd(); err != nil {
		return err
	}
	size, err = dec.ReadVectorStart()
	if err != nil {
		return err
	}
	if size < 0 {
		r.ExistWatches = nil
	} else {
		r.ExistWatches = make([]string, size)
		for i := 0; i < size; i++ {
			r.ExistWatches[i], err = dec.ReadString()
			if err != nil {
				return err
			}
		}
	}
	if err = dec.ReadVectorEnd(); err != nil {
		return err
	}
	size, err = dec.ReadVectorStart()
	if err != nil {
		return err
	}
	if size < 0 {
		r.ChildWatches = nil
	} else {
		r.ChildWatches = make([]string, size)
		for i := 0; i < size; i++ {
			r.ChildWatches[i], err = dec.ReadString()
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

func (r *SetWatches) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteLong(r.RelativeZxid); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.DataWatches), r.DataWatches == nil); err != nil {
		return err
	}
	for _, v := range r.DataWatches {
		if err := enc.WriteString(v); err != nil {
			return err
		}
	}
	if err := enc.WriteVectorEnd(); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.ExistWatches), r.ExistWatches == nil); err != nil {
		return err
	}
	for _, v := range r.ExistWatches {
		if err := enc.WriteString(v); err != nil {
			return err
		}
	}
	if err := enc.WriteVectorEnd(); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.ChildWatches), r.ChildWatches == nil); err != nil {
		return err
	}
	for _, v := range r.ChildWatches {
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

func (r *SetWatches) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SetWatches(%+v)", *r)
}
