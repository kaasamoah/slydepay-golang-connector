//slydepay_lib project apiresult.go
package slydepay_lib

type APIResult struct {
	//Flag indicating whether API call was successful or failed
	Success bool
	//Unique identifier of order defined by merchant
	OrderId string
	//Unique identifier used by client to reference order when making payment
	Token string
	//Short Code used by client to make payment via mobile
	PayCode string
	//Unique transaction identifier
	TransactionId string
	//Additional response returned by API. Usually contains error details
	Message string
	//URL customer can redirect to to complete payments online
	PayliveUrl string
}
