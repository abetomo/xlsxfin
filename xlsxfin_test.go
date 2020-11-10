package xlsxfin

import (
	"math"
	"testing"
)

func checkForRoundingError(actual float64, expected float64) bool {
	return math.Abs(actual-expected) < 0.0001
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
		if !checkForRoundingError(actual, expected) {
			t.Fatalf("got: %f\nwant: %f\n", actual, expected)
		}
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
			if !checkForRoundingError(actual, testCase.expected) {
				t.Fatalf("testCase: %#v\ngot: %f\n", testCase, actual)
			}
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
			if !checkForRoundingError(actual, testCase.expected) {
				t.Fatalf("testCase: %#v\ngot: %f\n", testCase, actual)
			}
		}
	})
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
		if actual != expected {
			t.Fatalf("got: %d\nwant: %d\n", actual, expected)
		}
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
			if actual != testCase.expected {
				t.Fatalf("testCase: %#v\ngot: %d\n", testCase, actual)
			}
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
			if actual != testCase.expected {
				t.Fatalf("testCase: %#v\ngot: %d\n", testCase, actual)
			}
		}
	})
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
		if actual != expected {
			t.Fatalf("got: %d\nwant: %d\n", actual, expected)
		}
	})

	t.Run("per is 0", func(t *testing.T) {
		actual := Ipmt(0.3, 0, 36, 100_000, 0, false)
		expected := 0
		if actual != expected {
			t.Fatalf("got: %d\nwant: %d\n", actual, expected)
		}
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
			if actual != testCase.expected {
				t.Fatalf("testCase: %#v\ngot: %d\n", testCase, actual)
			}
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
			if actual != testCase.expected {
				t.Fatalf("testCase: %#v\ngot: %d\n", testCase, actual)
			}
		}
	})
}

func TestFv(t *testing.T) {
	type testArgs struct {
		rate        float64
		nper        int
		pmt         int
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
				args:     testArgs{0.0, 12, 10_000, 0, false},
				expected: -120_000,
			},
			{
				args:     testArgs{0.0, 12, 10_000, 1_000, false},
				expected: -121_000,
			},
			{
				args:     testArgs{0.0, 12, 10_000, 0, true},
				expected: -120_000,
			},
			{
				args:     testArgs{0.0, 12, 10_000, 1_000, true},
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
			if actual != testCase.expected {
				t.Fatalf("testCase: %#v\ngot: %d\n", testCase, actual)
			}
		}
	})

	t.Run("rate > 0", func(t *testing.T) {
		testCases := []testData{
			{
				args:     testArgs{0.1, 12, 10_000, 0, false},
				expected: -213_843,
			},
			{
				args:     testArgs{0.1, 12, 10_000, 1_000, false},
				expected: -216_981,
			},
			{
				args:     testArgs{0.1, 12, 10_000, 0, true},
				expected: -235_227,
			},
			{
				args:     testArgs{0.1, 12, 10_000, 1_000, true},
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
			if actual != testCase.expected {
				t.Fatalf("testCase: %#v\ngot: %d\n", testCase, actual)
			}
		}
	})
}
