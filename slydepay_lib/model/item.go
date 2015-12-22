// item
package model

type OrderItem struct {
	itemCode  string
	itemName  string
	quantity  int16
	unitPrice float64
	subTotal  float64
}

func (i OrderItem) ItemCode() string {
	return i.itemCode
}

func (i *OrderItem) SetItemCode(itemCode string) {
	i.itemCode = itemCode
}

func (i OrderItem) ItemName() string {
	return i.itemName
}

func (i *OrderItem) SetItemName(itemName string) {
	i.itemName = itemName
}

func (i OrderItem) Quantity() int16 {
	return i.quantity
}

func (i *OrderItem) SetQuantity(quantity int16) {
	i.quantity = quantity
}

func (i OrderItem) UnitPrice() float64 {
	return i.unitPrice
}

func (i *OrderItem) SetUnitPrice(unitPrice float64) {
	i.unitPrice = unitPrice
}

func (i OrderItem) SubTotal() float64 {
	return i.subTotal
}

func (i *OrderItem) SetSubTotal(subTotal float64) {
	i.subTotal = subTotal
}
