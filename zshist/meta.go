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

// Convert func will metafy/unmetafy zsh_history file.
type Convert func(io.Reader, io.Writer) error

// Metafy marks and converts meta characters(0x83-0xa2).
func Metafy(in io.Reader, out io.Writer) error {
	reader := bufio.NewReader(in)
	writer := bufio.NewWriter(out)
	for {
		b, readErr := reader.ReadByte()
		switch readErr {
		case nil:
			// noop
		case io.EOF:
			return nil
		default:
			return readErr
		}

		if isMeta(b) {
			if writeErr := writer.WriteByte(meta); writeErr != nil {
				return writeErr
			}
			b = b ^ 0x20
		}
		if writeErr := writer.WriteByte(b); writeErr != nil {
			return writeErr
		}
	}
}

// Unmetafy trims marking characters and deconverts meta characters(0x83-0xa2).
func Unmetafy(in io.Reader, out io.Writer) error {
	reader := bufio.NewReader(in)
	writer := bufio.NewWriter(out)
	change := false
	for {
		b, readErr := reader.ReadByte()
		switch readErr {
		case nil:
			// noop
		case io.EOF:
			return nil
		default:
			return readErr
		}

		if b == meta {
			change = true
		} else {
			if change {
				b = b ^ 0x20
			}
			if writeErr := writer.WriteByte(b); writeErr != nil {
				return writeErr
			}
			change = false
		}
	}
}
