package utils

import (
	"fmt"
	"strconv"
)


func ValidateUserInput(firstName string,lastName string)(bool){
	//validate user input
	isValidName := len(firstName) >=2 && len(lastName) >=2
	
	return isValidName
}

func InterfaceToUint( usersDataDetail interface{})(uint){
	//convert interface type to string and to uint for arithmetic operations
	userDataDetailString := fmt.Sprint(usersDataDetail)
	userDataDetailToUint64, _ :=strconv.ParseUint(userDataDetailString,10,64)
	userDataDetail := uint(userDataDetailToUint64)
	return userDataDetail
}