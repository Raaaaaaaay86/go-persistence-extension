package util

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
)

type FieldInformation struct {
	Name        string
	ColumnName  string
	Value       any
	ReflectKind reflect.Kind
	ReflectType reflect.Type
}

func ParseTargetField(entity any, targetType reflect.Type) (*FieldInformation, error) {
	fields, err := ParseNoneZeroFields(reflect.ValueOf(entity))
	if err != nil {
		return nil, err
	}

	if len(fields) == 0 {
		return nil, errors.New("no target field specified")
	}

	var targetField *FieldInformation
	for i, count := 0, 0; i < len(fields); i++ {
		if fields[i].ReflectType.Name() == targetType.Name() {
			if count > 0 {
				return nil, errors.New("ambiguous target field. there are at most one target field allowed")
			}
			targetField = fields[i]
			count++
		}
	}

	if targetField == nil {
		return nil, fmt.Errorf("no target field found. expected field type (%s)", targetType.Name())
	}

	return targetField, nil
}

func ParseNoneZeroFields(v reflect.Value) ([]*FieldInformation, error) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("input must be a struct or a pointer to a struct")
	}

	var infos []*FieldInformation
	fields, values := FlattenStruct(v)
	for i := 0; i < len(fields); i++ {
		field := fields[i]
		value := values[i]

		if value.IsZero() {
			continue
		}

		info := &FieldInformation{
			Name:        field.Name,
			Value:       value.Interface(),
			ReflectKind: field.Type.Kind(),
			ReflectType: reflect.TypeOf(value.Interface()),
		}

		tagInfo := ParseGormTag(field.Tag.Get("gorm"))
		if tagInfo != nil {
			info.ColumnName = tagInfo.ColumnName
		} else {
			info.ColumnName = ToSnakeCase(info.Name)
		}

		infos = append(infos, info)
	}

	return infos, nil
}

type GormTagInfo struct {
	ColumnName string
}

func ParseGormTag(tag string) *GormTagInfo {
	info := &GormTagInfo{}
	for _, attr := range strings.Split(tag, ",") {
		attr = strings.TrimSpace(attr)
		split := strings.Split(attr, ":")
		switch split[0] {
		case "column":
			info.ColumnName = split[1]
		}
	}
	return nil
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func FlattenStruct(v reflect.Value) ([]*reflect.StructField, []*reflect.Value) {
	var fields []*reflect.StructField
	var values []*reflect.Value

	for i := 0; i < v.NumField(); i++ {
		structField := v.Type().Field(i)
		value := v.Field(i)

		if !structField.IsExported() {
			continue
		}

		if structField.Type.Kind() == reflect.Struct {
			switch structField.Type.Name() {
			case reflect.TypeOf(time.Time{}).Name(),
				reflect.TypeOf(gorm.DeletedAt{}).Name():
				fields = append(fields, &structField)
				values = append(values, &value)
			default:
				fs, vs := FlattenStruct(v.Field(i))
				fields = append(fields, fs...)
				values = append(values, vs...)
			}
			continue
		}

		fields = append(fields, &structField)
		values = append(values, &value)
	}

	return fields, values
}
