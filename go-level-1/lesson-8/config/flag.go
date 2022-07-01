package config

import (
	"flag"
	"strconv"
)

func getFlagConfig() map[string]string {
	var values = make(map[string]string)

	// Определим ожидаемые флаги
	for key, prp := range getDefaultPropertySet() {
		switch key {
		case "port":
			value, _ := strconv.Atoi(prp.Default)
			flag.Int(prp.FlagName, value, prp.Description)
		default:
			flag.String(prp.FlagName, prp.Default, prp.Description)
		}
	}

	// Соберем итоговую карту с флагами.
	flag.Parse()
	flag.Visit(func(flg *flag.Flag) {
		values[flg.Name] = flg.Value.String()
	})

	return values
}
