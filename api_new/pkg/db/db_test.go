package db_test

import (
	"testing"

	"github.com/nhannt315/real_estate_api/pkg/db"
	"github.com/nhannt315/real_estate_api/pkg/strings"
)

func TestBuildMySQLConnectionString(t *testing.T) {
	tests := []struct {
		name    string
		args    *db.Config
		want    string
		wantErr bool
	}{
		{"local", &db.Config{Username: strings.MaskedStringWithString("root"), Password: strings.MaskedStringWithString(""), Host: "localhost", Schema: "scotch", ParseTime: true, Port: 3306}, "root@tcp(localhost:3306)/scotch?parseTime=true", false},
		{"prod-like", &db.Config{Username: strings.MaskedStringWithString("scotch_update"), Password: strings.MaskedStringWithString("sEcRetPaSs"), Host: "prod-abc.cluster-xyz123.ap-northeast-1.rds.amazonaws.com", Schema: "scotch", ParseTime: true, Port: 3306}, "scotch_update:sEcRetPaSs@tcp(prod-abc.cluster-xyz123.ap-northeast-1.rds.amazonaws.com:3306)/scotch?parseTime=true", false},
		{"custom port", &db.Config{Username: strings.MaskedStringWithString("root"), Password: strings.MaskedStringWithString("password"), Host: "localhost", Port: 13306, Schema: "apple", ParseTime: true}, "root:password@tcp(localhost:13306)/apple?parseTime=true", false},
		{"host with port ignores port option", &db.Config{Username: strings.MaskedStringWithString("root"), Password: strings.MaskedStringWithString("password"), Host: "localhost:23306", Port: 9999, Schema: "mango", ParseTime: true}, "root:password@tcp(localhost:23306)/mango?parseTime=true", false},
		{"tls", &db.Config{Username: strings.MaskedStringWithString("user01"), Password: strings.MaskedStringWithString("secure-password"), Host: "example.com:443", Schema: "orange", TLS: "true", ParseTime: true}, "user01:secure-password@tcp(example.com:443)/orange?parseTime=true&tls=true", false},
		{"unknown tls option", &db.Config{Username: strings.MaskedStringWithString("root"), Password: strings.MaskedStringWithString("password"), Host: "example.com:443", Schema: "orange", TLS: "foo", ParseTime: true}, "", true},
		{"tz", &db.Config{Username: strings.MaskedStringWithString("root"), Password: strings.MaskedStringWithString(""), Host: "localhost", Schema: "scotch", Location: "Asia/Tokyo", ParseTime: true, Port: 3306}, "root@tcp(localhost:3306)/scotch?loc=Asia%2FTokyo&parseTime=true", false},
		{"collation", &db.Config{Username: strings.MaskedStringWithString("root"), Password: strings.MaskedStringWithString(""), Host: "localhost", Schema: "scotch", Collation: "utf8mb4_bin", ParseTime: true, Port: 3306}, "root@tcp(localhost:3306)/scotch?collation=utf8mb4_bin&parseTime=true", false},
		{"max allowed packet", &db.Config{Username: strings.MaskedStringWithString("root"), Password: strings.MaskedStringWithString(""), Host: "localhost", Schema: "scotch", MaxAllowedPacket: 20000, ParseTime: true, Port: 3306}, "root@tcp(localhost:3306)/scotch?parseTime=true&maxAllowedPacket=20000", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.BuildMySQLConnectionString(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildMySQLConnectionString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuildMySQLConnectionString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
