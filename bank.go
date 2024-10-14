package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64{
	data, _ := os.ReadFile(accountBalanceFile, )
	balanceText := string(data)
  balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644 )
}

func main() {
var accountBalance = getBalanceFromFile()

fmt.Println("Welcome to Go Bank!")

for  {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")


var choice int 
fmt.Print("Your choice ")
fmt.Scan(&choice)

// wantsCheckBalance := choice == 1
if choice == 1 {
	fmt.Println("Your balance is", accountBalance)

} else if choice  == 2 {
	fmt.Print("Your deposit: ")
	var depositAmount float64 
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		fmt.Println("Invalid amount. Must be greater than 0")
		// return 
		continue 

	}
	accountBalance +=  depositAmount
	fmt.Println("Balance updated! New amount:", accountBalance)
	writeBalanceToFile(accountBalance)
} else if choice == 3 {
	fmt.Print("Withdrawal amount: ")
	var withdrawalAmount float64 
	fmt.Scan(&withdrawalAmount)
	if withdrawalAmount <= 0 {
		fmt.Println("Invalid amount. Must be greater than 0.")
continue
	} else if withdrawalAmount > accountBalance {
			fmt.Println("Invalid amount. You can't withdraw more than you have.")
			return
		
	}
	accountBalance -=  withdrawalAmount
	fmt.Println("Balance updated! New amount:", accountBalance)
	writeBalanceToFile(accountBalance)

} else {
fmt.Println("Goodbye!")
// return
break
}


}
fmt.Println("Thanks for choosing our bank")


}





