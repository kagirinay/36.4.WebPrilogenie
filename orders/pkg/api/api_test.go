package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"orders/pkg/db"
	"testing"
)

func TestAPI_ordersHandler(t *testing.T) {
	// Создаем чистый объект API для теста.
	dbase := db.New()
	dbase.NewOrder(db.Order{})
	api := New(dbase)
	// Создаем HTTP-запрос.
	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	// Создаем объект для записи ответа обработчика.
	rr := httptest.NewRecorder()
	// Вызываем маршрутизатор. Маршрутизатор для пути и метода запроса
	// вызовет обработчик. Обработчик запишет ответ в созданный объект.
	api.r.ServeHTTP(rr, req)
	// Проверяем код ответа.
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	// Читаем тело ответа.
	b, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("не удалось раскодировать ответ сервера: %v", err)
	}
	// Раскодируем JSON в массив заказов.
	var data []db.Order
	err = json.Unmarshal(b, &data)
	if err != nil {
		t.Fatalf("не удалось раскодировать ответ сервера: %v", err)
	}
	// Проверяем, что в массиве ровно один элемент.
	const wantLen = 1
	if len(data) != wantLen {
		t.Fatalf("получено %d записей, ожидалось %d", len(data), wantLen)
	}
	// Также можно проверить совпадение заказов в результате
	// с добавленными в БД для теста.
}

func TestAPI_newOrderHandler(t *testing.T) {
	dbase := db.New()
	dbase.NewOrder(db.Order{})
	api := New(dbase)

	o := db.Order{}
	b, _ := json.Marshal(o)
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(b))
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}

func TestAPI_updateOrderHandler(t *testing.T) {
	dbase := db.New()
	dbase.NewOrder(db.Order{})
	api := New(dbase)

	o := db.Order{}
	b, _ := json.Marshal(o)
	req := httptest.NewRequest(http.MethodPatch, "/orders/100", bytes.NewBuffer(b))
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}

func TestAPI_deleteOrderHandler(t *testing.T) {
	dbase := db.New()
	dbase.NewOrder(db.Order{})
	api := New(dbase)

	req := httptest.NewRequest(http.MethodDelete, "/orders/100", nil)
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}
