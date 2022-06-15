package segmenter_test

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/blevesearch/segment"
	"github.com/clipperhouse/uax29/words"
)

func TestBleveUax29(t *testing.T) {
	file, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		t.Fatal(err)
	}

	allSpace := func(token []byte) bool {
		for _, r := range string(token) {
			if r != ' ' {
				return false
			}
		}

		return true
	}

	// Differences I found:
	// - Bleve splits a run of spaces into separate tokens,
	//   while uax29 returns a single token of multiple spaces
	// - Bleve appears to be Unicode 8.0.0, uax29 is 13.0.0,
	//   seems like a difference on emoji skin tone modifiers

	uax29 := words.NewSegmenter(file)
	var uax29Result [][]byte
	for uax29.Next() {
		token := uax29.Bytes()
		if allSpace(token) {
			// disregard space, see comment above
			continue
		}
		uax29Result = append(uax29Result, token)
	}
	if uax29.Err() != nil {
		t.Fatal(uax29.Err())
	}
	t.Logf("uax29 segmented %d tokens", len(uax29Result))

	bleve := segment.NewSegmenterDirect(file)
	var bleveResult [][]byte
	for bleve.Segment() {
		token := bleve.Bytes()
		if allSpace(token) {
			// disregard space, see comment above
			continue
		}
		bleveResult = append(bleveResult, token)
	}
	if bleve.Err() != nil {
		t.Fatal(bleve.Err())
	}
	t.Logf("bleve segmented %d tokens", len(bleveResult))

	if !reflect.DeepEqual(uax29Result, bleveResult) {
		// ok let's go spelunking
		var longer int
		if len(bleveResult) > len(uax29Result) {
			longer = len(bleveResult)
		}
		if len(bleveResult) < len(uax29Result) {
			longer = len(uax29Result)
		}

		if longer > 0 {
			for i := 0; i < longer; i++ {
				if !bytes.Equal(bleveResult[i], uax29Result[i]) {
					t.Logf("bleve at index %d: %s %v\n", i, bleveResult[i], bleveResult[i])
					t.Logf("uax at index %d: %s %v\n", i, uax29Result[i], uax29Result[i])

					t.Logf("bleve at index %d: %s %v\n", i+1, bleveResult[i+1], bleveResult[i+1])
					t.Logf("uax at index %d: %s %v\n", i+1, uax29Result[i+1], uax29Result[i+1])

					t.Fatal("see differences above. maybe emoji modifier, unicode version?")
				}
			}
		}

		t.Fatal("nope")
	}

	t.Log("confirmed identical results, modulo spaces")
}
