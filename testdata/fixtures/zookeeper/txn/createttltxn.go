// Autogenerated jute compiler
// @generated from 'testdata/zookeeper.jute'

package txn // github.com/go-zookeeper/zk/internal/txn

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
	"github.com/go-zookeeper/zk/internal/data"
)

type CreateTTLTxn struct {
	Path           *string     // path
	Data           []byte      // data
	Acl            []*data.ACL // acl
	ParentCVersion int32       // parentCVersion
	Ttl            int64       // ttl
}

func (r *CreateTTLTxn) GetPath() string {
	if r != nil && r.Path != nil {
		return *r.Path
	}
	return ""
}

func (r *CreateTTLTxn) GetData() []byte {
	if r != nil && r.Data != nil {
		return r.Data
	}
	return nil
}

func (r *CreateTTLTxn) GetAcl() []*data.ACL {
	if r != nil && r.Acl != nil {
		return r.Acl
	}
	return nil
}

func (r *CreateTTLTxn) GetParentCVersion() int32 {
	if r != nil {
		return r.ParentCVersion
	}
	return 0
}

func (r *CreateTTLTxn) GetTtl() int64 {
	if r != nil {
		return r.Ttl
	}
	return 0
}

func (r *CreateTTLTxn) Read(dec jute.Decoder) (err error) {
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
	r.ParentCVersion, err = dec.ReadInt()
	if err != nil {
		return err
	}
	r.Ttl, err = dec.ReadLong()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *CreateTTLTxn) Write(enc jute.Encoder) error {
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
	if err := enc.WriteInt(r.ParentCVersion); err != nil {
		return err
	}
	if err := enc.WriteLong(r.Ttl); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *CreateTTLTxn) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateTTLTxn(%+v)", *r)
}
