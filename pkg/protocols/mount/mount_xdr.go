// Code generated by go-xdr. DO NOT EDIT.
package mount

import (
	"context"
	"fmt"
	"io"

	"github.com/buildbarn/go-xdr/pkg/protocols/rpcv2"
	"github.com/buildbarn/go-xdr/pkg/runtime"
)

const FHSIZE = 32

// Type definition "fhandle".

type Fhandle = [32]byte

func ReadFhandle(r io.Reader, m *[32]byte) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.ReadFixedLengthOpaque(r, m[:])
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func WriteFhandle(w io.Writer, m *[32]byte) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteFixedLengthOpaque(w, m[:])
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

const FhandleEncodedSizeBytes = 32

// Type definition "fhstatus".

type Fhstatus interface {
	isFhstatus()
	GetStatus() uint32
	io.WriterTo
	GetEncodedSizeBytes() int
}

func ReadFhstatus(r io.Reader) (m Fhstatus, nTotal int64, err error) {
	var nField int64
	var discriminant uint32
	{
		var m uint32
		m, nField, err = runtime.ReadUnsignedInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		discriminant = m
	}
	switch discriminant {
	case 0:
		var mArm Fhstatus_0
		{
			m := &mArm
			{
				m := &m.Directory
				nField, err = runtime.ReadFixedLengthOpaque(r, m[:])
				nTotal += nField
				if err != nil {
					goto done
				}
			}
		}
		m = &mArm
	default:
		var mArm Fhstatus_default
		{
			m := &mArm
			m.Status = discriminant
			_ = m
		}
		m = &mArm
	}
done:
	return
}

