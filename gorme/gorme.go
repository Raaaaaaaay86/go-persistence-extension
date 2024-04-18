package gorme

import (
	"reflect"

	"github.com/raaaaaaaay86/go-persistence-extension/contract"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme/util"
)

func ToQueryMap(object any) contract.QueryMap {
	m := make(map[string]interface{})

	v := reflect.ValueOf(object)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return m
	}

	fields, values := util.FlattenStruct(v)
	for i := 0; i < len(fields); i++ {
		field := fields[i]
		value := values[i]

		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				continue
			}
			dereferenced := value.Elem()
			value = &dereferenced
		}

		tagInfo := util.ParseGormTag(field.Tag.Get("gorm"))
		if tagInfo != nil {
			m[tagInfo.ColumnName] = value.Interface()
		} else {
			m[util.ToSnakeCase(field.Name)] = value.Interface()
		}
	}

	return m
}