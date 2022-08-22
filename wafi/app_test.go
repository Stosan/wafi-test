package main

import (
	_"reflect"
	"testing"
)

// test function
func TestCreateAccount(t *testing.T) {
    var fname string = "Ade"
    var lname string = "Usman"
    var mobile uint = 8092789878
    var acctnm uint = 10
    CreateAccount(fname,lname,mobile,acctnm)
    
}


// test function
func TestCheckBalance(t *testing.T) {
    var mobile uint = 8092789878
    CheckBalance(mobile)
}


// test function
func TestTransfer(t *testing.T) {
    var mobile uint = 8092789878
    CheckBalance(mobile)
}