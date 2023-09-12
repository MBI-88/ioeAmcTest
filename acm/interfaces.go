package acm





type acmBuilder interface {
	ProcessData(employees []string) map[string]float64
	LoadFile(path string) []string
}