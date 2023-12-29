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

//export interface Account {
//	balance: number; // 여분 잔액
//	pending_charges: number; // 보류 중인 금액
//	name: string; // 유저 이름(본명)
//	email: string;
//	acls: string[]; // 접근 가능 권한들. billing, firewall...
//	last_payment_date: string; // 최근 결제일, ISO 포맷
//	last_payment_amount: number;
//}

// billing/history
//export interface BillingHistoryResponse {
//	billing_history: BillingHistory[];
//	meta: Meta;
//}
//
//export interface BillingHistory {
//	id: number
//	date: string; // 결제일. ISO 포맷
//	type: string; // 결제 타입. invoice, payment...
//	description: string; // 세부 정보
//	amount: number; // 금액, 달러 단위. 음수면 충전되었다는 것.
//	balance: number; // 결제 후 잔액
//}
