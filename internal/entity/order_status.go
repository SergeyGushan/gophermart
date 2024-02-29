package entity

type OrderStatus string

const (
	OrderStatusInvalid    OrderStatus = "INVALID"
	OrderStatusNew        OrderStatus = "NEW"
	OrderStatusError      OrderStatus = "ERROR"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusProcessed  OrderStatus = "PROCESSED"
)