func readFhstatusStatus(r io.Reader) (m uint32, nTotal int64, err error) {
	var nField int64
	m, nField, err = runtime.ReadUnsignedInt(r)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func writeFhstatusStatus(w io.Writer, m uint32) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteUnsignedInt(w, m)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

const fhstatusStatusEncodedSizeBytes = 4

type Fhstatus_0 struct {
	Directory [32]byte
}

func (m *Fhstatus_0) isFhstatus() {}

func (m *Fhstatus_0) GetStatus() uint32 {
	return 0
}

func (m *Fhstatus_0) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		var m uint32 = 0
		nField, err = runtime.WriteUnsignedInt(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := &m.Directory
		nField, err = runtime.WriteFixedLengthOpaque(w, m[:])
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Fhstatus_0) GetEncodedSizeBytes() (nTotal int) {
	nTotal += 4
	nTotal += 32
	return
}

type Fhstatus_default struct {
	Status uint32
}

func (m *Fhstatus_default) isFhstatus() {}

func (m *Fhstatus_default) GetStatus() uint32 {
	return m.Status
}

func (m *Fhstatus_default) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Status
		nField, err = runtime.WriteUnsignedInt(w, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Fhstatus_default) GetEncodedSizeBytes() (nTotal int) {
	nTotal += 4
	return
}

const MNTPATHLEN = 1024

const MNTNAMLEN = 255

const FHSIZE3 = 64

// Type definition "fhandle3".

type Fhandle3 = []byte

func ReadFhandle3(r io.Reader) (m []byte, nTotal int64, err error) {
	var nField int64
	m, nField, err = runtime.ReadVariableLengthOpaque(r, 64)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func WriteFhandle3(w io.Writer, m []byte) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteVariableLengthOpaque(w, 64, m)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func GetFhandle3EncodedSizeBytes(m []byte) (nTotal int) {
	nTotal += (len(m) + 7) &^ 3
	return
}

// Type definition "dirpath".

type Dirpath = string

func ReadDirpath(r io.Reader) (m string, nTotal int64, err error) {
	var nField int64
	m, nField, err = runtime.ReadASCIIString(r, 1024)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func WriteDirpath(w io.Writer, m string) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteASCIIString(w, 1024, m)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func GetDirpathEncodedSizeBytes(m string) (nTotal int) {
	nTotal += (len(m) + 7) &^ 3
	return
}

// Type definition "name".

type Name = string

func ReadName(r io.Reader) (m string, nTotal int64, err error) {
	var nField int64
	m, nField, err = runtime.ReadASCIIString(r, 255)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func WriteName(w io.Writer, m string) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteASCIIString(w, 255, m)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func GetNameEncodedSizeBytes(m string) (nTotal int) {
	nTotal += (len(m) + 7) &^ 3
	return
}

// Type definition "mountstat3".

type Mountstat3 int32

func (mParent *Mountstat3) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	var m Mountstat3
	*(*int32)(&m), nField, err = runtime.ReadInt(r)
	nTotal += nField
	if err != nil {
		goto done
	}
	*mParent = m
done:
	return
}

func (m Mountstat3) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteInt(w, int32(m))
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

const Mountstat3EncodedSizeBytes = 4

const MNT3ERR_ACCES Mountstat3 = 13

const MNT3ERR_INVAL Mountstat3 = 22

const MNT3ERR_IO Mountstat3 = 5

const MNT3ERR_NAMETOOLONG Mountstat3 = 63

const MNT3ERR_NOENT Mountstat3 = 2

const MNT3ERR_NOTDIR Mountstat3 = 20

const MNT3ERR_NOTSUPP Mountstat3 = 10004

const MNT3ERR_PERM Mountstat3 = 1

const MNT3ERR_SERVERFAULT Mountstat3 = 10006

const MNT3_OK Mountstat3 = 0

var Mountstat3_name = map[Mountstat3]string{
	13:    "MNT3ERR_ACCES",
	22:    "MNT3ERR_INVAL",
	5:     "MNT3ERR_IO",
	63:    "MNT3ERR_NAMETOOLONG",
	2:     "MNT3ERR_NOENT",
	20:    "MNT3ERR_NOTDIR",
	10004: "MNT3ERR_NOTSUPP",
	1:     "MNT3ERR_PERM",
	10006: "MNT3ERR_SERVERFAULT",
	0:     "MNT3_OK",
}

const MOUNT_PROGRAM_PROGRAM_NUMBER uint32 = 100005

type MountProgram interface {
	MountV1MountprocNull(context.Context) error
	MountV1MountprocMnt(context.Context, string) (Fhstatus, error)
	MountV1MountprocDump(context.Context) (*Mountbody, error)
	MountV1MountprocUmnt(context.Context, string) error
	MountV1MountprocUmntall(context.Context) error
	MountV1MountprocExport(context.Context) (*Exportnode, error)
	MountV3Mountproc3Null(context.Context) error
	MountV3Mountproc3Mnt(context.Context, string) (Mountres3, error)
	MountV3Mountproc3Dump(context.Context) (*Mountbody, error)
	MountV3Mountproc3Umnt(context.Context, string) error
	MountV3Mountproc3Umntall(context.Context) error
	MountV3Mountproc3Export(context.Context) (*Exportnode, error)
}

func NewMountProgramService(p MountProgram) func(context.Context, uint32, uint32, io.ReadCloser, io.Writer) (rpcv2.AcceptedReplyData, error) {
	return func(ctx context.Context, vers, proc uint32, r io.ReadCloser, w io.Writer) (rpcv2.AcceptedReplyData, error) {
		var err error
		switch vers {
		case 1:
			switch proc {
			case 0:
				r.Close()
				r = nil
				err := p.MountV1MountprocNull(ctx)
				if err != nil {
					return nil, err
				}
			case 1:
				var a0 string
				{
					var m string
					var nField, nTotal int64
					m, nField, err = runtime.ReadASCIIString(r, 1024)
					nTotal += nField
					if err != nil {
						goto done
					}
					a0 = m
				}
				r.Close()
				r = nil
				m, err := p.MountV1MountprocMnt(ctx, a0)
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
				r.Close()
				r = nil
				m, err := p.MountV1MountprocDump(ctx)
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
				var a0 string
				{
					var m string
					var nField, nTotal int64
					m, nField, err = runtime.ReadASCIIString(r, 1024)
					nTotal += nField
					if err != nil {
						goto done
					}
					a0 = m
				}
				r.Close()
				r = nil
				err := p.MountV1MountprocUmnt(ctx, a0)
				if err != nil {
					return nil, err
				}
			case 4:
				r.Close()
				r = nil
				err := p.MountV1MountprocUmntall(ctx)
				if err != nil {
					return nil, err
				}
			case 5:
				r.Close()
				r = nil
				m, err := p.MountV1MountprocExport(ctx)
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
			default:
				r.Close()
				return &rpcv2.AcceptedReplyData_default{Stat: rpcv2.PROC_UNAVAIL}, nil
			}
		case 3:
			switch proc {
			case 0:
				r.Close()
				r = nil
				err := p.MountV3Mountproc3Null(ctx)
				if err != nil {
					return nil, err
				}
			case 1:
				var a0 string
				{
					var m string
					var nField, nTotal int64
					m, nField, err = runtime.ReadASCIIString(r, 1024)
					nTotal += nField
					if err != nil {
						goto done
					}
					a0 = m
				}
				r.Close()
				r = nil
				m, err := p.MountV3Mountproc3Mnt(ctx, a0)
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
				r.Close()
				r = nil
				m, err := p.MountV3Mountproc3Dump(ctx)
				if err != nil {
					return nil, err
				}
				{
					var nField, nTotal int64
					if m == nil {
						nField, err = runtime.WriteBool(w, false)
						nTotal += nField
						if err != nil {
							goto done
						}
					} else {
						nField, err = runtime.WriteBool(w, true)
						nTotal += nField
						if err != nil {
							goto done
						}
						{
							nField, err = m.WriteTo(w)
							nTotal += nField
							if err != nil {
								goto done
							}
						}
					}
				}
			case 3:
				var a0 string
				{
					var m string
					var nField, nTotal int64
					m, nField, err = runtime.ReadASCIIString(r, 1024)
					nTotal += nField
					if err != nil {
						goto done
					}
					a0 = m
				}
				r.Close()
				r = nil
				err := p.MountV3Mountproc3Umnt(ctx, a0)
				if err != nil {
					return nil, err
				}
			case 4:
				r.Close()
				r = nil
				err := p.MountV3Mountproc3Umntall(ctx)
				if err != nil {
					return nil, err
				}
			case 5:
				r.Close()
				r = nil
				m, err := p.MountV3Mountproc3Export(ctx)
				if err != nil {
					return nil, err
				}
				{
					var nField, nTotal int64
					if m == nil {
						nField, err = runtime.WriteBool(w, false)
						nTotal += nField
						if err != nil {
							goto done
						}
					} else {
						nField, err = runtime.WriteBool(w, true)
						nTotal += nField
						if err != nil {
							goto done
						}
						{
							nField, err = m.WriteTo(w)
							nTotal += nField
							if err != nil {
								goto done
							}
						}
					}
				}
			default:
				r.Close()
				return &rpcv2.AcceptedReplyData_default{Stat: rpcv2.PROC_UNAVAIL}, nil
			}
		default:
			r.Close()
			var replyData rpcv2.AcceptedReplyData_PROG_MISMATCH
			replyData.MismatchInfo.Low = 1
			replyData.MismatchInfo.High = 3
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

// Type definition "mountres3_ok".

type Mountres3Ok struct {
	Fhandle     []byte
	AuthFlavors []int32
}

func (m *Mountres3Ok) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.Fhandle
		var m []byte
		m, nField, err = runtime.ReadVariableLengthOpaque(r, 64)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.AuthFlavors
		var m []int32
		var nElements uint32
		nElements, nField, err = runtime.ReadUnsignedInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		if nElements > 4294967295 {
			err = fmt.Errorf("size of %d elements exceeds mountres3_ok.auth_flavors's maximum of 4294967295 elements", nElements)
			goto done
		}
		for nElements > 0 {
			nElements--
			mParent := &m
			var m int32
			m, nField, err = runtime.ReadInt(r)
			nTotal += nField
			if err != nil {
				goto done
			}
			*mParent = append(*mParent, m)
		}
		*mSave = m
	}
done:
	return
}

func (m *Mountres3Ok) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.Fhandle
		nField, err = runtime.WriteVariableLengthOpaque(w, 64, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.AuthFlavors
		if uint(len(m)) > 4294967295 {
			err = fmt.Errorf("size of %d elements exceeds mountres3_ok.auth_flavors's maximum of 4294967295 elements", len(m))
			goto done
		}
		nField, err = runtime.WriteUnsignedInt(w, uint32(len(m)))
		nTotal += nField
		if err != nil {
			goto done
		}
		for _, m := range m {
			nField, err = runtime.WriteInt(w, m)
			nTotal += nField
			if err != nil {
				goto done
			}
		}
	}
done:
	return
}

func (m *Mountres3Ok) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.Fhandle
		nTotal += (len(m) + 7) &^ 3
	}
	{
		m := m.AuthFlavors
		nTotal += 4 + 4*len(m)
	}
	return
}

// Type definition "mountres3".

type Mountres3 interface {
	isMountres3()
	GetFhsStatus() Mountstat3
	io.WriterTo
	GetEncodedSizeBytes() int
}

func ReadMountres3(r io.Reader) (m Mountres3, nTotal int64, err error) {
	var nField int64
	var discriminant Mountstat3
	{
		var m Mountstat3
		*(*int32)(&m), nField, err = runtime.ReadInt(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		discriminant = m
	}
	switch discriminant {
	case 0:
		var mArm Mountres3_MNT3_OK
		{
			m := &mArm
			{
				m := &m.Mountinfo
				nField, err = m.ReadFrom(r)
				nTotal += nField
				if err != nil {
					goto done
				}
			}
		}
		m = &mArm
	default:
		var mArm Mountres3_default
		{
			m := &mArm
			m.FhsStatus = discriminant
			_ = m
		}
		m = &mArm
	}
done:
	return
}

func readMountres3FhsStatus(r io.Reader) (m Mountstat3, nTotal int64, err error) {
	var nField int64
	*(*int32)(&m), nField, err = runtime.ReadInt(r)
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

func writeMountres3FhsStatus(w io.Writer, m Mountstat3) (nTotal int64, err error) {
	var nField int64
	nField, err = runtime.WriteInt(w, int32(m))
	nTotal += nField
	if err != nil {
		goto done
	}
done:
	return
}

const mountres3FhsStatusEncodedSizeBytes = 4

type Mountres3_MNT3_OK struct {
	Mountinfo Mountres3Ok
}

func (m *Mountres3_MNT3_OK) isMountres3() {}

func (m *Mountres3_MNT3_OK) GetFhsStatus() Mountstat3 {
	return 0
}

func (m *Mountres3_MNT3_OK) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		var m Mountstat3 = 0
		nField, err = runtime.WriteInt(w, int32(m))
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := &m.Mountinfo
		nField, err = m.WriteTo(w)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Mountres3_MNT3_OK) GetEncodedSizeBytes() (nTotal int) {
	nTotal += 4
	{
		m := &m.Mountinfo
		nTotal += m.GetEncodedSizeBytes()
	}
	return
}

type Mountres3_default struct {
	FhsStatus Mountstat3
}

func (m *Mountres3_default) isMountres3() {}

func (m *Mountres3_default) GetFhsStatus() Mountstat3 {
	return m.FhsStatus
}

func (m *Mountres3_default) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.FhsStatus
		nField, err = runtime.WriteInt(w, int32(m))
		nTotal += nField
		if err != nil {
			goto done
		}
	}
done:
	return
}

