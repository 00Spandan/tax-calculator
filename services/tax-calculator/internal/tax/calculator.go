package tax

type Bracket struct {
	UpperBound float64
	Rate       float64
}

type Calculator struct {
	brackets []Bracket
}

type Result struct {
	TaxableIncome float64
	TaxOwed       float64
}

func NewCalculator() *Calculator {
	return &Calculator{
		brackets: []Bracket{
			{UpperBound: 18200, Rate: 0.00},
			{UpperBound: 45000, Rate: 0.19},
			{UpperBound: 120000, Rate: 0.325},
			{UpperBound: 180000, Rate: 0.37},
			{UpperBound: 0, Rate: 0.45},
		},
	}
}

func (c *Calculator) CalculateAnnualTax(annualIncome float64) Result {
	if annualIncome <= 0 {
		return Result{}
	}

	taxable := annualIncome
	var tax float64
	lowerBound := 0.0

	for _, bracket := range c.brackets {
		upperBound := bracket.UpperBound
		if upperBound == 0 {
			tax += (taxable - lowerBound) * bracket.Rate
			break
		}

		if taxable <= upperBound {
			tax += (taxable - lowerBound) * bracket.Rate
			break
		}

		tax += (upperBound - lowerBound) * bracket.Rate
		lowerBound = upperBound
	}

	return Result{
		TaxableIncome: taxable,
		TaxOwed:       tax,
	}
}
