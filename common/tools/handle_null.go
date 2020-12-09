package tools

import (
	"reflect"
)

type VariantStruct struct {
	Variant interface{}
}

func GetStructTags(str interface{}, tagName string) map[string]string {
	strType := reflect.TypeOf(str)

	numFields := strType.NumField()

	tags := make(map[string]string)

	for i := 0; i < numFields; i++ {
		field := strType.Field(i)
		tags[field.Name] = field.Tag.Get(tagName)
	}
	return tags
}
