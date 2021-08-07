package proxy

import (
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"time"
)

type Proxy struct {
	host       string
	url        string
	credential string
	Client     *http.Client
}

func Create(host string, credential string) *Proxy {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	return &Proxy{
		host:       host,
		credential: credential,
		Client:     client,
	}
}

func (i *Proxy) BuildUrl() url2.URL {
	url := url2.URL{
		Scheme: "https",
		Host:   i.host,
	}
	q := url.Query()
	q.Add("apikey", i.credential)
	url.RawQuery = q.Encode()
	return url
}

func (i *Proxy) Query(params map[string]string) *Proxy {
	url := i.BuildUrl()
	q := url.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	url.RawQuery = q.Encode()
	i.url = url.String()
	return i
}

func (i Proxy) Get() (result []byte, err error) {
	resp, err := i.Client.Get(i.url)
	if err != nil {
		return nil, err
	}

	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return
}
