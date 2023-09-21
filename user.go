package user

import (
	"bankingapp/accounts"
	"fmt"
)

// type User struct {
// 	 CustomerId  int    
// 	 FirstName   string  
// 	 LastName    string
// 	 TotalBalance float64
// 	 Accounts    []accounts.Account  
// 	 IsAdmin     bool 
// }

// func createAdmin() *User {
// 	return &User{
// 		IsAdmin: true,
// 	}
// }

// func (c *User) AddCustomer(user []User, user User) []User {
// 	users = append(users, user)
// 	return users
// }

// func (c *Customer) UpdateCustomer(customers []Customer, customer Customer) []Customer {
// 	for i, cust := range customers {
// 		if cust.CustomerId == customer.CustomerId {
// 			customers[i] = customer 
// 			break
// 		}    
// 	}    

// 	return customers  
// }

// func (c *Customer) DeleteCustomer(customers []Customer, id int) []Customer {
// 	for i, cust := range customers {
// 		if cust.CustomerId == id {
// 			customers = append(customers[:i], customers[i+1:]...)
// 			break
// 		}    
// 	}   
// 	return customers
// }

// func (c *Customer) AddBank(banks []bank.Bank, bank bank.Bank) []bank.Bank {
// 	banks = append(banks, bank)
// 	return banks
// }

// func (c *Customer) UpdateBank(banks []bank.Bank, bank bank.Bank) []bank.Bank {
// 	for i, b := range banks {
// 		if b.BankId == bank.BankId {
// 			banks[i] = bank  
// 			break
// 		}    
// 	}
// 	return banks    
// }

// func (c *Customer) DeleteBank(banks []bank.Bank, id int) []bank.Bank {
// 	for i, b := range banks {
// 		if b.BankId == id {
// 			banks = append(banks[:i], banks[i+1:]...)
// 			break
// 		}    
// 	}
// 	return banks
// }

var users []*User
var admin *User

type User struct {
	userId    int
	firstName string
	lastName  string
	isAdmin   bool
	isActive  bool
	account  []*accounts.Account
}

// User factory
func (u *User) newUser(firstName, lastName string) *User {
	newUserId := len(users) + 1
	var newUser = &User{
		firstName: firstName,
		lastName:  lastName,
		userId:    newUserId,
		isAdmin:   false,
		isActive:  true,
		account:  []*accounts.Account{},
	}
	return newUser
}

func newAdmin(firstName, lastName string) *User {
	var newAdmin = &User{
		firstName: firstName,
		lastName:  lastName,
		userId:    -1,
		isAdmin:   true,
		isActive:  true,
		account:  []*accounts.Account{},
	}

	return newAdmin
}

func GetAdmin() *User {
	if admin != nil {
		return admin
	}
	return nil
}

func (u *User) CreateUser(firstName,lastName string) {
	if !u.isAdmin {
		return
	}

	var newUser *User = u.newUser(firstName, lastName)

	users = append(users, newUser)

	fmt.Println("User created successfully")

	
	}

func CreateAdmin(firstName,lastName string) {
	if admin != nil {
		fmt.Println("Admin exists")
		return
	}
	var newAdmin *User = newAdmin(firstName, lastName)
	admin = newAdmin
	fmt.Println("Admin created successfully")
}

func (u *User) PrintUser() {
	fmt.Println("-----------------")
	fmt.Println("User Id: ", u.userId)
	fmt.Println("First Name: ", u.firstName)
	fmt.Println("Last Name: ", u.lastName)
	fmt.Println("Is Admin: ", u.isAdmin)
	fmt.Println("Is Active: ", u.isActive)
	// fmt.Println("Contacts: ", u.contacts)
	fmt.Println("-----------------")
}

func (u *User) ReadAllUsers() {
	if !u.isAdmin {
		return
	}
	for i := 0; i < len(users); i++ {
		if users[i].isActive {
			users[i].PrintUser()
		}
	}
}

func (u *User) ReadUserById(id int) (*User, string) {
	if !u.isAdmin && !u.isActive {
		return nil, "Not an Admin"
	}
	for i := 0; i < len(users); i++ {
		if users[i].userId == id && users[i].isActive {
			return users[i], "User Found"
		}
	}
	return nil, "UserId does not exist"
}

// Get user by id for Staff
func GetUserById(uId int) (*User, string) {
	for i := 0; i < len(users); i++ {
		if users[i].userId == uId && users[i].isActive {
			return users[i], "Found"
		}
	}
	return nil, "UserId does not exist"
}

// Update functions
func (u *User) updateFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) updateLastName(lastName string) {
	u.lastName = lastName
}


// Delete User
func (u *User) DeleteUserById(userId int) string {
	if !u.isAdmin && !u.isActive {
		return "Not an admin/ Not Active Admin"
	}
	user, msgString := u.ReadUserById(userId)
	if user == nil {
		return msgString
	}

	user.isActive = false
	return "Deleted successfully"
}





