package util

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
)

type FieldInformation struct {
	Name       string
	ColumnName string
	Value      any
}

func ParseNoneZeroFields(v reflect.Value) ([]*FieldInformation, error) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("input must be a struct or a pointer to a struct")
	}

	var infos []*FieldInformation
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.IsZero() {
			continue
		}

		info := &FieldInformation{
			Name:  v.Type().Field(i).Name,
			Value: field.Interface(),
		}

		tagInfo := ParseGormTag(v.Type().Field(i).Tag.Get("gorm"))
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
