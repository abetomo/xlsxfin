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

func Ipmt(rate float64, per int, nper int, pv int, fv int, paymentFlag bool) int {
	if nper == 0 {
		return 0
	}

	if per == 0 {
		return 0
	}

	pmt := PmtFloat64(rate, nper, pv, fv, false)
	perSub1Float64 := float64(per - 1)

	n := 0.0
	if math.Abs(rate) > 0.5 {
		n = math.Pow(1.0+rate, perSub1Float64)
	} else {
		n = math.Exp(perSub1Float64 * math.Log(1.0+float64(rate)))
	}

	m := 0.0
	if rate <= -1 {
		m = math.Pow(1.0+rate, perSub1Float64) - 1
	} else {
		m = math.Exp(perSub1Float64*math.Log(1.0+float64(rate))) - 1
	}

	ip := -(float64(pv)*n*rate + float64(pmt)*m)
	if !paymentFlag {
		return round(ip)
	}
	return round(ip / (1.0 + rate))
}
