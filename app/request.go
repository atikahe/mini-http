package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Request struct {
	Headers map[string]string
	Method  string
	Path    string
	Version string
}

func DecodeRequest(reader *bufio.Reader) (*Request, error) {
	req, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("error reading request content: %v", err)
	}

	method, path, version, err := decodeDefinition(req)
	if err != nil {
		return nil, fmt.Errorf("error decoding request content: %v", err)
	}

	headers := make(map[string]string)
	for {
		header, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("failed decoding header: %v\n", err)
		}

		if header[:len(header)-2] == "" {
			break
		}

		key, value, err := decodeHeader(header)
		if err != nil {
			fmt.Printf("failed decoding header: %v\n", err)
		}

		headers[key] = value
	}

	return &Request{
		Headers: headers,
		Method:  method,
		Path:    path,
		Version: version,
	}, nil
}

func decodeDefinition(req string) (method, path, version string, err error) {
	if !strings.HasSuffix(req, "\r\n") {
		err = fmt.Errorf("failed decoding request: error reading EOF")
		return
	}

	req = req[:len(req)-2]

	reqArr := strings.Split(req, " ")
	if len(reqArr) < 3 {
		err = fmt.Errorf("failed decoding request: length invalid")
		return
	}

	if len(strings.Split(reqArr[2], "/")) < 2 {
		err = fmt.Errorf("failed decoding request: need http version")
		return
	}

	method = reqArr[0]
	path = reqArr[1]
	version = strings.Split(reqArr[2], "/")[1]
	return
}

func decodeHeader(header string) (key, value string, err error) {
	if !strings.HasSuffix(header, "\r\n") {
		err = fmt.Errorf("failed decoding header: error reading EOF")
		return
	}

	header = header[:len(header)-2]

	headerArr := strings.Split(header, ": ")
	if len(headerArr) < 2 {
		err = fmt.Errorf("failed decoding header: length invalid")
		return
	}

	key = headerArr[0]
	value = headerArr[1]
	return
}
