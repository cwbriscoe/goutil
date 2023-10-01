// Copyright 2023 Christopher Briscoe.  All rights reserved.
package compress

import (
	"bytes"
	"io"
	"sync"

	"github.com/andybalholm/brotli"
)

// BrotliPool contains the structure for a brotli pool
type BrotliPool struct {
	pool sync.Pool
}

// NewBrotliPool creates a new gzip pool at the specified compression level
func NewBrotliPool(level int) *BrotliPool {
	p := &BrotliPool{}
	p.pool = sync.Pool{
		New: func() any {
			brotli := brotli.NewWriterLevel(io.Discard, level)
			return brotli
		},
	}
	return p
}

// Compress compresses the supplied []bytes
func (p *BrotliPool) Compress(src []byte) ([]byte, error) {
	w := p.pool.Get().(*brotli.Writer)
	defer p.pool.Put(w)

	dest := &bytes.Buffer{}
	w.Reset(dest)

	return brotliInternal(src, dest, w)
}

// Brotli compresses an []byte given the supplied compression level
func Brotli(src []byte, level int) ([]byte, error) {
	dest := &bytes.Buffer{}
	w := brotli.NewWriterLevel(dest, level)

	return brotliInternal(src, dest, w)
}

func brotliInternal(src []byte, dest *bytes.Buffer, w *brotli.Writer) ([]byte, error) {
	_, err := w.Write(src)
	if err != nil {
		return nil, err
	}

	err = w.Close()
	if err != nil {
		return nil, err
	}

	return dest.Bytes(), nil
}
