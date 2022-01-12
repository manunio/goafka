package client

import (
	"bytes"
	"errors"
)

var errBufTooSmall = errors.New("buffer too small to fit a single message")

const defaultScratchSize = 64 * 1024

// Simple represents an instance of client connected to a set of goafka servers.
type Simple struct {
	addrs   []string
	buf     bytes.Buffer
	restBuf bytes.Buffer
}

// NewSimple creates a new client for goafka server.
func NewSimple(addrs []string) *Simple {
	return &Simple{
		addrs: addrs,
	}
}

// Send sends the message to goafka servers.
func (s *Simple) Send(msgs []byte) error {
	_, err := s.buf.Write(msgs)
	return err
}

// Receive will either wait for new messages or return,
// an error in-case something gets wrong.
// The scratch buffer can be used to read the data.
func (s *Simple) Receive(scratch []byte) ([]byte, error) {
	if scratch == nil {
		scratch = make([]byte, defaultScratchSize)
	}

	off := 0

	if s.restBuf.Len() > 0 {
		if s.restBuf.Len() >= len(scratch) {
			return nil, errBufTooSmall
		}

		n, err := s.restBuf.Read(scratch)
		if err != nil {
			return nil, err
		}

		s.restBuf.Reset()
		off += n
	}

	n, err := s.buf.Read(scratch[off:])
	if err != nil {
		return nil, err
	}

	// example: truncated = "100\n101\n", rest = "10"
	truncated, rest, err := cutToLastMessage(scratch[0 : n+off])
	if err != nil {
		return nil, err
	}

	s.restBuf.Reset()
	s.restBuf.Write(rest)

	return truncated, nil
}

func cutToLastMessage(res []byte) (truncated []byte, rest []byte, err error) {
	n := len(res)

	if n == 0 {
		return res, nil, nil
	}

	if res[n-1] == '\n' {
		return res, nil, nil
	}

	lastpos := bytes.LastIndexByte(res, '\n')
	if lastpos < 0 {
		return res, nil, errBufTooSmall
	}
	return res[0 : lastpos+1], res[lastpos+1:], nil
}
