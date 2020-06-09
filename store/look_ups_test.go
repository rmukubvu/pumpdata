package store

import (
	"github.com/rmukubvu/pumpdata/model"
	"reflect"
	"testing"
)

func TestAddPumpType(t *testing.T) {
	type args struct {
		p model.PumpTypes
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddPumpType(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("AddPumpType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFetchPumpTypes(t *testing.T) {
	testData := make([]model.PumpTypes, 4)
	testData[0] = model.PumpTypes(struct {
		Id   int
		Name string
	}{Id: 1, Name: "Electrical"})
	testData[1] = model.PumpTypes(struct {
		Id   int
		Name string
	}{Id: 2, Name: "Mechanical"})
	testData[2] = model.PumpTypes(struct {
		Id   int
		Name string
	}{Id: 3, Name: "Actuator"})
	testData[3] = model.PumpTypes(struct {
		Id   int
		Name string
	}{Id: 4, Name: "Test Jig"})

	tests := []struct {
		name    string
		want    []model.PumpTypes
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test-1", testData, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchPumpTypes()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchPumpTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchPumpTypes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
