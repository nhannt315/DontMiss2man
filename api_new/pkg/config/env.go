package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/nhannt315/real_estate_api/pkg/errors"
)

// loadFromEnv os.Env の値を コンフィグにセットする。
// Envのキーは fieldの "env" タグで指定する。
// また、configの中に、要素が初期化されているmapがある場合、${map_key}_${"env" タグの値} をキーとして、map valueにos.Envの値をセットします。
func loadFromEnv(config interface{}) error {
	return setValuesFromEnv(config, "")
}

func setValuesFromEnv(obj interface{}, prefix string) error {
	configValue := reflect.Indirect(reflect.ValueOf(obj))

	if configValue.Kind() != reflect.Struct {
		return errors.New("invalid config, should be struct")
	}

	configType := configValue.Type()
	for i := 0; i < configType.NumField(); i++ {
		var (
			fieldStruct = configType.Field(i)
			field       = configValue.Field(i)
		)

		if err := setFiledValue(field, fieldStruct, prefix); err != nil {
			return err
		}
	}

	return nil
}

func setFiledValue(field reflect.Value, fieldStruct reflect.StructField, prefix string) error {
	if !field.CanAddr() || !field.CanInterface() {
		return nil
	}

	envName := fieldStruct.Tag.Get("env") // read configuration from shell env
	if prefix != "" {
		envName = prefix + "_" + envName
	}

	kind := field.Kind()

	value := os.Getenv(envName)
	if value != "" {
		return bindFieldFromEnvValue(kind, value, field, fieldStruct)
	}
	//
	err := setStructField(field, prefix)
	if err != nil {
		return err
	}

	if isBlank := reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()); isBlank {
		// Set default configuration if blank
		if err := setDefaultValue(field, fieldStruct, kind); err != nil {
			return err
		}
	}
	return nil
}

func setDefaultValue(field reflect.Value, fieldStruct reflect.StructField, kind reflect.Kind) error {
	if value := fieldStruct.Tag.Get("default"); value != "" {
		if err := yaml.Unmarshal([]byte(value), field.Addr().Interface()); err != nil {
			return errors.Errorf("fail to unmarshal. field: %s type: %s, default value: \"%s\"  ", fieldStruct.Name, kind, value)
		}
	}

	if fieldStruct.Tag.Get("required") == "true" {
		// return error if it is required but blank
		return errors.Errorf("%s is required, but blank", fieldStruct.Name)
	}
	return nil
}

func setStructField(field reflect.Value, prefix string) error {

	if field.Type().Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct {
		if field.IsNil() {
			return nil
		}
		return setValuesFromEnv(field.Interface(), prefix)
	}

	if field.Kind() == reflect.Struct {
		return setValuesFromEnv(field.Addr().Interface(), prefix)
	}

	if field.Kind() == reflect.Map {
		return setMapValue(field, prefix)
	}
	return nil
}

func setMapValue(field reflect.Value, prefix string) error {
	iter := field.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		keyString := stringValue(k)
		if v.Type().Kind() == reflect.Ptr &&
			v.Type().Elem().Kind() == reflect.Struct &&
			v.IsNil() {
			newValue := reflect.New(v.Type().Elem())
			field.SetMapIndex(k, newValue)

			v = newValue
		}

		if err := setStructField(v, joinPrefix(keyString, prefix)); err != nil {
			return err
		}
	}
	return nil
}

func stringValue(v reflect.Value) string {
	return fmt.Sprintf("%v", v.Interface())
}

func joinPrefix(p1, p2 string) string {
	upper1 := strings.ToUpper(p1)
	upper2 := strings.ToUpper(p2)
	if upper1 == "" {
		return upper2
	}
	if upper2 == "" {
		return upper1
	}

	return upper1 + "_" + upper2
}

func bindFieldFromEnvValue(kind reflect.Kind, value string, field reflect.Value, fieldStruct reflect.StructField) error {
	switch kind {
	case reflect.Bool:
		switch strings.ToLower(value) {
		case "", "0", "f", "false":
			field.Set(reflect.ValueOf(false))
		default:
			field.Set(reflect.ValueOf(true))
		}
	case reflect.String:
		field.Set(reflect.ValueOf(value))
	default:
		if err := yaml.Unmarshal([]byte(value), field.Addr().Interface()); err != nil {
			return errors.Errorf("fail to unmarshal. field: %s type: %s, env value: \"%s\"  ", fieldStruct.Name, kind.String(), value)
		}
	}
	return nil
}
