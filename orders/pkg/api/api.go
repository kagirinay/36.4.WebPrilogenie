package api

import (
	"encoding/json"
	"net/http"
	"orders/pkg/db"

	"github.com/gorilla/mux"
)

// API приложения.
type API struct {
	r  *mux.Router // Маршрутизатор запросов
	db *db.DB
}

// Конструктор API
func New(db *db.DB) *API {
	api := API{}
	api.db = db
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

// Router возвращает маршрутизатор запросов.
func (api *API) Router() *mux.Router {
	return api.r
}

// Регистрация методов API в маршрутизаторе запросов.
func (api *API) endpoints() {
	api.r.HandleFunc("/orders", api.ordersHandler).Methods(http.MethodGet)
	api.r.HandleFunc("/orders", api.newOrdersHandler).Methods(http.MethodPost)
	api.r.HandleFunc("/orders/{id}", api.updateOrdersHandler).Methods(http.MethodPatch)
	api.r.HandleFunc("/orders/{id}", api.deleteOrdersHandler).Methods(http.MethodDelete)
}

// ordersHandler создаёт новый заказ
func (api *API) ordersHandler(w http.ResponseWriter, r http.Request) {
	// Получение данных из БД.
	orders := api.db.Orders()
	// Отправка данных клиенту в формате JSON
	json.NewEncoder(w).Encode(orders)
}
