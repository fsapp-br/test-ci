package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(3, 5)

	if total != 8 {
		t.Errorf("REsultado da soma é inválido. Resultado %d. Esperado %d", total, 8)
	}
}

func TestAdicao(t *testing.T) {
	total := adicao(1, 1)

	if total != 2 {
		t.Errorf("REsultado da soma é inválido. Resultado %d. Esperado %d", total, 8)
	}
}

func TestSubtracao(t *testing.T) {
	total := subtracao(2, 1)

	if total != 1 {
		t.Errorf("REsultado da soma é inválido. Resultado %d. Esperado %d", total, 8)
	}
}

func TestDivisao(t *testing.T) {
	total := divisao(4, 2)

	if total != 2 {
		t.Errorf("REsultado da soma é inválido. Resultado %d. Esperado %d", total, 8)
	}
}

func TestMultiplicacao(t *testing.T) {
	total := multiplicacao(3, 3)

	if total != 9 {
		t.Errorf("REsultado da soma é inválido. Resultado %d. Esperado %d", total, 8)
	}
}

func TestMultiplicacaoX(t *testing.T) {
	total := multiplicacaoX(3, 3)

	if total != 9 {
		t.Errorf("REsultado da soma é inválido. Resultado %d. Esperado %d", total, 8)
	}
}

func TestMultiplicacaoY(t *testing.T) {
	total := multiplicacaoY(3, 3)

	if total != 9 {
		t.Errorf("REsultado da soma é inválido. Resultado %d. Esperado %d", total, 8)
	}
}
