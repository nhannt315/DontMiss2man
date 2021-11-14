package logs

// Config ロギングに関する設定を保持
type Config struct {
	Level  Level  `yaml:"level" required:"true"`
	Format Format `yaml:"format" default:"json" required:"true"`
}