func (m *Mountres3_default) GetEncodedSizeBytes() (nTotal int) {
	nTotal += 4
	return
}

// Type definition "mountlist".

type Mountlist = *Mountbody

func ReadMountlist(r io.Reader) (m *Mountbody, nTotal int64, err error) {
	var nField int64
	var isSet bool
	isSet, nField, err = runtime.ReadBool(r)
	nTotal += nField
	if err != nil {
		goto done
	}
	if isSet {
		mParent := &m
		var m Mountbody
		nField, err = m.ReadFrom(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mParent = &m
	}
done:
	return
}

func WriteMountlist(w io.Writer, m *Mountbody) (nTotal int64, err error) {
	var nField int64
	if m == nil {
		nField, err = runtime.WriteBool(w, false)
		nTotal += nField
		if err != nil {
			goto done
		}
	} else {
		nField, err = runtime.WriteBool(w, true)
		nTotal += nField
		if err != nil {
			goto done
		}
		{
			nField, err = m.WriteTo(w)
			nTotal += nField
			if err != nil {
				goto done
			}
		}
	}
done:
	return
}

func GetMountlistEncodedSizeBytes(m *Mountbody) (nTotal int) {
	nTotal += 4
	if m != nil {
		nTotal += m.GetEncodedSizeBytes()
	}
	return
}

// Type definition "mountbody".

type Mountbody struct {
	MlHostname  string
	MlDirectory string
	MlNext      *Mountbody
}

func (m *Mountbody) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.MlHostname
		var m string
		m, nField, err = runtime.ReadASCIIString(r, 255)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.MlDirectory
		var m string
		m, nField, err = runtime.ReadASCIIString(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.MlNext
		var m *Mountbody
		var isSet bool
		isSet, nField, err = runtime.ReadBool(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		if isSet {
			mParent := &m
			var m Mountbody
			nField, err = m.ReadFrom(r)
			nTotal += nField
			if err != nil {
				goto done
			}
			*mParent = &m
		}
		*mSave = m
	}
done:
	return
}

func (m *Mountbody) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.MlHostname
		nField, err = runtime.WriteASCIIString(w, 255, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.MlDirectory
		nField, err = runtime.WriteASCIIString(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.MlNext
		if m == nil {
			nField, err = runtime.WriteBool(w, false)
			nTotal += nField
			if err != nil {
				goto done
			}
		} else {
			nField, err = runtime.WriteBool(w, true)
			nTotal += nField
			if err != nil {
				goto done
			}
			{
				nField, err = m.WriteTo(w)
				nTotal += nField
				if err != nil {
					goto done
				}
			}
		}
	}
done:
	return
}

func (m *Mountbody) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.MlHostname
		nTotal += (len(m) + 7) &^ 3
	}
	{
		m := m.MlDirectory
		nTotal += (len(m) + 7) &^ 3
	}
	{
		m := m.MlNext
		nTotal += 4
		if m != nil {
			nTotal += m.GetEncodedSizeBytes()
		}
	}
	return
}

