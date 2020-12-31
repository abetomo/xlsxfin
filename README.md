# ExcelFormulasFinance

## Example

```go
package main

import (
	"fmt"
	"github.com/abetomo/xlsxfin"
)

func main() {
	fmt.Println(xlsxfin.Pmt(0.08/12, 10, 1_000_000, 0, false))
	// Output:-103703
}
```

## Functions

## [PMT](https://support.microsoft.com/en-us/office/pmt-function-0214da64-9a63-4996-bc20-214433fa6441)

```go
func Pmt(rate float64, nper int, pv int, fv int, paymentFlag bool) int
```

## [IPMT](https://support.microsoft.com/en-us/office/ipmt-function-5cce0ad6-8402-4a41-8d29-61a0b054cb6f)

```go
func Ipmt(rate float64, per int, nper int, pv int, fv int, paymentFlag bool) int
```

## [FV](https://support.microsoft.com/en-us/office/fv-function-2eef9f44-a084-4c61-bdd8-4fe4bb1b71b3)

```go
func Fv(rate float64, nper int, pmt float64, pv int, paymentFlag bool) int
```

## [PPMT](https://support.microsoft.com/en-us/office/ppmt-function-c370d9e3-7749-4ca4-beea-b06c6ac95e1b)

```go
func Ppmt(rate float64, per int, nper int, pv int, fv int, paymentFlag bool) int
```

## [CUMIPMT](https://support.microsoft.com/en-us/office/cumipmt-function-61067bb0-9016-427d-b95b-1a752af0e606)

```go
func Cumipmt(rate float64, nper int, pv int, start int, end int, paymentFlag bool) int
```
