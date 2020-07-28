package iyzipay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type restClient struct {
}

func jsonToModel(data []byte, result interface{}) error {
	err := json.Unmarshal(data, result)
	return err
}

func (r *restClient) Init() IClient {
	return r
}

func (r *restClient) Get(url string, headers map[string]string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return jsonToModel(body, result)
}

func (r *restClient) option(method, url string, headers map[string]string, request interface{}, result interface{}) error {
	reqyestBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqyestBody))

	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return jsonToModel(body, result)
}

func (r *restClient) Post(url string, headers map[string]string, request interface{}, result interface{}) error {
	return r.option("POST", url, headers, request, result)
}

func (r *restClient) Put(url string, headers map[string]string, request interface{}, result interface{}) error {
	return r.option("PUT", url, headers, request, result)
}

func (r *restClient) Delete(url string, headers map[string]string, request interface{}, result interface{}) error {
	return r.option("DELETE", url, headers, request, result)
}
