// slydepay_lib project slydepay_lib.go
package slydepay_lib

import (
	"log"
	"slydepay_lib/client/soap"
	"slydepay_lib/model"
	"strconv"
	"time"
)

type SlydepayInApp struct {
}

func main() {
	creds := new(model.PayliveCredentials)
	creds.SetMerchantEmail("iwallet@dreamoval.com")
	creds.SetMerchantKey("bdVI+jtRl80PG4x6NMvYOwfZTZtwfN")

	order := new(model.PaymentOrder)
	order.SetOrderId(strconv.FormatInt(time.Now().UnixNano(), 16))
	order.SetSubTotal(6)
	order.SetShipping(0)
	order.SetTax(0)
	order.SetTotal(6)
	order.SetComment("This is a test")

	item1 := new(model.OrderItem)
	item1.SetItemCode("001")
	item1.SetItemName("Item One")
	item1.SetQuantity(1)
	item1.SetSubTotal(3.5)
	item1.SetUnitPrice(3.5)

	item2 := new(model.OrderItem)
	item2.SetItemCode("002")
	item2.SetItemName("Item Two")
	item2.SetQuantity(2)
	item2.SetSubTotal(2.5)
	item2.SetUnitPrice(1.25)

	var items = make([]model.OrderItem, 2)
	items[0] = *item1
	items[1] = *item2

	order.SetItems(items)

	//result, success := soap.CreateOrder(*creds, *order)
	result, success := soap.VerifyPayment(*creds, "2015122201")
	//result, success := soap.ConfirmOrder(*creds, "4603cfba-9ece-4807-9bea-6d555ba23f85", "64c73321-d429-45ac-b037-d93b61f18127")

	if !success {
		log.Println("Failed")
	}
	log.Println(result)

	//Next Steps: Check if response message contains (starts with) 'Error'.
	//If so, return appropriate response to calling app (false?)
	//Otherwise, generation was successful. Call Slydepay mobile app to process
}

func CreateOrder(merchantEmail string, merchantKey string, orderId string, subTotal float64, shipping float64, tax float64, total float64, comment string, itemCode string, description string) string {
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
	item1.SetItemCode(itemCode)
	item1.SetItemName(description)
	item1.SetQuantity(1)
	item1.SetSubTotal(subTotal)
	item1.SetUnitPrice(subTotal)

	var items = make([]model.OrderItem, 1)
	items[0] = *item1

	result, success := soap.CreateOrder(*creds, *order)
	if !success {
		log.Fatalf("Order generation failed")
	}
	return result
}

func VerifyPayment(merchantEmail string, merchantKey string, orderId string) string {
	creds := new(model.PayliveCredentials)
	creds.SetMerchantEmail(merchantEmail)
	creds.SetMerchantKey(merchantKey)

	result, success := soap.VerifyPayment(*creds, orderId)
	if !success {
		log.Fatalf("Error verifying payment")
	}
	return result
}

func ConfirmOrder(merchantEmail string, merchantKey string, token string, transactionId string) string {
	creds := new(model.PayliveCredentials)
	creds.SetMerchantEmail(merchantEmail)
	creds.SetMerchantKey(merchantKey)

	result, success := soap.ConfirmOrder(*creds, token, transactionId)

	if !success {
		log.Fatalf("Transaction confirmation failed")
	}
	return result
}
