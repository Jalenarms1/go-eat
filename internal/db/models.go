package db

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

type FoodType string

var (
	FoodAmerican      FoodType = "American"
	FoodMexican       FoodType = "Mexican"
	FoodItalian       FoodType = "Italian"
	FoodChinese       FoodType = "Chinese"
	FoodJapanese      FoodType = "Japanese"
	FoodIndian        FoodType = "Indian"
	FoodMediterranean FoodType = "Mediterranean"
	FoodThai          FoodType = "Thai"
	FoodKorean        FoodType = "Korean"
	FoodFrench        FoodType = "French"
	FoodVietnamese    FoodType = "Vietnamese"
	FoodGreek         FoodType = "Greek"
	FoodCaribbean     FoodType = "Caribbean"
	FoodBBQ           FoodType = "BBQ"
	FoodVegan         FoodType = "Vegan"
	FoodVegetarian    FoodType = "Vegetarian"
	FoodSeafood       FoodType = "Seafood"
	FoodFastFood      FoodType = "Fast Food"
	FoodDeli          FoodType = "Deli"
	FoodSteakhouse    FoodType = "Steakhouse"
	FoodSoulFood      FoodType = "Soul Food"
	FoodBrunch        FoodType = "Brunch"
	FoodDessert       FoodType = "Dessert"
	FoodBakery        FoodType = "Bakery"
	FoodLatin         FoodType = "Latin"
	FoodTexMex        FoodType = "Tex-Mex"
	FoodMiddleEastern FoodType = "Middle Eastern"
	FoodAfrican       FoodType = "African"
	FoodFusion        FoodType = "Fusion"
	FoodHealthy       FoodType = "Healthy"
	FoodStreetFood    FoodType = "Street Food"
	FoodComfort       FoodType = "Comfort Food"
)

type FoodShop struct {
	Id          string   `db:"Id" json:"id"`
	Slug        string   `db:"Slug" json:"slug"`
	Name        string   `db:"Name" json:"name"`
	Bio         string   `db:"Bio" json:"bio"`
	FoodType    FoodType `db:"FoodType" json:"foodType"`
	Logo        string   `db:"Logo" json:"logo"`
	Address     string   `db:"Address" json:"address"`
	HasDelivery bool     `db:"HasDelivery" json:"hasDelivery"`
}

func (f *FoodShop) Validate() error {
	if f.Slug == "" || len(f.Slug) < 3 {
		return errors.New("please enter a valid value for your food shop's url slug")
	}
	f.Slug = strings.Replace(f.Slug, " ", "-", -1)

	if f.Name == "" || len(f.Name) < 3 {
		return errors.New("please enter a valid value for your food shop's name")
	}
	if f.Bio == "" || len(f.Bio) < 3 {
		return errors.New("please enter a valid value for your food shop's bio")
	}

	if f.FoodType == "" || len(f.FoodType) < 3 {
		return errors.New("please enter a valid value for your food shop's food type")
	}

	if f.Logo == "" || len(f.Logo) < 3 {
		return errors.New("please enter a valid value for your food shop's logo")
	}

	if f.Address == "" || len(f.Address) < 3 {
		return errors.New("please enter a valid value for your food shop's address")
	}

	return nil
}

func NewFoodShop(name string, bio string, foodType FoodType, logo string, address string, hasDelivery bool) (*FoodShop, error) {
	if name == "" || len(name) < 3 {
		return nil, errors.New("please enter a valid value for your food shop's name")
	}
	if bio == "" || len(bio) < 3 {
		return nil, errors.New("please enter a valid value for your food shop's bio")
	}

	if foodType == "" || len(foodType) < 3 {
		return nil, errors.New("please enter a valid value for your food shop's foodType")
	}

	if logo == "" || len(logo) < 3 {
		return nil, errors.New("please enter a valid value for your food shop's logo")
	}

	if address == "" || len(address) < 3 {
		return nil, errors.New("please enter a valid value for your food shop's bio")
	}

	uid, _ := uuid.NewV4()
	return &FoodShop{
		Id:          uid.String(),
		Name:        name,
		Bio:         bio,
		FoodType:    foodType,
		Logo:        logo,
		Address:     address,
		HasDelivery: hasDelivery,
	}, nil
}
