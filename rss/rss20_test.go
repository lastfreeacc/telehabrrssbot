package rss

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_readRssRdr(t *testing.T) {
	_, err := readRssRdr(nil)
	assert.NotNil(t, err, "read from nil reader")

	file, _ := os.Open("../testdata/rss20.rss")
	feed, _ := readRssRdr(file)

	assert.Equal(t, "Хабрахабр / Все публикации", feed.Channel.Title, "wrong title for rss channel")
	assert.Equal(t, 20, len(feed.Channel.Items), "wrong rss items count")
	assert.Equal(t, "https://habrahabr.ru/post/318878/", feed.Channel.Items[0].GUID, "wrong rss item guid")
}
