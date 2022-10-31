package config

var (
	C *Config
)

type Config struct {
	YoutuberBotConfig `yaml:"youtuber_bot_config"`
}

type BotConfig struct {
	Token string `yaml:"token"`
}

type YoutuberBotConfig struct {
	Name      string    `yaml:"name"`
	BotConfig BotConfig `yaml:"bot_config"`
}
