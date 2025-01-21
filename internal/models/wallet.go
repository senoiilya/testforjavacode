package models

type Wallet struct {
	WalletID      int64   `json:"wallet_id"`
	UserID        int64   `json:"user_id"`
	OperationType string  `json:"operationType"`
	Amount        float64 `json:"amount"`
}
