package mark_test

import (
	"testing"

	"github.com/raaaaaaaay86/go-persistence-extension/mark"
	"github.com/stretchr/testify/assert"
)

func TestMarks(t *testing.T) {
	assert.False(t, mark.TargetTime.IsZero())
}