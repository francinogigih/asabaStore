package request

type Product struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Total       int64  `json:"total"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}
