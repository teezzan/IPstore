package ipstore

import (
	"reflect"
	"sort"
	"testing"

	"github.com/jaswdr/faker"
)

func TestRequestHandled(t *testing.T) {
	type args struct {
		ip_address string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test incorrect ip address input",
			args: args{
				ip_address: "evisit.com",
			},
			wantErr: true,
		},
		{
			name: "test correct ipv4 address input",
			args: args{
				ip_address: "102.89.32.126",
			},
			wantErr: false,
		},
		{
			name: "test correct ipv6 address input",
			args: args{
				ip_address: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RequestHandled(tt.args.ip_address); (err != nil) != tt.wantErr {
				t.Errorf("RequestHandled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

