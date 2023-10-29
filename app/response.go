package main

type StatusCode int

const (
	StatusOK       StatusCode = 200
	StatusNotFound StatusCode = 400
)

var mapStatusString = map[StatusCode]string{
	StatusOK:       "OK",
	StatusNotFound: "Not Found",
}

func (s StatusCode) String() string {
	if _, ok := mapStatusString[s]; !ok {
		return ""
	}

	return mapStatusString[s]
}
