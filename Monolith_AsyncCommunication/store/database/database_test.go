//Package Database emulates a Database.

//In fact all the functions are implemented, but the data is simply stored

//in RAM and will be resetted when the application is closed.

package database

import (
	"reflect"
	"testing"
	"time"

	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/store/model"
)

func TestGetItemById(t *testing.T) {

	item := model.DBItem{ProductID: 1, ItemID: "1", ReceivingDate: time.Now()}
	item2 := model.DBItem{ProductID: 1, ItemID: "2", ReceivingDate: time.Now()}
	item3 := model.DBItem{ProductID: 1, ItemID: "3", ReceivingDate: time.Now()}

	items = append(items, item)
	items = append(items, item2)
	items = append(items, item3)

	type args struct {
		itemID string
	}
	tests := []struct {
		name string
		args args
		want model.DBItem
	}{
		{
			name: "Get 1",
			args: args{
				itemID: item.ItemID,
			},
			want: item,
		},
		{
			name: "Get 2",
			args: args{
				itemID: item2.ItemID,
			},
			want: item2,
		},
		{
			name: "Get 3",
			args: args{
				itemID: item3.ItemID,
			},
			want: item3,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetItemById(tt.args.itemID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemById() = %v, want %v", got, tt.want)
			}
		})
	}
}
