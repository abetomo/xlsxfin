package xlsxfin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const DELTA = 0.0001

func ExamplePmtFloat64() {
	v := PmtFloat64(0.3, 36, 100_000, 0, false)
	fmt.Println(v)
	// Output:-30002.37243823623
}

func TestPmtFloat64(t *testing.T) {
	type testArgs struct {
		rate        float64
		nper        int
		pv          int
		fv          int
		paymentFlag bool
	}

	type testData struct {
		args     testArgs
		expected float64
	}

	t.Run("nper is 0", func(t *testing.T) {
		actual := PmtFloat64(0.3, 0, 100_000, 0, false)
		expected := 0.0
		assert.InDelta(t, expected, actual, DELTA)
	})

	t.Run("rate is 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.0, 36, 100_000, 0, false},
				expected: -2777.777778,
			},
			{
				args:     testArgs{0.0, 36, 100_000, 0, true},
				expected: -2777.777778,
			},
			{
				args:     testArgs{0.0, 36, 100_000, 1_000, false},
				expected: -2805.555556,
			},
			{
				args:     testArgs{0.0, 36, 100_000, 1_000, true},
				expected: -2805.555556,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := PmtFloat64(
				args.rate,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.InDelta(t, testCase.expected, actual, DELTA, testCase)
		}
	})

	t.Run("rate > 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.3, 36, 100_000, 0, false},
				expected: -30_002.372438,
			},
			{
				args:     testArgs{0.3, 36, 100_000, 0, true},
				expected: -23078.748029,
			},
			{
				args:     testArgs{0.3, 36, 100_000, 1_000, false},
				expected: -30_002.396163,
			},
			{
				args:     testArgs{0.3, 36, 100_000, 1_000, true},
				expected: -23078.766279,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := PmtFloat64(
				args.rate,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.InDelta(t, testCase.expected, actual, DELTA, testCase)
		}
	})
}

func ExamplePmt() {
	v := Pmt(0.3, 36, 100_000, 0, false)
	fmt.Println(v)
	// Output:-30002
}

func TestPmt(t *testing.T) {
	type testArgs struct {
		rate        float64
		nper        int
		pv          int
		fv          int
		paymentFlag bool
	}

	type testData struct {
		args     testArgs
		expected int
	}

	t.Run("nper is 0", func(t *testing.T) {
		actual := Pmt(0.3, 0, 100_000, 0, false)
		expected := 0
		assert.Equal(t, expected, actual)
	})

	t.Run("rate is 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.0, 36, 100_000, 0, false},
				expected: -2778,
			},
			{
				args:     testArgs{0.0, 36, 100_000, 0, true},
				expected: -2778,
			},
			{
				args:     testArgs{0.0, 36, 100_000, 1_000, false},
				expected: -2806,
			},
			{
				args:     testArgs{0.0, 36, 100_000, 1_000, true},
				expected: -2806,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Pmt(
				args.rate,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual)
		}
	})

	t.Run("rate > 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.3, 36, 100_000, 0, false},
				expected: -30_002,
			},
			{
				args:     testArgs{0.3, 36, 100_000, 0, true},
				expected: -23_079,
			},
			{
				args:     testArgs{0.3, 36, 100_000, 1_000, false},
				expected: -30_002,
			},
			{
				args:     testArgs{0.3, 36, 100_000, 1_000, true},
				expected: -23_079,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Pmt(
				args.rate,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual)
		}
	})
}

func ExampleIpmtFloat64() {
	v := IpmtFloat64(0.1, 2, 36, 800_000, 0, false)
	fmt.Println(v)
	// Output: -79732.55489453014
}

