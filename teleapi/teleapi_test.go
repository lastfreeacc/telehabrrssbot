package teleapi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeURL(t *testing.T) {
	secretToken := "secretToken"
	bot := NewBot(secretToken)
	mthd := sendMessageMthd
	url := bot.makeURL(mthd)
	assert.Equal(t, "https://api.telegram.org/bot"+secretToken+"/sendMessage", url, fmt.Sprintf("wrong url for: %s", mthd))
}
