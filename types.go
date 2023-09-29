package main

import (
	"time"
	"github.com/google/uuid"
)

// SQL QUERY -> SELECT * FROM accounts JOIN portfolios ON portfolios.account_uuid = accounts.uuid
// SQL QUERY -> SELECT * FROM portfolios WHERE portfolios.account_uuid = $s, account_uuid -> THis output is then put into the account as Portfolios?
// SQL QUERY -> SELECT * FROM transactions WHERE transactions.portfolio_uuid = $s, portfolio_uuid -> This then will be put into Portoflio

// And we build it up from the top to bottom -> Frist we lok at the Account get all Portfolios associated with it in the portfolio table then get all transactions associated with those portfolios. Building the portfolios that then get added to the Account
// -> Then account gets returned ?

// ======== GENERAL STRUCT DEFINITIONS ==============

type Account struct {
	AccountUUID uuid.UUID   `json:"account_uuid,omitempty"`
	FirstName   string      `json:"first_name,omitempty"`
	LastName    string      `json:"last_name,omitempty"`
	Balance     float64     `json:"balance,omitempty"`
	Portfolios  []Portfolio `json:"portfolios,omitempty"` // Do we really need this?
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
}

type ScannedAccount struct {
	AccountUUID string      `json:"account_uuid,omitempty"`
	FirstName   string      `json:"first_name,omitempty"`
	LastName    string      `json:"last_name,omitempty"`
	Balance     float64     `json:"balance,omitempty"`
	Portfolios  []Portfolio `json:"portfolios,omitempty"` // Do we really need this?
	UpdatedAt   string      `json:"updated_at,omitempty"`
	CreatedAt   string      `json:"created_at,omitempty"`
}
type Portfolio struct {
	PortfolioUUID     uuid.UUID     `json:"portfolio_uuid,omitempty"`
	Name              string        `json:"name,omitempty"`
	AssociatedAccount uuid.UUID     `json:"associated_account,omitempty"`
	Balance           float64       `json:"balance,omitempty"`
	Ledger            []Transaction `json:"ledger,omitempty"`
	UpdatedAt         time.Time     `json:"updated_at,omitempty"`
	CreatedAt         time.Time     `json:"created_at,omitempty"`
}

type Transaction struct {
	TransactionUUID      uuid.UUID `json:"transaction_uuid,omitempty"`
	Name                 string    `json:"name,omitempty"`                  // Name of the Equity (Stock)
	Amount               float32   `json:"amount,omitempty"`                // How much the transaction was
	AssociatedPortofolio uuid.UUID `json:"associated_portofolio,omitempty"` // Portfolio uuid that the transaction belongs to
	Type                 string    `json:"type,omitempty"`                  // Dividends, Sell, Buy
	Assets               float32   `json:"assets,omitempty"`                // Amount of the equity recieved / sold in this transaction
	Date                 time.Time `json:"date,omitempty"`                  // Time that the transaction took place
	CreatedAt            time.Time `json:"created_at,omitempty"`
}

type Equity struct {
	EquityUUID  uuid.UUID `json:"equity_uuid,omitempty"`
	Name        string    `json:"name,omitempty"`         // Full name of the equity
	Ticker      string    `json:"ticker,omitempty"`       // Ticker of the Equity
	Price       float32   `json:"price,omitempty"`        // Current price at the last updated Time
	Dividends   float32   `json:"dividends,omitempty"`    // last dividends
	Payouts     int8      `json:"payouts,omitempty"`      // how many months the stock pays dividends
	MarketCap   int       `json:"market_cap,omitempty"`   // MarketCap of the Stock in Billions of local currency
	Industry    string    `json:"industry,omitempty"`     // What industry the stock is in
	Sector      string    `json:"sector,omitempty"`       // The sector of the stock
	LastUpdated time.Time `json:"last_updated,omitempty"` // When the ticker was last updated (mainly price field)
}

// ========= CREATING =============

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		AccountUUID: uuid.New(),
		FirstName:   firstName,
		LastName:    lastName,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
}

func NewPortfolio(name string, account_uuid uuid.UUID) *Portfolio {
	return &Portfolio{
		PortfolioUUID:     uuid.New(),
		AssociatedAccount: account_uuid,
		Name:              name,
		UpdatedAt:         time.Now(),
		CreatedAt:         time.Now(),
	}
}

func NewTransaction(name, transaction_type string, amount, assets float32, date time.Time, portfolio_uuid uuid.UUID) *Transaction {
	return &Transaction{
		TransactionUUID:      uuid.New(),
		Name:                 name,
		Amount:               amount,
		AssociatedPortofolio: portfolio_uuid,
		Type:                 transaction_type,
		Assets:               assets,
		Date:                 time.Now(),
		CreatedAt:            time.Now(),
	}
}

func NewEquity(name, ticker, industry, sector string, price, dividends float32, payouts int8) *Equity {
	return &Equity{
		EquityUUID:  uuid.New(),
		Name:        name,
		Ticker:      ticker,
		Price:       price,
		Dividends:   dividends,
		Payouts:     payouts,
		Industry:    industry,
		Sector:      sector,
		LastUpdated: time.Now(),
	}
}
