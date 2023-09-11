package acm

type acmRange struct {
	StartRange string 
	EndRange   string 
	Payment    int    
}

var payments = make(map[string][]acmRange)

func rangePayments(key string) []acmRange {
	payments["Week"] = []acmRange{
		{
			StartRange: "00:01",
			EndRange:   "09:00",
			Payment:    25,
		},
		{
			StartRange: "09:01",
			EndRange:   "18:00",
			Payment:    15,
		},
		{
			StartRange: "18:01",
			EndRange:   "23:00",
			Payment:    20,
		},
	}
	payments["Weekend"] = []acmRange{
		{
			StartRange: "00:01",
			EndRange:   "09:00",
			Payment:    30,
		},
		{
			StartRange: "09:01",
			EndRange:   "18:00",
			Payment:    20,
		},
		{
			StartRange: "18:01",
			EndRange:   "23:00",
			Payment:    25,
		},
	}

	return payments[key]
}
