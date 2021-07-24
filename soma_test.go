package iteracao

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {

	t.Run("coleção de 5 numeros", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}
		resultado := Soma(numeros)
		esperado := 15

		if resultado != esperado {
			t.Errorf("esperado '%d' mas obteve '%v',dado '%v'", esperado, resultado, numeros)
		}
	})

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}
		resultado := Soma(numeros)
		esperado := 6

		if resultado != esperado {
			t.Errorf("esperado '%d' mas obteve '%v',dado '%v'", esperado, resultado, numeros)
		}
	})
}

func TestSomaTudo(t *testing.T) {
	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}
	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado '%v',esperado '%v'", resultado, esperado)
	}
}

func TestSomaTodoOResto(t *testing.T) {

	verificaSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v, esperado %v", resultado, esperado)
		}
	}

	t.Run("faz as somas de alguns slices", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}

		verificaSomas(t, resultado, esperado)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}

		verificaSomas(t, resultado, esperado)
	})

}
