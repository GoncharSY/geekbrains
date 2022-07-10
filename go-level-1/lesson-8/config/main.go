package config

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

type property struct {
	EnvName     string
	FlagName    string
	Default     string
	Description string
}

type config struct {
	props map[string]property
	flags map[string]string
	envrs map[string]string
}

// Получить значение параметра из конфигурации.
func (c *config) GetValue(name string) (string, bool) {
	var prop property
	var value string
	var ok bool

	if prop, ok = c.props[name]; !ok {
		return "", false
	}

	if value, ok = c.flags[name]; ok {
		return value, true
	}

	if value, ok = c.envrs[name]; ok {
		return value, true
	}

	return prop.Default, prop.Default != ""
}

// Получить конфигурацию из флагов.
func (c *config) GetFlagConfig() map[string]string {
	return c.flags
}

// Получить конфигурацию из переменных окружения.
func (c *config) GetEnvConfig() map[string]string {
	return c.envrs
}

// Получить конфигурацию по умолчанию.
func (c *config) GetDefaultConfig(full bool) map[string]string {
	var values = make(map[string]string)

	for name, prop := range c.props {
		if prop.Default != "" || full {
			values[name] = prop.Default
		}
	}

	return values
}

// Получить итоговую конфигурацию.
func (c *config) GetResultConfig(full bool) map[string]string {
	var values = make(map[string]string)

	for name := range c.props {
		if value, ok := c.GetValue(name); ok || full {
			values[name] = value
		}
	}

	return values
}

// Проверить корректность параметров конфигурации.
func (c *config) Validate() map[string]error {
	var errs = make(map[string]error)

	for name, value := range c.GetResultConfig(true) {
		switch name {
		case "port":
			if port, err := strconv.Atoi(value); err != nil {
				errs[name] = fmt.Errorf("incorrect port: %w", err)
			} else if port < 0 || port > 65535 {
				errs[name] = errors.New("incorrect port (0-65535 only)")
			}
		case "db":
			fallthrough
		case "kafka":
			fallthrough
		case "sentry":
			fallthrough
		case "jaeger":
			if _, err := url.Parse(value); err != nil {
				errs[name] = fmt.Errorf("incorrect url of '%v': %w", name, err)
			}
		}
	}

	return errs
}

// Создать новую конфигурацию.
func New() *config {
	return &config{
		props: getDefaultPropertySet(),
		flags: getFlagConfig(),
		envrs: getEnvConfig(),
	}
}
