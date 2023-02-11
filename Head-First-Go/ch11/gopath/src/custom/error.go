package custom

import (
	"fmt"
)

// 自定义错误类型
type OverheatError float64

// 具有一个返回 string 的 Error 方法，它就满足 error 接口
func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %.2f degrees!", o)
}

// 使用 error 接口作为函数返回值
func CheckTemperature(actual float64, safe float64) error {
	if actual > safe {
		return OverheatError(actual - safe)
	} else {
		return nil
	}
}
