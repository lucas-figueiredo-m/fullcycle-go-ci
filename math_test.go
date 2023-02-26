package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("\nResultado da soma é inválido:\nResultado %d.\nEsperado: %d", total, 30)
	}
}
