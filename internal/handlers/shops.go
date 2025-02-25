package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Jalenarms1/go-eat/internal/db"
	"github.com/gofrs/uuid"
)

func HandleNewShop(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	fmt.Println(string(body))

	var foodShop db.FoodShop
	err = json.Unmarshal(body, &foodShop)
	if err != nil {
		return err
	}

	err = foodShop.Validate()
	if err != nil {
		return err
	}

	uid, _ := uuid.NewV4()

	foodShop.Id = uid.String()

	return nil
}
