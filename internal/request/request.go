package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// Represents a fully parsed HTTP request
type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	dat, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	reqLine, err := parseRequestLine(dat)
	if err != nil {
		return nil, err
	}
	return &Request{
		RequestLine: reqLine,
	}, nil
}

func parseRequestLine(data []byte) (RequestLine, error) {
	strData := string(data)
	reqLine, _, found := strings.Cut(strData, "\r\n")
	if !found {
		return RequestLine{}, fmt.Errorf("malformed data: %s", strData)
	}
	// GET /coffee HTTP/1.1
	parts := strings.Fields(reqLine)
	if len(parts) != 3 {
		return RequestLine{}, fmt.Errorf("malformed request-line: %s", reqLine)
	}
	method := parts[0]
	if strings.ToUpper(method) != method {
		return RequestLine{}, fmt.Errorf("malformed method: %s", method)
	}
	path := parts[1]
	httpVer := parts[2]
	if httpVer != "HTTP/1.1" {
		return RequestLine{}, fmt.Errorf("malformed http version string: %s", httpVer)
	}
	ver, found := strings.CutPrefix(httpVer, "HTTP/")
	if !found {
		return RequestLine{}, errors.New("should not happen")
	}
	return RequestLine{
		HttpVersion:   ver,
		RequestTarget: path,
		Method:        method,
	}, nil
}
