package main

import (
	"fmt"
	"math/big"
)

func Add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func Mult(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func Div(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		panic("деление на ноль")
	}
	return new(big.Int).Div(a, b)
}

func AddAndSub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(Add(a, b), Sub(a, b))
}

func MultAndDiv(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		panic("деление на ноль")
	}
	return new(big.Int).Div(Mult(a, b), Div(a, b))
}

func main() {
	a := new(big.Int).Exp(big.NewInt(2), big.NewInt(30), nil)
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil)
	fmt.Println("Значение a:", a)
	fmt.Println("Значение b:", b)

	fmt.Println("Сложение:", Add(a, b))
	fmt.Println("Вычитание:", Sub(a, b))
	fmt.Println("Умножение:", Mult(a, b))
	fmt.Println("Деление:", Div(a, b))

	fmt.Println("Сложение и вычитание:", AddAndSub(a, b))
	fmt.Println("Умножение и деление:", MultAndDiv(a, b))

}
