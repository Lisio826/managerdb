package utils
// https://www.cnblogs.com/thinkeridea/p/10324806.html

import (
	"bytes"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/thinkeridea/go-extend/exbytes"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

type Adapter struct {
	pool sync.Pool
}

func New() *Adapter {
	return &Adapter{
		pool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 2048))
			},
		},
	}
}

func (api *Adapter) GetReq(r *http.Request) (*http.Request, error) {
	buffer := api.pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer func() {
		if buffer != nil {
			api.pool.Put(buffer)
			buffer = nil
		}
	}()

	_, err := io.Copy(buffer, r.Body)
	if err != nil {
		return nil, err
	}

	request := &http.Request{}
	if err = jsoniter.Unmarshal(buffer.Bytes(), request); err != nil {
		logrus.WithFields(logrus.Fields{
			"json": exbytes.ToString(buffer.Bytes()),
		}).Errorf("jsoniter.UnmarshalJSON fail. error:%v", err)
		return nil, err
	}
	api.pool.Put(buffer)
	buffer = nil

	// ....

	return request, nil
}

func (api *Adapter) GetRequest(r *http.Request) (*http.Response, error) {
	var err error
	buffer := api.pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer func() {
		if buffer != nil {
			api.pool.Put(buffer)
			buffer = nil
		}
	}()

	e := jsoniter.NewEncoder(buffer)
	err = e.Encode(r)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"request": r,
		}).Errorf("jsoniter.Marshal failure: %v", err)
		return nil, fmt.Errorf("jsoniter.Marshal failure: %v", err)
	}

	data := buffer.Bytes()
	req, err := http.NewRequest("POST", "http://xxx.com", buffer)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"data": exbytes.ToString(data),
		}).Errorf("http.NewRequest failed: %v", err)
		return nil, fmt.Errorf("http.NewRequest failed: %v", err)
	}

	req.Header.Set("User-Agent", "xxx")

	httpResponse, err := http.DefaultClient.Do(req)
	if httpResponse != nil {
		defer func() {
			io.Copy(ioutil.Discard, httpResponse.Body)
			httpResponse.Body.Close()
		}()
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url": "http://xxx.com",
		}).Errorf("query service failed %v", err)
		return nil, fmt.Errorf("query service failed %v", err)
	}

	if httpResponse.StatusCode != 200 {
		logrus.WithFields(logrus.Fields{
			"url":         "http://xxx.com",
			"status":      httpResponse.Status,
			"status_code": httpResponse.StatusCode,
		}).Errorf("invalid http status code")
		return nil, fmt.Errorf("invalid http status code")
	}

	buffer.Reset()
	_, err = io.Copy(buffer, httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("adapter io.copy failure error:%v", err)
	}

	respData := buffer.Bytes()
	logrus.WithFields(logrus.Fields{
		"response_json": exbytes.ToString(respData),
	}).Debug("response json")

	res := &http.Response{}
	err = jsoniter.Unmarshal(respData, res)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"data": exbytes.ToString(respData),
			"url":  "http://xxx.com",
		}).Errorf("adapter jsoniter.Unmarshal failed, error:%v", err)
		return nil, fmt.Errorf("adapter jsoniter.Unmarshal failed, error:%v", err)
	}

	api.pool.Put(buffer)
	buffer = nil

	// ...
	return res, nil
}
