package bank

// Bank struct
import (
	"bankingapp/accounts"
)

var id = uint(0)

type Bank struct {
	id uint 
	Name string
	Abbreviation string 
	accounts []*accounts.Account
}


func CreateBank(name, abbreviation string) *Bank {
	id++
 
	return &Bank {
		id: id,
		Name: name,
		Abbreviation: abbreviation,
		accounts: []*accounts.Account{},
	}
	
}

func (b *Bank) getBank(id uint) uint {
	return b.id
}

func (bank *Bank) NewAccount(holderName,accountType string) *accounts.Account {
      newAccount := accounts.CreateAccount(bank.Name,holderName,accountType)

	  bank.accounts = append(bank.accounts, newAccount)
	  return newAccount

	  
}


func (bank *Bank) DepositMoneyIntoAccount(acId uint,amount int){
	for _,i := range bank.accounts{
		if(i.GetId()==acId){
			i.Deposit(amount)
			return
		}
	}
	return
}


func (bank *Bank) WithdrawMoneyFromAccount(acId uint,amount int){
	for _,i := range bank.accounts{
		if(i.GetId()==acId){
			i.Withdraw(amount)
			return
		}
	}
	return
}


func (b *Bank) TransferMoney(srcAcc, desAcc uint, amount int) {
	var srcAccount *accounts.Account
	var desAccount *accounts.Account

	for _, i := range b.accounts {
		if i.GetId() == uint(srcAcc) {
			srcAccount = i
		}
	}

	for _, i := range b.accounts {
		if i.GetId() == uint(desAcc) {
			desAccount = i
		}
	}

	srcAccount.Withdraw(amount)
	desAccount.Deposit(amount)
}






