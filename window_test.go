package STWindow

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewArray (t *testing.T) {

	arr := NewArray(time.Hour)

	for i:=0; i<1000; i++{
		arr.Append(i)
	}

	list := arr.List()

	assert.Equal(t, 1000, len(list))
}

func TestMoveOut(t *testing.T) {
	arr := NewArray(2 * time.Second)

	for i:=0; i<10; i++{
		arr.Append(i)
	}

	time.Sleep(3 * time.Second)

	list := arr.List()

	assert.Equal(t, 0, len(list))

	for i:=0; i<10; i++{
		arr.Append(i)
	}

	list = arr.List()

	assert.Equal(t, 10, len(list))
}
