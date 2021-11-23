package palabra

import (
	"fmt"
	"testing"

	"github.com/carmu14/go_binnarium/Nivel13/Ejercicio2/cita"
)

func TestConteoUso(t *testing.T) {
	m := ConteoUso("uno dos tres tres tres")
	for k, v := range m {
		switch k {
		case "uno":
			if v != 1 {
				t.Error("Esperaba", 1, "Obtuvo", v)
			}
		case "dos":
			if v != 1 {
				t.Error("Esperaba", 1, "Obtuvo", v)
			}
		case "tres":
			if v != 3 {
				t.Error("Esperaba", 3, "Obtuvo", v)
			}
		}
	}
}

func TestConteo(t *testing.T) {
	n := Conteo(cita.Cuando)
	if n != 355 {
		t.Error("Esperana", 355, "palabras", "pero se obtuvo", n)
	}
}

func ExampleConteo() {
	fmt.Println(Conteo("uno dos tres"))
	//Output:
	//3
}

func BenchmarkConteo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conteo(cita.Cuando)
	}
}

func BenchmarkConteoUso(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConteoUso(cita.Cuando)
	}
}
