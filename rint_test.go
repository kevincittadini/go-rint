package rint

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReversedInt(t *testing.T) {
	e := 987654321
	a := rint(123456789)
	assert.Equal(t, e, a)
}