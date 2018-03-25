package logger

import (
	"errors"
	"fmt"
	"strconv"
)

func ForLog(v interface{}) (string, error) {
	switch t := v.(type) {
	case string:
		return string(v.(string)), nil
	case int:
		return strconv.Itoa(v.(int)), nil
	case uint:
		return strconv.FormatUint(v.(uint64), 10), nil
	case bool:
		return strconv.FormatBool(v.(bool)), nil
	case float32:
		return strconv.FormatFloat(v.(float64), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64), nil
	case error:
		return v.(error).Error(), nil
	default:
		_ = t
		return "", errors.New(fmt.Sprintf("unknown type %T", v))
	}
}
