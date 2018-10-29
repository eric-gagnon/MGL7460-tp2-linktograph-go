package links

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	x := []string{}

	x = append(x, "https://www.blazegraph.com")
	x = append(x, "http://ontodia.org")
	x = append(x, "https://www.blazegraph.com")

	filteredx := removeDuplicates(x)

	xLen := len(filteredx)
	if xLen != 2 {
		t.Errorf("removeDuplicates failed %v, expected %v\nBefore:\n %v\nAfter:\n %v\n", xLen, 2, x, filteredx)
	}
}
