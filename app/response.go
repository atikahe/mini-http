package main

import (
	"fmt"
	"strings"
)

// type ResponseHeader struct {
// 	ContentType string
// }

const (
	ContentType   string = "Content-Type"
	ContentLength string = "Content-Length"
	TextPlain     string = "text/plain"
)

type Response struct {
	Version string
	Status  StatusCode
	Headers map[string]string
	Body    interface{}
}

func BuildResponse(req *Request) (*Response, error) {
	if req.Path == "/" {
		return &Response{
			Version: req.Version,
			Status:  StatusOK,
		}, nil
	}

	pathArr := strings.Split(req.Path, "/")
	if len(pathArr) < 3 {
		return nil, fmt.Errorf("error parsing path: invalid length")
	}

	path := pathArr[1]
	if path != "echo" {
		return &Response{
			Version: req.Version,
			Status:  StatusNotFound,
		}, nil
	}

	content := req.Path[5:]

	headers := map[string]string{
		ContentType:   TextPlain,
		ContentLength: fmt.Sprintf("%d", len(content)),
	}

	return &Response{
		Version: req.Version,
		Status:  StatusOK,
		Headers: headers,
		Body:    content,
	}, nil
}

func EncodeResponse(res *Response) string {
	spec := fmt.Sprintf("HTTP/%s %d %s\r\n", res.Version, res.Status, res.Status)
	headers := ""

	for key, value := range res.Headers {
		headers += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	body := fmt.Sprintf("\r\n%v\r\n", res.Body)

	return fmt.Sprintf("%s%s%s", spec, headers, body)
}
