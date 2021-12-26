package client

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestCutToLastMessage(t *testing.T) {
	assert := assert2.New(t)

	res := []byte("100\n101\n10")

	wantTruncated, wantRest := []byte("100\n101\n"), []byte("10")
	gotTruncated, gotRest, err := cutToLastMessage(res)

	assert.NoError(err)
	assert.Equal(string(wantTruncated), string(gotTruncated))
	assert.Equal(string(wantRest), string(gotRest))
}

func TestCutToLastMessageErrors(t *testing.T) {
	assert := assert2.New(t)

	res := []byte("100000")
	_, _, err := cutToLastMessage(res)

	assert.Error(err)
}
