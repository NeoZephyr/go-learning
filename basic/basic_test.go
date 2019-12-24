package main

import "testing"

func TestBasic(t *testing.T) {
	t.Log("basic test...")
	t.Log(12, 45)
}

func TestPow(t *testing.T) {
	testData := []struct{
		a, b, c int
	}{
		{2, 3, 8},
		{3, 5, 81},
		{4, 2, 16},
		{5, 6, 125},
		{9, 3, 729},
	}

	for _, data := range testData {
		if actual := pow(data.a, data.b); actual != data.c {
			t.Errorf("test pow(%d, %d) got %d, expected %d\n", data.a, data.b, actual, data.c)
		}
	}
}

func BenchmarkPow(b *testing.B) {

	// b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := pow(12, 12)
		if actual != 8916100448256 {
			b.Errorf("test pow(%d, %d) got %d, expected %d\n", 5, 6, actual, 8916100448256)
		}
	}
}