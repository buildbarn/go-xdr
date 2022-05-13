package runtime_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/buildbarn/go-xdr/internal/mock"
	"github.com/buildbarn/go-xdr/pkg/runtime"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestReadUTF8String(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("IOFailure", func(t *testing.T) {
		r := mock.NewMockReader(ctrl)
		r.EXPECT().Read(gomock.Any()).Return(0, errors.New("disk failure"))

		_, n, err := runtime.ReadUTF8String(r, 10)
		require.Equal(t, errors.New("disk failure"), err)
		require.Equal(t, int64(0), n)
	})

	t.Run("InvalidCharacters", func(t *testing.T) {
		_, n, err := runtime.ReadUTF8String(
			bytes.NewBuffer([]byte{0, 0, 0, 5, 0x48, 0x65, 0x8a, 0x6c, 0x6f, 0, 0, 0}),
			10)
		require.Equal(t, errors.New("string is not valid UTF-8"), err)
		require.Equal(t, int64(12), n)
	})

	t.Run("Success", func(t *testing.T) {
		s, n, err := runtime.ReadUTF8String(
			bytes.NewBuffer([]byte{0, 0, 0, 5, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0, 0, 0}),
			10)
		require.NoError(t, err)
		require.Equal(t, int64(12), n)
		require.Equal(t, "Hello", s)
	})
}

func TestWriteUTF8String(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("IOFailure", func(t *testing.T) {
		w := mock.NewMockWriter(ctrl)
		w.EXPECT().Write(gomock.Any()).Return(0, errors.New("disk failure"))

		n, err := runtime.WriteUTF8String(w, 10, "Hello")
		require.Equal(t, errors.New("disk failure"), err)
		require.Equal(t, int64(0), n)
	})

	t.Run("InvalidCharacters", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteUTF8String(b, 30, string([]byte{0xff, 0xfe, 0xfd}))
		require.Equal(t, errors.New("string is not valid UTF-8"), err)
		require.Equal(t, int64(0), n)
	})

	t.Run("Success", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteUTF8String(b, 10, "Hello")
		require.NoError(t, err)
		require.Equal(t, int64(12), n)
		require.Equal(t, []byte{0, 0, 0, 5, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0, 0, 0}, b.Bytes())
	})
}
