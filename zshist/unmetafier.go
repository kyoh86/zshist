package zshist

import (
	"bufio"
	"io"
)

// Unmetafier marks and converts meta characters(0x83-0xa2).
type Unmetafier struct {
	dest   *bufio.Writer
	change bool
}

func NewUnmetafier(w io.Writer) *Unmetafier {
	return &Unmetafier{dest: bufio.NewWriter(w)}
}

func (u *Unmetafier) Write(p []byte) (n int, err error) {
	for _, b := range p {
		if b == meta {
			u.change = true
		} else {
			if u.change {
				b = b ^ 0x20
			}
			if err := u.dest.WriteByte(b); err != nil {
				return 0, err
			}
			u.change = false
		}
	}
	return len(p), nil
}

func (u *Unmetafier) Close() error {
	return u.dest.Flush()
}

var _ io.WriteCloser = (*Unmetafier)(nil)
