package serialization

import (
	"bytes"
	"encoding/json"
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

func SerializeToJson(req http.Request) ([]byte, error) {
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

func WriteToFile(filename string, req http.Request) error {
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
