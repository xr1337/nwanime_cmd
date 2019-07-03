package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	block, err := GetWebBlock()
	assert.NotNil(t, block)
	assert.Nil(t, err)
	assert.NotEmpty(t, block)
}

func TestExtract(t *testing.T) {
	data := "<HTML><HEAD></HEAD><BODY><DIV class=\"tes\"/><a>Cool</a></DIV></BODY></HTML>"
	mock := strings.NewReader(data)
	sut := Extract(mock)
	assert.NotContains(t, sut, "HEAD")
	assert.NotContains(t, sut, "BODY")
}

func TestExtractValid(t *testing.T) {
	data := `<HTML><HEAD></HEAD><BODY>
		<DIV class="video-item"/>
			<div class="video-title"><span><a>Cool</a></span></DIV>
			<span class="time">fire</span></DIV>
		</DIV></BODY></HTML>`
	mock := strings.NewReader(data)
	sut := Extract(mock)
	assert.Contains(t, sut, "Cool - fire")
}
