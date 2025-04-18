package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var num float64
	var x, y []float64
	i := 1

	for {
		_, err := fmt.Fscan(os.Stdin, &num)
		if err != nil {
			os.Exit(0)
		}

		x = append(x, float64(i))
		y = append(y, num)

		if len(x) < 2 {
			i++
			continue
		}

		a := LinearRegression(x, y)

		meanX := Median(x)
		meanY := Median(y)

		b := meanY - a*meanX

		nextY := a*float64(i) + b

		fmt.Printf("%.0f %.0f\n", nextY-10, nextY+10)

		i++
	}
}

func Median(arr []float64) float64 {
	sort.Float64s(arr)
	n := len(arr)

	if n%2 == 1 {
		return arr[n/2]
	}
	return (arr[n/2-1] + arr[n/2]) / 2
}

func LinearRegression(x, y []float64) float64 {
	n := float64(len(x))

	var sumX, sumY, sumXY, sumX2 float64
	for i := range x {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	den := n*sumX2 - sumX*sumX

	if den == 0 {
		return 0
	}
	return (n*sumXY - sumX*sumY) / den

}
