package accrualrepo

import (
	"encoding/json"
	"gophermart/internal/entity"
	"io"
	"net/http"
)

func (r *Repository) GetOrderAccrual(orderID string) (entity.Accrual, error) {
	var accrual entity.Accrual

	request, errNewRequest := http.NewRequest("GET", r.baseURL+"/api/orders/"+orderID, nil)

	if errNewRequest != nil {
		return accrual, errNewRequest
	}
	resp, errSendRequest := r.client.Do(request)

	if errSendRequest != nil {
		return accrual, errSendRequest
	}

	accrual.Code = resp.StatusCode

	if resp.StatusCode == http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		_ = json.Unmarshal(respBody, &accrual.Item)
		resp.Body.Close()
	}

	return accrual, nil
}
