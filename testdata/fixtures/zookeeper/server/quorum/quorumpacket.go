// Autogenerated jute compiler
// @generated from 'testdata/zookeeper.jute'

package quorum // github.com/go-zookeeper/zk/internal/server/quorum

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
	"github.com/go-zookeeper/zk/internal/data"
)

type QuorumPacket struct {
	Type     int32      // type
	Zxid     int64      // zxid
	Data     []byte     // data
	Authinfo []*data.Id // authinfo
}

func (r *QuorumPacket) GetType() int32 {
	if r != nil {
		return r.Type
	}
	return 0
}

func (r *QuorumPacket) GetZxid() int64 {
	if r != nil {
		return r.Zxid
	}
	return 0
}

func (r *QuorumPacket) GetData() []byte {
	if r != nil && r.Data != nil {
		return r.Data
	}
	return nil
}

func (r *QuorumPacket) GetAuthinfo() []*data.Id {
	if r != nil && r.Authinfo != nil {
		return r.Authinfo
	}
	return nil
}

func (r *QuorumPacket) Read(dec jute.Decoder) (err error) {
	var size int
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Type, err = dec.ReadInt()
	if err != nil {
		return err
	}
	r.Zxid, err = dec.ReadLong()
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
		r.Authinfo = nil
	} else {
		r.Authinfo = make([]*data.Id, size)
		for i := 0; i < size; i++ {
			if err = dec.ReadRecord(r.Authinfo[i]); err != nil {
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

func (r *QuorumPacket) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteInt(r.Type); err != nil {
		return err
	}
	if err := enc.WriteLong(r.Zxid); err != nil {
		return err
	}
	if err := enc.WriteBuffer(r.Data); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.Authinfo), r.Authinfo == nil); err != nil {
		return err
	}
	for _, v := range r.Authinfo {
		if err := enc.WriteRecord(v); err != nil {
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

func (r *QuorumPacket) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("QuorumPacket(%+v)", *r)
}
