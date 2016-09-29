package model

type LineItem struct {
	ItemCode  string
	ItemName  string
	Quantity  int16
	UnitPrice float64
	SubTotal  float64
}
