package reader

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

type Rss struct {
	Title string `xml:"channel>title"`
	Items []Item `xml:"channel>item"`
}

func Read(url string) (*Rss, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	rss := &Rss{}
	err = xml.Unmarshal(b, &rss)

	if err != nil {
		return nil, err
	}

	return rss, nil
}
