// Autogenerated jute compiler
// @generated from '/home/pmazzini/repos/jute/testdata/zookeeper.jute'

package data // github.com/go-zookeeper/zk/internal/data

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type Id struct {
	Scheme *string // scheme
	Id     *string // id
}

func (r *Id) GetScheme() string {
	if r != nil && r.Scheme != nil {
		return *r.Scheme
	}
	return ""
}

func (r *Id) GetId() string {
	if r != nil && r.Id != nil {
		return *r.Id
	}
	return ""
}

func (r *Id) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Scheme, err = dec.ReadString()
	if err != nil {
		return err
	}
	r.Id, err = dec.ReadString()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *Id) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteString(r.Scheme); err != nil {
		return err
	}
	if err := enc.WriteString(r.Id); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *Id) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Id(%+v)", *r)
}
