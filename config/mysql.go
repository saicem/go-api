package config

import "time"

type MySQL struct {
	Addr            string        `toml:"addr"`
	User            string        `toml:"user"`
	Pass            string        `toml:"pass"`
	Name            string        `toml:"name"`
	Config          string        `toml:"config"`
	MaxOpenConn     int           `toml:"maxOpenConn"`
	MaxIdleConn     int           `toml:"maxIdleConn"`
	ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
}

func (m *MySQL) Dsn() string {
	return m.User + ":" + m.Pass + "@tcp(" + m.Addr + ")/" + m.Name + "?" + m.Config
}
