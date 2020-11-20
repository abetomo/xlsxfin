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

func IpmtFloat64(rate float64, per int, nper int, pv int, fv int, paymentFlag bool) float64 {
	if nper == 0 {
		return 0.0
	}

	if per == 0 {
		return 0.0
	}

	if rate < 0 {
		return 0.0
	}

	pmt := PmtFloat64(rate, nper, pv, fv, false)
	perSub1Float64 := float64(per - 1)

	n := 0.0
	if math.Abs(rate) > 0.5 {
		n = math.Pow(1.0+rate, perSub1Float64)
	} else {
		n = math.Exp(perSub1Float64 * math.Log(1.0+rate))
	}

	m := math.Exp(perSub1Float64*math.Log(1.0+rate)) - 1

	ip := -(float64(pv)*n*rate + pmt*m)
	if !paymentFlag {
		return ip
	}
	return ip / (1.0 + rate)
}

func Ipmt(rate float64, per int, nper int, pv int, fv int, paymentFlag bool) int {
	return round(IpmtFloat64(rate, per, nper, pv, fv, paymentFlag))
}

func FvFloat64(rate float64, nper int, pmt float64, pv int, paymentFlag bool) float64 {
	pvFloat64 := float64(pv)
	nperFloat64 := float64(nper)
	if rate == 0 {
		return -(pvFloat64 + pmt*nperFloat64)
	}
	term := math.Pow(1.0+rate, nperFloat64)
	if paymentFlag {
		return -(pvFloat64*term + (pmt*(1+rate)*(term-1))/rate)
	}
	return -(pvFloat64*term + (pmt*(term-1))/rate)
}

func Fv(rate float64, nper int, pmt float64, pv int, paymentFlag bool) int {
	return round(FvFloat64(rate, nper, pmt, pv, paymentFlag))
}

func Ppmt(rate float64, per int, nper int, pv int, fv int, paymentFlag bool) int {
	if per < 1 || per >= nper+1 {
		return 0
	}
	pmt := PmtFloat64(rate, nper, pv, fv, paymentFlag)
	ipmt := IpmtFloat64(rate, per, nper, pv, fv, paymentFlag)
	return round(pmt - ipmt)
}

func CumipmtFloat64(rate float64, nper int, pv int, start int, end int, paymentFlag bool) float64 {
	if rate <= 0.0 || nper <= 0 || pv <= 0 {
		return 0.0
	}

	if start < 1 || end < 1 || start > end {
		return 0.0
	}

	pmt := PmtFloat64(rate, nper, pv, 0, paymentFlag)
	interest := 0.0
	if start == 1 {
		if !paymentFlag {
			interest = -float64(pv)
			start++
		}
	}
	for i := start; i <= end; i++ {
		if paymentFlag {
			interest += FvFloat64(rate, i-2, pmt, pv, true) - pmt
		} else {
			interest += FvFloat64(rate, i-1, pmt, pv, false)
		}
	}
	return interest * rate
}

func Cumipmt(rate float64, nper int, pv int, start int, end int, paymentFlag bool) int {
	return round(CumipmtFloat64(rate, nper, pv, start, end, paymentFlag))
}
