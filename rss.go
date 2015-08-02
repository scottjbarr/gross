package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Client struct {
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

func (c *Client) Get(url *string) (*Rss, error) {
	resp, err := http.Get(*url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	rss := new(Rss)
	buf := bytes.NewBuffer(body)
	decoded := xml.NewDecoder(buf)

	err = decoded.Decode(rss)

	if err != nil {
		return nil, err
	}

	return rss, nil
}
