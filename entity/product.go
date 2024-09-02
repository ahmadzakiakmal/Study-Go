package entity

type Product struct {
	Id    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	status := ""
	if p.Stock < 10 {
		status = "Stock is low"
	}
	if p.Stock < 3 {
		status = "Stock is nearly out, order now!"
	}
	return status
}
