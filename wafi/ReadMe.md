# peer to peer payment app Challenge
For this problem we want to write a small piece of the backend for a peer to peer payment app. Please note that we want an in memory solution and that we do not need to implement an actual API. The functionality that we want for this problem are: adding users to the app, users depositing money into the app, users sending money to other app users, users checking their balance in the app, and users transferring their money out of the app. Please also write unit tests for your solution that ensures all functionality as described above are working properly.

## Project Description
Example Run through of the app:
User A is added to the app
User A deposits 10 dollars
User B is added to the app
User B deposits 20 dollars
User B sends 15 dollars to User A
User A checks their balance and has 25 dollars
User B checks their balance and has 5 dollars
User A transfers 25 dollars from their account
User A checks their balance and has 0 dollars

Time Limit: 45 minutes

### Documentation:

1. To start the app
```
go run .
```
#### Features
##### Create Account

This feature requires two account to demonstrate the app and takes in the following fields:

Field definition
 - 'firstname' user first name or initials
 - 'last name' user last name or surname
 - 'mobile number' user mobile number without the leading 0 e.g. 08093111111 should be 8093111111
 - 'initial depost' user must make an initial deposit above $10
 if successful the response should be:
 ```
 
 ```
 
 ##### Check Account Balance

This feature requires an account owner to check their balance in accounts using in the following field:

Field definition
 - 'mobile number' user mobile number without the leading 0 e.g. 08093111111 should be 8093111111
 if successful the response should be:
 ```
 
 ```
 
  ##### Check Account Balance

This feature requires an account owner to transfer from their balance to another account using in the following field:

Field definition
 - 'mobile number' user mobile number without the leading 0 e.g. 08093111111 should be 8093111111
 if successful the response should be:
 ```
 
 ```
 
2. To run test

```
go test .
```
