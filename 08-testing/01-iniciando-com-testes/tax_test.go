package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	t.Run("when amount is greater than 1000", func(t *testing.T) {
		amount := 1001.0
		expected := 10.0

		result := CalculateTax(amount)

		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("when amount is less than or equal to 1000", func(t *testing.T) {
		amount := 1000.0
		expected := 5.0

		result := CalculateTax(amount)

		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}

func TestCalculateTaxTableDriven(t *testing.T) {
	//tests := []struct {
	//	name string
	//	amount, expected float64
	//}{
	//	{"when amount is greater than 1000", 1001.0, 10.0},
	//	{"when amount is less than or equal to 1000", 1000.0, 5.0},
	//}

	type taxesCalc struct {
		name             string
		amount, expected float64
	}

	tests := []taxesCalc{
		{"when amount is greater than 1000", 1001.0, 10.0},
		{"when amount is less than or equal to 1000", 1000.0, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateTax(tt.amount)

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Received %f but expected expected 20", result)
		}

	})
}
