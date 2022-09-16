package main

import (
	"errors"
	"fmt"
	"reflect"
)

// Присвоить значения полей из мапы в структуру.
// in - обязательно должен быть структурой.
func assign(in any, values map[string]any) error {
	if structure, err := extractValue(in); err != nil {
		return err
	} else {
		return assignValues(structure, values)
	}
}

// Извлечь значение с которым можно работать (структуру).
func extractValue(in any) (reflect.Value, error) {
	var value reflect.Value

	if in == nil {
		return value, errors.New("'in' parameter is nil")
	}
	if value = reflect.ValueOf(in); value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return value, errors.New("'in' isn't structure pointer")
	}

	return value, nil
}

// Присвоить значения полей структуры из мапы.
// Не следует использовать данный метод вне текущего файла!
func assignValues(structure reflect.Value, values map[string]any) error {
	for i := 0; i < structure.NumField(); i++ {
		var field = structure.Field(i)
		var fname = structure.Type().Field(i).Name
		var value, isNew = values[fname]

		if !isNew || !field.CanSet() {
			continue
		}

		if err := assignField(field, value); err != nil {
			return fmt.Errorf("assigning field '%v': %w", fname, err)
		}
	}

	return nil
}

// Присвоить новое значение полю структуры. NIL-значения игнорируются.
// Если нельзя присвоить как есть, то будет попытка конвертировать.
func assignField(field reflect.Value, value any) error {
	if value == nil {
		return nil
	}

	var newV = reflect.ValueOf(value)
	var newT = newV.Type()
	var actT = field.Type()

	if newT.AssignableTo(actT) {
		field.Set(newV)
	} else if newV.CanConvert(actT) {
		field.Set(newV.Convert(actT))
	} else {
		return fmt.Errorf("%v can't be assigned to field with type %v", value, actT)
	}

	return nil
}
