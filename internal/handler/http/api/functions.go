package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"gophermart/pkg/middleware"
	"net/http"
)

func getDataFromRequest(req *http.Request, requestData interface{}) error {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)

	if err != nil {
		return errors.New("")
	}

	if err = json.Unmarshal(buf.Bytes(), &requestData); err != nil {
		return errors.New("")
	}

	return nil
}

func GetUserIDFromContext(ctx context.Context) (int64, error) {
	value := ctx.Value(middleware.UserIDKey)

	if value != nil {
		userID, ok := value.(int64)
		if ok {
			return userID, nil
		}
	}

	return 0, errors.New("userID not found")
}
