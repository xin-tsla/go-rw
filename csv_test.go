package go_rw_test

import (
	"io"
	"testing"
	go_rw "xin-tsla/go-rw"

	"github.com/stretchr/testify/assert"
)

func TestCSV_ReadOneLine(t *testing.T) {
	c, err := go_rw.NewCSV("./test.csv")
	assert.NoError(t, err)
	defer c.Close()
	firstLine, err := c.ReadOneLine()
	assert.NoError(t, err)
	assert.Len(t, firstLine, 2)
	cnt := 0
	for {
		_, err = c.ReadOneLine()
		if err != nil {
			break
		}
		cnt++
	}
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, 7, cnt)
}
