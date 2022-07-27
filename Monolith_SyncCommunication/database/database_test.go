package database

import (
	"reflect"
	"testing"
	"time"

	"github.com/DavidWenkemann/Masterarbeit/Monolith_SyncCommunication/model"
)

func Test_oldItem(t *testing.T) {
	ti := time.Now()
	dbi := model.DBItem{
		ProductID:     4,
		ItemID:        "1",
		ReceivingDate: ti,
	}
	dbi2 := model.DBItem{
		ProductID:     5,
		ItemID:        "2",
		ReceivingDate: ti,
		SellingDate:   ti.Add(time.Hour),
	}
	type args struct {
		pID          int
		timeReceived time.Time
		timeSelled   *time.Time
	}
	tests := []struct {
		name string
		args args
		want model.DBItem
	}{
		{
			name: "happycase",
			args: args{
				pID:          4,
				timeReceived: ti,
				timeSelled:   nil,
			},
			want: dbi,
		},

		{
			name: "happycase2",
			args: args{
				pID:          5,
				timeReceived: ti,
				timeSelled:   timePtr(ti.Add(time.Hour)),
			},
			want: dbi2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oldItem(tt.args.pID, tt.args.timeReceived, tt.args.timeSelled); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oldItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetItemsInStockByEan(t *testing.T) {

	//newProduct("4011803092174", "Spezi", 0.75)
	dbp := newProduct("4011803092174", "Spezi", 0.75)
	dbp2 := newProduct("4029764001807", "Mate", 1)

	items = append(items, oldItem(dbp.ProductID, time.Now().Add(-24*time.Hour), nil))
	items = append(items, oldItem(dbp2.ProductID, time.Now().Add(-24*time.Hour), nil))
	items = append(items, oldItem(dbp2.ProductID, time.Now().Add(-24*time.Hour), nil))

	type args struct {
		ean string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get1Spezi",
			args: args{
				ean: "4011803092174",
			},
			want: 1,
		},
		{
			name: "get2Mate",
			args: args{
				ean: "4029764001807",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetItemsInStockByEan(tt.args.ean); got != tt.want {
				t.Errorf("GetItemsInStockByEan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProductByID(t *testing.T) {

	dbp := newProduct("4011803092174", "Spezi", 0.75)
	newProduct("4066600641919", "Paulaner Hefeweizen", 1.39)
	dbp3 := newProduct("4029764001807", "Clubmate", 2.50)

	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want model.BProduct
	}{
		// TODO: Add test cases.

		{
			name: "getSpezi",
			args: args{
				id: dbp.ProductID,
			},
			want: mapDBProductToBProduct(dbp),
		},
		{
			name: "getMate",
			args: args{
				id: dbp3.ProductID,
			},
			want: mapDBProductToBProduct(dbp3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProductByID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProductByEan(t *testing.T) {

	dbp := newProduct("4011803092174", "Spezi", 0.75)
	newProduct("4066600641919", "Paulaner Hefeweizen", 1.39)
	dbp3 := newProduct("4029764001807", "Clubmate", 2.50)

	type args struct {
		ean string
	}
	tests := []struct {
		name string
		args args
		want model.BProduct
	}{
		{
			name: "getSpezi",
			args: args{
				ean: dbp.EAN,
			},
			want: mapDBProductToBProduct(dbp),
		},
		{
			name: "getMate",
			args: args{
				ean: dbp3.EAN,
			},
			want: mapDBProductToBProduct(dbp3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProductByEan(tt.args.ean); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductByEan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveProductByEan(t *testing.T) {
	newProduct("4011803092174", "Spezi", 0.75)
	newProduct("4066600641919", "Paulaner Hefeweizen", 1.39)
	newProduct("4029764001807", "Clubmate", 2.50)

	type args struct {
		ean string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "DeleteWeizen",
			args: args{
				ean: "4066600641919",
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveProductByEan(tt.args.ean)
		})
	}
}
