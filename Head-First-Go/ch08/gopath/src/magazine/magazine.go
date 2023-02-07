package magazine

// 被导出的 struct 的名字必须以大写字母开头
// 被导出的 struct 字段也必须大写
type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

type Employee struct {
	Name   string
	Salary float64
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
