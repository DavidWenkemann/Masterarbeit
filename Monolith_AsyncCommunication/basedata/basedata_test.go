//basedata package is responsible for adding and removing basedata.
//It is connected to DB and UI

package basedata

import (
	"reflect"
	"testing"

	"github.com/DavidWenkemann/Masterarbeit/Monolith_AsyncCommunication/model"
)

func TestAddProduct(t *testing.T) {

	NewProduct("4011803092174", "Spezi", 0.75)

	type args struct {
		ean   string
		name  string
		price float64
	}
	tests := []struct {
		name string
		args args
		want model.BProduct
	}{
		{
			name: "AddMate",
			args: args{
				ean:   "4029764001807",
				name:  "Mate",
				price: 1.0,
			},
			want: GetProductByEan("4029764001807"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddProduct(tt.args.ean, tt.args.name, tt.args.price); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveProduct(t *testing.T) {

	type args struct {
		ean string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveProductByEan(tt.args.ean)
		})
	}
}
