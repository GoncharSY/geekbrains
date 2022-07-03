package config

// Получить набор полей конфигурации по умоланию.
func getDefaultPropertySet() map[string]property {
	return map[string]property{
		"port": {
			EnvName:     "port",
			FlagName:    "port",
			Default:     "1000",
			Description: "Порт приложения на сервере",
		},
		"someId": {
			EnvName:     "some_id",
			FlagName:    "someId",
			Default:     "default_id",
			Description: "Какой-то идентификатор в приложении",
		},
		"someKey": {
			EnvName:     "some_key",
			FlagName:    "someKey",
			Default:     "default_key",
			Description: "Какой-то ключ внутри приложения",
		},
		"db": {
			EnvName:     "db_url",
			FlagName:    "db",
			Default:     "postgres://default:default@db:1111",
			Description: "Адрес базы данных приложения",
		},
		"jaeger": {
			EnvName:     "jaeger_url",
			FlagName:    "jaeger",
			Default:     "http://jaeger:16686",
			Description: "Адрес системы трассировки Jaeger",
		},
		"sentry": {
			EnvName:     "sentry_url",
			FlagName:    "sentry",
			Default:     "http://sentry:9000",
			Description: "Адрес монитора ошибок Sentry",
		},
		"kafka": {
			EnvName:     "kafka_broker",
			FlagName:    "kafka",
			Default:     "kafka:9092",
			Description: "Адрес брокера обмена сообщениями Kafka",
		},
	}
}
