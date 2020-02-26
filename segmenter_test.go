package segmenter_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/blevesearch/segment"
)

func TestLong(t *testing.T) {
	b, err := ioutil.ReadFile("wikinyc-long.txt")
	if err != nil {
		t.Error(err)
	}

	r := bytes.NewReader(b)
	segmenter := segment.NewSegmenter(r)

	segmentcount := 0
	for segmenter.Segment() {
		segmentcount++
	}

	if err := segmenter.Err(); err != nil {
		t.Error(err)
	}

	t.Logf("saw %d segments\n", segmentcount)
}

func TestShort(t *testing.T) {
	b, err := ioutil.ReadFile("wikinyc-short.txt")
	if err != nil {
		t.Error(err)
	}

	r := bytes.NewReader(b)
	segmenter := segment.NewSegmenter(r)

	segmentcount := 0
	for segmenter.Segment() {
		segmentcount++
	}

	if err := segmenter.Err(); err != nil {
		t.Error(err)
	}

	t.Logf("saw %d segments\n", segmentcount)
}
