package request

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestLineParse(t *testing.T) {
	// Test: Good GET Request line
	r, err := RequestFromReader(strings.NewReader("GET / HTTP/1.1\r\n" +
		"Host: localhost:42069\r\n" +
		"User-Agent: curl/7.81.0\r\n" +
		"Accept: */*\r\n\r\n"))
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "GET", r.RequestLine.Method)
	assert.Equal(t, "/", r.RequestLine.RequestTarget)
	assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

	// Test: Good GET Request line with path
	r, err = RequestFromReader(strings.NewReader("GET /coffee HTTP/1.1\r\n" +
		"Host: localhost:42069\r\n" +
		"User-Agent: curl/7.81.0\r\n" +
		"Accept: */*\r\n\r\n"))
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "GET", r.RequestLine.Method)
	assert.Equal(t, "/coffee", r.RequestLine.RequestTarget)
	assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

	// Test: Good POST method with path
	r, err = RequestFromReader(strings.NewReader("POST /coffee HTTP/1.1\r\n" +
		"Host: localhost:42069\r\n" +
		"User-Agent: curl/7.81.0\r\n" +
		"Accept: */*\r\n" +
		"Content-Type: application/json\r\n" +
		"Content-Length: 22\r\n\r\n" +
		`{"flavor":"dark mode"}`))
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "POST", r.RequestLine.Method)
	assert.Equal(t, "/coffee", r.RequestLine.RequestTarget)
	assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

	// Test: Invalid number of parts in request line
	_, err = RequestFromReader(strings.NewReader("/coffee HTTP/1.1\r\n" +
		"Host: localhost:42069\r\n" +
		"User-Agent: curl/7.81.0\r\n" +
		"Accept: */*\r\n\r\n"))
	require.Error(t, err)

	// Test: Invalid method (out of order) Request line
	_, err = RequestFromReader(strings.NewReader("/coffee GET HTTP/1.1\r\n" +
		"Host: localhost:42069\r\n" +
		"User-Agent: curl/7.81.0\r\n" +
		"Accept: */*\r\n\r\n"))
	require.Error(t, err)

	// Test: Invalid HTTP version number
	_, err = RequestFromReader(strings.NewReader("GET /coffee HTTP/4.5\r\n" +
		"Host: localhost:42069\r\n" +
		"User-Agent: curl/7.81.0\r\n" +
		"Accept: */*\r\n\r\n"))
	require.Error(t, err)

	// Test: lowercase method name
	_, err = RequestFromReader(strings.NewReader("get /coffee HTTP/1.1\r\n" +
		"Host: localhost:42069\r\n" +
		"User-Agent: curl/7.81.0\r\n" +
		"Accept: */*\r\n\r\n"))
	require.Error(t, err)

	// Test: invalid HTTP version name
	_, err = RequestFromReader(strings.NewReader("GET /coffee TOTO/1.1\r\n" +
		"Host: localhost:42069\r\n" +
		"User-Agent: curl/7.81.0\r\n" +
		"Accept: */*\r\n\r\n"))
	require.Error(t, err)
}
