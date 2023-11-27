package db

import "sync"

// Заказ на доставку товаров.
type Order struct {
	ID              int       // Номер заказа.
	IsOpen          bool      // Открыт или закрыт заказ.
	DeliveryTime    int64     // Срок доставки заказа.
	DeliveryAddress string    // Адрес доставки заказа.
	Products        []Product // Состав заказа.
}

// Товар.
type Product struct {
	ID    int     // Артикул товара.
	Name  string  // Название товара.
	Price float64 // Цена товара.
}

// База данных заказов.
type DB struct {
	m     sync.Mutex // Мьютекс для синхронизации доступа.
	id    int        // Текущее значение ID для нового заказа.
	store map[int]Order
}

// Конструктор БД.
func New() *DB {
	db := DB{
		id:    1, // Первый номер заказа.
		store: map[int]Order{},
	}
	return &db
}

// Orders возвращает все заказы.
func (db *DB) Orders() []Order {
	db.m.Lock()
	defer db.m.Unlock()
	var data []Order
	for _, v := range db.store {
		data = append(data, v)
	}
	return data
}

// NewOrder создаёт новый заказ.
func (db *DB) NewOrder(o Order) int {
	db.m.Lock()
	defer db.m.Unlock()
	o.ID = db.id
	db.store[o.ID] = o
	db.id++
	return o.ID
}

// UpdateOrder обновляет данные заказа по ID.
func (db *DB) UpdateOrder(o Order) {
	db.m.Lock()
	defer db.m.Unlock()
	db.store[o.ID] = o
}

// DeleteOrder удаляет заказ по IDю
func (db *DB) DeleteOrder(id int) {
	db.m.Lock()
	defer db.m.Unlock()
	delete(db.store, id)
}
