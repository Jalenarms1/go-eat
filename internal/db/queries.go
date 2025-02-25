package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func (f *FoodShop) Insert() error {
	_, err := db.Exec("insert into FoodShop (Id, Slug, Name, Bio, FoodType, Logo, Address, HasDelivery) values (?,?,?,?,?,?,?,?)", f.Id, f.Slug, f.Name, f.Bio, f.FoodType, f.Logo, f.Address, f.HasDelivery)
	if err != nil {
		return err
	}

	if os.Getenv("IS_DEV") == "true" {
		fmt.Print("adding to host")
		cmd := exec.Command("sh", "-c", fmt.Sprintf("echo '%s' | sudo -S sh -c 'echo 127.0.0.1 %s.localhost >> /etc/hosts'", os.Getenv("SH_PASS"), f.Slug))
		err := cmd.Run()
		if err != nil {
			fmt.Println("error adding subdomain to hosts")
		}
	}

	return nil
}

func GetFoodShop(slug string) (*FoodShop, error) {
	var foodShop FoodShop
	row := db.QueryRow("select Id, Slug, Name, Bio, FoodType, Logo, Address, HasDelivery from FoodShop where Slug = ?", slug)
	if row == nil {
		return nil, errors.New("food shop not found")
	}

	fmt.Println(row == nil)

	err := row.Scan(&foodShop.Id, &foodShop.Slug, &foodShop.Name, &foodShop.Bio, &foodShop.FoodType, &foodShop.Logo, &foodShop.Address, &foodShop.HasDelivery)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
	}

	return &foodShop, nil
}
