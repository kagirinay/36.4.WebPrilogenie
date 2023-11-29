package orders

import (
	"net/http"
	"orders/pkg/db"
	"time"
)

func main() {
	// Инициализация БД в памяти.
	dbase := db.New()
	// Добавление тестового заказа в БД.
	p := []db.Product{
		{
			Name:  "Яблоки",
			Price: 20,
		},
		{
			Name:  "Груши",
			Price: 30,
		},
	}
	o := db.Order{
		IsOpen:       true,
		DeliveryTime: time.Now().Unix(),
		Products:     p,
	}
	dbase.NewOrder(o)
	// Создание объекта API, использующего БД в памяти.
	api := api.New(dbase)
	// Запуск сетевой службы и HTTP-сервера
	// на всех локальных IP-адресах на порту 80.
	http.ListenAndServe(":80", api.Router())
}
