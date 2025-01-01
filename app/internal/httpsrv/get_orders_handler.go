package httpsrv

import "net/http"

type Orders struct {
	orders []Order
}

type Order struct {
	id       int
	userId   int
	products []Product
}

type Product struct {
	productID int
	count     int
	price     float64
}

func (a *API) getOrders(w http.ResponseWriter, r *http.Request) {
	return

}
