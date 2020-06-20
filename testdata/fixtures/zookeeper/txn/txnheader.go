// Autogenerated jute compiler
// @generated from '/home/pmazzini/repos/jute/testdata/zookeeper.jute'

package txn // github.com/go-zookeeper/zk/internal/txn

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type TxnHeader struct {
	ClientId int64 // clientId
	Cxid     int32 // cxid
	Zxid     int64 // zxid
	Time     int64 // time
	Type     int32 // type
}

func (r *TxnHeader) GetClientId() int64 {
	if r != nil {
		return r.ClientId
	}
	return 0
}

func (r *TxnHeader) GetCxid() int32 {
	if r != nil {
		return r.Cxid
	}
	return 0
}

func (r *TxnHeader) GetZxid() int64 {
	if r != nil {
		return r.Zxid
	}
	return 0
}

func (r *TxnHeader) GetTime() int64 {
	if r != nil {
		return r.Time
	}
	return 0
}

func (r *TxnHeader) GetType() int32 {
	if r != nil {
		return r.Type
	}
	return 0
}

func (r *TxnHeader) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.ClientId, err = dec.ReadLong()
	if err != nil {
		return err
	}
	r.Cxid, err = dec.ReadInt()
	if err != nil {
		return err
	}
	r.Zxid, err = dec.ReadLong()
	if err != nil {
		return err
	}
	r.Time, err = dec.ReadLong()
	if err != nil {
		return err
	}
	r.Type, err = dec.ReadInt()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *TxnHeader) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteLong(r.ClientId); err != nil {
		return err
	}
	if err := enc.WriteInt(r.Cxid); err != nil {
		return err
	}
	if err := enc.WriteLong(r.Zxid); err != nil {
		return err
	}
	if err := enc.WriteLong(r.Time); err != nil {
		return err
	}
	if err := enc.WriteInt(r.Type); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *TxnHeader) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TxnHeader(%+v)", *r)
}
