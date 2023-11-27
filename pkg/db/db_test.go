package db

import (
	"reflect"
	"sync"
	"testing"
)

func TestDB_DeleteOrder(t *testing.T) {
	type fields struct {
		m     sync.Mutex
		id    int
		store map[int]Order
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				m:     tt.fields.m,
				id:    tt.fields.id,
				store: tt.fields.store,
			}
			db.DeleteOrder(tt.args.id)
		})
	}
}

func TestDB_NewOrder(t *testing.T) {
	type fields struct {
		m     sync.Mutex
		id    int
		store map[int]Order
	}
	type args struct {
		o Order
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				m:     tt.fields.m,
				id:    tt.fields.id,
				store: tt.fields.store,
			}
			if got := db.NewOrder(tt.args.o); got != tt.want {
				t.Errorf("NewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Orders(t *testing.T) {
	type fields struct {
		m     sync.Mutex
		id    int
		store map[int]Order
	}
	tests := []struct {
		name   string
		fields fields
		want   []Order
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				m:     tt.fields.m,
				id:    tt.fields.id,
				store: tt.fields.store,
			}
			if got := db.Orders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Orders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_UpdateOrder(t *testing.T) {
	type fields struct {
		m     sync.Mutex
		id    int
		store map[int]Order
	}
	type args struct {
		o Order
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				m:     tt.fields.m,
				id:    tt.fields.id,
				store: tt.fields.store,
			}
			db.UpdateOrder(tt.args.o)
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
