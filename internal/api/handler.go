package api

import (
	"encoding/json"
	"net/http"
	"unicode"

	"github.com/Teddy-Bear-IT/calc-service/internal/calculator"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type CalculateResponse struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !isValidExpression(req.Expression) {
		sendErrorResponse(w, "Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	result, err := calculator.Calc(req.Expression)
	if err != nil {
		if err.Error() == "invalid expression" || err.Error() == "mismatched parentheses" {
			sendErrorResponse(w, "Expression is not valid", http.StatusUnprocessableEntity)
		} else {
			sendErrorResponse(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	sendSuccessResponse(w, result)
}

func isValidExpression(expression string) bool {
	for _, char := range expression {
		if !unicode.IsDigit(char) && !isAllowedSymbol(char) {
			return false
		}
	}
	return true
}

func isAllowedSymbol(char rune) bool {
	allowedSymbols := []rune{'+', '-', '*', '/', '(', ')', '.', ' '}
	for _, symbol := range allowedSymbols {
		if char == symbol {
			return true
		}
	}
	return false
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(CalculateResponse{Error: message})
}

func sendSuccessResponse(w http.ResponseWriter, result float64) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CalculateResponse{Result: result})
}