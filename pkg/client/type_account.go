package client

type AccountResponse struct {
	Account *Account `json:"account"`
}

type Account struct {
	Balance           float64  `json:"balance"`
	PendingCharges    float64  `json:"pending_charges"`
	Name              string   `json:"name"`
	Email             string   `json:"email"`
	Acls              []string `json:"acls"`
	LastPaymentDate   string   `json:"last_payment_date"`
	LastPaymentAmount float64  `json:"last_payment_amount"`
}
