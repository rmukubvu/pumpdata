package notifications

import "testing"

func TestSendEmail(t *testing.T) {
	type args struct {
		toAddress   string
		pumpMessage string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{
			toAddress:   "rmukubvu@gmail.com",
			pumpMessage: "This is a good test",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
