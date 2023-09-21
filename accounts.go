package accounts

import "fmt"

var id = uint(10000)

type Account struct {
	id            uint   
	bankName      string
	balance       float64  
	holderName    string
	accountType string
	isActive bool
}

func CreateAccount(bankName,holderName,accountType string,) *Account {

	id++

	return &Account{
		id: id, 
		bankName: bankName,
		balance: 0,
		holderName : holderName,
		accountType : accountType,
		isActive: true,
	}
}

func Withdraw(account *Account, amount float64) bool {
	if account.balance < amount {
	fmt.Println("insufficient funds")
	}

	account.balance -= amount

	return true
}

func Deposit(account *Account, amount float64) {
		account.balance += amount 
}

func (account *Account) GetId() uint{
	return account.id
}