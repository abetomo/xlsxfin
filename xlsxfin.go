package xlsxfin

import "math"

func round(f float64) int {
	return int(math.Floor(f + .5))
}

func PmtFloat64(rate float64, nper int, pv int, fv int, paymentFlag bool) float64 {
	if nper == 0 {
		return 0
	}
	if rate == 0.0 {
		return -float64(pv+fv) / float64(nper)
	}

	pvif := math.Pow(1.0+rate, float64(nper))
	pmt := (rate / (pvif - 1)) * -(float64(pv)*pvif + float64(fv))

	if !paymentFlag {
		return pmt
	}
	return pmt / (1 + rate)
}

func Pmt(rate float64, nper int, pv int, fv int, paymentFlag bool) int {
	return round(PmtFloat64(rate, nper, pv, fv, paymentFlag))
}
