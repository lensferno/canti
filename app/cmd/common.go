package cmd

import (
	"github.com/urfave/cli/v2"
	"reflect"
)

func _readCliParam[T any](c *cli.Context, key string, defaultVal T) T {
	if !c.IsSet(key) {
		return defaultVal
	}

	var (
		v  T
		ok bool
	)

	switch reflect.TypeOf(defaultVal).Kind() {
	case reflect.Bool:
		v, ok = any(c.Bool(key)).(T)
	case reflect.String:
		v, ok = any(c.String(key)).(T)
	case reflect.Int:
		v, ok = any(c.Int(key)).(T)
	case reflect.Int64:
		v, ok = any(c.Int64(key)).(T)
	}

	if ok {
		return v
	} else {
		return defaultVal
	}
}
