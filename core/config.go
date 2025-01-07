package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"path/filepath"
)

type ConfigInterface interface {
	Bind(instance any) error
}
type Config struct {
	instance *viper.Viper
	opts     ConfigOptions
}

type ConfigOptions struct {
	BasePath  string
	FileName  string
	FileType  string
	WatchAble bool
	OnChange  func(e fsnotify.Event)
}

func DefaultConfigOptions() ConfigOptions {
	return ConfigOptions{
		BasePath:  "./config",
		FileName:  "config",
		FileType:  "yaml",
		WatchAble: true,
		OnChange:  nil,
	}
}

func NewConfig(optsArr ...ConfigOptions) (*Config, error) {
	var opts ConfigOptions
	if len(optsArr) == 0 {
		opts = DefaultConfigOptions()
	}
	instance, err := CreateConfig(optsArr...)
	if err != nil {
		return nil, err
	}
	return &Config{
		instance: instance,
		opts:     opts,
	}, nil
}

func (c *Config) Bind(instance any) (err error) {
	if err = c.instance.Unmarshal(&instance); err != nil {
		return fmt.Errorf("❌ Failed to unmarshal(bind) config: %w", err)
	}
	if c.opts.WatchAble {
		c.instance.WatchConfig()
		c.instance.OnConfigChange(func(e fsnotify.Event) {
			err = c.instance.Unmarshal(&instance)
			if err != nil {
				fmt.Println(err)
			}
			if c.opts.OnChange != nil {
				c.opts.OnChange(e)
			}
		})
	}
	return nil
}

func CreateConfig(optsArr ...ConfigOptions) (*viper.Viper, error) {
	var opts ConfigOptions
	if len(optsArr) == 0 {
		opts = DefaultConfigOptions()
	}
	configPaths := getConfigFilePaths(optsArr...)
	if len(configPaths) == 0 {
		return nil, fmt.Errorf("❌ No valid configuration files found")
	}
	viper.SetConfigType(opts.FileType)
	v := viper.New()
	for _, configPath := range configPaths {
		v.SetConfigFile(configPath)
		if err := v.MergeInConfig(); err != nil {
			return nil, fmt.Errorf("❌ Error merging config file %s: %w", configPath, err)
		}
	}
	v.AutomaticEnv()
	return v, nil
}

func getConfigFilePaths(optsArr ...ConfigOptions) (configFiles []string) {
	var opts ConfigOptions
	if len(optsArr) == 0 {
		opts = DefaultConfigOptions()
	}
	env := Mode()
	fileNames := []string{
		opts.FileName,
		fmt.Sprintf("%s.local", opts.FileName),
		fmt.Sprintf("%s.%s", opts.FileName, env),
		fmt.Sprintf("%s.%s.local", opts.FileName, env),
	}
	switch env {
	case DevMode:
		fileNames = append(fileNames, fmt.Sprintf("%s.dev", opts.FileName))
		fileNames = append(fileNames, fmt.Sprintf("%s.dev.local", opts.FileName))
		fileNames = append(fileNames, fmt.Sprintf("%s.development", opts.FileName))
		fileNames = append(fileNames, fmt.Sprintf("%s.development.local", opts.FileName))
	case ProMode:
		fileNames = append(fileNames, fmt.Sprintf("%s.pro", opts.FileName))
		fileNames = append(fileNames, fmt.Sprintf("%s.pro.local", opts.FileName))
		fileNames = append(fileNames, fmt.Sprintf("%s.prod", opts.FileName))
		fileNames = append(fileNames, fmt.Sprintf("%s.prod.local", opts.FileName))
		fileNames = append(fileNames, fmt.Sprintf("%s.production", opts.FileName))
		fileNames = append(fileNames, fmt.Sprintf("%s.production.local", opts.FileName))
	case TestMode:
		fileNames = append(fileNames, fmt.Sprintf("%s.test", opts.FileName))
		fileNames = append(fileNames, fmt.Sprintf("%s.test.local", opts.FileName))
	}
	for _, fileName := range fileNames {
		file := filepath.Join(opts.BasePath, fmt.Sprintf("%s.%s", fileName, opts.FileType))
		if isDir, exists, _ := Exists(file); exists && !isDir {
			configFiles = append(configFiles, file)
		}
	}
	return configFiles
}
