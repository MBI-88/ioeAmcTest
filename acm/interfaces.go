package acm





type acmBuilder interface {
	ProcessData(employees []string) map[string]int
}