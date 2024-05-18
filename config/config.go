package config

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// app struct config
type app = struct {
	Name        string        `toml:"name"`
	Port        string        `toml:"port"`
	PrintRoutes bool          `toml:"print-routes"`
	Prefork     bool          `toml:"prefork"`
	Production  bool          `toml:"production"`
	IdleTimeout time.Duration `toml:"idle-timeout"`
	TLS         struct {
		Enable   bool   `toml:"enable"`
		CertFile string `toml:"cert-file"`
		KeyFile  string `toml:"key-file"`
	}
}

// chain struct config
type chain = struct {
	Url     string `toml:"url"`
	Filekey string `toml:"filekey"`
	Jwt     struct {
		Key                string        `toml:"key"`
		ExpirationDuration time.Duration `toml:"expiration_duration"`
	}
}

// log struct config
type logger = struct {
	Level    zerolog.Level `toml:"level"`
	Prettier bool          `toml:"prettier"`
}

// middleware
type middleware = struct {
	Compress struct {
		Enable bool
	}

	Recover struct {
		Enable bool
	}

	Monitor struct {
		Enable bool
		Path   string
	}

	Pprof struct {
		Enable bool
	}

	Limiter struct {
		Enable     bool
		Max        int
		Expiration time.Duration `toml:"expiration_seconds"`
	}

	FileSystem struct {
		Enable bool
		Browse bool
		MaxAge int `toml:"max_age"`
		Index  string
		Root   string
	}

	Jwt struct {
		Enable     bool          `toml:"enable"`
		Secret     string        `toml:"secret"`
		Expiration time.Duration `toml:"expiration_seconds"`
	}
}

type Config struct {
	App   app
	Chain chain

	Logger     logger
	Middleware middleware
}

func getConfigPath(name string, debug ...bool) string {
	if len(debug) > 0 {
		return name
	} else {
		return filepath.Join("./config/", name+".toml") // default path ./config/***.toml
	}
}

func ParseConfigFromToml(configToml []byte) (*Config, error) {
	var contents Config
	err := toml.Unmarshal(configToml, &contents)
	if err != nil {
		return &Config{}, err
	}
	return &contents, err
}

// ParseConfig func to parse config
func ParseConfig(name string, debug ...bool) (*Config, error) {
	file, err := os.ReadFile(getConfigPath(name, debug...))
	if err != nil {
		return &Config{}, err
	}
	return ParseConfigFromToml(file)
}

// NewConfig initialize config
func NewConfig() *Config {
	config, err := ParseConfig("config") // read from default path ./config/config.toml
	if err != nil {
		// panic if config is not found
		log.Panic().Err(err).Msg("config not found")
	}
	log.Info().Msg("config loaded")
	return config
}

func (c *Config) WriteConfig(name string, debug ...bool) error {
	data, err := toml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(getConfigPath(name, debug...), data, 0666)
}

// ParseAddress func to parse address
func ParseAddress(raw string) (host, port string) {
	if i := strings.LastIndex(raw, ":"); i > 0 {
		return raw[:i], raw[i+1:]
	}

	return raw, ""
}
