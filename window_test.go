package STWindow

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestNewArray (t *testing.T) {

	arr := NewArray(time.Hour)

	for i:=0; i<1000; i++{
		arr.AppendInt(i)
	}

	list := arr.List()

	log.Println(list)

	assert.Equal(t, 1000, len(list))
}

func TestMoveOut(t *testing.T) {
	arr := NewArray(2 * time.Second)

	for i:=0; i<10; i++{
		arr.AppendInt(i)
	}

	time.Sleep(3 * time.Second)

	list := arr.List()

	assert.Equal(t, 0, len(list))

	for i:=0; i<10; i++{
		arr.AppendInt(i)
	}

	list = arr.List()

	assert.Equal(t, 10, len(list))
}

func TestMemory (t *testing.T) {

	arr := NewArray(time.Second)

	for i:=0; i<10000000; i++{
		arr.AppendInt(i)
		time.Sleep(time.Millisecond)
	}

	list := arr.List()

	time.Sleep(10 * time.Second)

	//list := arr.List()
	//
	//log.Println(list)

	assert.Equal(t, 0, len(list))
}
