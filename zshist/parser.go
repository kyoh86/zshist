package zshist

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// Parser marks and converts meta characters(0x83-0xa2).
type Parser struct {
	scan *bufio.Scanner
	err  error
	next Entry
}

type Entry struct {
	Time    time.Time `json:"time"`
	Seconds int64     `json:"seconds"`
	Command string    `json:"command"`
}

func NewParser(r io.Reader) *Parser {
	return &Parser{scan: bufio.NewScanner(r)}
}

func (p *Parser) scanNextLine() (bool, string) {
	found := p.scan.Scan()
	if !found {
		return false, ""
	}
	if err := p.scan.Err(); err != nil {
		return false, ""
	}
	text := p.scan.Text()
	if strings.HasSuffix(text, "\\") {
		_, next := p.scanNextLine()
		return true, strings.Join([]string{text, next}, "\n")
	}
	return true, text
}

func (p *Parser) Scan() bool {
	found, text := p.scanNextLine()
	if !found {
		return false
	}
	line := strings.TrimPrefix(text, ": ")
	sects := strings.SplitN(line, ";", 2)
	if len(sects) != 2 {
		p.err = fmt.Errorf("invalid parts(%d) in %s: %w", len(sects), text, io.ErrUnexpectedEOF)
		return false
	}
	words := strings.SplitN(sects[0], ":", 2)
	x, err := strconv.ParseInt(words[0], 10, 64)
	if err != nil {
		p.err = fmt.Errorf("invalid time: %w", err)
		return false
	}
	d, err := strconv.ParseInt(words[1], 10, 64)
	if err != nil {
		p.err = fmt.Errorf("invalid seconds: %w", err)
		return false
	}
	p.next.Time = time.Unix(x, 0)
	p.next.Seconds = d
	p.next.Command = sects[1]
	return true
}

func (p *Parser) Entry() Entry {
	return p.next
}

func (p *Parser) Err() error {
	if p.err != nil {
		return p.err
	}
	if err := p.scan.Err(); err != nil {
		return err
	}
	return nil
}
