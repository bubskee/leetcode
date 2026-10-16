package math

import (
	"reflect"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{48, 18, 6},
		{18, 48, 6},
		{0, 5, 5},
		{5, 0, 5},
		{7, 7, 7},
	}
	for _, tt := range tests {
		if got := GCD(tt.a, tt.b); got != tt.want {
			t.Errorf("GCD(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{4, 6, 12},
		{21, 6, 42},
		{0, 5, 0},
	}
	for _, tt := range tests {
		if got := LCM(tt.a, tt.b); got != tt.want {
			t.Errorf("LCM(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMobiusSieve(t *testing.T) {
	got := MobiusSieve(10)
	want := []int{0, 1, -1, -1, 0, -1, 1, -1, 0, 0, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MobiusSieve(10) = %v, want %v", got, want)
	}
}

func TestPowTable(t *testing.T) {
	got := PowTable(2, 5, 1_000_000_007)
	want := []int{1, 2, 4, 8, 16, 32}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PowTable(2, 5, mod) = %v, want %v", got, want)
	}
}
