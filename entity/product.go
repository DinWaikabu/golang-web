package entity

type Product struct {
	ID    int
	Nama  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	var status string
	if p.Stock > 2 {
		status = "Stock hampir habis"
	} else if p.Stock < 10 {
		status = "Stok sedikit tersedia"
	} else {
		status = "Tersedia"
	}
	return status
}