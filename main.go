package main

import (
	"io"
	"log"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/kyoh86/zshist/zshist"
)

func main() {
	app := kingpin.New("zshist", "Encode(metafy) / decode(unmetafy) .zsh_history file")
	enc := app.Command("encode", "Encode(metafy) .zsh_history file").Alias("metafy")
	dec := app.Command("decode", "Decode(unmetafy) .zsh_history file").Alias("unmetafy")

	var i input
	var o output
	defer func() {
		if err := i.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	defer func() {
		if err := o.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	inputFlag(enc, &i)
	inputFlag(dec, &i)
	outputFlag(enc, &o)
	outputFlag(dec, &o)

	command, err := app.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var convert zshist.Convert
	switch command {
	case enc.FullCommand():
		convert = zshist.Metafy
	case dec.FullCommand():
		convert = zshist.Unmetafy
	}

	if err := convert(i.Reader(), o.Writer()); err != nil {
		log.Fatal(err)
	}
}

func inputFlag(c *kingpin.CmdClause, v *input) {
	c.Flag("input", "A file to read. If - is used as file, file will be read from standard input.").Default("-").PlaceHolder("file").Short('i').SetValue(v)
}

type input struct {
	value string
	io.ReadCloser
}

func (i *input) Close() error {
	if c := i.ReadCloser; c != nil {
		return c.Close()
	}
	return nil
}

func (i *input) Reader() io.Reader {
	if c := i.ReadCloser; c != nil {
		return c
	}
	return os.Stdin
}

func (i *input) Set(v string) error {
	i.value = v
	if v == "-" {
		i.ReadCloser = os.Stdin
		return nil
	}
	f, err := os.Open(v)
	if err != nil {
		return err
	}
	i.ReadCloser = f
	return nil
}

func (i *input) String() string {
	return i.value
}

func outputFlag(c *kingpin.CmdClause, v *output) {
	c.Flag("output", "A file to write result. If - is used as file, result will be write to standard output.").Default("-").PlaceHolder("file").Short('o').SetValue(v)
}

type output struct {
	value string
	io.WriteCloser
}

func (o *output) Close() error {
	if c := o.WriteCloser; c != nil {
		return c.Close()
	}
	return nil
}

func (o *output) Writer() io.Writer {
	if c := o.WriteCloser; c != nil {
		return c
	}
	return os.Stdout
}

func (o *output) Set(v string) error {
	o.value = v
	if v == "-" {
		o.WriteCloser = os.Stdout
		return nil
	}
	f, err := os.OpenFile(v, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	o.WriteCloser = f
	return nil
}

func (o *output) String() string {
	return o.value
}
