package basic

import "testing"

func TestAddOne(t *testing.T) {
	var (
		input = 1
		ouput = 3
	)

	actual := AddOne(1)
	if actual != ouput {
		t.Errorf("input %d, ouput %d, actual %d", input, ouput, actual)
	}
}

// run: go test -v
// Thực tế: go test -coverprofile=coverage.out
// Sau đó: go tool cover -html=coverage.out -o coverage.html
// Mở file: open coverage.html
