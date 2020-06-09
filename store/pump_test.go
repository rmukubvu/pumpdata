package store

import (
	"github.com/rmukubvu/pumpdata/model"
	"testing"
	"time"
)

func TestAddPump(t *testing.T) {
	type args struct {
		p model.Pump
	}
	model := model.Pump{
		TypeId:       1,
		SerialNumber: "123456789",
		NickName:     "Fourways-Pump",
		Lat:          -26.0189066,
		Lng:          28.0043687,
		CreatedDate:  time.Now().Unix(),
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test-1", args{p: model}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddPump(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("AddPump() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
