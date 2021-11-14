package config

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestLoadFromEmv(t *testing.T) {
	type args struct {
		config *testConf
	}
	tests := []struct {
		name    string
		envs    map[string]string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "normal",
			envs: map[string]string{
				"TEST_ACCESS_KEY_ID":     "a",
				"TEST_SECRET_ACCESS_KEY": "b",
				"TEST_REGION":            "c",
				"TEST_DEBUG":             "true",
				"TEST_RETRY_COUNT":       "10",
				"TEST_TIME_OUT":          "100s",
				"MASTER_TEST_URL":        "master://",
				"MASTER_TEST_TIMEOUT":    "1m",
				"SLAVE_TEST_URL":         "slave://",
				"SLAVE_TEST_TIMEOUT":     "10m",
			},
			args: args{
				config: &testConf{
					Child: &testConfChild{},
					Sessions: map[string]*testSessionCnf{
						"master": nil,
						"slave":  nil,
					},
				},
			},
			want: &testConf{
				AccessKeyID:     "a",
				SecretAccessKey: "b",
				TEST_DEBUG:      true,
				RetryCount:      10,
				Timeout:         100 * time.Second,
				Child:           &testConfChild{Region: "c"},
				ChildVal:        testConfChild{Region: "c"},
				Sessions: map[string]*testSessionCnf{
					"master": &testSessionCnf{
						URL:     "master://",
						Timeout: 1 * time.Minute,
					},
					"slave": &testSessionCnf{
						URL:     "slave://",
						Timeout: 10 * time.Minute,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "normal default value",
			envs: map[string]string{
				"TEST_ACCESS_KEY_ID": "mid",
				// "TEST_SECRET_ACCESS_KEY": "b",
				"TEST_REGION":      "mid",
				"TEST_DEBUG":       "true",
				"TEST_RETRY_COUNT": "11",
				"TEST_TIME_OUT":    "110s",
			},
			args: args{
				config: &testConf{
					Child: &testConfChild{},
				},
			},
			want: &testConf{
				AccessKeyID:     "mid",
				SecretAccessKey: "defaultAccessKey",
				TEST_DEBUG:      true,
				RetryCount:      11,
				Timeout:         110 * time.Second,
				Child:           &testConfChild{Region: "mid"},
				ChildVal:        testConfChild{Region: "mid"},
			},
			wantErr: false,
		},
		{
			name: "normal: default is not overwrite existing value",
			envs: map[string]string{
				"TEST_ACCESS_KEY_ID": "mid",
				// "TEST_SECRET_ACCESS_KEY": "b",
				"TEST_REGION":      "mid",
				"TEST_DEBUG":       "true",
				"TEST_RETRY_COUNT": "11",
				"TEST_TIME_OUT":    "110s",
			},
			args: args{
				config: &testConf{
					SecretAccessKey: "existsValue",
					Child:           &testConfChild{},
				},
			},
			want: &testConf{
				AccessKeyID:     "mid",
				SecretAccessKey: "existsValue",
				TEST_DEBUG:      true,
				RetryCount:      11,
				Timeout:         110 * time.Second,
				Child:           &testConfChild{Region: "mid"},
				ChildVal:        testConfChild{Region: "mid"},
			},
			wantErr: false,
		},

		{
			name: "error required filed is empty",
			args: args{
				config: &testConf{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "required fields has already set value",
			envs: map[string]string{
				"TEST_SECRET_ACCESS_KEY": "ca",
				"TEST_REGION":            "ca_region",
				"TEST_DEBUG":             "true",
				"TEST_RETRY_COUNT":       "11",
				"TEST_TIME_OUT":          "110s",
			},
			args: args{
				config: &testConf{AccessKeyID: "existingKeyID"},
			},
			want: &testConf{
				AccessKeyID:     "existingKeyID",
				SecretAccessKey: "ca",
				TEST_DEBUG:      true,
				RetryCount:      11,
				Timeout:         110 * time.Second,
				ChildVal:        testConfChild{Region: "ca_region"},
			},
			wantErr: false,
		},
		{
			name: "error parse int",
			envs: map[string]string{
				"TEST_ACCESS_KEY_ID":     "a",
				"TEST_SECRET_ACCESS_KEY": "b",
				"TEST_REGION":            "c",
				"TEST_DEBUG":             "true",
				"TEST_RETRY_COUNT":       "not integer value",
				"TEST_TIME_OUT":          "100s",
				"MASTER_TEST_URL":        "master://",
				"MASTER_TEST_TIMEOUT":    "1m",
				"SLAVE_TEST_URL":         "slave://",
				"SLAVE_TEST_TIMEOUT":     "10m",
			},
			args: args{
				config: &testConf{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(tt.envs, t)
			defer func() {
				unsetTestEnv(tt.envs, t)
			}()
			err := loadFromEnv(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadFromEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			got := tt.args.config
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func setEnv(m map[string]string, t *testing.T) {
	for k, v := range m {
		if err := os.Setenv(k, v); err != nil {
			t.Error(err)
			return
		}
	}
}

func unsetTestEnv(m map[string]string, t *testing.T) {
	for k, _ := range m {
		if err := os.Unsetenv(k); err != nil {
			t.Error(err)
			return
		}
	}
}

type testConf struct {
	AccessKeyID     string `required:"true" env:"TEST_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"TEST_SECRET_ACCESS_KEY" default:"defaultAccessKey"`
	TEST_DEBUG      bool   `env:"TEST_DEBUG"`

	Timeout    time.Duration `env:"TEST_TIME_OUT"`
	RetryCount uint64        `env:"TEST_RETRY_COUNT"`

	Child    *testConfChild `yaml: child`
	ChildVal testConfChild

	Sessions map[string]*testSessionCnf
}

type testConfChild struct {
	Region string `env:"TEST_REGION"`
}

type testSessionCnf struct {
	URL     string        `env:"TEST_URL"`
	Timeout time.Duration `env:"TEST_TIMEOUT"`
}
