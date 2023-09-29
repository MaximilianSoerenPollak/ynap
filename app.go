package main

import (
	"context"
	"fmt"
	"log"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}
// startup is called when the app starts. The context is saved
// so we can call the runtime methods

// Getting a DB conection making it not so repetitive with error handling
func getDBConnection() *Sqlite3Store {
	s, err := InitSqlite3Store()
	if err != nil {
		log.Fatal(err)
	}
	return s
}


func (a *App) startup(ctx context.Context) {
	// Creating Database file
	err := CreateDatabaseFile()
	if err != nil {
		log.Fatal(err)
	}
	s := getDBConnection()
	err = CreateTables(s)
	if err != nil {
		log.Fatal(err)
	}
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) MakeAccount(firstName, lastName string) (string, error) {
	s := getDBConnection()
	account := NewAccount(firstName, lastName)
	err := s.CreateAccount(account)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("YOU just made an account with %s, %s", firstName, lastName), nil
}

func (a *App) GetAllAccounts() ([]Account, error) {
	s := getDBConnection()
	accounts, err := s.GetAccounts()
	if err != nil {
		fmt.Print(err)
		return []Account{}, err
	}
	fmt.Println("Here is all accounts: ", accounts)
	return accounts, nil
}
