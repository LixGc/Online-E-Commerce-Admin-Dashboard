package entity

type User struct {
	ID                 uint
	FullName           string
	Email              string
	Password           string
	Role               string
	Balance            int
	TransactionHistory []TransactionHistory
}

type Category struct {
	ID                uint
	Type              string
	SoldProductAmount int
	Products          []Product
}
type Product struct {
	ID                 string
	Title              string
	Price              int
	Stock              int
	CategoryID         int
	TransactionHistory []TransactionHistory
}
type TransactionHistory struct {
	ID         string
	ProductID  uint
	UserID     uint
	Quantity   int
	TotalPrice int
	Product    Product
}
