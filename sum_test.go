package vector

import "testing"

func TestSum(t *testing.T) {
	array := [8]int16{1, 1, 2, 3, 5, 8, 13, 21}
	for i := 0; i <= len(array); i++ {
		slice := array[:i]
		want := int16(0)
		for _, v := range slice {
			want += v
		}
		got := Sum(slice)
		if want != got {
			t.Fatalf("Sum: wanted %v, got %v", want, got)
		}
	}
}

func TestSumGeneric(t *testing.T) {
	array := [8]int16{1, 1, 2, 3, 5, 8, 13, 21}
	for i := 0; i <= len(array); i++ {
		slice := array[:i]
		want := int16(0)
		for _, v := range slice {
			want += v
		}
		got := SumGeneric(slice)
		if want != got {
			t.Fatalf("SumGeneric: wanted %v, got %v", want, got)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	slice := make([]int16, 4096)
	for i := 0; i < b.N; i++ {
		Sum(slice)
	}
}

func BenchmarkSumGeneric(b *testing.B) {
	slice := make([]int16, 4096)
	for i := 0; i < b.N; i++ {
		SumGeneric(slice)
	}
}
