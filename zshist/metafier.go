package zshist

import (
	"bufio"
	"io"
)

const (
	null   = byte(0)
	meta   = byte(0x83)
	marker = byte(0xa2)
)

func isMeta(b byte) bool {
	return null == b || meta <= b && b <= marker
}

// Metafier marks and converts meta characters(0x83-0xa2).
type Metafier struct {
	dest *bufio.Writer
}

func NewMetafier(w io.Writer) *Metafier {
	return &Metafier{dest: bufio.NewWriter(w)}
}

func (m *Metafier) Write(p []byte) (n int, err error) {
	for _, b := range p {
		if isMeta(b) {
			if err = m.dest.WriteByte(meta); err != nil {
				return 0, err
			}
			b = b ^ 0x20
		}
		if err = m.dest.WriteByte(b); err != nil {
			return 0, err
		}
	}
	return len(p), nil
}

func (m *Metafier) Close() error {
	return m.dest.Flush()
}

var _ io.WriteCloser = (*Metafier)(nil)
