package config

type Session struct {
	MaxAge int    `toml:"max_age"`
	Domain string `toml:"domain"`
}
