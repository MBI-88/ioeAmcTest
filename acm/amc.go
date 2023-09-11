package acm

import (
	"strings"
	"time"
)

type acmCompany struct {
	toPay map[string]int
}

func (ac *acmCompany) operation(payrange []acmRange, name string, start, end time.Time) {
	for i, acm := range payrange {
		if start.Compare(ac.formater(acm.EndRange)) > 0 {
			continue
		} else {
			if (start.Compare(ac.formater(acm.StartRange)) > -1) &&
				(start.Compare(ac.formater(acm.EndRange)) < 1) && (end.Compare(ac.formater(acm.EndRange)) < 1) {
				hours := end.Sub(start).Abs().Hours()
				ac.toPay[name] += int(hours) * acm.Payment
			} else if (start.Compare(ac.formater(acm.StartRange)) > -1) &&
				(start.Compare(ac.formater(acm.EndRange)) < 1) &&
				(end.Compare(ac.formater(acm.EndRange)) > 0) {
				acmEnd := ac.formater(acm.EndRange)
				hours := acmEnd.Sub(start).Abs().Hours()
				if hours == 0.0 {
					hours = 1.0
				}
				ac.toPay[name] += int(hours) * acm.Payment
				for _, acmk := range payrange[i+1:] {
					if (end.Compare(ac.formater(acmk.StartRange)) > -1) &&
						(end.Compare(ac.formater(acmk.EndRange)) < 1) {
						acmkStart := ac.formater(acmk.StartRange)
						hours := end.Sub(acmkStart).Abs().Hours()
						ac.toPay[name] += int(hours) * acmk.Payment
					}else {
						acmkStart := ac.formater(acmk.StartRange)
						acmkEnd := ac.formater(acmk.EndRange)
						hours := acmkEnd.Sub(acmkStart).Abs().Hours()
						ac.toPay[name] += int(hours) * acmk.Payment
					}
				}
			}

		}
	}

}

// Convert from string to time (hours)
func (ac acmCompany) formater(hours string) time.Time {
	value, err := time.Parse("10:06", hours)
	if err != nil {
		panic(err)
	}
	return value
}

// Process data from an array
func (ac *acmCompany) ProcessData(employees []string) map[string]int {
	
	for _, emp := range employees {
		name, data := strings.Split(emp, "=")[0], strings.Split(strings.Split(emp, "=")[1], ",")
		ac.toPay[name] = 0
		for _, dt := range data {
			day := dt[0:2]
			timeRange := strings.Split(dt[2:], "-")
			start := ac.formater(timeRange[0])
			end := ac.formater(timeRange[1])
			switch day {
			case "MO", "TU", "ME", "FR":
				payrange := rangePayments("Week")
				ac.operation(payrange, name, start, end)

			case "SA", "SU":
				payrange := rangePayments("Weekend")
				ac.operation(payrange, name, start, end)

			default:
				continue
			}

		}
	}
	return ac.toPay

}

// AcmCompany builder
func NewAcmCompany() acmBuilder {
	return &acmCompany{}
}
