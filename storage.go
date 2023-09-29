package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// Storage interface to make it agnositc on what we implement
type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)
	// CreatePortfolio(*Account) error
	// DeletePortfolio(int) error
	// UpdatePortfolio(*Account) error
	// GetPortfolios() ([]*Account, error)
	// GetPortfolioById(int) (*Account, error)
}

func CreateDatabaseFile() error {
	path := "./ynap.db"
	if _, err := os.Stat(path); err != nil {
		file, err := os.Create(path)
		if err != nil {
			fmt.Println("could not create database")
			return err
		}
		file.Close()
	}
	return nil
}

type Sqlite3Store struct {
	db *sql.DB
}

func InitSqlite3Store() (*Sqlite3Store, error) {
	db, err := sql.Open("sqlite3", "./ynap.db")
	if err != nil {
		return nil, err
	}
	// Testing if we can ping the database
	if err := db.Ping(); err != nil {
		return nil, err
	}
	// make sure FOREIGN KEYS ARE ON
	_, err = db.Exec(`PRAGMA foreign_keys = ON;`)
	if err != nil {
		return nil, err
	}
	fmt.Println("could open database")
	return &Sqlite3Store{
		db: db,
	}, nil
}

// ==== CREATING / DROPPING TABLES =========
func CreateTables(s *Sqlite3Store) error {
	// LEAVING this here so I can easily comment it back in
	err := DropAllTables(s)
	if err != nil {
		return err
	}
	err = CreateAccountsTable(s)
	if err != nil {
		return err
	}
	err = CreatePortfoliosTable(s)
	if err != nil {
		return err
	}
	err = CreateTransactionsTable(s)
	if err != nil {
		return err
	}
	err = CreateEquitiesTable(s)
	if err != nil {
		return err
	}
	return nil
}

// FUNCTION TO DROP TABLES IF NEEDED
func DropAllTables(s *Sqlite3Store) error {
	query := `DROP TABLE IF EXISTS accounts;`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	query = `DROP TABLE IF EXISTS portfolios;`
	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}
	query = `DROP TABLE IF EXISTS transactions;`
	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}
	query = `DROP TABLE IF EXISTS equities;`
	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func CreateAccountsTable(s *Sqlite3Store) error {
	query := `CREATE TABLE IF NOT EXISTS accounts(
		account_uuid TEXT PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		balance REAL,
		updated_at TEXT,
		created_at TEXT
	);`
	_, err := s.db.Exec(query)
	if err != nil {
		fmt.Println("Could not create accounts table")
		return err
	}
	fmt.Println("Created accounts Table successfully")
	return nil
}

func CreatePortfoliosTable(s *Sqlite3Store) error {
	query := `CREATE TABLE IF NOT EXISTS portfolios(
		portfolio_uuid TEXT PRIMARY KEY,
		name TEXT,
		balance REAL,
		associated_account TEXT,
		updated_at TEXT,
		created_at TEXT,
		FOREIGN KEY(associated_account) REFERENCES accounts(account_uuid));`
	_, err := s.db.Exec(query)
	if err != nil {
		fmt.Println("Could not create portfolio table")
		return err
	}
	fmt.Println("Created portfolios Table successfully")
	return nil
}

func CreateTransactionsTable(s *Sqlite3Store) error {
	// SHould name be actually the UUID of the equitie not the name?
	// If so we have to make one more query to get the name for the display
	// But we should be faster in querying it?
	// Don think it makes a difference tbh.
	query := `CREATE TABLE IF NOT EXISTS transactions(
		transaction_uuid TEXT PRIMARY KEY,
		name TEXT,
		amount REAL,
		associated_portfolio TEXT,
		type TEXT,
		assets REAL,
		date TEXT,
		created_at TEXT,
		FOREIGN KEY(associated_portfolio) REFERENCES portfolios(portfolio_uuid),
		FOREIGN KEY(name) REFERENCES equities(name));`
	_, err := s.db.Exec(query)
	if err != nil {
		fmt.Println("Could not create transactions table")
		return err
	}
	fmt.Println("Created transactions Table successfully")
	return nil
}

