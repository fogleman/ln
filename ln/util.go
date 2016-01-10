package ln

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"strconv"
	"time"
)

func Radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Degrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

func SavePNG(path string, im image.Image) error {
	fmt.Printf("Writing %s... ", path)
	defer fmt.Println("OK")
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, im)
}

func Median(items []float64) float64 {
	n := len(items)
	switch {
	case n == 0:
		return 0
	case n%2 == 1:
		return items[n/2]
	default:
		a := items[n/2-1]
		b := items[n/2]
		return (a + b) / 2
	}
}

func DurationString(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%d:%02d:%02d", h, m, s)
}

func NumberString(x float64) string {
	suffixes := []string{"", "k", "M", "G"}
	for _, suffix := range suffixes {
		if x < 1000 {
			return fmt.Sprintf("%.1f%s", x, suffix)
		}
		x /= 1000
	}
	return fmt.Sprintf("%.1f%s", x, "T")
}

func ParseFloats(items []string) []float64 {
	result := make([]float64, len(items))
	for i, item := range items {
		f, _ := strconv.ParseFloat(item, 64)
		result[i] = f
	}
	return result
}
