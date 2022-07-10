package config

import (
	"os"
	"strings"
)

// Получить конфигурацию из переменных окружения.
func getEnvConfig() map[string]string {
	var values = map[string]string{}

	// Получаю значения переменных.
	for key, prop := range getDefaultPropertySet() {
		name := makeEnvName(prop.EnvName)
		if value, ok := os.LookupEnv(name); ok {
			values[key] = value
		}
	}

	return values
}

// Создать имя переменной окружения.
func makeEnvName(name string) string {
	var separator = "_"
	var prefix = "WEBAPP"
	var suffix = strings.ToUpper(name)
	return strings.Join([]string{prefix, suffix}, separator)
}
