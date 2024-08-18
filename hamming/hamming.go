package hamming

import (
	"fmt"
	"strings"
)

func Distance(a, b string) (int, error) {
	aSlc := strings.Split(a, "")
	bSlc := strings.Split(b, "")
	hammingDistance := 0

	if len(aSlc) != len(bSlc) {
		return 0, fmt.Errorf("hamming cannot be calculated since the length of sequence is not equal")
	}

	for i := 0; i < len(aSlc); i++ {
		if aSlc[i] != bSlc[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
