package hellogo_test

import (
	"testing"

	"github.com/hsmtkk/cuddly-potato/hellogo"
	"github.com/stretchr/testify/assert"
)

func TestGetMessage(t *testing.T) {
	want := "Hello,World!"
	got := hellogo.GetMessage()
	assert.Equal(t, want, got)
}
