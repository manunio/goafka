package client

import (
	"bytes"
	"testing"
)

func TestCutToLastMessage(t *testing.T) {
	// assert := assert2.New(t)
	// assert.NoError(err)
	// assert.Equal(wantTruncated, gotTruncated)
	// assert.Equal(wantRest, gotRest)

	res := []byte("100\n101\n10")
	wantTruncated, wantRest := []byte("100\n101\n"), []byte("10")
	gotTruncated, gotRest, err := cutToLastMessage(res)
	if err != nil {
		t.Errorf("cutToLastMessage(%q): got error %v; want no errors", res, err)
	}
	if !bytes.Equal(wantTruncated, gotTruncated) || !bytes.Equal(wantRest, gotRest) {
		t.Errorf("cutToLastMessage(%q): got %q, %q; want %q %q ", res, gotTruncated, gotRest, wantTruncated, wantRest)
	}
}

func TestCutToLastMessageErrors(t *testing.T) {
	res := []byte("100000")
	_, _, err := cutToLastMessage(res)
	if err != nil {
		t.Errorf("cutToLastMessage(%q): got no errors, want an error", res, err)
	}
}
