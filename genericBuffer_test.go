package genericbuffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericBuffer(t *testing.T) {
	b := NewGenericBuffer[int]()

	assert.Equal(t, 0, b.Len())
	assert.EqualValues(t, []int{}, b.Peek(4))

	b.Append([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	assert.Equal(t, 10, b.Len())

	assert.EqualValues(t, []int{0, 1, 2, 3}, b.Peek(4))
	assert.Equal(t, 10, b.Len())

	assert.EqualValues(t, []int{0, 1, 2, 3}, b.Peek(4))
	assert.Equal(t, 10, b.Len())

	assert.EqualValues(t, []int{0, 1, 2, 3}, b.Next(4))
	assert.Equal(t, 6, b.Len())

	assert.EqualValues(t, []int{4, 5, 6, 7}, b.Peek(4))
	assert.Equal(t, 6, b.Len())

	assert.EqualValues(t, []int{4, 5, 6, 7}, b.Next(4))
	assert.Equal(t, 2, b.Len())

	assert.EqualValues(t, []int{8, 9}, b.Peek(4))
	assert.Equal(t, 2, b.Len())

	assert.EqualValues(t, []int{8, 9}, b.Next(4))
	assert.Equal(t, 0, b.Len())

	assert.EqualValues(t, []int{}, b.Peek(4))
	assert.Equal(t, 0, b.Len())

	assert.EqualValues(t, []int{}, b.Next(4))
	assert.Equal(t, 0, b.Len())

	b.Append([]int{10, 11, 12, 13, 14})
	assert.Equal(t, 5, b.Len())
	assert.EqualValues(t, []int{10, 11, 12, 13, 14}, b.Peek(20))

	b.Append([]int{15, 16, 17, 18, 19, 20, 21, 22, 23, 24})
	assert.Equal(t, 15, b.Len())
	assert.EqualValues(t, []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}, b.Peek(20))

	assert.EqualValues(t, []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}, b.Next(20))
	assert.Equal(t, 0, b.Len())

	b.Append([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	b.Skip(5)
	assert.Equal(t, []int{5, 6, 7}, b.Next(3))
	assert.Equal(t, []int{8, 9}, b.Next(3))
}
