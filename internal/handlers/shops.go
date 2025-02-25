package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Jalenarms1/go-eat/internal/db"
	"github.com/gofrs/uuid"
)

func HandleNewShop(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return nil

	}

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

	fmt.Println(foodShop)

	_, err = db.GetFoodShop(foodShop.Slug)
	if err == nil {
		return errors.New("a food shop with this slug already exists, please use a unique url slug")
	}

	err = foodShop.Insert()
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}
