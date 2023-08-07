package entity

import "time"

type Account struct {
	account string
}

func (a *Account) NewFormString(account string) {
	a.account = account
}

func (a *Account) ToString() string {
	return a.account
}

type Password struct {
	password string
}

func (a *Password) NewFormString(password string) {
	a.password = password
}

type Accounts struct {
	Account  Account
	Password Password
	RegTime  time.Time
}

func (a *Accounts) Create(account Account, password Password) {
	a.Account = account
	a.Password = password
}
