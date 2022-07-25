package database

import (
	"reflect"
	"testing"
	"time"

	"github.com/DavidWenkemann/Masterarbeit/model"
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
	type args struct {
		ean string
	}
	tests := []struct {
		name string
		args args
		want int
	}{


		
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetItemsInStockByEan(tt.args.ean); got != tt.want {
				t.Errorf("GetItemsInStockByEan() = %v, want %v", got, tt.want)
			}
		})
	}
}
