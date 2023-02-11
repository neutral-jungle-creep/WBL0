package main

import "fmt"

type User struct{}

func (u *User) ChargePhoneInEurope(c CBlock) { // функция зарядки телефона принимает интерфейс зарядный блок
	c.UseEuroSocket()
}

type CBlock interface {
	UseEuroSocket() // есть только евро розетка
}

type EuroBlock struct{}

func (b *EuroBlock) UseEuroSocket() { // евро блок подходит
	fmt.Println("зарядка устройства началась")
}

type AsianBlock struct{}

func (b *AsianBlock) UseAsianSocket() { // блок азиатского типа не может использовать евро розетку
	fmt.Println("зарядка устройства началась")
}

type EuroAdapter struct {
	aBlock *AsianBlock // на помощь приходит переходник на евро
}

func (a *EuroAdapter) UseEuroSocket() {
	fmt.Println("к блоку азиатского типа добавлен евро адаптер")
	a.aBlock.UseAsianSocket() // теперь устройство можно зарядить
}

func main() {
	var (
		user       = &User{}
		eurBlock   = &EuroBlock{}
		asianBlock = &AsianBlock{}
	)

	user.ChargePhoneInEurope(eurBlock) // пример без переходника

	euroAdapter := &EuroAdapter{ // создание адаптера
		aBlock: asianBlock,
	}

	user.ChargePhoneInEurope(euroAdapter) // пример работы с переходником
}
