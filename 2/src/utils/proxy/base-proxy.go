package proxy

import (
	"net/http"
	url2 "net/url"
	"time"
)

type Proxy struct {
	url        string
	credential string
	client     *http.Client
}

func Create(url string, credential string) Proxy {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	return Proxy{
		url:        url,
		credential: credential,
		client:     client,
	}
}

func (i *Proxy) BuildUrl() url2.URL {
	url := url2.URL{
		Scheme: "https",
		Host:   i.url,
	}
	q := url.Query()
	q.Add("apikey", i.credential)
	url.RawQuery = q.Encode()
	return url
}
