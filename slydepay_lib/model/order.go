// order
package model

type PaymentOrder struct {
	orderId  string
	subTotal float64
	shipping float64
	tax      float64
	total    float64
	comment  string
	items    []OrderItem
}

func (o PaymentOrder) OrderId() string {
	return o.orderId
}

func (o *PaymentOrder) SetOrderId(orderId string) {
	o.orderId = orderId
}

func (o PaymentOrder) SubTotal() float64 {
	return o.subTotal
}

func (o *PaymentOrder) SetSubTotal(subTotal float64) {
	o.subTotal = subTotal
}

func (o PaymentOrder) Shipping() float64 {
	return o.shipping
}

func (o *PaymentOrder) SetShipping(shipping float64) {
	o.shipping = shipping
}

func (o PaymentOrder) Tax() float64 {
	return o.tax
}

func (o *PaymentOrder) SetTax(tax float64) {
	o.tax = tax
}

func (o PaymentOrder) Total() float64 {
	return o.total
}

func (o *PaymentOrder) SetTotal(total float64) {
	o.total = total
}

func (o PaymentOrder) Comment() string {
	return o.comment
}

func (o *PaymentOrder) SetComment(comment string) {
	o.comment = comment
}

func (o PaymentOrder) Items() []OrderItem {
	return o.items
}

func (o *PaymentOrder) SetItems(items []OrderItem) {
	o.items = items
}
