package utils

import "reflect"

// Usage: utils.Copy(&expertiseTvInfoModel, &expertiseTvInfo)
func Copy(dst, src interface{}) {
	dstVal := reflect.ValueOf(dst).Elem()
	srcVal := reflect.ValueOf(src).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		dstVal.Field(i).Set(srcVal.Field(i))
	}
}