func CreateEquitiesTable(s *Sqlite3Store) error {
	query := `CREATE TABLE IF NOT EXISTS equities(
		equity_uuid TEXT,
		name TEXT,
		ticker TEXT,
		price REAL,
		dividends REAL,
		payouts INT,
		industry TEXT,
		sector TEXT,
		market_cap INT,
		last_updated TEXT
);`
	_, err := s.db.Exec(query)
	if err != nil {
		fmt.Println("Could not create equities table")
		return err
	}
	fmt.Println("Created equities Table successfully")
	return nil
}

// ======= INSERTING DATA =========

func (s *Sqlite3Store) CreateAccount(a *Account) error {
	query := `INSERT INTO accounts(
		account_uuid,
		first_name,
		last_name,
		balance,
		updated_at,
		created_at) values ($1, $2, $3, $4, $5, $6);`
	_, err := s.db.Exec(
		query,
		a.AccountUUID.String(),
		a.FirstName,
		a.LastName,
		a.Balance,
		a.UpdatedAt.Format("2006.01.02 15:04:05"),
		a.CreatedAt.Format("2006.01.02 15:04:05"),
	)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Something went wrong with inserting Account: ", a)
		return err
	}
	fmt.Println("Successfully inserted Account into database")
	fmt.Println("Account inserted: ", a)
	return nil
}

func (s *Sqlite3Store) DeleteAccount(id int) error {
	query := `DELETE FROM accounts WHERE id = $1;`
	_, err := s.db.Exec(
		query,
		id)
	if err != nil {
		fmt.Println("Something went wrong with deleting account: ", id)
		return err
	}
	return nil
}

func (s *Sqlite3Store) UpdateAccount(a *Account) error {
	query := `UPDATE accounts
		SET first_name = $1,
			last_name = $2,
			balance = $3, 
			updated_at = $4
		WHERE account_uuid = $5;`
	_, err := s.db.Exec(
		query,
		a.FirstName,
		a.LastName,
		a.Balance,
		a.UpdatedAt.Format("2006.01.02 15:04:05"),
		a.AccountUUID,
	)
	if err != nil {
		fmt.Println("Could not update account: ", a)
		return err
	}
	fmt.Println("Account updated, account now is: ", a)
	return nil
}

func (s *Sqlite3Store) GetAccounts() ([]Account, error) {
	query := `SELECT * FROM accounts;`
	rows, err := s.db.Query(
		query)
	if err != nil {
		fmt.Println("Could not get all accounts")
		return nil, err
	}
	accounts := []Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		fmt.Println("Account: ", *account)
		accounts = append(accounts, *account)
	}
	fmt.Println("This is all accounts: ", accounts)
	return accounts, nil
}

func (s *Sqlite3Store) GetAccountById(id int) (*Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account with id: %d not found", id)
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	scanAccount := new(ScannedAccount)
	// account := new(Account)
	// Probably  have to convert text back to date time objects once I read it. lets see
	err := rows.Scan(
		&scanAccount.AccountUUID,
		&scanAccount.FirstName,
		&scanAccount.LastName,
		&scanAccount.Balance,
		&scanAccount.UpdatedAt,
		&scanAccount.CreatedAt,
	)
	if err != nil {
		fmt.Println("Could not read rows from account table")
		return nil, err 
	}
	uuidConv, err := uuid.Parse(scanAccount.AccountUUID)
	if err != nil {
		fmt.Println("Could not parse UUID from string to UUID")
		return nil , err 
	}
	updatedAtParsed, err := time.Parse("2006.01.02 15:04:05", scanAccount.UpdatedAt)
	if err != nil {
		fmt.Println("Could not parse Updated at from string to DateTime ")
		return nil , err 
	}
	createdAtParsed, err := time.Parse("2006.01.02 15:04:05", scanAccount.CreatedAt)
	if err != nil {
		fmt.Println("Could not parse created at from string to DateTime")
		return nil , err 
	}
	account := Account{
		AccountUUID: uuidConv,
		FirstName:   scanAccount.FirstName,
		LastName:    scanAccount.LastName,
		Balance:     scanAccount.Balance,
		UpdatedAt:   updatedAtParsed,
		CreatedAt:   createdAtParsed,
	}
	if err != nil {
		return nil, err
	}
	fmt.Println("This is account: ", account)
	return &account, nil
}
