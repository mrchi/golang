package scanner

type Checker interface {
	Check(host string, port int) *Result
}

type Result struct {
	Vulnerable bool
	Details    string
}
