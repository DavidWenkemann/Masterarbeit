package userinterface

import "testing"

func Test_scanItem(t *testing.T) {

	type args struct {
		itemID string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanItem(tt.args.itemID)
		})
	}
}
