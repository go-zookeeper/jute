// Autogenerated jute compiler
// @generated from 'testdata/zookeeper.jute'

package proto // github.com/go-zookeeper/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
	"github.com/go-zookeeper/zk/internal/data"
)

type SetACLResponse struct {
	Stat *data.Stat // stat
}

func (r *SetACLResponse) GetStat() *data.Stat {
	if r != nil && r.Stat != nil {
		return r.Stat
	}
	return nil
}

func (r *SetACLResponse) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	if err = dec.ReadRecord(r.Stat); err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *SetACLResponse) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteRecord(r.Stat); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *SetACLResponse) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SetACLResponse(%+v)", *r)
}
