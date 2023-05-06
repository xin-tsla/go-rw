package go_rw

import (
	"encoding/csv"
	"os"
)

type CSV struct {
	path string
	f    *os.File
	r    *csv.Reader
}

func NewCSV(path string) (*CSV, error) {
	c := &CSV{path: path}
	// open file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	c.f = f
	c.r = csv.NewReader(f)
	return c, nil
}

func (c *CSV) Close() error {
	err := c.f.Close()
	if err != nil {
		return err
	}
	return nil
}

// ReadOneLine Returns column(s)'s value(s) nil err for each line, return EOF error or other IO error otherwise.
func (c *CSV) ReadOneLine() ([]string, error) {
	return c.r.Read()
}
