// paylive project client/soap/paylive.go
package soap

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"slydepay_lib/model"
	"strconv"
	"strings"

	"gopkg.in/xmlpath.v1"
)

type PayliveClient struct {
}

func CreateOrder(credentials model.PayliveCredentials, order model.PaymentOrder, isLive bool) (success bool, orderId string, token string, code string, payliveUrl string, message string) {
	var envelope = "<soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:pay=\"http://www.i-walletlive.com/payLIVE\">"
	envelope += GenerateHeaderXML(credentials)
	envelope += GenerateBodyXML(order)
	envelope += "</soapenv:Envelope>"

	response, success := CallPaylive(envelope, isLive)

	if !success {
		return false, order.OrderId(), "", "", "", response
	}

	msg, success := ParseSuccess(response)
	if !success {
		msg, success = ParseError(response)
		return false, order.OrderId(), "", "", "", msg
	}

	token, success = ParseToken(response)
	if !success {
		return false, order.OrderId(), "", "", "", token
	}
	//return token, success

	payCode, success := ParsePayCode(response)

	if !success {
		return false, order.OrderId(), "", "", "", payCode
	}

	var url string = "https://test.slydepay.com/webservices/paymentservice.asmx?pay_token=" + token
	if isLive {
		url = "https://app.slydepay.com/webservices/paymentservice.asmx?pay_token=" + token
	}

	return true, order.OrderId(), token, code, url, ""
}

func VerifyPayment(credentials model.PayliveCredentials, orderId string, isLive bool) (result string, success bool) {
	var envelope = "<soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:pay=\"http://www.i-walletlive.com/payLIVE\">"
	envelope += GenerateHeaderXML(credentials)
	envelope += fmt.Sprintf("<soapenv:Body><pay:verifyMobilePayment><pay:orderId>%s</pay:orderId></pay:verifyMobilePayment></soapenv:Body>", orderId)
	envelope += "</soapenv:Envelope>"

	response, success := CallPaylive(envelope, isLive)
	if !success {
		return response, success
	}

	path := xmlpath.MustCompile("//status")
	root, err := xmlpath.Parse(strings.NewReader(response))
	if err != nil {
		log.Fatalf("Error reading response from server: %s", err.Error())
		return "", false
	}

	status, ok := path.String(root)
	if !ok {
		return "Server returned unexpected response", false
	}
	if status == "false" {
		return "Unable to verify order status", false
	}

	path = xmlpath.MustCompile("//transactionId")
	root, err = xmlpath.Parse(strings.NewReader(response))
	if err != nil {
		log.Fatalf("Error reading response from server: %s", err.Error())
		return "", false
	}

	transactionId, ok := path.String(root)
	if !ok {
		return "Server returned unexpected response", false
	}
	if len(transactionId) < 1 {
		return "Payment pending", false
	}
	return transactionId, true
}

func ConfirmOrder(credentials model.PayliveCredentials, token string, transactionId string, isLive bool) (result string, success bool) {
	var envelope = "<soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:pay=\"http://www.i-walletlive.com/payLIVE\">"
	envelope += GenerateHeaderXML(credentials)
	envelope += fmt.Sprintf("<soapenv:Body><pay:ConfirmTransaction><pay:payToken>%s</pay:payToken><pay:transactionId>%s</pay:transactionId></pay:ConfirmTransaction></soapenv:Body>", token, transactionId)
	envelope += "</soapenv:Envelope>"

	response, success := CallPaylive(envelope, isLive)
	if !success {
		return response, success
	}

	path := xmlpath.MustCompile("//ConfirmTransactionResult")
	root, err := xmlpath.Parse(strings.NewReader(response))
	if err != nil {
		log.Fatalf("Error reading response from server: %s", err.Error())
		return "", false
	}

	confirmed, ok := path.String(root)
	if !ok {
		return "Server returned unexpected response", false
	}
	if confirmed == "0" {
		return "Invalid transaction Id", false
	}
	if confirmed == "-1" {
		return "Invalid token", false
	}
	return "Transaction completed successfully", true
}

func CancelOrder(credentials model.PayliveCredentials, token string, transactionId string, isLive bool) (result string, success bool) {
	var envelope = "<soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\" xmlns:pay=\"http://www.i-walletlive.com/payLIVE\">"
	envelope += GenerateHeaderXML(credentials)
	envelope += fmt.Sprintf("<soapenv:Body><pay:CancelTransaction><pay:payToken>%s</pay:payToken><pay:transactionId>%s</pay:transactionId></pay:CancelTransaction></soapenv:Body>", token, transactionId)
	envelope += "</soapenv:Envelope>"

	response, success := CallPaylive(envelope, isLive)
	if !success {
		return response, success
	}

	path := xmlpath.MustCompile("//CancelTransactionResult")
	root, err := xmlpath.Parse(strings.NewReader(response))
	if err != nil {
		log.Fatalf("Error reading response from server: %s", err.Error())
		return "", false
	}

	confirmed, ok := path.String(root)
	if !ok {
		return "Server returned unexpected response", false
	}
	if confirmed == "0" {
		return "Invalid transaction Id", false
	}
	if confirmed == "-1" {
		return "Invalid token", false
	}
	return "Transaction cancelled successfully", true
}

