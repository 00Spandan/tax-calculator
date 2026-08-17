package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/00Spandan/financial-tools/services/bff/internal/taxclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	taxClient *taxclient.Client
}

type calculateTaxRequest struct {
	AnnualIncome float64 `json:"annualIncome"`
	TaxYear      string  `json:"taxYear"`
	CountryCode  string  `json:"countryCode"`
}

type calculateTaxResponse struct {
	TaxableIncome    float64 `json:"taxableIncome"`
	TaxOwed          float64 `json:"taxOwed"`
	EffectiveTaxRate float64 `json:"effectiveTaxRate"`
}

func New(taxClient *taxclient.Client) *Server {
	return &Server{taxClient: taxClient}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", s.handleHealth)
	mux.HandleFunc("/api/tax/calculate", s.handleCalculateTax)
	return withCORS(mux)
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) handleCalculateTax(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request calculateTaxRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if request.AnnualIncome < 0 {
		writeError(w, http.StatusBadRequest, "annualIncome must be greater than or equal to 0")
		return
	}

	result, err := s.taxClient.CalculateTax(r.Context(), taxclient.CalculateTaxParams{
		AnnualIncome: request.AnnualIncome,
		TaxYear:      request.TaxYear,
		CountryCode:  request.CountryCode,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.InvalidArgument {
			writeError(w, http.StatusBadRequest, st.Message())
			return
		}
		log.Printf("tax calculator call failed: %v", err)
		writeError(w, http.StatusBadGateway, "failed to calculate tax")
		return
	}

	writeJSON(w, http.StatusOK, calculateTaxResponse{
		TaxableIncome:    result.TaxableIncome,
		TaxOwed:          result.TaxOwed,
		EffectiveTaxRate: result.EffectiveTaxRate,
	})
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeJSON(w, statusCode, map[string]string{"error": message})
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		if !errors.Is(err, http.ErrHandlerTimeout) {
			log.Printf("failed to write JSON response: %v", err)
		}
	}
}
