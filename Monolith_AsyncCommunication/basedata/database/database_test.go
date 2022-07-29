//Package Database emulates a Database.

//In fact all the functions are implemented, but the data is simply stored

//in RAM and will be resetted when the application is closed.

package database

import "testing"

func TestRemoveProductByEan(t *testing.T) {

	p1 := NewProduct("1234", "Cola", 1.2)
	p2 := NewProduct("12345", "Cola", 1.2)
	NewProduct("123456", "Cola", 1.2)
	NewProduct("1234567", "Cola", 1.2)

	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Delete 1",
			args: args{
				id: p1.ProductID,
			},
		},
		{
			name: "Delete 2",
			args: args{
				id: p2.ProductID,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveProductByEan(tt.args.id)
		})
	}
}