// Type definition "groups".

type Groups = *Groupnode

func ReadGroups(r io.Reader) (m *Groupnode, nTotal int64, err error) {
	var nField int64
	var isSet bool
	isSet, nField, err = runtime.ReadBool(r)
	nTotal += nField
	if err != nil {
		goto done
	}
	if isSet {
		mParent := &m
		var m Groupnode
		nField, err = m.ReadFrom(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mParent = &m
	}
done:
	return
}

func WriteGroups(w io.Writer, m *Groupnode) (nTotal int64, err error) {
	var nField int64
	if m == nil {
		nField, err = runtime.WriteBool(w, false)
		nTotal += nField
		if err != nil {
			goto done
		}
	} else {
		nField, err = runtime.WriteBool(w, true)
		nTotal += nField
		if err != nil {
			goto done
		}
		{
			nField, err = m.WriteTo(w)
			nTotal += nField
			if err != nil {
				goto done
			}
		}
	}
done:
	return
}

func GetGroupsEncodedSizeBytes(m *Groupnode) (nTotal int) {
	nTotal += 4
	if m != nil {
		nTotal += m.GetEncodedSizeBytes()
	}
	return
}

// Type definition "groupnode".

type Groupnode struct {
	GrName string
	GrNext *Groupnode
}

func (m *Groupnode) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.GrName
		var m string
		m, nField, err = runtime.ReadASCIIString(r, 255)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.GrNext
		var m *Groupnode
		var isSet bool
		isSet, nField, err = runtime.ReadBool(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		if isSet {
			mParent := &m
			var m Groupnode
			nField, err = m.ReadFrom(r)
			nTotal += nField
			if err != nil {
				goto done
			}
			*mParent = &m
		}
		*mSave = m
	}
done:
	return
}

