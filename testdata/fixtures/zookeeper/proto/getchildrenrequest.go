// Autogenerated jute compiler
// @generated from '/home/pmazzini/repos/jute/testdata/zookeeper.jute'

package proto // github.com/go-zookeeper/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type GetChildrenRequest struct {
	Path  *string // path
	Watch bool    // watch
}

func (r *GetChildrenRequest) GetPath() string {
	if r != nil && r.Path != nil {
		return *r.Path
	}
	return ""
}

func (r *GetChildrenRequest) GetWatch() bool {
	if r != nil {
		return r.Watch
	}
	return false
}

func (r *GetChildrenRequest) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Path, err = dec.ReadString()
	if err != nil {
		return err
	}
	r.Watch, err = dec.ReadBoolean()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *GetChildrenRequest) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteString(r.Path); err != nil {
		return err
	}
	if err := enc.WriteBoolean(r.Watch); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *GetChildrenRequest) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetChildrenRequest(%+v)", *r)
}
