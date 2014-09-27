package poco

import "testing"
import "github.com/nowk/assert"

func TestPoco(t *testing.T) {
	r, w := Pipe()
	go w.Write([]byte("Hello World!"))

	b := make([]byte, 1024)
	n, err := r.Read(b)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World!", string(b[:n]))
}

func TestReadBuffers(t *testing.T) {
	r, w := Pipe()
	go func() {
		w.Write([]byte("Hello World!"))
		w.Write([]byte("Wat!"))
	}()

	for _, c := range []byte("Hello World!") {
		b := make([]byte, 1)
		n, err := r.Read(b)
		assert.Nil(t, err)
		assert.Equal(t, 1, n)
		assert.Equal(t, string(c), string(b[:n]))
	}

	b := make([]byte, 1024)
	n, err := r.Read(b)
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, "Wat!", string(b[:n]))
}
