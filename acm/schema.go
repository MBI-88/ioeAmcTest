package acm

type acmRange struct {
	StartRange string 
	EndRange   string 
	Payment    float64    
}

var payments = make(map[string][]acmRange)

func rangePayments(key string) []acmRange {
	payments["Week"] = []acmRange{
		{
			StartRange: "00:01",
			EndRange:   "09:00",
			Payment:    25.0,
		},
		{
			StartRange: "09:01",
			EndRange:   "18:00",
			Payment:    15.0,
		},
		{
			StartRange: "18:01",
			EndRange:   "23:00",
			Payment:    20.0,
		},
	}
	payments["Weekend"] = []acmRange{
		{
			StartRange: "00:01",
			EndRange:   "09:00",
			Payment:    30.0,
		},
		{
			StartRange: "09:01",
			EndRange:   "18:00",
			Payment:    20.0,
		},
		{
			StartRange: "18:01",
			EndRange:   "23:00",
			Payment:    25.0,
		},
	}

	return payments[key]
}
