package httpsrv

import "net/http"

type productStatus string

const (
	created    productStatus = "новый"
	processing productStatus = "в обработке"
	cancelled  productStatus = "отменён"
	delivered  productStatus = "доставлен"
)

func (a *API) getStatus(w http.ResponseWriter, r *http.Request) {

	return
}
