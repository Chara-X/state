package state

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	Client  *http.Client
	Address string
}

func (c *Client) Get(key string) (value []byte, ok bool, etag string) {
	var resp, err = c.Client.Get(c.Address + "/state/get?key=" + url.QueryEscape(key))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		value, _ = io.ReadAll(resp.Body)
		return value, true, resp.Header.Get("ETag")
	}
	return nil, false, ""
}
func (c *Client) Set(key string, value []byte, duration time.Duration, etag string) bool {
	var req, _ = http.NewRequest(http.MethodPost, c.Address+"/state/set?key="+url.QueryEscape(key)+"&duration="+duration.String(), bytes.NewReader(value))
	req.Header.Set("If-Match", etag)
	var resp, err = c.Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
