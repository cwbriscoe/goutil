// Copyright 2023 Christopher Briscoe.  All rights reserved.
package compress

import (
	"bytes"
	"errors"
	"io"
	"sync"

	"github.com/klauspost/compress/gzip"
)

// GzipPool contains the structure for a gzip pool
type GzipPool struct {
	pool sync.Pool
}

// NewGzipPool creates a new gzip pool at the specified compression level
func NewGzipPool(level int) *GzipPool {
	p := &GzipPool{}
	p.pool = sync.Pool{
		New: func() any {
			gzip, _ := gzip.NewWriterLevel(io.Discard, level)
			return gzip
		},
	}
	return p
}

// Compress compresses the supplied []bytes
func (p *GzipPool) Compress(src []byte) ([]byte, error) {
	val := p.pool.Get()
	w, ok := val.(*gzip.Writer)
	if !ok {
		return nil, errors.New("invalid type returned from gzip pool")
	}
	defer p.pool.Put(w)

	dest := &bytes.Buffer{}
	w.Reset(dest)

	return gzipInternal(src, dest, w)
}

// Gzip compresses an []byte given the supplied compression level
func Gzip(src []byte, level int) ([]byte, error) {
	dest := &bytes.Buffer{}
	w, err := gzip.NewWriterLevel(dest, level)
	if err != nil {
		return nil, err
	}

	return gzipInternal(src, dest, w)
}

func gzipInternal(src []byte, dest *bytes.Buffer, w *gzip.Writer) ([]byte, error) {
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
