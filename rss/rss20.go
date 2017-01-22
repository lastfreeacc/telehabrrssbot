package rss

import (
	"encoding/xml"
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// ReadRssURL ...
func ReadRssURL(url string) (*Rss, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[Warning] error in req: %s\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	return readRssRdr(resp.Body)
}

func readRssRdr(rdr io.Reader) (*Rss, error) {
	if rdr == nil {
		log.Printf("[Warning] Error Nil argument")
		return nil, errors.New("Nil argument")
	}
	data, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Printf("[Warning] Error in data body: %s\n", err.Error())
		return nil, err
	}
	var feed Rss
	if err := xml.Unmarshal(data, &feed); err != nil {
		log.Printf("[Warning] Error in unmarshal: %s\n", err.Error())
		return nil, err
	}
	return &feed, nil
}

// Rss ...
type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel ...
type Channel struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	Items       []Item        `xml:"item"`
}

// Item ...
type Item struct {
	Title string `xml:"title"`
	GUID  string `xml:"guid"` // isPermaLink="true" mb use struct here like RssGuid {isPermaLink, Guid}
	Link  string `xml:"link"`
}
