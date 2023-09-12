package tests

import (
	"acm-payments/acm"
	"testing"
)

func TestProcessData(t *testing.T) {
	acmTest := acm.NewAcmCompany()
	test := make(map[string]int)
	sample := []string{
		"JOSE=MO08:00-19:00",
		"JUAN=TH12:00-17:00,FR09:01-16:00",
		"ANA=SA09:00-13:00,MO08:10-16:00",
	}

	test["JOSE"] = 179
	test["JUAN"] = 179
	test ["ANA"] = 209

	restul := acmTest.ProcessData(sample)

	for key, _ := range restul {
		if int(restul[key]) != test[key] {
			t.Errorf("Name %s resutl => %d and test => %d\n",key, int(restul[key]), test[key])
		}
	}

}
