// Autogenerated jute compiler
// @generated from 'testdata/zookeeper.jute'

package txn // github.com/go-zookeeper/zk/internal/txn

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
	"github.com/go-zookeeper/zk/internal/data"
)

type CreateTxn struct {
	Path           *string     // path
	Data           []byte      // data
	Acl            []*data.ACL // acl
	Ephemeral      bool        // ephemeral
	ParentCVersion int32       // parentCVersion
}

func (r *CreateTxn) GetPath() string {
	if r != nil && r.Path != nil {
		return *r.Path
	}
	return ""
}

func (r *CreateTxn) GetData() []byte {
	if r != nil && r.Data != nil {
		return r.Data
	}
	return nil
}

func (r *CreateTxn) GetAcl() []*data.ACL {
	if r != nil && r.Acl != nil {
		return r.Acl
	}
	return nil
}

func (r *CreateTxn) GetEphemeral() bool {
	if r != nil {
		return r.Ephemeral
	}
	return false
}

func (r *CreateTxn) GetParentCVersion() int32 {
	if r != nil {
		return r.ParentCVersion
	}
	return 0
}

func (r *CreateTxn) Read(dec jute.Decoder) (err error) {
	var size int
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Path, err = dec.ReadString()
	if err != nil {
		return err
	}
	r.Data, err = dec.ReadBuffer()
	if err != nil {
		return err
	}
	size, err = dec.ReadVectorStart()
	if err != nil {
		return err
	}
	if size < 0 {
		r.Acl = nil
	} else {
		r.Acl = make([]*data.ACL, size)
		for i := 0; i < size; i++ {
			if err = dec.ReadRecord(r.Acl[i]); err != nil {
				return err
			}
		}
	}
	if err = dec.ReadVectorEnd(); err != nil {
		return err
	}
	r.Ephemeral, err = dec.ReadBoolean()
	if err != nil {
		return err
	}
	r.ParentCVersion, err = dec.ReadInt()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *CreateTxn) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteString(r.Path); err != nil {
		return err
	}
	if err := enc.WriteBuffer(r.Data); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.Acl), r.Acl == nil); err != nil {
		return err
	}
	for _, v := range r.Acl {
		if err := enc.WriteRecord(v); err != nil {
			return err
		}
	}
	if err := enc.WriteVectorEnd(); err != nil {
		return err
	}
	if err := enc.WriteBoolean(r.Ephemeral); err != nil {
		return err
	}
	if err := enc.WriteInt(r.ParentCVersion); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *CreateTxn) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateTxn(%+v)", *r)
}
