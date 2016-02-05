package serialization

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Host             string
	URL              string
	Method           string //(defaults to "GET")
	Header           map[string][]string
	ContentLength    int64
	TransferEncoding string
	Body             []byte
}

func (req Request) String() string {
	return fmt.Sprintf("%s %s %s %v %d %s", req.Host, req.URL, req.Method, req.Header, req.ContentLength, req.TransferEncoding)
}

func ToRequest(req *http.Request) (Request, error) {
	if req == nil {
		return Request{}, fmt.Errorf("req arg to ToRequest is nil")
	}
	bodyBuf := new(bytes.Buffer)
	bodyBuf.ReadFrom(req.Body)
	return Request{
		Host:          req.Host,
		URL:           req.URL.String(),
		Method:        req.Method,
		Header:        req.Header,
		ContentLength: req.ContentLength,
		Body:          bodyBuf.Bytes(),
	}, nil
}

func SerializeToJson(req *http.Request) ([]byte, error) {
	if req == nil {
		return nil, fmt.Errorf("req arg to SerializeToJson is nil")
	}
	bodyBuf := new(bytes.Buffer)
	bodyBuf.ReadFrom(req.Body)
	reqSerial := Request{
		Host:          req.Host,
		URL:           req.URL.String(),
		Method:        req.Method,
		Header:        req.Header,
		ContentLength: req.ContentLength,
		Body:          bodyBuf.Bytes(),
	}

	reqBytes, err := json.Marshal(reqSerial)
	if err != nil {
		return nil, err
	}
	return reqBytes, nil
}

func WriteToFile(filename string, req *http.Request) error {
	if req == nil {
		return fmt.Errorf("req arg to WriteToFile is nil")
	}
	reqBytes, err := SerializeToJson(req)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, reqBytes, 0777)
	if err != nil {
		return err
	}

	return nil
}
