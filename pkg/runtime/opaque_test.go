package runtime_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/buildbarn/go-xdr/internal/mock"
	"github.com/buildbarn/go-xdr/pkg/runtime"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestReadFixedLengthOpaque(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("IOFailure", func(t *testing.T) {
		r := mock.NewMockReader(ctrl)
		r.EXPECT().Read(gomock.Any()).Return(0, errors.New("disk failure"))

		var b [10]byte
		n, err := runtime.ReadFixedLengthOpaque(r, b[:])
		require.Equal(t, errors.New("disk failure"), err)
		require.Equal(t, int64(0), n)
	})

	t.Run("EOF", func(t *testing.T) {
		var b [10]byte
		n, err := runtime.ReadFixedLengthOpaque(bytes.NewBuffer(nil), b[:])
		require.Equal(t, io.EOF, err)
		require.Equal(t, int64(0), n)
	})

	t.Run("UnexpectedEOF", func(t *testing.T) {
		input := [...]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		for i := 1; i <= len(input); i++ {
			var b [10]byte
			n, err := runtime.ReadFixedLengthOpaque(bytes.NewBuffer(input[:i]), b[:])
			require.Equal(t, io.ErrUnexpectedEOF, err)
			require.Equal(t, int64(i), n)
		}
	})

	t.Run("InvalidPadding", func(t *testing.T) {
		var b [2]byte
		n, err := runtime.ReadFixedLengthOpaque(bytes.NewBuffer([]byte{5, 2, 3, 7}), b[:])
		require.Equal(t, errors.New("padding contains non-zero values [3 7]"), err)
		require.Equal(t, int64(4), n)
	})

	t.Run("SuccessPadding0Bytes", func(t *testing.T) {
		var b [4]byte
		n, err := runtime.ReadFixedLengthOpaque(
			bytes.NewBuffer([]byte{0x2d, 0x11, 0x8b, 0xa4}),
			b[:])
		require.NoError(t, err)
		require.Equal(t, int64(4), n)
		require.Equal(t, [...]byte{0x2d, 0x11, 0x8b, 0xa4}, b)
	})

	t.Run("SuccessPadding1Byte", func(t *testing.T) {
		var b [5]byte
		n, err := runtime.ReadFixedLengthOpaque(
			bytes.NewBuffer([]byte{0x2d, 0x11, 0x8b, 0xa4, 0x72, 0x00, 0x00, 0x00}),
			b[:])
		require.NoError(t, err)
		require.Equal(t, int64(8), n)
		require.Equal(t, [...]byte{0x2d, 0x11, 0x8b, 0xa4, 0x72}, b)
	})

	t.Run("SuccessPadding2Bytes", func(t *testing.T) {
		var b [6]byte
		n, err := runtime.ReadFixedLengthOpaque(
			bytes.NewBuffer([]byte{0x2d, 0x11, 0x8b, 0xa4, 0x72, 0x35, 0x00, 0x00}),
			b[:])
		require.NoError(t, err)
		require.Equal(t, int64(8), n)
		require.Equal(t, [...]byte{0x2d, 0x11, 0x8b, 0xa4, 0x72, 0x35}, b)
	})

	t.Run("SuccessPadding3Bytes", func(t *testing.T) {
		var b [7]byte
		n, err := runtime.ReadFixedLengthOpaque(
			bytes.NewBuffer([]byte{0x2d, 0x11, 0x8b, 0xa4, 0x72, 0x35, 0xaa, 0x00}),
			b[:])
		require.NoError(t, err)
		require.Equal(t, int64(8), n)
		require.Equal(t, [...]byte{0x2d, 0x11, 0x8b, 0xa4, 0x72, 0x35, 0xaa}, b)
	})
}