func GenerateHeaderXML(credentials model.PayliveCredentials) (headerXML string) {
	var xml = "<soapenv:Header>"
	xml += "<pay:PaymentHeader>"
	xml += "<pay:APIVersion>1.3</pay:APIVersion>"
	xml += "<pay:MerchantKey>" + credentials.MerchantKey() + "</pay:MerchantKey>"
	xml += "<pay:MerchantEmail>" + credentials.MerchantEmail() + "</pay:MerchantEmail>"
	xml += "<pay:SvcType>C2B</pay:SvcType>"
	xml += "<pay:UseIntMode>0</pay:UseIntMode>"
	xml += "</pay:PaymentHeader>"
	xml += "</soapenv:Header>"

	return xml
}

func GenerateBodyXML(order model.PaymentOrder) (bodyXML string) {
	var body = "<soapenv:Body>"
	body += "<pay:mobilePaymentOrder>"
	body += "<pay:orderId>" + order.OrderId() + "</pay:orderId>"
	body += "<pay:subtotal>" + strconv.FormatFloat(order.SubTotal(), 'f', -1, 32) + "</pay:subtotal>"
	body += "<pay:shippingCost>" + strconv.FormatFloat(order.Shipping(), 'f', -1, 32) + "</pay:shippingCost>"
	body += "<pay:taxAmount>" + strconv.FormatFloat(order.Tax(), 'f', -1, 32) + "</pay:taxAmount>"
	body += "<pay:total>" + strconv.FormatFloat(order.Total(), 'f', -1, 32) + "</pay:total>"
	body += "<pay:comment1>" + order.Comment() + "</pay:comment1>"
	body += "<pay:orderItems>"

	for i := 0; i < len(order.Items()); i++ {
		item := GenerateItemXML(order.Items()[i])
		body += item
	}

	body += "</pay:orderItems>"
	body += "</pay:mobilePaymentOrder>"
	body += "</soapenv:Body>"

	return body
}

func GenerateItemXML(item model.OrderItem) (itemXML string) {
	var orderItem = "<pay:OrderItem>"
	orderItem += "<pay:ItemCode>" + item.ItemCode + "</pay:ItemCode>"
	orderItem += "<pay:ItemName>" + item.ItemName + "</pay:ItemName>"
	orderItem += "<pay:UnitPrice>" + strconv.FormatFloat(item.UnitPrice, 'f', -1, 64) + "</pay:UnitPrice>"
	orderItem += "<pay:Quantity>" + strconv.FormatInt(int64(item.Quantity), 10) + "</pay:Quantity>"
	orderItem += "<pay:SubTotal>" + strconv.FormatFloat(item.SubTotal, 'f', -1, 64) + "</pay:SubTotal>"
	orderItem += "</pay:OrderItem>"

	return orderItem
}

func CallPaylive(envelope string, isLive bool) (result string, success bool) {
	// log.Printf("Calling server with payload: %s", envelope)

	var url string = "https://test.slydepay.com/webservices/paymentservice.asmx"
	if isLive {
		url = "https://app.slydepay.com/webservices/paymentservice.asmx"
	}

	resp, err := http.Post(url, "application/json", strings.NewReader(envelope))

	if err != nil {
		log.Fatalf("Error creating order on Paylive: %s", err.Error())
		result := "Sorry, an error occurred"
		return result, false
	}

	if resp.StatusCode != 200 {
		result := fmt.Sprintln("Server returned HTTP status ", resp.StatusCode)
		return result, false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result = string(body)

	// log.Println("Server returned response: ", result)
	return result, true
}

func ParseSuccess(response string) (result string, success bool) {
	var successful bool = false

	isSuccess := xmlpath.MustCompile("//success")

	root, err := xmlpath.Parse(strings.NewReader(response))
	if err != nil {
		log.Fatalf("Error reading response from server: %s", err.Error())
		return "Unexpected response from server", false
	}
	if value, ok := isSuccess.String(root); ok {
		successful, err = strconv.ParseBool(value)
	}
	if err != nil {
		return "Unexpected response from server", false
	}
	return "Success", successful
}

func ParseToken(response string) (result string, success bool) {
	token := xmlpath.MustCompile("//token")

	root, err := xmlpath.Parse(strings.NewReader(response))
	if err != nil {
		log.Fatalf("Error reading response from server: %s", err.Error())
		return "Unexpected response from server", false
	}
	if value, ok := token.String(root); ok {
		return value, true
	}
	return "Error contacting server", false
}

func ParsePayCode(response string) (result string, success bool) {
	orderCode := xmlpath.MustCompile("//orderCode")

	root, err := xmlpath.Parse(strings.NewReader(response))
	if err != nil {
		log.Fatalf("Error reading response from server: %s", err.Error())
		return "Unexpected response from server", false
	}
	if value, ok := orderCode.String(root); ok {
		return value, true
	}
	return "Error contacting server", false
}

func ParseError(response string) (result string, success bool) {
	error := xmlpath.MustCompile("//error")

	root, err := xmlpath.Parse(strings.NewReader(response))
	if err != nil {
		log.Fatalf("Error reading response from server: %s", err.Error())
		return "Unexpected response from server", false
	}
	if value, ok := error.String(root); ok {
		return value, true
	}
	return "Error contacting server", false
}