func TestIpmtFloat64(t *testing.T) {
	type testArgs struct {
		rate        float64
		per         int
		nper        int
		pv          int
		fv          int
		paymentFlag bool
	}

	type testData struct {
		args     testArgs
		expected float64
	}

	t.Run("nper is 0", func(t *testing.T) {
		actual := IpmtFloat64(0.3, 3, 0, 100_000, 0, false)
		expected := 0.0
		assert.InDelta(t, expected, actual, DELTA)
	})

	t.Run("per is 0", func(t *testing.T) {
		actual := IpmtFloat64(0.3, 0, 36, 100_000, 0, false)
		expected := 0.0
		assert.InDelta(t, expected, actual, DELTA)
	})

	t.Run("rate < 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{-1.0, 24, 36, 100_000, 0, false},
				expected: 0.0,
			},
			{
				args:     testArgs{-1.0, 24, 36, 100_000, 0, true},
				expected: 0.0,
			},
			{
				args:     testArgs{-1.0, 24, 36, 100_000, 1_000, false},
				expected: 0.0,
			},
			{
				args:     testArgs{-1.0, 24, 36, 100_000, 1_000, true},
				expected: 0.0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := IpmtFloat64(
				args.rate,
				args.per,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.InDelta(t, testCase.expected, actual, DELTA, testCase)
		}
	})

	t.Run("rate is 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.0, 24, 36, 100_000, 0, false},
				expected: 0.0,
			},
			{
				args:     testArgs{0.0, 24, 36, 100_000, 0, true},
				expected: 0.0,
			},
			{
				args:     testArgs{0.0, 24, 36, 100_000, 1_000, false},
				expected: 0.0,
			},
			{
				args:     testArgs{0.0, 24, 36, 100_000, 1_000, true},
				expected: 0.0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := IpmtFloat64(
				args.rate,
				args.per,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.InDelta(t, testCase.expected, actual, DELTA, testCase)
		}
	})

	t.Run("rate > 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 2, 36, 800_000, 0, false},
				expected: -79_732.554895,
			},
			{
				args:     testArgs{0.6, 2, 36, 800_000, 0, false},
				expected: -479_999.987086,
			},
			{
				args:     testArgs{0.1, 2, 36, 800_000, 0, true},
				expected: -72_484.140813,
			},
			{
				args:     testArgs{0.6, 2, 36, 800_000, 0, true},
				expected: -299_999.991929,
			},
			{
				args:     testArgs{0.1, 2, 36, 800_000, 1_000, false},
				expected: -79_732.220588,
			},
			{
				args:     testArgs{0.6, 2, 36, 800_000, 1_000, false},
				expected: -479_999.987069,
			},
			{
				args:     testArgs{0.1, 2, 36, 800_000, 1_000, true},
				expected: -72_483.836898,
			},
			{
				args:     testArgs{0.6, 2, 36, 800_000, 1_000, true},
				expected: -299_999.991918,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := IpmtFloat64(
				args.rate,
				args.per,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.InDelta(t, testCase.expected, actual, DELTA, testCase)
		}
	})
}

func ExampleIpmt() {
	v := Ipmt(0.1, 2, 36, 800_000, 0, false)
	fmt.Println(v)
	// Output: -79733
}

