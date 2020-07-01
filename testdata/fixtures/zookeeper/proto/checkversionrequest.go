// Autogenerated jute compiler
// @generated from 'testdata/zookeeper.jute'

package proto // github.com/go-zookeeper/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type CheckVersionRequest struct {
	Path    *string // path
	Version int32   // version
}

func (r *CheckVersionRequest) GetPath() string {
	if r != nil && r.Path != nil {
		return *r.Path
	}
	return ""
}

func (r *CheckVersionRequest) GetVersion() int32 {
	if r != nil {
		return r.Version
	}
	return 0
}

func (r *CheckVersionRequest) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Path, err = dec.ReadString()
	if err != nil {
		return err
	}
	r.Version, err = dec.ReadInt()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *CheckVersionRequest) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteString(r.Path); err != nil {
		return err
	}
	if err := enc.WriteInt(r.Version); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *CheckVersionRequest) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CheckVersionRequest(%+v)", *r)
}
