// Autogenerated jute compiler
// @generated from 'testdata/zookeeper.jute'

package data // github.com/go-zookeeper/zk/internal/data

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type ACL struct {
	Perms int32 // perms
	Id    Id    // id
}

func (r *ACL) GetPerms() int32 {
	if r != nil {
		return r.Perms
	}
	return 0
}

func (r *ACL) GetId() Id {
	if r != nil {
		return r.Id
	}
	return Id{}
}

func (r *ACL) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Perms, err = dec.ReadInt()
	if err != nil {
		return err
	}
	if err = dec.ReadRecord(&r.Id); err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *ACL) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteInt(r.Perms); err != nil {
		return err
	}
	if err := enc.WriteRecord(&r.Id); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *ACL) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ACL(%+v)", *r)
}