func TestIpmt(t *testing.T) {
	type testArgs struct {
		rate        float64
		per         int
		nper        int
		pv          int
		fv          int
		paymentFlag bool
	}

	type testData struct {
		args     testArgs
		expected int
	}

	t.Run("nper is 0", func(t *testing.T) {
		actual := Ipmt(0.3, 3, 0, 100_000, 0, false)
		expected := 0
		assert.Equal(t, expected, actual)
	})

	t.Run("per is 0", func(t *testing.T) {
		actual := Ipmt(0.3, 0, 36, 100_000, 0, false)
		expected := 0
		assert.Equal(t, expected, actual)
	})

	t.Run("rate is 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.0, 24, 36, 100_000, 0, false},
				expected: 0,
			},
			{
				args:     testArgs{0.0, 24, 36, 100_000, 0, true},
				expected: 0,
			},
			{
				args:     testArgs{0.0, 24, 36, 100_000, 1_000, false},
				expected: 0,
			},
			{
				args:     testArgs{0.0, 24, 36, 100_000, 1_000, true},
				expected: 0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Ipmt(
				args.rate,
				args.per,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("rate > 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 2, 36, 800_000, 0, false},
				expected: -79_733,
			},
			{
				args:     testArgs{0.6, 2, 36, 800_000, 0, false},
				expected: -480_000,
			},
			{
				args:     testArgs{0.1, 2, 36, 800_000, 0, true},
				expected: -72_484,
			},
			{
				args:     testArgs{0.6, 2, 36, 800_000, 0, true},
				expected: -300_000,
			},
			{
				args:     testArgs{0.1, 2, 36, 800_000, 1_000, false},
				expected: -79_732,
			},
			{
				args:     testArgs{0.6, 2, 36, 800_000, 1_000, false},
				expected: -480_000,
			},
			{
				args:     testArgs{0.1, 2, 36, 800_000, 1_000, true},
				expected: -72_484,
			},
			{
				args:     testArgs{0.6, 2, 36, 800_000, 1_000, true},
				expected: -300_000,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Ipmt(
				args.rate,
				args.per,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})
}

func ExampleFvFloat64() {
	v := FvFloat64(0.1, 12, 10_000.0, 0, false)
	fmt.Println(v)
	// Output: -213842.83767210023
}

func TestFvFloat64(t *testing.T) {
	type testArgs struct {
		rate        float64
		nper        int
		pmt         float64
		pv          int
		paymentFlag bool
	}

	type testData struct {
		args     testArgs
		expected float64
	}

	t.Run("rate is 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.0, 12, 10_000.0, 0, false},
				expected: -120_000.0,
			},
			{
				args:     testArgs{0.0, 12, 10_000.0, 1_000, false},
				expected: -121_000.0,
			},
			{
				args:     testArgs{0.0, 12, 10_000.0, 0, true},
				expected: -120_000.0,
			},
			{
				args:     testArgs{0.0, 12, 10_000.0, 1_000, true},
				expected: -121_000.0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := FvFloat64(
				args.rate,
				args.nper,
				args.pmt,
				args.pv,
				args.paymentFlag,
			)
			assert.InDelta(t, testCase.expected, actual, DELTA, testCase)
		}
	})

	t.Run("rate > 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 12, 10_000.0, 0, false},
				expected: -213_842.837672,
			},
			{
				args:     testArgs{0.1, 12, 10_000.0, 1_000, false},
				expected: -216_981.266049,
			},
			{
				args:     testArgs{0.1, 12, 10_000.0, 0, true},
				expected: -235_227.121439,
			},
			{
				args:     testArgs{0.1, 12, 10_000.0, 1_000, true},
				expected: -238365.549816,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := FvFloat64(
				args.rate,
				args.nper,
				args.pmt,
				args.pv,
				args.paymentFlag,
			)
			assert.InDelta(t, testCase.expected, actual, DELTA, testCase)
		}
	})
}

func ExampleFv() {
	v := Fv(0.1, 12, 10_000.0, 0, false)
	fmt.Println(v)
	// Output: -213843
}

