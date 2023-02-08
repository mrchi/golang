package fuel

type Liters float64
type Gallons float64
type Milliliters float64

// 定义的方法与此类型值都关联
// 按照惯例，Go开发者通常使用小写的接收器类型名称的首字母作为名称。
// 方法和类型必须定义在同一包中
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}
