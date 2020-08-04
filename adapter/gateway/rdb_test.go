package gateway

import (
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

func Test_rdbSpecific_GetDriver(t *testing.T) {
	type fields struct {
		driver string
		source string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test",
			fields: fields{
				driver: "driver",
			},
			want: "driver",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rdbSpecific{
				driver: tt.fields.driver,
				source: tt.fields.source,
			}
			if got := s.GetDriver(); got != tt.want {
				t.Errorf("rdbSpecific.GetDriver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rdbSpecific_GetSource(t *testing.T) {
	type fields struct {
		driver string
		source string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test",
			fields: fields{
				source: "source",
			},
			want: "source",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &rdbSpecific{
				driver: tt.fields.driver,
				source: tt.fields.source,
			}
			if got := s.GetSource(); got != tt.want {
				t.Errorf("rdbSpecific.GetSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRDBClient(t *testing.T) {
	type args struct {
		config PostgresConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *rdbClient
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRDBClient(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRDBClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRDBClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
