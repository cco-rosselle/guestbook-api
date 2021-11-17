package settings

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.jtlabs.io/settings"
)

type Settings struct {
	Data struct {
		Cluster					string 	`yaml:"cluster"`
		Database				string 	`yaml:"database"`
		DefaultLimit		int			`yaml:"defaultLimit"`
		Host						string 	`yaml:"host"`
		Query						string	`yaml:"query"`
		MaxLimit				int			`yaml:"maxLimit"`
		Pass						string	`yaml:"pass"`
		Protocol				string	`yaml:"protocol"`
		TimeoutSeconds	int			`yaml:"timeoutSeconds"`
		User						string	`yaml:"user"`

	} `yaml:"data"`

	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`

	Name string `yaml:"name"`
	
	Server struct {
		Address string `yaml:"address"`
	} `yaml:"server"`
}

func (s Settings) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("data.cluster", s.Data.Cluster).
		Str("data.database", s.Data.Database).
		Int("data.defaultLimit", s.Data.DefaultLimit).
		Str("data.host", s.Data.Host).
		Str("data.query", s.Data.Query).
		Int("data.maxLimit", s.Data.MaxLimit).
		Str("data.pass", "sensitive").
		Str("data.protocol", s.Data.Protocol).
		Int("data.timeoutSeconds", s.Data.TimeoutSeconds).
		Str("data.user", s.Data.User).
		
		Str("logging.level", s.Logging.Level).

		Str("name", s.Name).

		Str("server.address", s.Server.Address)
}

var (
	defaultSettings *Settings
	l zerolog.Logger = log.With().Str("package", "settings").Logger()
)

func Default() *Settings {
	if defaultSettings == nil {
		l.Warn().Msg("settings are empty")
		defaultSettings = &Settings{}
	}
	
	return defaultSettings
}

func Load() error {
	opts := settings.
		Options().
		EnvDefault().
		SetBasePath("./settings/defaults.yml").
		SetVarsMap(map[string]string {
			"DATA_CLUSTER":					"Data.Cluster",
			"DATA_DATABASE":				"Data.Database",
			"DATA_DEFAULTLIMIT":		"Data.DefaultLimit",
			"DATA_HOST":						"Data.Host",
			"DATA_MAXLIMIT":				"Data.MaxLimit",
			"DATA_QUERY":						"Data.Query",
			"DATA_PASS":						"Data.Pass",
			"DATA_PROTOCOL":				"Data.Protocol",
			"DATA_TIMEOUTSECONDS":	"Data.TimeoutSeconds",
			"DATA_USER":						"Data.User",

			"LOGGING_LEVEL":				"Logging.Level",
			"NAME":									"Name",
			"SERVER_ADDRESS":				"Server.Address",
		})
	
	if err := settings.Gather(opts, &defaultSettings); err != nil {
		l.Fatal().
			Stack().
			Err(err).
			Msg("unable to load settings")
		return err
	}

	l.Info().
		Object("Settings", defaultSettings).
		Msg("settings loaded")

	return nil
}