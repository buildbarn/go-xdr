// Code generated by go-xdr. DO NOT EDIT.
package nlm

import (
	"context"
	"io"

	"github.com/buildbarn/go-xdr/pkg/protocols/rpcv2"
	"github.com/buildbarn/go-xdr/pkg/runtime"
)

// Type definition "uint64".

type Uint64 = uint64

func ReadUint64(r io.Reader) (m uint64, nTotal int64, err error) {
	var nField int64
	m, nField, err = runtime.ReadUnsignedHyper(r)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func WriteUint64(w io.Writer, m uint64) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteUnsignedHyper(w, m)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

const Uint64EncodedSizeBytes = 8

// Type definition "int64".

type Int64 = int64

func ReadInt64(r io.Reader) (m, nTotal int64, err error) {
	var nField int64
	m, nField, err = runtime.ReadHyper(r)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func WriteInt64(w io.Writer, m int64) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteHyper(w, m)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

const Int64EncodedSizeBytes = 8

// Type definition "uint32".

type Uint32 = uint32

func ReadUint32(r io.Reader) (m uint32, nTotal int64, err error) {
	var nField int64
	m, nField, err = runtime.ReadUnsignedInt(r)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func WriteUint32(w io.Writer, m uint32) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteUnsignedInt(w, m)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

const Uint32EncodedSizeBytes = 4

// Type definition "int32".

type Int32 = int32

func ReadInt32(r io.Reader) (m int32, nTotal int64, err error) {
	var nField int64
	m, nField, err = runtime.ReadInt(r)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func WriteInt32(w io.Writer, m int32) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteInt(w, m)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

const Int32EncodedSizeBytes = 4

// Type definition "nlm4_stats".

type Nlm4Stats int32

func (mParent *Nlm4Stats) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	var m Nlm4Stats
	*(*int32)(&m), nField, err = runtime.ReadInt(r)
	nTotal += nField
	if err != nil {
		goto done
	}
	*mParent = m
done:
	return
}

func (m Nlm4Stats) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteInt(w, int32(m))
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

const Nlm4StatsEncodedSizeBytes = 4

const NLM4_BLOCKED Nlm4Stats = 3

const NLM4_DEADLCK Nlm4Stats = 5

const NLM4_DENIED Nlm4Stats = 1

const NLM4_DENIED_GRACE_PERIOD Nlm4Stats = 4

const NLM4_DENIED_NOLOCKS Nlm4Stats = 2

const NLM4_FAILED Nlm4Stats = 9

const NLM4_FBIG Nlm4Stats = 8

const NLM4_GRANTED Nlm4Stats = 0

const NLM4_ROFS Nlm4Stats = 6

const NLM4_STALE_FH Nlm4Stats = 7

var Nlm4Stats_name = map[Nlm4Stats]string{
	3: "NLM4_BLOCKED",
	5: "NLM4_DEADLCK",
	1: "NLM4_DENIED",
	4: "NLM4_DENIED_GRACE_PERIOD",
	2: "NLM4_DENIED_NOLOCKS",
	9: "NLM4_FAILED",
	8: "NLM4_FBIG",
	0: "NLM4_GRANTED",
	6: "NLM4_ROFS",
	7: "NLM4_STALE_FH",
}

// Type definition "nlm4_holder".

type Nlm4Holder struct {
	Exclusive bool
	Svid      int32
	Oh        []byte
	LOffset   uint64
	LLen      uint64
}

func (m *Nlm4Holder) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.Exclusive
		var m bool
		m, nField, err = runtime.ReadBool(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Svid
		var m int32
		m, nField, err = runtime.ReadInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Oh
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.LOffset
		var m uint64
		m, nField, err = runtime.ReadUnsignedHyper(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.LLen
		var m uint64
		m, nField, err = runtime.ReadUnsignedHyper(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
done:
	return
}

func (m *Nlm4Holder) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Exclusive
		nField, err = runtime.WriteBool(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Svid
		nField, err = runtime.WriteInt(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Oh
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.LOffset
		nField, err = runtime.WriteUnsignedHyper(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.LLen
		nField, err = runtime.WriteUnsignedHyper(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Holder) GetEncodedSizeBytes() (nTotal int) {
	nTotal += 4
	nTotal += 4
	{
		m := m.Oh
		nTotal += (len(m) + 7) &^ 3
	}
	nTotal += 8
	nTotal += 8
	return
}

// Type definition "nlm4_lock".

type Nlm4Lock struct {
	CallerName string
	Fh         []byte
	Oh         []byte
	Svid       int32
	LOffset    uint64
	LLen       uint64
}

func (m *Nlm4Lock) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.CallerName
		var m string
		m, nField, err = runtime.ReadASCIIString(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Fh
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Oh
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Svid
		var m int32
		m, nField, err = runtime.ReadInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.LOffset
		var m uint64
		m, nField, err = runtime.ReadUnsignedHyper(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.LLen
		var m uint64
		m, nField, err = runtime.ReadUnsignedHyper(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
done:
	return
}

func (m *Nlm4Lock) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.CallerName
		nField, err = runtime.WriteASCIIString(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Fh
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Oh
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Svid
		nField, err = runtime.WriteInt(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.LOffset
		nField, err = runtime.WriteUnsignedHyper(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.LLen
		nField, err = runtime.WriteUnsignedHyper(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Lock) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.CallerName
		nTotal += (len(m) + 7) &^ 3
	}
	{
		m := m.Fh
		nTotal += (len(m) + 7) &^ 3
	}
	{
		m := m.Oh
		nTotal += (len(m) + 7) &^ 3
	}
	nTotal += 4
	nTotal += 8
	nTotal += 8
	return
}

// Type definition "nlm4_share".

type Nlm4Share struct {
	CallerName string
	Fh         []byte
	Oh         []byte
	Mode       Fsh4Mode
	Access     Fsh4Access
}

func (m *Nlm4Share) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.CallerName
		var m string
		m, nField, err = runtime.ReadASCIIString(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Fh
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Oh
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Mode
		var m Fsh4Mode
		*(*int32)(&m), nField, err = runtime.ReadInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Access
		var m Fsh4Access
		*(*int32)(&m), nField, err = runtime.ReadInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
done:
	return
}

func (m *Nlm4Share) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.CallerName
		nField, err = runtime.WriteASCIIString(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Fh
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Oh
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Mode
		nField, err = runtime.WriteInt(w, int32(m))
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Access
		nField, err = runtime.WriteInt(w, int32(m))
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Share) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.CallerName
		nTotal += (len(m) + 7) &^ 3
	}
	{
		m := m.Fh
		nTotal += (len(m) + 7) &^ 3
	}
	{
		m := m.Oh
		nTotal += (len(m) + 7) &^ 3
	}
	nTotal += 4
	nTotal += 4
	return
}

const NLM_PROG_PROGRAM_NUMBER uint32 = 100021

type NlmProg interface {
	Nlm4VersNlmproc4Null(context.Context) error
	Nlm4VersNlmproc4Test(context.Context, *Nlm4Testargs) (*Nlm4Testres, error)
	Nlm4VersNlmproc4Lock(context.Context, *Nlm4Lockargs) (*Nlm4Res, error)
	Nlm4VersNlmproc4Cancel(context.Context, *Nlm4Cancargs) (*Nlm4Res, error)
	Nlm4VersNlmproc4Unlock(context.Context, *Nlm4Unlockargs) (*Nlm4Res, error)
	Nlm4VersNlmproc4Granted(context.Context, *Nlm4Testargs) (*Nlm4Res, error)
	Nlm4VersNlmproc4TestMsg(context.Context, *Nlm4Testargs) error
	Nlm4VersNlmproc4LockMsg(context.Context, *Nlm4Lockargs) error
	Nlm4VersNlmproc4CancelMsg(context.Context, *Nlm4Cancargs) error
	Nlm4VersNlmproc4UnlockMsg(context.Context, *Nlm4Unlockargs) error
	Nlm4VersNlmproc4GrantedMsg(context.Context, *Nlm4Testargs) error
	Nlm4VersNlmproc4TestRes(context.Context, *Nlm4Testres) error
	Nlm4VersNlmproc4LockRes(context.Context, *Nlm4Res) error
	Nlm4VersNlmproc4CancelRes(context.Context, *Nlm4Res) error
	Nlm4VersNlmproc4UnlockRes(context.Context, *Nlm4Res) error
	Nlm4VersNlmproc4GrantedRes(context.Context, *Nlm4Res) error
	Nlm4VersNlmproc4Share(context.Context, *Nlm4Shareargs) (*Nlm4Shareres, error)
	Nlm4VersNlmproc4Unshare(context.Context, *Nlm4Shareargs) (*Nlm4Shareres, error)
	Nlm4VersNlmproc4NmLock(context.Context, *Nlm4Lockargs) (*Nlm4Res, error)
	Nlm4VersNlmproc4FreeAll(context.Context, *Nlm4Notify) error
}

func NewNlmProgService(p NlmProg) func(context.Context, uint32, uint32, io.ReadCloser, io.Writer) (rpcv2.AcceptedReplyData, error) {
	return func(ctx context.Context, vers, proc uint32, r io.ReadCloser, w io.Writer) (rpcv2.AcceptedReplyData, error) {
		var err error
		switch vers {
		case 4:
			switch proc {
			case 0:
				r.Close()
				r = nil
				err := p.Nlm4VersNlmproc4Null(ctx)
				if err != nil {
					return nil, err
				}
			case 1:
				var a0 Nlm4Testargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				m, err := p.Nlm4VersNlmproc4Test(ctx, &a0)
				if err != nil {
					return nil, err
				}
				{
					var nField, nTotal int64
					nField, err = m.WriteTo(w)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
			case 2:
				var a0 Nlm4Lockargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				m, err := p.Nlm4VersNlmproc4Lock(ctx, &a0)
				if err != nil {
					return nil, err
				}
				{
					var nField, nTotal int64
					nField, err = m.WriteTo(w)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
			case 3:
				var a0 Nlm4Cancargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				m, err := p.Nlm4VersNlmproc4Cancel(ctx, &a0)
				if err != nil {
					return nil, err
				}
				{
					var nField, nTotal int64
					nField, err = m.WriteTo(w)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
			case 4:
				var a0 Nlm4Unlockargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				m, err := p.Nlm4VersNlmproc4Unlock(ctx, &a0)
				if err != nil {
					return nil, err
				}
				{
					var nField, nTotal int64
					nField, err = m.WriteTo(w)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
			case 5:
				var a0 Nlm4Testargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				m, err := p.Nlm4VersNlmproc4Granted(ctx, &a0)
				if err != nil {
					return nil, err
				}
				{
					var nField, nTotal int64
					nField, err = m.WriteTo(w)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
			case 6:
				var a0 Nlm4Testargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				err := p.Nlm4VersNlmproc4TestMsg(ctx, &a0)
				if err != nil {
					return nil, err
				}
			case 7:
				var a0 Nlm4Lockargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				err := p.Nlm4VersNlmproc4LockMsg(ctx, &a0)
				if err != nil {
					return nil, err
				}
			case 8:
				var a0 Nlm4Cancargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				err := p.Nlm4VersNlmproc4CancelMsg(ctx, &a0)
				if err != nil {
					return nil, err
				}
			case 9:
				var a0 Nlm4Unlockargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				err := p.Nlm4VersNlmproc4UnlockMsg(ctx, &a0)
				if err != nil {
					return nil, err
				}
			case 10:
				var a0 Nlm4Testargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				err := p.Nlm4VersNlmproc4GrantedMsg(ctx, &a0)
				if err != nil {
					return nil, err
				}
			case 11:
				var a0 Nlm4Testres
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				err := p.Nlm4VersNlmproc4TestRes(ctx, &a0)
				if err != nil {
					return nil, err
				}
			case 12:
				var a0 Nlm4Res
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				err := p.Nlm4VersNlmproc4LockRes(ctx, &a0)
				if err != nil {
					return nil, err
				}
			case 13:
				var a0 Nlm4Res
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				err := p.Nlm4VersNlmproc4CancelRes(ctx, &a0)
				if err != nil {
					return nil, err
				}
			case 14:
				var a0 Nlm4Res
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				err := p.Nlm4VersNlmproc4UnlockRes(ctx, &a0)
				if err != nil {
					return nil, err
				}
			case 15:
				var a0 Nlm4Res
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				err := p.Nlm4VersNlmproc4GrantedRes(ctx, &a0)
				if err != nil {
					return nil, err
				}
			case 20:
				var a0 Nlm4Shareargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				m, err := p.Nlm4VersNlmproc4Share(ctx, &a0)
				if err != nil {
					return nil, err
				}
				{
					var nField, nTotal int64
					nField, err = m.WriteTo(w)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
			case 21:
				var a0 Nlm4Shareargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				m, err := p.Nlm4VersNlmproc4Unshare(ctx, &a0)
				if err != nil {
					return nil, err
				}
				{
					var nField, nTotal int64
					nField, err = m.WriteTo(w)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
			case 22:
				var a0 Nlm4Lockargs
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				m, err := p.Nlm4VersNlmproc4NmLock(ctx, &a0)
				if err != nil {
					return nil, err
				}
				{
					var nField, nTotal int64
					nField, err = m.WriteTo(w)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
			case 23:
				var a0 Nlm4Notify
				{
					m := &a0
					var nField, nTotal int64
					nField, err = m.ReadFrom(r)
					nTotal += nField
					if err != nil {
						goto done
					}
				}
				r.Close()
				r = nil
				err := p.Nlm4VersNlmproc4FreeAll(ctx, &a0)
				if err != nil {
					return nil, err
				}
			default:
				r.Close()
				return &rpcv2.AcceptedReplyData_default{Stat: rpcv2.PROC_UNAVAIL}, nil
			}
		default:
			r.Close()
			var replyData rpcv2.AcceptedReplyData_PROG_MISMATCH
			replyData.MismatchInfo.Low = 4
			replyData.MismatchInfo.High = 4
			return &replyData, nil
		}
		return &rpcv2.AcceptedReplyData_SUCCESS{}, nil
	done:
		if r != nil {
			r.Close()
		}
		return nil, err
	}
}

const LM_MAXSTRLEN = 1024

const LM_MAXNAMELEN = 1025

const MAXNETOBJ_SZ = 1024

// Type definition "netobj".

type Netobj = []byte

func ReadNetobj(r io.Reader) (m []byte, nTotal int64, err error) {
	var nField int64
	m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func WriteNetobj(w io.Writer, m []byte) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func GetNetobjEncodedSizeBytes(m []byte) (nTotal int) {
	nTotal += (len(m) + 7) &^ 3
	return
}

// Type definition "nlm4_stat".

type Nlm4Stat struct {
	Stat Nlm4Stats
}

func (m *Nlm4Stat) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.Stat
		var m Nlm4Stats
		*(*int32)(&m), nField, err = runtime.ReadInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
done:
	return
}

func (m *Nlm4Stat) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Stat
		nField, err = runtime.WriteInt(w, int32(m))
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

const Nlm4StatEncodedSizeBytes = 4

// Type definition "nlm4_res".

type Nlm4Res struct {
	Cookie []byte
	Stat   Nlm4Stat
}

func (m *Nlm4Res) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.Cookie
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		m := &m.Stat
		nField, err = m.ReadFrom(r)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Res) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Cookie
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := &m.Stat
		nField, err = m.WriteTo(w)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Res) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.Cookie
		nTotal += (len(m) + 7) &^ 3
	}
	nTotal += 4
	return
}

// Type definition "nlm4_testrply".

type Nlm4Testrply interface {
	isNlm4Testrply()
	GetStat() Nlm4Stats
	io.WriterTo
	GetEncodedSizeBytes() int
}

func ReadNlm4Testrply(r io.Reader) (m Nlm4Testrply, nTotal int64, err error) {
	var nField int64
	var discriminant Nlm4Stats
	{
		var m Nlm4Stats
		*(*int32)(&m), nField, err = runtime.ReadInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		discriminant = m
	}
	switch discriminant {
	case 1:
		var mArm Nlm4Testrply_NLM4_DENIED
		{
			m := &mArm
			{
				m := &m.Holder
				nField, err = m.ReadFrom(r)
				nTotal += nField
				if err != nil {
					goto done
				}
			}
		}
		m = &mArm
	default:
		var mArm Nlm4Testrply_default
		{
			m := &mArm
			m.Stat = discriminant
			_ = m
		}
		m = &mArm
	}
done:
	return
}

func readNlm4TestrplyStat(r io.Reader) (m Nlm4Stats, nTotal int64, err error) {
	var nField int64
	*(*int32)(&m), nField, err = runtime.ReadInt(r)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func writeNlm4TestrplyStat(w io.Writer, m Nlm4Stats) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteInt(w, int32(m))
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

const nlm4TestrplyStatEncodedSizeBytes = 4

type Nlm4Testrply_NLM4_DENIED struct {
	Holder Nlm4Holder
}

func (m *Nlm4Testrply_NLM4_DENIED) isNlm4Testrply() {}

func (m *Nlm4Testrply_NLM4_DENIED) GetStat() Nlm4Stats {
	return 1
}

func (m *Nlm4Testrply_NLM4_DENIED) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		var m Nlm4Stats = 1
		nField, err = runtime.WriteInt(w, int32(m))
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := &m.Holder
		nField, err = m.WriteTo(w)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Testrply_NLM4_DENIED) GetEncodedSizeBytes() (nTotal int) {
	nTotal += 4
	{
		m := &m.Holder
		nTotal += m.GetEncodedSizeBytes()
	}
	return
}

type Nlm4Testrply_default struct {
	Stat Nlm4Stats
}

func (m *Nlm4Testrply_default) isNlm4Testrply() {}

func (m *Nlm4Testrply_default) GetStat() Nlm4Stats {
	return m.Stat
}

func (m *Nlm4Testrply_default) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Stat
		nField, err = runtime.WriteInt(w, int32(m))
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Testrply_default) GetEncodedSizeBytes() (nTotal int) {
	nTotal += 4
	return
}

// Type definition "nlm4_testres".

type Nlm4Testres struct {
	Cookie   []byte
	TestStat Nlm4Testrply
}

func (m *Nlm4Testres) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.Cookie
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.TestStat
		var m Nlm4Testrply
		m, nField, err = ReadNlm4Testrply(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
done:
	return
}

func (m *Nlm4Testres) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Cookie
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.TestStat
		nField, err = m.WriteTo(w)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Testres) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.Cookie
		nTotal += (len(m) + 7) &^ 3
	}
	{
		m := m.TestStat
		nTotal += m.GetEncodedSizeBytes()
	}
	return
}

// Type definition "nlm4_lockargs".

type Nlm4Lockargs struct {
	Cookie    []byte
	Block     bool
	Exclusive bool
	Alock     Nlm4Lock
	Reclaim   bool
	State     int32
}

func (m *Nlm4Lockargs) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.Cookie
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Block
		var m bool
		m, nField, err = runtime.ReadBool(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Exclusive
		var m bool
		m, nField, err = runtime.ReadBool(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		m := &m.Alock
		nField, err = m.ReadFrom(r)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		mSave := &m.Reclaim
		var m bool
		m, nField, err = runtime.ReadBool(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.State
		var m int32
		m, nField, err = runtime.ReadInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
done:
	return
}

func (m *Nlm4Lockargs) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Cookie
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Block
		nField, err = runtime.WriteBool(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Exclusive
		nField, err = runtime.WriteBool(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := &m.Alock
		nField, err = m.WriteTo(w)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Reclaim
		nField, err = runtime.WriteBool(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.State
		nField, err = runtime.WriteInt(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Lockargs) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.Cookie
		nTotal += (len(m) + 7) &^ 3
	}
	nTotal += 4
	nTotal += 4
	{
		m := &m.Alock
		nTotal += m.GetEncodedSizeBytes()
	}
	nTotal += 4
	nTotal += 4
	return
}

// Type definition "nlm4_cancargs".

type Nlm4Cancargs struct {
	Cookie    []byte
	Block     bool
	Exclusive bool
	Alock     Nlm4Lock
}

func (m *Nlm4Cancargs) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.Cookie
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Block
		var m bool
		m, nField, err = runtime.ReadBool(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Exclusive
		var m bool
		m, nField, err = runtime.ReadBool(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		m := &m.Alock
		nField, err = m.ReadFrom(r)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Cancargs) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Cookie
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Block
		nField, err = runtime.WriteBool(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Exclusive
		nField, err = runtime.WriteBool(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := &m.Alock
		nField, err = m.WriteTo(w)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Cancargs) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.Cookie
		nTotal += (len(m) + 7) &^ 3
	}
	nTotal += 4
	nTotal += 4
	{
		m := &m.Alock
		nTotal += m.GetEncodedSizeBytes()
	}
	return
}

// Type definition "nlm4_testargs".

type Nlm4Testargs struct {
	Cookie    []byte
	Exclusive bool
	Alock     Nlm4Lock
}

func (m *Nlm4Testargs) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.Cookie
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Exclusive
		var m bool
		m, nField, err = runtime.ReadBool(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		m := &m.Alock
		nField, err = m.ReadFrom(r)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Testargs) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Cookie
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Exclusive
		nField, err = runtime.WriteBool(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := &m.Alock
		nField, err = m.WriteTo(w)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Testargs) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.Cookie
		nTotal += (len(m) + 7) &^ 3
	}
	nTotal += 4
	{
		m := &m.Alock
		nTotal += m.GetEncodedSizeBytes()
	}
	return
}

// Type definition "nlm4_unlockargs".

type Nlm4Unlockargs struct {
	Cookie []byte
	Alock  Nlm4Lock
}

func (m *Nlm4Unlockargs) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.Cookie
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		m := &m.Alock
		nField, err = m.ReadFrom(r)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Unlockargs) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Cookie
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := &m.Alock
		nField, err = m.WriteTo(w)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Unlockargs) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.Cookie
		nTotal += (len(m) + 7) &^ 3
	}
	{
		m := &m.Alock
		nTotal += m.GetEncodedSizeBytes()
	}
	return
}

// Type definition "fsh4_mode".

type Fsh4Mode int32

func (mParent *Fsh4Mode) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	var m Fsh4Mode
	*(*int32)(&m), nField, err = runtime.ReadInt(r)
	nTotal += nField
	if err != nil {
		goto done
	}
	*mParent = m
done:
	return
}

func (m Fsh4Mode) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteInt(w, int32(m))
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

const Fsh4ModeEncodedSizeBytes = 4

const FSM_DN Fsh4Mode = 0

const FSM_DR Fsh4Mode = 1

const FSM_DRW Fsh4Mode = 3

const FSM_DW Fsh4Mode = 2

var Fsh4Mode_name = map[Fsh4Mode]string{
	0: "fsm_DN",
	1: "fsm_DR",
	3: "fsm_DRW",
	2: "fsm_DW",
}

// Type definition "fsh4_access".

type Fsh4Access int32

func (mParent *Fsh4Access) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	var m Fsh4Access
	*(*int32)(&m), nField, err = runtime.ReadInt(r)
	nTotal += nField
	if err != nil {
		goto done
	}
	*mParent = m
done:
	return
}

func (m Fsh4Access) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteInt(w, int32(m))
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

const Fsh4AccessEncodedSizeBytes = 4

const FSA_NONE Fsh4Access = 0

const FSA_R Fsh4Access = 1

const FSA_RW Fsh4Access = 3

const FSA_W Fsh4Access = 2

var Fsh4Access_name = map[Fsh4Access]string{
	0: "fsa_NONE",
	1: "fsa_R",
	3: "fsa_RW",
	2: "fsa_W",
}

// Type definition "nlm4_shareargs".

type Nlm4Shareargs struct {
	Cookie  []byte
	Share   Nlm4Share
	Reclaim bool
}

func (m *Nlm4Shareargs) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.Cookie
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		m := &m.Share
		nField, err = m.ReadFrom(r)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		mSave := &m.Reclaim
		var m bool
		m, nField, err = runtime.ReadBool(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
done:
	return
}

func (m *Nlm4Shareargs) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Cookie
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := &m.Share
		nField, err = m.WriteTo(w)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Reclaim
		nField, err = runtime.WriteBool(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Shareargs) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.Cookie
		nTotal += (len(m) + 7) &^ 3
	}
	{
		m := &m.Share
		nTotal += m.GetEncodedSizeBytes()
	}
	nTotal += 4
	return
}

// Type definition "nlm4_shareres".

type Nlm4Shareres struct {
	Cookie   []byte
	Stat     Nlm4Stats
	Sequence int32
}

func (m *Nlm4Shareres) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.Cookie
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Stat
		var m Nlm4Stats
		*(*int32)(&m), nField, err = runtime.ReadInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.Sequence
		var m int32
		m, nField, err = runtime.ReadInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
done:
	return
}

func (m *Nlm4Shareres) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Cookie
		nField, err = runtime.WriteVariableLengthOpaque(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Stat
		nField, err = runtime.WriteInt(w, int32(m))
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.Sequence
		nField, err = runtime.WriteInt(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Shareres) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.Cookie
		nTotal += (len(m) + 7) &^ 3
	}
	nTotal += 4
	nTotal += 4
	return
}

// Type definition "nlm4_notify".

type Nlm4Notify struct {
	Name  string
	State int32
}

func (m *Nlm4Notify) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.Name
		var m string
		m, nField, err = runtime.ReadASCIIString(r, 1025)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.State
		var m int32
		m, nField, err = runtime.ReadInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
done:
	return
}

func (m *Nlm4Notify) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Name
		nField, err = runtime.WriteASCIIString(w, 1025, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.State
		nField, err = runtime.WriteInt(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Nlm4Notify) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.Name
		nTotal += (len(m) + 7) &^ 3
	}
	nTotal += 4
	return
}
