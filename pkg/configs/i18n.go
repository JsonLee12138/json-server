package configs

type I18nConfig struct {
	RootPath         string   `mapstructure:"root" json:"root" yaml:"root" toml:"root"`
	AcceptLanguages  []string `mapstructure:"accept-languages" json:"accept-languages" yaml:"accept-languages" toml:"accept-languages"`
	FormatBundleFile string   `mapstructure:"format-bundle-file" json:"format-bundle-file" yaml:"format-bundle-file" toml:"format-bundle-file"`
	DefaultLanguage  string   `mapstructure:"default-language" json:"default-language" yaml:"default-language" toml:"default-language"`
}
