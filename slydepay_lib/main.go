// slydepay_lib project main.go
package slydepay_lib

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slydepay_lib/client/soap"
	"slydepay_lib/model"
	//	"strconv"
	//	"time"
)

const AppVersion = "1.0.0 beta"

type SlydepayInApp struct {
}

func main() {
	version := flag.Bool("v", false, "prints current roxy version")
	flag.Parse()
	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}
	//	creds := new(model.PayliveCredentials)
	//	creds.SetMerchantEmail("iwallet@dreamoval.com")
	//	creds.SetMerchantKey("bdVI+jtRl80PG4x6NMvYOwfZTZtwfN")

	//	order := new(model.PaymentOrder)
	//	order.SetOrderId(strconv.FormatInt(time.Now().UnixNano(), 16))
	//	order.SetSubTotal(6)
	//	order.SetShipping(0)
	//	order.SetTax(0)
	//	order.SetTotal(6)
	//	order.SetComment("This is a test")

	//	item1 := new(model.OrderItem)
	//	item1.SetItemCode("001")
	//	item1.SetItemName("Item One")
	//	item1.SetQuantity(1)
	//	item1.SetSubTotal(3.5)
	//	item1.SetUnitPrice(3.5)

	//	item2 := new(model.OrderItem)
	//	item2.SetItemCode("002")
	//	item2.SetItemName("Item Two")
	//	item2.SetQuantity(2)
	//	item2.SetSubTotal(2.5)
	//	item2.SetUnitPrice(1.25)

	//	var items = make([]model.OrderItem, 2)
	//	items[0] = *item1
	//	items[1] = *item2

	//	order.SetItems(items)

	//result, success := soap.CreateOrder(*creds, *order)
	//result, success := soap.VerifyPayment(*creds, "2015122201")
	//result, success := soap.ConfirmOrder(*creds, "4603cfba-9ece-4807-9bea-6d555ba23f85", "64c73321-d429-45ac-b037-d93b61f18127")

	//	if !success {
	//		log.Println("Failed")
	//	}

	//	result := CreateOrder("iwallet@dreamoval.com", "bdVI+jtRl80PG4x6NMvYOwfZTZtwfN", strconv.FormatInt(time.Now().UnixNano(), 16), 2.5, 0, 0, 2.5, "Another test", "003", "Third Item", false)
	//	log.Println(result)

	//	result := VerifyPayment("iwallet@dreamoval.com", "bdVI+jtRl80PG4x6NMvYOwfZTZtwfN", "142944ca154a58c4", false)
	//	log.Println(result)

	//	result := CancelOrder("iwallet@dreamoval.com", "bdVI+jtRl80PG4x6NMvYOwfZTZtwfN", "6e53478d-5339-47c4-8712-338f963f72ec", "59bbc1aa-83aa-41fa-a5a9-2f626d7931a8", false)
	//	log.Println(result)

	//Next Steps: Check if response message contains (starts with) 'Error'.
	//If so, return appropriate response to calling app (false?)
	//Otherwise, generation was successful. Call Slydepay mobile app to process

}

func CreateOrder(merchantEmail string, merchantKey string, orderId string, subTotal float64, shipping float64, tax float64, total float64, comment string, itemCode string, description string, isLive bool) (response *APIResult) {
	creds := new(model.PayliveCredentials)
	creds.SetMerchantEmail(merchantEmail)
	creds.SetMerchantKey(merchantKey)

	order := new(model.PaymentOrder)
	order.SetOrderId(orderId)
	order.SetSubTotal(subTotal)
	order.SetShipping(shipping)
	order.SetTax(tax)
	order.SetTotal(total)
	order.SetComment(comment)

	item1 := new(model.OrderItem)
	item1.ItemCode = itemCode
	item1.ItemName = description
	item1.Quantity = 1
	item1.SubTotal = subTotal
	item1.UnitPrice = subTotal

	var items = make([]model.OrderItem, 1)
	items[0] = *item1

	order.SetItems(items)

	success, orderId, token, paycode, message := soap.CreateOrder(*creds, *order, isLive)

	result := new(APIResult)
	result.Success = success
	result.Message = message
	result.OrderId = orderId
	result.PayCode = paycode
	result.Token = token

	//	result, success := soap.CreateOrder(*creds, *order)
	//	if !success {
	//		log.Fatalf("Order generation failed")
	//		return "Error: " + result
	//	}
	return result
}

//func CreateOrderWithItems(merchantEmail string, merchantKey string, orderId string, subTotal float64, shipping float64, tax float64, total float64, comment string, items *[]model.OrderItem) string {
//	creds := new(model.PayliveCredentials)
//	creds.SetMerchantEmail(merchantEmail)
//	creds.SetMerchantKey(merchantKey)

//	order := new(model.PaymentOrder)
//	order.SetOrderId(orderId)
//	order.SetSubTotal(subTotal)
//	order.SetShipping(shipping)
//	order.SetTax(tax)
//	order.SetTotal(total)
//	order.SetComment(comment)

//	order.SetItems(*items)

//	result, success := soap.CreateOrder(*creds, *order)
//	if !success {
//		log.Fatalf("Order generation failed")
//		return "Error: " + result
//	}
//	return result
//}

func VerifyPayment(merchantEmail string, merchantKey string, orderId string, isLive bool) (apiResult *APIResult) {
	creds := new(model.PayliveCredentials)
	creds.SetMerchantEmail(merchantEmail)
	creds.SetMerchantKey(merchantKey)

	response := new(APIResult)
	response.OrderId = orderId

	result, success := soap.VerifyPayment(*creds, orderId, isLive)

	response.Success = success
	if !success {
		log.Fatalf("Error verifying payment: %s", result)
		response.Message = result
		return response
	}
	response.TransactionId = result
	return response
}

func ConfirmOrder(merchantEmail string, merchantKey string, token string, transactionId string, isLive bool) (apiResult *APIResult) {
	creds := new(model.PayliveCredentials)
	creds.SetMerchantEmail(merchantEmail)
	creds.SetMerchantKey(merchantKey)

	response := new(APIResult)
	response.Token = token
	response.TransactionId = transactionId

	result, success := soap.ConfirmOrder(*creds, token, transactionId, isLive)
	response.Success = success

	if !success {
		log.Fatalf("Transaction confirmation failed: %s", result)
		response.Message = result
		return response
	}
	return response
}

func CancelOrder(merchantEmail string, merchantKey string, token string, transactionId string, isLive bool) (apiResult *APIResult) {
	creds := new(model.PayliveCredentials)
	creds.SetMerchantEmail(merchantEmail)
	creds.SetMerchantKey(merchantKey)

	response := new(APIResult)
	response.Token = token
	response.TransactionId = transactionId

	result, success := soap.CancelOrder(*creds, token, transactionId, isLive)
	response.Success = success

	if !success {
		log.Fatalf("Transaction cancellation failed: %s", result)
		response.Message = result
		return response
	}
	return response
}