func TestWriteFixedLengthOpaque(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("IOFailure", func(t *testing.T) {
		w := mock.NewMockWriter(ctrl)
		w.EXPECT().Write(gomock.Any()).Return(0, errors.New("disk failure"))

		n, err := runtime.WriteFixedLengthOpaque(
			w,
			[]byte{1, 2, 3, 4, 5, 6, 7, 8})
		require.Equal(t, errors.New("disk failure"), err)
		require.Equal(t, int64(0), n)
	})

	t.Run("SuccessPadding0Bytes", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteFixedLengthOpaque(
			b,
			[]byte{1, 2, 3, 4})
		require.NoError(t, err)
		require.Equal(t, int64(4), n)
		require.Equal(t, []byte{1, 2, 3, 4}, b.Bytes())
	})

	t.Run("SuccessPadding1Byte", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteFixedLengthOpaque(
			b,
			[]byte{1, 2, 3, 4, 5})
		require.NoError(t, err)
		require.Equal(t, int64(8), n)
		require.Equal(t, []byte{1, 2, 3, 4, 5, 0, 0, 0}, b.Bytes())
	})

	t.Run("SuccessPadding2Bytes", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteFixedLengthOpaque(
			b,
			[]byte{1, 2, 3, 4, 5, 6})
		require.NoError(t, err)
		require.Equal(t, int64(8), n)
		require.Equal(t, []byte{1, 2, 3, 4, 5, 6, 0, 0}, b.Bytes())
	})

	t.Run("SuccessPadding3Bytes", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteFixedLengthOpaque(
			b,
			[]byte{1, 2, 3, 4, 5, 6, 7})
		require.NoError(t, err)
		require.Equal(t, int64(8), n)
		require.Equal(t, []byte{1, 2, 3, 4, 5, 6, 7, 0}, b.Bytes())
	})
}

