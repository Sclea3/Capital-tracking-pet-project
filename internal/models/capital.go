package models

import "time"

type CapitalData struct {
	UUID           string    `database:"uuid"`
	InstrumentName string    `database:"instrument"`
	Timestamp      time.Time `database:"timestamp"`
	BidPrice       float64   `database:"bid_price"`
	AskPrice       float64   `database:"ask_price"`
	LastPrice      float64   `database:"last_price"`
	Volume         float64   `database:"volume"`
}

type AccountInfoResponse struct {
	AccountType           string    `json:"account_type"`
	AccountInfo           Balance   `json:"accountInfo"`
	CurrencyIsoCode       string    `json:"currencyIsoCode"`
	CurrencySymbol        string    `json:"currencySymbol"`
	CurrentAccountID      string    `json:"currentAccountId"`
	StreamingHost         string    `json:"streamingHost"`
	Accounts              []Account `json:"accounts"`
	ClientID              string    `json:"clientId"`
	TimezoneOffset        int       `json:"timezoneOffset"`
	HasActiveDemoAccounts bool      `json:"hasActiveDemoAccounts"`
	HasActiveLiveAccounts bool      `json:"hasActiveLiveAccounts"`
	TrailingStopsEnabled  bool      `json:"trailingStopsEnabled"`
}

type Balance struct {
	Balance    float64 `json:"balance"`
	Deposit    float64 `json:"deposit"`
	ProfitLoss float64 `json:"profitLoss"`
	Available  float64 `json:"available"`
}

type Account struct {
	AccountID   string  `json:"accountId"`
	AccountName string  `json:"accountName"`
	Preferred   bool    `json:"preferred"`
	AccountType string  `json:"accountType"`
	Currency    string  `json:"currency"`
	Symbol      string  `json:"symbol"`
	Balance     Balance `json:"balance"`
}
