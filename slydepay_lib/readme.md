  
Description
================
GoLang library for slydePay payment platform.Works for mobile currently and should be fine for WEB as well.


Prequistes
=================
* Get GOLANG setup on development machine, https://golang.org/dl/

* GET and SETUP GOMOBILE , https://godoc.org/golang.org/x/mobile/cmd/gomobile

* GET ANDROID / IOS SDK Installed

* Package golang code to an android or ios library
  Download source into $GOPATH(basically,your GO workspace)

gomobile bind [-target android|ios] [-o output]  [package]
Example : gomobile bind -target android -o slydepay-lib.arr slydepay_lib

* Methods

- `CreateOrder(merchantEmail string, merchantKey string, orderId string, subTotal float64, shipping float64, tax float64, total float64, comment string, itemCode string, description string) returns string `
Generates an order invoice to slydepay app,which return an invoice code

- `VerifyPayment(merchantEmail string, merchantKey string, orderId string) returns string`

- `ConfirmOrder(merchantEmail string, merchantKey string, token string, transactionId string) returns string `