func TestReadVariableLengthOpaque(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("IOFailure", func(t *testing.T) {
		r := mock.NewMockReader(ctrl)
		r.EXPECT().Read(gomock.Any()).Return(0, errors.New("disk failure"))

		_, n, err := runtime.ReadVariableLengthOpaque(r, 20)
		require.Equal(t, errors.New("disk failure"), err)
		require.Equal(t, int64(0), n)
	})

	t.Run("EOF", func(t *testing.T) {
		_, n, err := runtime.ReadVariableLengthOpaque(bytes.NewBuffer(nil), 20)
		require.Equal(t, io.EOF, err)
		require.Equal(t, int64(0), n)
	})

	t.Run("UnexpectedEOF", func(t *testing.T) {
		input := [...]byte{0, 0, 0, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		for i := 1; i <= len(input); i++ {
			_, n, err := runtime.ReadVariableLengthOpaque(bytes.NewBuffer(input[:i]), 20)
			require.Equal(t, io.ErrUnexpectedEOF, err)
			require.Equal(t, int64(i), n)
		}
	})

	t.Run("TooSmall", func(t *testing.T) {
		_, n, err := runtime.ReadVariableLengthOpaque(
			bytes.NewBuffer([]byte{0, 0, 0, 11, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}),
			10)
		require.Equal(t, errors.New("size of 11 bytes exceeds field's maximum of 10 bytes"), err)
		require.Equal(t, int64(4), n)
	})

	t.Run("InvalidPadding", func(t *testing.T) {
		_, n, err := runtime.ReadVariableLengthOpaque(
			bytes.NewBuffer([]byte{0, 0, 0, 2, 5, 2, 3, 7}),
			20)
		require.Equal(t, errors.New("padding contains non-zero values [3 7]"), err)
		require.Equal(t, int64(8), n)
	})

	t.Run("SuccessPadding0Bytes", func(t *testing.T) {
		b, n, err := runtime.ReadVariableLengthOpaque(
			bytes.NewBuffer([]byte{0, 0, 0, 4, 0x2d, 0x11, 0x8b, 0xa4}),
			20)
		require.NoError(t, err)
		require.Equal(t, int64(8), n)
		require.Equal(t, []byte{0x2d, 0x11, 0x8b, 0xa4}, b)
	})

	t.Run("SuccessPadding1Byte", func(t *testing.T) {
		b, n, err := runtime.ReadVariableLengthOpaque(
			bytes.NewBuffer([]byte{0, 0, 0, 5, 0x2d, 0x11, 0x8b, 0xa4, 0x72, 0x00, 0x00, 0x00}),
			20)
		require.NoError(t, err)
		require.Equal(t, int64(12), n)
		require.Equal(t, []byte{0x2d, 0x11, 0x8b, 0xa4, 0x72}, b)
	})

	t.Run("SuccessPadding2Bytes", func(t *testing.T) {
		b, n, err := runtime.ReadVariableLengthOpaque(
			bytes.NewBuffer([]byte{0, 0, 0, 6, 0x2d, 0x11, 0x8b, 0xa4, 0x72, 0x35, 0x00, 0x00}),
			20)
		require.NoError(t, err)
		require.Equal(t, int64(12), n)
		require.Equal(t, []byte{0x2d, 0x11, 0x8b, 0xa4, 0x72, 0x35}, b)
	})

	t.Run("SuccessPadding3Bytes", func(t *testing.T) {
		b, n, err := runtime.ReadVariableLengthOpaque(
			bytes.NewBuffer([]byte{0, 0, 0, 7, 0x2d, 0x11, 0x8b, 0xa4, 0x72, 0x35, 0xaa, 0x00}),
			20)
		require.NoError(t, err)
		require.Equal(t, int64(12), n)
		require.Equal(t, []byte{0x2d, 0x11, 0x8b, 0xa4, 0x72, 0x35, 0xaa}, b)
	})
}

func TestWriteVariableLengthOpaque(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("IOFailure", func(t *testing.T) {
		w := mock.NewMockWriter(ctrl)
		w.EXPECT().Write(gomock.Any()).Return(0, errors.New("disk failure"))

		n, err := runtime.WriteVariableLengthOpaque(
			w,
			16,
			[]byte{1, 2, 3, 4, 5, 6, 7, 8})
		require.Equal(t, errors.New("disk failure"), err)
		require.Equal(t, int64(0), n)
	})

	t.Run("TooLarge", func(t *testing.T) {
		w := mock.NewMockWriter(ctrl)

		n, err := runtime.WriteVariableLengthOpaque(
			w,
			7,
			[]byte{1, 2, 3, 4, 5, 6, 7, 8})
		require.Equal(t, errors.New("size of 8 bytes exceeds field's maximum of 7 bytes"), err)
		require.Equal(t, int64(0), n)
	})

	t.Run("SuccessPadding0Bytes", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteVariableLengthOpaque(
			b,
			16,
			[]byte{1, 2, 3, 4})
		require.NoError(t, err)
		require.Equal(t, int64(8), n)
		require.Equal(t, []byte{0, 0, 0, 4, 1, 2, 3, 4}, b.Bytes())
	})

	t.Run("SuccessPadding1Byte", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteVariableLengthOpaque(
			b,
			16,
			[]byte{1, 2, 3, 4, 5})
		require.NoError(t, err)
		require.Equal(t, int64(12), n)
		require.Equal(t, []byte{0, 0, 0, 5, 1, 2, 3, 4, 5, 0, 0, 0}, b.Bytes())
	})

	t.Run("SuccessPadding2Bytes", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteVariableLengthOpaque(
			b,
			16,
			[]byte{1, 2, 3, 4, 5, 6})
		require.NoError(t, err)
		require.Equal(t, int64(12), n)
		require.Equal(t, []byte{0, 0, 0, 6, 1, 2, 3, 4, 5, 6, 0, 0}, b.Bytes())
	})

	t.Run("SuccessPadding3Bytes", func(t *testing.T) {
		b := bytes.NewBuffer(nil)
		n, err := runtime.WriteVariableLengthOpaque(
			b,
			16,
			[]byte{1, 2, 3, 4, 5, 6, 7})
		require.NoError(t, err)
		require.Equal(t, int64(12), n)
		require.Equal(t, []byte{0, 0, 0, 7, 1, 2, 3, 4, 5, 6, 7, 0}, b.Bytes())
	})
}
