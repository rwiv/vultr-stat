package client

import (
	"fmt"

	"vultr-stat/pkg/lib/date"
)

type AccountResponse struct {
	Account *Account `json:"account"`
}

type Account struct {
	Balance           float64  `json:"balance"`         // 여분 잔액
	PendingCharges    float64  `json:"pending_charges"` // 보류 중인 금액
	Name              string   `json:"name"`            // 유저 이름(본명)
	Email             string   `json:"email"`
	Acls              []string `json:"acls"`
	LastPaymentDate   string   `json:"last_payment_date"` // 최근 결제일, RFC3339 포맷
	LastPaymentAmount float64  `json:"last_payment_amount"`
}

func (a *Account) ToInfo() *AccountInfo {
	lpdTime, err := date.ByRFC3339String(a.LastPaymentDate)
	if err != nil {
		panic(err)
	}
	return &AccountInfo{
		Name:              a.Name,
		Email:             a.Email,
		Balance:           a.Balance,
		PendingCharges:    a.PendingCharges,
		LastPaymentDate:   date.ToPrettyString(lpdTime),
		LastPaymentAmount: a.LastPaymentAmount,
	}
}

type AccountInfo struct {
	Name              string  `json:"name"` // 유저 이름(본명)
	Email             string  `json:"email"`
	Balance           float64 `json:"balance"`           // 여분 잔액
	PendingCharges    float64 `json:"pending_charges"`   // 보류 중인 금액
	LastPaymentDate   string  `json:"last_payment_date"` // 최근 결제일, RFC3339 포맷
	LastPaymentAmount float64 `json:"last_payment_amount"`
}

func (a *AccountInfo) ToPretty() string {
	result := ""
	result += fmt.Sprintf("Name:                   %s\n", a.Name)
	result += fmt.Sprintf("Email:                  %s\n", a.Email)
	result += fmt.Sprintf("Balance:                %f\n", a.Balance)
	result += fmt.Sprintf("Pending Charges         %f\n", a.PendingCharges)
	result += fmt.Sprintf("Last Payment Date:      %s\n", a.LastPaymentDate)
	result += fmt.Sprintf("Last Payment Ammount:   %f\n", a.LastPaymentAmount)
	return result
}

// billing/history
//export interface BillingHistoryResponse {
//	billing_history: BillingHistory[];
//	meta: Meta;
//}
//
//export interface BillingHistory {
//	id: number
//	date: string; // 결제일. RFC3339 포맷
//	type: string; // 결제 타입. invoice, payment...
//	description: string; // 세부 정보
//	amount: number; // 금액, 달러 단위. 음수면 충전되었다는 것.
//	balance: number; // 결제 후 잔액
//}
