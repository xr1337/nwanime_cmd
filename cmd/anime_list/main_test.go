package main

import (
	"fmt"
	"io"
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

func TestExtractValid(t *testing.T) {
	got := Extract(sampleReader())
	want := []Anime{
		Anime{"fire", "abc", "https://somelink"},
	}
	assert.Equal(t, want, got)
}

func TestNicePrint(t *testing.T) {
	items := []Anime{
		Anime{"fire", "abc", "https://somelink"},
		Anime{"fire2", "abc2", "https://somelink2"},
	}
	got := AnimeFormat(items)
	want := fmt.Sprintf("%s%s", items[0].String(), items[1].String())
	assert.Equal(t, want, got)
}

func sampleReader() io.Reader {
	data := `<HTML><HEAD></HEAD><BODY>
		<DIV class="video-item"/>
			<div class="video-title"><span>
				<a title="Carole Tuesday Episode 24" href="https://somelink" class="video-title-left">abc</a>
			</span></DIV>
			<span class="time">fire</span></DIV>
		</DIV></BODY></HTML>`
	return strings.NewReader(data)
}
