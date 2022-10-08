package buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingWrite(t *testing.T) {
	buf := NewRingBuffer[byte](10)

	buf.Write([]byte{1, 2, 3, 4})

	assert.Equal(t, 4, buf.Readable())
}

func TestRingRead(t *testing.T) {
	buf := NewRingBuffer[byte](10)
	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, buf.Readable())

	readedData := buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+1), readedData[i])
	}

	assert.Equal(t, 0, buf.Readable())
}

func TestRingOverwrite(t *testing.T) {
	buf := NewRingBuffer[byte](5)
	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, buf.Readable())
	assert.Equal(t, 1, buf.Writable())

	buf.Write([]byte{5})
	assert.Equal(t, 5, buf.Readable())
	assert.Equal(t, 0, buf.Writable())

	writed := buf.Write([]byte{6})
	assert.Equal(t, 0, writed)
	assert.Equal(t, 5, buf.Readable())
	assert.Equal(t, 0, buf.Writable())

	readedData := buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+1), readedData[i])
	}
	assert.Equal(t, 1, buf.Readable())
	assert.Equal(t, 4, buf.Writable())

	writed = buf.Write([]byte{6, 7, 8})
	assert.Equal(t, 3, writed)
	assert.Equal(t, 3, buf.writePt)
	assert.Equal(t, 4, buf.Readable())
	assert.Equal(t, 1, buf.Writable())

	writed = buf.Write([]byte{6, 7, 8})
	assert.Equal(t, 1, writed)
	assert.Equal(t, 5, buf.Readable())
	assert.Equal(t, 0, buf.Writable())

	readedData = buf.Read(4)
	assert.Equal(t, 4, len(readedData))
	assert.Equal(t, 1, buf.Readable())
	assert.Equal(t, 4, buf.Writable())
}
