package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(3, 5)

	if total != 8 {
		t.Errorf("REsultado da soma é inválido. Resultado %d. Esperado %d", total, 8)
	}
}