func (m *Groupnode) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.GrName
		nField, err = runtime.WriteASCIIString(w, 255, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.GrNext
		if m == nil {
			nField, err = runtime.WriteBool(w, false)
			nTotal += nField
			if err != nil {
				goto done
			}
		} else {
			nField, err = runtime.WriteBool(w, true)
			nTotal += nField
			if err != nil {
				goto done
			}
			{
				nField, err = m.WriteTo(w)
				nTotal += nField
				if err != nil {
					goto done
				}
			}
		}
	}
done:
	return
}

func (m *Groupnode) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.GrName
		nTotal += (len(m) + 7) &^ 3
	}
	{
		m := m.GrNext
		nTotal += 4
		if m != nil {
			nTotal += m.GetEncodedSizeBytes()
		}
	}
	return
}

// Type definition "exports".

type Exports = *Exportnode

func ReadExports(r io.Reader) (m *Exportnode, nTotal int64, err error) {
	var nField int64
	var isSet bool
	isSet, nField, err = runtime.ReadBool(r)
	nTotal += nField
	if err != nil {
		goto done
	}
	if isSet {
		mParent := &m
		var m Exportnode
		nField, err = m.ReadFrom(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mParent = &m
	}
done:
	return
}

func WriteExports(w io.Writer, m *Exportnode) (nTotal int64, err error) {
	var nField int64
	if m == nil {
		nField, err = runtime.WriteBool(w, false)
		nTotal += nField
		if err != nil {
			goto done
		}
	} else {
		nField, err = runtime.WriteBool(w, true)
		nTotal += nField
		if err != nil {
			goto done
		}
		{
			nField, err = m.WriteTo(w)
			nTotal += nField
			if err != nil {
				goto done
			}
		}
	}
done:
	return
}

func GetExportsEncodedSizeBytes(m *Exportnode) (nTotal int) {
	nTotal += 4
	if m != nil {
		nTotal += m.GetEncodedSizeBytes()
	}
	return
}

// Type definition "exportnode".

type Exportnode struct {
	ExDir    string
	ExGroups *Groupnode
	ExNext   *Exportnode
}

func (m *Exportnode) ReadFrom(r io.Reader) (nTotal int64, err error) {
	var nField int64
	{
		mSave := &m.ExDir
		var m string
		m, nField, err = runtime.ReadASCIIString(r, 1024)
		nTotal += nField
		if err != nil {
			goto done
		}
		*mSave = m
	}
	{
		mSave := &m.ExGroups
		var m *Groupnode
		var isSet bool
		isSet, nField, err = runtime.ReadBool(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		if isSet {
			mParent := &m
			var m Groupnode
			nField, err = m.ReadFrom(r)
			nTotal += nField
			if err != nil {
				goto done
			}
			*mParent = &m
		}
		*mSave = m
	}
	{
		mSave := &m.ExNext
		var m *Exportnode
		var isSet bool
		isSet, nField, err = runtime.ReadBool(r)
		nTotal += nField
		if err != nil {
			goto done
		}
		if isSet {
			mParent := &m
			var m Exportnode
			nField, err = m.ReadFrom(r)
			nTotal += nField
			if err != nil {
				goto done
			}
			*mParent = &m
		}
		*mSave = m
	}
done:
	return
}

func (m *Exportnode) WriteTo(w io.Writer) (nTotal int64, err error) {
	var nField int64
	{
		m := m.ExDir
		nField, err = runtime.WriteASCIIString(w, 1024, m)
		nTotal += nField
		if err != nil {
			goto done
		}
	}
	{
		m := m.ExGroups
		if m == nil {
			nField, err = runtime.WriteBool(w, false)
			nTotal += nField
			if err != nil {
				goto done
			}
		} else {
			nField, err = runtime.WriteBool(w, true)
			nTotal += nField
			if err != nil {
				goto done
			}
			{
				nField, err = m.WriteTo(w)
				nTotal += nField
				if err != nil {
					goto done
				}
			}
		}
	}
	{
		m := m.ExNext
		if m == nil {
			nField, err = runtime.WriteBool(w, false)
			nTotal += nField
			if err != nil {
				goto done
			}
		} else {
			nField, err = runtime.WriteBool(w, true)
			nTotal += nField
			if err != nil {
				goto done
			}
			{
				nField, err = m.WriteTo(w)
				nTotal += nField
				if err != nil {
					goto done
				}
			}
		}
	}
done:
	return
}

func (m *Exportnode) GetEncodedSizeBytes() (nTotal int) {
	{
		m := m.ExDir
		nTotal += (len(m) + 7) &^ 3
	}
	{
		m := m.ExGroups
		nTotal += 4
		if m != nil {
			nTotal += m.GetEncodedSizeBytes()
		}
	}
	{
		m := m.ExNext
		nTotal += 4
		if m != nil {
			nTotal += m.GetEncodedSizeBytes()
		}
	}
	return
}