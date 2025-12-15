package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type сategoryType string

const (
	Income  сategoryType = "income"
	Expense сategoryType = "expense"
)

type Category struct {
	ID          string
	Title       string
	Type        сategoryType
	Edit        bool
	Description string
	// Budget float64
}

var (
	defaultExpenseCategories = []Category{
		{ID: "sys_exp_food", Title: "Продукты", Type: Expense, Edit: true},
		{ID: "sys_exp_transport", Title: "Транспорт", Type: Expense, Edit: true},
		{ID: "sys_exp_housing", Title: "Жилье", Type: Expense, Edit: true},
		{ID: "sys_exp_entertainment", Title: "Развлечения", Type: Expense, Edit: true},
	}

	defaultIncomeCategories = []Category{
		{ID: "sys_inc_salary", Title: "Зарплата", Type: Income, Edit: true},
		{ID: "sys_inc_gifts", Title: "Подарки", Type: Income, Edit: true},
		{ID: "sys_inc_other", Title: "Прочие доходы", Type: Income, Edit: true},
	}
)

const filename = "./data/categories.json"

func loadFromFile() {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&defaultExpenseCategories)
	if err != nil {
		fmt.Println("Ошибка загрузки данных, начнём заново.")
		defaultExpenseCategories = []Category{}
	}

	decoder = json.NewDecoder(file)
	err = decoder.Decode(&defaultIncomeCategories)
	if err != nil {
		fmt.Println("Ошибка загрузки данных, начнём заново.")
		defaultIncomeCategories = []Category{}
	}
}

func saveToFile() {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Не могу сохранить файл.")
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(defaultExpenseCategories)
	if err != nil {
		fmt.Println("Ошибка при сохранении.")
	}

	encoder = json.NewEncoder(file)
	err = encoder.Encode(defaultIncomeCategories)
	if err != nil {
		fmt.Println("Ошибка при сохранении.")
	}
}

func GetDefaultCategories() []Category {

	fileInfo, err := os.ReadFile("./data/categories.json")

	if err != nil {
		fmt.Printf("Возникла ошабка при инициализации базы данных\n%s", err)
	}

	if fileInfo == nil {
		for _, v := range defaultExpenseCategories {
			fmt.Printf("ID: %s Название: %s  Тип транзакции: %s  Подлежит редактированию: %v", v.ID, v.Title, v.Type, v.Edit)
		}
		for _, v := range defaultIncomeCategories {
			fmt.Printf("ID: %s Название: %s  Тип транзакции: %s  Подлежит редактированию: %v", v.ID, v.Title, v.Type, v.Edit)
		}

	}
	var inf []Category

	if err := json.Unmarshal(fileInfo, &inf); err != nil {
		fmt.Printf("Ошибка парсинга JSON: %v\n", err)
		return inf
	}

	for _, v := range inf {
		fmt.Printf("ID: %s Название: %s  Тип транзакции: %s  Подлежит редактированию: %v", v.ID, v.Title, v.Type, v.Edit)
	}

	return inf
}

func FindCategoryByID(id string) (*Category, error) {
	defaultCategories := GetDefaultCategories()
	for _, category := range defaultCategories {
		if category.ID == id {
			return &category, nil
		}
	}

}
