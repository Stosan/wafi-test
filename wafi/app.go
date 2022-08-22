package main

import (
	"fmt"
	"wafi-test/utils"

)

const orgname string = "wafi cash" // declare organisation name
type UserData map[string]interface{} //declare map interface 
var NewUser = []UserData{} //declare array

//declare operation variables
var firstName string
var lastName string
var mobile uint
var accountbalance uint
var amounttoTransfer uint

func main(){
	/**
	entry function
	**/
	
	WelcomeUser() //welcome message
	
	for{
		/**
		intiate a while loop to keep console active
		*/
		i := wafiMenu() //initialize menu options
		switch i {
			case 1:
				for i := 0; i <= 1; i++{
					fmt.Println("Enter your first name: ")
					fmt.Scan(&firstName)
					fmt.Println("Enter your last name: ")
					fmt.Scan(&lastName)
					fmt.Println("Enter your mobile number: ")
					fmt.Scan(&mobile)
					fmt.Println("Make an initial deposit (minimum of $10): ")
					fmt.Scan(&accountbalance)
					CreateAccount(firstName, lastName,mobile,accountbalance)
					if i == 0{
						fmt.Println("*******************************")
						fmt.Println("Create another account *****")
						fmt.Println("*******************************")
					}
				}
			case 2:
				fmt.Println("Enter your mobile number: ")
				fmt.Scan(&mobile)
				CheckBalance(mobile);
			case 3:
				fmt.Println("Enter your mobile number: ")
				fmt.Scan(&mobile)
				TransferAmount(mobile)
		}
	}

}



func WelcomeUser(){
	/*
	This function welcomes users to the app
	*/
	fmt.Println("********** Welcome to **********")
	fmt.Printf("         %v \n",orgname)
	fmt.Println("*******************************")

}

func CreateAccount(firstName string, lastName string, mobile uint, accountbalance uint){
	
		/*
		This function creates a new user acount using field parameters
		*/
		
		// assign input data as maps to the interface
		isValidName := utils.ValidateUserInput(firstName,lastName)
		if isValidName{
			userData:=UserData{"firstName":firstName,"lastName":lastName,"mobile":mobile,"accountbalance":accountbalance}
			//append each map to the array
			NewUser = append(NewUser, userData)
			fmt.Printf("Thank you for creating an account, your account details are \n first name: %v \n last name: %v \n mobile number: %v \n account balance %v \n", firstName,lastName,mobile,accountbalance)
			
		}else{
			fmt.Println("Please enter a real name")
		}
}

func wafiMenu()(uint){
	/**
	this function provides the wafi user menu and returns a section type of unsigned integers
	*/
	fmt.Println( "\n **********MENU**********")
	fmt.Println( "Select an option")
	fmt.Println( "1. Create an account")
	fmt.Println( "2. Check balance")
	fmt.Println( "3. Transfer")
	fmt.Println( "************************")
	var userselection uint
	fmt.Scan(&userselection)
	return userselection
}

func CheckBalance(mobile uint){
	/**
	this function allows the users check their balance using the mobile number
	*/
	
	for _, user := range  NewUser{
		if mobile == user["mobile"]{
			fmt.Printf("your account balance is: $%v",user["accountbalance"])
		}
	}
	
	
}

func TransferAmount(mobile uint){
	/**
	this function allows the users transfer money out of their account using the mobile number
	*/
	
	for _, user := range  NewUser{
		if mobile == user["mobile"]{
			useraccountbalance := utils.InterfaceToUint( user["accountbalance"])
			if useraccountbalance != 0{
				otherUser := NewUser[1]
				otherUserName:= otherUser["firstName"]
				fmt.Printf( "enter an amount to transfer to %v\n",otherUserName)
				fmt.Scan(&amounttoTransfer)
				if amounttoTransfer < useraccountbalance{
					user["accountbalance"] = useraccountbalance - amounttoTransfer
					useraccountbalance := utils.InterfaceToUint( otherUser["accountbalance"]) + amounttoTransfer
					otherUser["accountbalance"] = useraccountbalance
					fmt.Println("Transfer successful!")
					fmt.Printf("your balance %v",user["accountbalance"])
				}
				
			}
			
		}
	}
	

}
