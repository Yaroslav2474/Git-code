package models

import (
	"fmt"
	"time"
)

type TransactionType string

var Pending, Completed, Canceled TransactionType

type Transaction struct {
	ID          string
	Amount      float64
	Type        TransactionType
	Category    string
	Date        time.Time
	Description string
}

func (tran Transaction) IsValid() (bool, error) {

	if tran.Amount <= 0 {
		return false, fmt.Errorf("сумма должна быть положительной")
	}

	if tran.Type != TransactionType(Income) || tran.Type != TransactionType(Expense) {
		return false, fmt.Errorf("неизветсный тип транзакции")
	}

	if len(tran.Category) == 0 {
		return false, fmt.Errorf("категория не может быть пустой")
	}

	if len(tran.Description) == 0 {
		return false, fmt.Errorf("описание не должно быть пустым")
	}

	return true, fmt.Errorf("")
}

func (tran Transaction) GetFormattedAmount() string {

}

func (tran Transaction) GetDisplayDate() string {

}

func NewTransaction(amount float64, tType TransactionType, category, description string) (*Transaction, error) {

}

func ParseTransactionType(s string) (TransactionType, error) {

}

func AddTrannsaction() {
	var choise int

	fmt.Print("----------------------Меню----------------------")
	fmt.Print("1. Добавить транзакцию\n2. Показать транзакции\n0. Выход")

	fmt.Scan(&choise)

	for {
		if choise == 1 {

			fmt.Println("Выберете вид транзакции: \n1.Доходы\n2.Расходы")
			fmt.Scan(&choise)

			if choise == 1 {

			} else if choise == 2 {

			} else {
				fmt.Print("Только два вида транзакций.")
			}

		} else if choise == 2 {
			GetDefaultCategories()
		} else if choise == 0 {
			fmt.Println("До свидания!")
			break
		} else {
			fmt.Println("Только 3 действия в меню.")
		}
	}
}