func TestFv(t *testing.T) {
	type testArgs struct {
		rate        float64
		nper        int
		pmt         float64
		pv          int
		paymentFlag bool
	}

	type testData struct {
		args     testArgs
		expected int
	}

	t.Run("rate is 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.0, 12, 10_000.0, 0, false},
				expected: -120_000,
			},
			{
				args:     testArgs{0.0, 12, 10_000.0, 1_000, false},
				expected: -121_000,
			},
			{
				args:     testArgs{0.0, 12, 10_000.0, 0, true},
				expected: -120_000,
			},
			{
				args:     testArgs{0.0, 12, 10_000.0, 1_000, true},
				expected: -121_000,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Fv(
				args.rate,
				args.nper,
				args.pmt,
				args.pv,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("rate > 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 12, 10_000.0, 0, false},
				expected: -213_843,
			},
			{
				args:     testArgs{0.1, 12, 10_000.0, 1_000, false},
				expected: -216_981,
			},
			{
				args:     testArgs{0.1, 12, 10_000.0, 0, true},
				expected: -235_227,
			},
			{
				args:     testArgs{0.1, 12, 10_000.0, 1_000, true},
				expected: -238_366,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Fv(
				args.rate,
				args.nper,
				args.pmt,
				args.pv,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})
}

func ExamplePpmtFloat64() {
	v := PpmtFloat64(0.1, 12, 36, 800_000, 0, false)
	fmt.Printf("%.5f\n", v)
	// Output: -7630.52098
}

func TestPpmtFloat64(t *testing.T) {
	type testArgs struct {
		rate        float64
		per         int
		nper        int
		pv          int
		fv          int
		paymentFlag bool
	}

	type testData struct {
		args     testArgs
		expected float64
	}

	t.Run("per < 1", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 0, 10, 800_000, 0, false},
				expected: 0.0,
			},
			{
				args:     testArgs{0.1, -1, 10, 800_000, 0, false},
				expected: 0.0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := PpmtFloat64(
				args.rate,
				args.per,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("per >= nper + 1", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 11, 10, 800_000, 0, false},
				expected: 0.0,
			},
			{
				args:     testArgs{0.1, 15, 10, 800_000, 0, false},
				expected: 0.0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := PpmtFloat64(
				args.rate,
				args.per,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("per >= 1 && per < nper + 1", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 12, 36, 800_000, 0, false},
				expected: -7_630.520984,
			},
			{
				args:     testArgs{0.1, 12, 36, 800_000, 1_000, false},
				expected: -7_640.059135,
			},
			{
				args:     testArgs{0.1, 12, 36, 800_000, 0, true},
				expected: -6_936.837258,
			},
			{
				args:     testArgs{0.1, 12, 36, 800_000, 1_000, true},
				expected: -6_945.508305,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := PpmtFloat64(
				args.rate,
				args.per,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.InDelta(t, testCase.expected, actual, DELTA, testCase)
		}
	})
}

func ExamplePpmt() {
	v := Ppmt(0.1, 12, 36, 800_000, 0, false)
	fmt.Println(v)
	// Output: -7631
}

func TestPpmt(t *testing.T) {
	type testArgs struct {
		rate        float64
		per         int
		nper        int
		pv          int
		fv          int
		paymentFlag bool
	}

	type testData struct {
		args     testArgs
		expected int
	}

	t.Run("per < 1", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 0, 10, 800_000, 0, false},
				expected: 0,
			},
			{
				args:     testArgs{0.1, -1, 10, 800_000, 0, false},
				expected: 0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Ppmt(
				args.rate,
				args.per,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("per >= nper + 1", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 11, 10, 800_000, 0, false},
				expected: 0,
			},
			{
				args:     testArgs{0.1, 15, 10, 800_000, 0, false},
				expected: 0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Ppmt(
				args.rate,
				args.per,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("per >= 1 && per < nper + 1", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 12, 36, 800_000, 0, false},
				expected: -7_631,
			},
			{
				args:     testArgs{0.1, 12, 36, 800_000, 1_000, false},
				expected: -7_640,
			},
			{
				args:     testArgs{0.1, 12, 36, 800_000, 0, true},
				expected: -6_937,
			},
			{
				args:     testArgs{0.1, 12, 36, 800_000, 1_000, true},
				expected: -6_946,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Ppmt(
				args.rate,
				args.per,
				args.nper,
				args.pv,
				args.fv,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})
}

func ExampleCumipmtFloat64() {
	v := CumipmtFloat64(0.1, 36, 800_000, 6, 12, true)
	fmt.Println(v)
	// Output: -488961.5711288557
}

func TestCumipmtFloat64(t *testing.T) {
	type testArgs struct {
		rate        float64
		nper        int
		pv          int
		start       int
		end         int
		paymentFlag bool
	}

	type testData struct {
		args     testArgs
		expected float64
	}

	t.Run("rate <= 0.0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0, 36, 800_000, 6, 12, false},
				expected: 0.0,
			},
			{
				args:     testArgs{-1, -1, 10, 800_000, 0, false},
				expected: 0.0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := CumipmtFloat64(
				args.rate,
				args.nper,
				args.pv,
				args.start,
				args.end,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("nper <= 0.0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 0, 800_000, 6, 12, false},
				expected: 0.0,
			},
			{
				args:     testArgs{0.1, -1, 800_000, 6, 12, false},
				expected: 0.0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := CumipmtFloat64(
				args.rate,
				args.nper,
				args.pv,
				args.start,
				args.end,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("pv <= 0.0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 36, 0, 6, 12, false},
				expected: 0.0,
			},
			{
				args:     testArgs{0.1, 36, -1, 6, 12, false},
				expected: 0.0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := CumipmtFloat64(
				args.rate,
				args.nper,
				args.pv,
				args.start,
				args.end,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("start < 1 || end < 1 || start > end", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 36, 800_000, 0, 12, false},
				expected: 0.0,
			},
			{
				args:     testArgs{0.1, 36, 800_000, 1, 0, false},
				expected: 0.0,
			},
			{
				args:     testArgs{0.1, 36, 800_000, 10, 9, false},
				expected: 0.0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := CumipmtFloat64(
				args.rate,
				args.nper,
				args.pv,
				args.start,
				args.end,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("Calculate", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 36, 800_000, 6, 12, true},
				expected: -488_961.571129,
			},
			{
				args:     testArgs{0.1, 36, 800_000, 6, 12, false},
				expected: -537_857.728242,
			},
			{
				args:     testArgs{0.1, 36, 800_000, 1, 12, true},
				expected: -777_183.8112556307,
			},
			{
				args:     testArgs{0.1, 36, 800_000, 1, 12, false},
				expected: -934_902.192381,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := CumipmtFloat64(
				args.rate,
				args.nper,
				args.pv,
				args.start,
				args.end,
				args.paymentFlag,
			)
			assert.InDelta(t, testCase.expected, actual, DELTA, testCase)
		}
	})
}

func ExampleCumipmt() {
	v := Cumipmt(0.1, 36, 800_000, 6, 12, true)
	fmt.Println(v)
	// Output: -488962
}

func TestCumipmt(t *testing.T) {
	type testArgs struct {
		rate        float64
		nper        int
		pv          int
		start       int
		end         int
		paymentFlag bool
	}

	type testData struct {
		args     testArgs
		expected int
	}

	t.Run("rate <= 0.0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0, 36, 800_000, 6, 12, false},
				expected: 0,
			},
			{
				args:     testArgs{-1, -1, 10, 800_000, 0, false},
				expected: 0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Cumipmt(
				args.rate,
				args.nper,
				args.pv,
				args.start,
				args.end,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("nper <= 0.0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 0, 800_000, 6, 12, false},
				expected: 0,
			},
			{
				args:     testArgs{0.1, -1, 800_000, 6, 12, false},
				expected: 0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Cumipmt(
				args.rate,
				args.nper,
				args.pv,
				args.start,
				args.end,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("pv <= 0.0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 36, 0, 6, 12, false},
				expected: 0,
			},
			{
				args:     testArgs{0.1, 36, -1, 6, 12, false},
				expected: 0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Cumipmt(
				args.rate,
				args.nper,
				args.pv,
				args.start,
				args.end,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("start < 1 || end < 1 || start > end", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 36, 800_000, 0, 12, false},
				expected: 0,
			},
			{
				args:     testArgs{0.1, 36, 800_000, 1, 0, false},
				expected: 0,
			},
			{
				args:     testArgs{0.1, 36, 800_000, 10, 9, false},
				expected: 0,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Cumipmt(
				args.rate,
				args.nper,
				args.pv,
				args.start,
				args.end,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})

	t.Run("Calculate", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 36, 800_000, 6, 12, true},
				expected: -488_962,
			},
			{
				args:     testArgs{0.1, 36, 800_000, 6, 12, false},
				expected: -537_858,
			},
			{
				args:     testArgs{0.1, 36, 800_000, 1, 12, true},
				expected: -777_184,
			},
			{
				args:     testArgs{0.1, 36, 800_000, 1, 12, false},
				expected: -934_902,
			},
		}
		for _, testCase := range testCases {
			args := testCase.args
			actual := Cumipmt(
				args.rate,
				args.nper,
				args.pv,
				args.start,
				args.end,
				args.paymentFlag,
			)
			assert.Equal(t, testCase.expected, actual, testCase)
		}
	})
}
