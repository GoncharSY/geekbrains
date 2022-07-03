package config

import (
	"encoding/json"
	"os"
	"strconv"
)

type jsonConfig struct {
	Port        int    `json:"port"`
	DbURL       string `json:"db_url"`
	JaegerURL   string `json:"jaeger_url"`
	SentryURL   string `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	SomeAppID   string `json:"some_app_id"`
	SomeAppKey  string `json:"some_app_key"`
}

// Загрузить данные из json-файла и разорабрать их в структуру.
func (j *jsonConfig) load(path string) error {
	var data []byte
	var err error

	if data, err = os.ReadFile(path); err != nil {
		return err
	}
	if err = json.Unmarshal(data, j); err != nil {
		return err
	}

	return nil
}

// Преобразовать данные структуры в мапу.
func (j *jsonConfig) toMap() map[string]string {
	var data = make(map[string]string)

	setValueToMap(j.Port, "port", &data)
	setValueToMap(j.DbURL, "db", &data)
	setValueToMap(j.JaegerURL, "jaeger", &data)
	setValueToMap(j.SentryURL, "sentry", &data)
	setValueToMap(j.KafkaBroker, "kafka", &data)
	setValueToMap(j.SomeAppID, "someId", &data)
	setValueToMap(j.SomeAppKey, "someKey", &data)
	return data
}

// Добавить значение в карту, если оно не исходное (нулевое).
func setValueToMap(value any, name string, data *map[string]string) {
	if *data == nil || name == "" {
		return
	}

	switch v := value.(type) {
	case int:
		if v != 0 {
			(*data)[name] = strconv.Itoa(v)
		}
	case string:
		if v != "" {
			(*data)[name] = v
		}
	}
}
