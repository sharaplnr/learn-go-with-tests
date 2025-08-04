package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("sum 2+2", func(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, expected)
	}
	})

	t.Run("zero plus zero", func(t *testing.T) {
		sum := Add(0, 0)
		expected := 0

		if sum != expected {
			t.Errorf("expected '%d', but got '%d'", expected, sum)
		}
	})
	
}

func ExampleAdd() {
	sum := Add(3, 5)
	fmt.Println(sum)
	// Output: 8
}