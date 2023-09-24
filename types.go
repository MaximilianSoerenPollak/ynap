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
	AccountUUID uuid.UUID
	FirstName   string
	LastName    string
	Balance     float64
	Portfolios  []Portfolio // Do we really need this?
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

type ScannedAccount struct {
	AccountUUID string
	FirstName   string
	LastName    string
	Balance     float64
	Portfolios  []Portfolio // Do we really need this?
	UpdatedAt   string 
	CreatedAt   string 
}
type Portfolio struct {
	PortfolioUUID     uuid.UUID
	Name              string
	AssociatedAccount uuid.UUID
	Balance           float64
	Ledger            []Transaction
	UpdatedAt         time.Time
	CreatedAt         time.Time
}

type Transaction struct {
	TransactionUUID      uuid.UUID
	Name                 string    // Name of the Equity (Stock)
	Amount               float32   // How much the transaction was
	AssociatedPortofolio uuid.UUID // Portfolio uuid that the transaction belongs to
	Type                 string    // Dividends, Sell, Buy
	Assets               float32   // Amount of the equity recieved / sold in this transaction
	Date                 time.Time    // Time that the transaction took place
	CreatedAt            time.Time
}

type Equity struct {
	EquityUUID  uuid.UUID
	Name        string    // Full name of the equity
	Ticker      string    // Ticker of the Equity
	Price       float32   // Current price at the last updated Time
	Dividends   float32   // last dividends
	Payouts     int8      // how many months the stock pays dividends
	MarketCap   int       // MarketCap of the Stock in Billions of local currency
	Industry    string    // What industry the stock is in
	Sector      string    // The sector of the stock
	LastUpdated time.Time // When the ticker was last updated (mainly price field)
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
