package tsv

import (
	"bufio"
	"io"
	"reflect"
	"sync"

	"strings"

	"github.com/fatih/structs"
	"github.com/miku/span/xio"
)

// A Decoder reads and decodes TSV rows from an input stream.
type Decoder struct {
	Header    []string        // Column names.
	Separator string          // Field separator.
	r         *xio.SkipReader // The underlying reader.
	once      sync.Once
}

// NewDecoder returns a new decoder with tab as field separator.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: xio.NewSkipReader(bufio.NewReader(r)), Separator: "\t"}
}

// NewDecoderSeparator creates a new decoder with a given separator.
func NewDecoderSeparator(r io.Reader, sep string) *Decoder {
	return &Decoder{r: xio.NewSkipReader(bufio.NewReader(r)), Separator: sep}
}

// readHeader attempts to read the first row and store the column names. If the
// header has been already set manually, the values won't be overwritten.
func (dec *Decoder) readHeader() (err error) {
	dec.once.Do(func() {
		if len(dec.Header) > 0 {
			return
		}
		var line string
		if line, err = dec.r.ReadString('\n'); err != nil {
			return
		}
		dec.Header = strings.Split(line, dec.Separator)
	})
	return
}

// Decode a single entry, reuse csv struct tags.
func (dec *Decoder) Decode(v interface{}) error {
	if err := dec.readHeader(); err != nil {
		return err
	}
	if reflect.TypeOf(v).Elem().Kind() != reflect.Struct {
		return nil
	}
	line, err := dec.r.ReadString('\n')
	if err == io.EOF {
		return io.EOF
	}
	record := strings.Split(line, dec.Separator)

	s := structs.New(v)
	for _, f := range s.Fields() {
		tag := f.Tag("csv")
		if tag == "" || tag == "-" {
			continue
		}
		for i, header := range dec.Header {
			if tag != header {
				continue
			}
			if i >= len(record) {
				break // Record has too few columns.
			}
			if err := f.Set(record[i]); err != nil {
				return err
			}
		}
	}
	return nil
}
