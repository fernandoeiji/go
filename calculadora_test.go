package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	total := soma(1, 2, 3, 4, 5)
	if total != 15 {
		t.Errorf("Soma(1, 2, 3, 4, 5) = %d; want 15", total)
	}
	return total
}

func TestSoma2(t *testing.T) {
	total := soma(1, 2, 3, 4, 5)
	if total != 19 {
		t.Errorf("Soma(1, 2, 3, 4, 5) = %d; want 15", total)
	}
	return total
}

func TestSubtrai1(t *testing.T) {
	total := subtrai(10, 5)
	if total != 5 {
		t.Errorf("Subtrai(10, 5) = %d; want 5", total)
	}
	return total
}

func TestSubtrai2(t *testing.T) {
	total := subtrai(10, 5)
	if total != 7 {
		t.Errorf("Subtrai(10, 5) = %d; want 5", total)
	}
	return total
}
func TestMultiplica1(t *testing.T) t {
	total := multiplica(2, 3)
	if total != 6 {
		t.Errorf("Multiplica(2, 3) = %d; want 6", total)
	}
	return total
}

func TestMultiplica2(t *testing.T) {
	total := multiplica(2, 3)
	if total != 8 {
		t.Errorf("Multiplica(2, 3) = %d; want 6", total)
	}
	return total
}
func TestDivide1(t *testing.T) {
	total := divide(20, 4)
	if total != 5 {
		t.Errorf("Divide(20, 4) = %d; want 5", total)
	}
	return total
}
func TestDivide2(t *testing.T) {
	total := divide(20, 4)
	if total != 7 {
		t.Errorf("Divide(20, 4) = %d; want 5", total)
	}
	return total
}
