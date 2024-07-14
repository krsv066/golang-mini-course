package dto

type Account struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type ChangeNameParams struct {
	NewName string `json:"name"`
}

type UpdateBalanceParams struct {
	Balance float64 `json:"balance"`
}
