package tax

import "testing"

func TestCalculateAnnualTax(t *testing.T) {
	calculator := NewCalculator()

	tests := []struct {
		name         string
		annualIncome float64
		expectedTax  float64
	}{
		{name: "no tax", annualIncome: 18000, expectedTax: 0},
		{name: "second bracket", annualIncome: 30000, expectedTax: 2242},
		{name: "top bracket", annualIncome: 200000, expectedTax: 60667},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := calculator.CalculateAnnualTax(tc.annualIncome)
			if int(result.TaxOwed) != int(tc.expectedTax) {
				t.Fatalf("expected tax %v, got %v", tc.expectedTax, result.TaxOwed)
			}
		})
	}
}
