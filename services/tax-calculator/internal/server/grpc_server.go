package server

import (
	"context"

	taxv1 "github.com/00Spandan/financial-tools/gen/go/tax"
	"github.com/00Spandan/financial-tools/services/tax-calculator/internal/tax"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaxServer struct {
	taxv1.UnimplementedTaxCalculatorServer
	calculator *tax.Calculator
}

func NewTaxServer(calculator *tax.Calculator) *TaxServer {
	return &TaxServer{calculator: calculator}
}

func (s *TaxServer) CalculateTax(_ context.Context, req *taxv1.CalculateTaxRequest) (*taxv1.CalculateTaxResponse, error) {
	if req.GetAnnualIncome() < 0 {
		return nil, status.Error(codes.InvalidArgument, "annual_income must be greater than or equal to 0")
	}

	result := s.calculator.CalculateAnnualTax(req.GetAnnualIncome())
	effectiveTaxRate := 0.0
	if result.TaxableIncome > 0 {
		effectiveTaxRate = (result.TaxOwed / result.TaxableIncome) * 100
	}

	return &taxv1.CalculateTaxResponse{
		TaxableIncome:    result.TaxableIncome,
		TaxOwed:          result.TaxOwed,
		EffectiveTaxRate: effectiveTaxRate,
	}, nil
}
