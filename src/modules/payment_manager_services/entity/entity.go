package entity

import "database/sql"

type HistoricalTransaction struct {
	Id                       string
	AccountPaymentIdPayed    sql.NullString // Using sql.NullString to handle nullable UUIDs
	AccountPaymentIdReceived sql.NullString
	TotalTransaction         float64
	Currency                 string
	StatusTransaction        string
	TypeTransaction          string
}

type AccountPayment struct {
	Id     string
	Name   string
	Type   string
	UserId string
	Amount float64
}

type WithdrawReq struct {
	AccountPaymentId string
	TotalTransaction float64
	Type             string
	Currency         string
}

type SendReq struct {
	AccountPaymentIdPayed    string
	AccountPaymentIdReceived string
	TotalTransaction         float64
	Type                     string
	Currency                 string
}
