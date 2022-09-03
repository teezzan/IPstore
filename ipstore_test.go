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

func TestTop100(t *testing.T) {
	isList1 := generateIPList(100)
	isListLen50 := generateIPList(50)
	isListLen70 := generateIPList(70)

	tests := []struct {
		name string
		seed []string
		want []string
	}{
		{
			name: "output should be of length 100",
			seed: isList1,
			want: isList1,
		},
		{
			name: "output should be of length 100",
			seed: isListLen50,
			want: isListLen50,
		},
		{
			name: "output should be of length 100",
			seed: isListLen70,
			want: isListLen70,
		},
	}
	for _, tt := range tests {
		Clear()
		insertIntoIPstore(tt.seed)
		sort.Strings(tt.seed)
		t.Run(tt.name, func(t *testing.T) {
			got := Top100()
			sort.Strings(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Top100() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateIPList(limit int) []string {
	faker := faker.New()
	var outputIP []string
	for i := 0; i < limit; i++ {
		outputIP = append(outputIP, faker.Internet().Ipv4())
	}
	return outputIP
}

func insertIntoIPstore(ips []string) {
	for _, ip_address := range ips {
		RequestHandled(ip_address)
	}
}

func TestClear(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "should return empty list",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Clear()
			if got := Top100(); len(got) != 0 {
				t.Errorf("Top100() = %v, want []", got)
			}
		})
	}
}
