package tool

import (
	"fmt"
	"strconv"
)

// ToInt 尝试将任何类型转换为int
func AnyToInt(value any) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		return int(v), nil
	case float32:
		return int(v), nil // 注意这会截断小数部分
	case float64:
		return int(v), nil // 注意这会截断小数部分
	case string:
		// 尝试将字符串转换为int
		return strconv.Atoi(v)
	default:
		return 0, fmt.Errorf("无法将类型 %T 转换为 int", value)
	}
}
