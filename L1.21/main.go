package main

import "fmt"

// OrderManager - интерфейс для управления заказами, который ожидает клиент
type OrderManager interface {
	CreateOrder(orderID, item string) string
	CancelOrder(orderID string) bool
}

// KFCOrderManager - совместимая реализация интерфейса OrderManager для KFC
type KFCOrderManager struct {
	orders map[string]string
}

// Конструктор KFCOrderManager
func NewKFCOrderManager() *KFCOrderManager {
	return &KFCOrderManager{orders: make(map[string]string)}
}

func (kom *KFCOrderManager) CreateOrder(orderID, item string) string {
	fmt.Printf("[KFC] Создан заказ %s: %s\n", orderID, item)
	kom.orders[orderID] = item
	return orderID
}

func (kom *KFCOrderManager) CancelOrder(orderID string) bool {
	if _, exists := kom.orders[orderID]; exists {
		delete(kom.orders, orderID)
		fmt.Printf("[KFC] Заказ %s отменён\n", orderID)
		return true
	}
	fmt.Printf("[KFC] Заказ %s не найден\n", orderID)
	return false
}

// McDonaldsOrderManager - несовместимая реализация менеджера заказов для McDonalds (не реализует OrderManager)
type McDonaldsOrderManager struct {
	orders map[string]string
	price  map[string]float64
}

// Конструктор McDonaldsOrderManager
func NewMcDonaldsOrderManager() *McDonaldsOrderManager {
	return &McDonaldsOrderManager{
		orders: make(map[string]string),
		price:  make(map[string]float64),
	}
}

func (mom *McDonaldsOrderManager) CreateOrder(orderID, item string, price float64) string {
	fmt.Printf("[McDonalds] Создан заказ %s: %s по цене %.2f\n", orderID, item, price)
	mom.orders[orderID] = item
	mom.price[orderID] = price
	return orderID
}

func (mom *McDonaldsOrderManager) CancelOrder(orderID string) bool {
	if _, exists := mom.orders[orderID]; exists {
		delete(mom.orders, orderID)
		delete(mom.price, orderID)
		fmt.Printf("[McDonalds] Заказ %s отменён\n", orderID)
		return true
	}
	fmt.Printf("[McDonalds] Заказ %s не найден\n", orderID)
	return false
}

// McDonaldsOrderManagerAdapter - адаптер для McDonaldsOrderManager, реализующий интерфейс OrderManager
type McDonaldsOrderManagerAdapter struct {
	mcd   *McDonaldsOrderManager
	price float64
}

// Проверка соответствия интерфейсу OrderManager
var _ OrderManager = (*McDonaldsOrderManagerAdapter)(nil)

// NewOrderManagerAdapter - конструктор адаптера
func NewOrderManagerAdapter(mcd *McDonaldsOrderManager, price float64) *McDonaldsOrderManagerAdapter {
	return &McDonaldsOrderManagerAdapter{mcd: mcd, price: price}
}

// Реализация методов интерфейса OrderManager через адаптер
func (oma *McDonaldsOrderManagerAdapter) CreateOrder(orderID, item string) string {
	// Маппим вызов: CreateOrder(orderID, item) -> CreateOrder(orderID, item, oma.price)
	// Используем фиксированную цену для всех заказов через адаптер
	return oma.mcd.CreateOrder(orderID, item, oma.price)
}

func (oma *McDonaldsOrderManagerAdapter) CancelOrder(orderID string) bool {
	return oma.mcd.CancelOrder(orderID)
}

// checkout - клиентский код для оформления заказа
func checkout(orderManager OrderManager, id, item string) {
	orderManager.CreateOrder(id, item)
}

func main() {
	// kfcManager и mcdManager - менеджеры заказов для KFC и McDonalds соответственно
	kfcManager := NewKFCOrderManager()
	mcdManager := NewMcDonaldsOrderManager()
	// mcDonaldsOrderManagerAdapter адаптирует mcdManager к интерфейсу OrderManager
	mcDonaldsOrderManagerAdapter := NewOrderManagerAdapter(mcdManager, 4.99)

	// managers - срез менеджеров заказов, реализующих интерфейс OrderManager
	var managers = []OrderManager{kfcManager, mcDonaldsOrderManagerAdapter}

	checkout(managers[0], "KFC123", "Ведро курицы")
	checkout(managers[1], "MCD456", "Биг Мак")

	managers[0].CancelOrder("KFC123")
	managers[1].CancelOrder("MCD457") // Попытка отменить несуществующий заказ
}
