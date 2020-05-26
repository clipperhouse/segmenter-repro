package segmenter_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/blevesearch/segment"
	"github.com/clipperhouse/uax29/words"
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

func TestCompare(t *testing.T) {
	file, _ := ioutil.ReadFile("wikinyc-short.txt")

	r := bytes.NewReader(file)
	sc := words.NewScanner(r)

	r2 := bytes.NewReader(file)
	segmenter := segment.NewSegmenter(r2)

	for {
		if !sc.Scan() {
			break
		}
		if !segmenter.Segment() {
			break
		}

		if sc.Text() != segmenter.Text() {
			t.Fatalf("segmenter: %q, uax29: %q", segmenter.Text(), sc.Text())
		}
	}
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

func BenchmarkUAX29(b *testing.B) {
	file, err := ioutil.ReadFile("wikinyc-short.txt")

	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()

	count := 0
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(file)
		sc := words.NewScanner(r)

		c := 0
		for sc.Scan() {
			c++
		}
		if err := sc.Err(); err != nil {
			b.Error(err)
		}

		count = c
	}
	b.Logf("%d tokens\n", count)
}

func BenchmarkSegment(b *testing.B) {
	file, err := ioutil.ReadFile("wikinyc-short.txt")

	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()

	count := 0
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(file)
		sc := bufio.NewScanner(r)
		sc.Split(segment.SplitWords)

		c := 0
		for sc.Scan() {
			c++
		}
		if err := sc.Err(); err != nil {
			b.Error(err)
		}

		count = c
	}
	b.Logf("%d tokens\n", count)
}
