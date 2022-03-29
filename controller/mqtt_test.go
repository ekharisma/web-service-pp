package controller_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLast2Element(t *testing.T) {
	test := []struct {
		name     string
		expected int
	}{{
		name:     "Test Get 2 Last Element",
		expected: 2},
	}
	for range test {
		arr := []float32{1, 2, 3, 4, 5}
		slicedArr := arr[len(arr)-2:]
		assert.Equal(t, len(slicedArr), 2)
	}

}
