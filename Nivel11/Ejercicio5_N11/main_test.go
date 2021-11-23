package main_test

import "testing"

func TestSumar(t *testing.T) { // Recibir struct de tipo testing.T
	resultado := sumar(5, 2)
	resultadoEsperado := 7
	if resultado != resultadoEsperado {
		t.Errorf("Error en %s. Se esperaba %d pero el resultado fue %d", t.Name(), resultadoEsperado, resultado)
	}
	// Y si no, o sea, si todo va bien, no hacemos nada
}
