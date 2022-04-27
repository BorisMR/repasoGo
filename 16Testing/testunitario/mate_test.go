package testunitario

import "testing"

func TestSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 1, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)

		if total != item.c {
			t.Errorf("ERROR -> %d + %d no es %d", item.a, item.b, item.c)
		}
	}
}
