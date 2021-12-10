package xreflect

import (
    "reflect"
)

/**
 * @author Rancho
 * @date 2021/12/11
 */

func IsPtr(i interface{}) bool {
    return reflect.ValueOf(i).Type().Kind() == reflect.Ptr
}

// DeReferType will return t itself or t.Elem() when t is a pointer
func DeReferType(t reflect.Type) reflect.Type {
    if t.Kind() == reflect.Ptr {
        return t.Elem()
    }
    return t
}

func IsBasedOn(t reflect.Type, expected reflect.Kind) bool {
    t = DeReferType(t)
    return t.Kind() == expected
}
