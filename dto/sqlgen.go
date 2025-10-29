package dto

type SQLGenResponseDTO struct {
	SQL   string `json:"sql"`
	Error string `json:"error,omitempty"`
}
