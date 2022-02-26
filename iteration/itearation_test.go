package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 4)
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("hello", 3)
	fmt.Println(repeated)
	// Output: hellohellohello
}

func TestStringsRepeat(t *testing.T) {
	repeated := strings.Repeat("a", 2)
	expected := "aa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
