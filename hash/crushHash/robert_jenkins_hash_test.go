package crushHash

import (
	"testing"
)

var c *CrushHash

func init() {
	c = NewCrushHash(0)
}

func benchmarkCrushHash(i int64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		c.CrushHash32RJenkins1(i)
	}
}

func benchmarkCrushHash2(i, j int64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		c.CrushHash32RJenkins2(i, j)
	}
}

func benchmarkCrushHash3(i, j, g int64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		c.CrushHash32RJenkins3(i, j, g)
	}
}

func benchmarkCrushHash4(i, j, g, a int64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		c.CrushHash32RJenkins4(i, j, g, a)
	}
}

func benchmarkCrushHash5(i, j, g, a, d int64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		c.CrushHash32RJenkins5(i, j, g, a, d)
	}
}

func BenchmarkCrushHash1(t *testing.B) { benchmarkCrushHash(1231231231, t) }
func BenchmarkCrushHash2(t *testing.B) { benchmarkCrushHash2(1231231231, 123213123, t) }
func BenchmarkCrushHash3(t *testing.B) { benchmarkCrushHash3(1231231231, 123213123, 122121312, t) }
func BenchmarkCrushHash4(t *testing.B) {
	benchmarkCrushHash4(1231231231, 123213123, 122121312, 3213123, t)
}
func BenchmarkCrushHash5(t *testing.B) {
	benchmarkCrushHash5(1231231231, 123213123, 122121312, 3312312312, 321312, t)
}
