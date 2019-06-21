package STWindow

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestExample (t *testing.T) {

	arr := NewArray(time.Second)

	for i:=0; i<1000; i++{
		arr.Append(rand.Int() % 100)
	}

	list := arr.List()

	assert.Equal(t, 1000, len(list))
}
