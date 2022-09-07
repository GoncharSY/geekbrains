package testbox

import (
	"fmt"
	"testing"
)

var box = New()
var cases = []struct {
	read  int
	write int
}{
	{
		write: 10,
		read:  90,
	}, {
		write: 50,
		read:  50,
	}, {
		write: 90,
		read:  10,
	},
}

func BenchmarkMutexSeries(b *testing.B) {
	for _, cs := range cases {
		b.Run(fmt.Sprintf("W%v-R%v", cs.write, cs.read), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				for i := 0; i < cs.write; i++ {
					box.SetValueM(i%100, 0.123)
				}
				for i := 0; i < cs.read; i++ {
					_ = box.GetValueM(i % 100)
				}
			}
		})
	}
}

func BenchmarkRWMutexSeries(b *testing.B) {
	for _, cs := range cases {
		b.Run(fmt.Sprintf("W%v-R%v", cs.write, cs.read), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				for i := 0; i < cs.write; i++ {
					box.SetValueRW(i%100, 0.123)
				}
				for i := 0; i < cs.read; i++ {
					_ = box.GetValueRW(i % 100)
				}
			}
		})
	}
}

func BenchmarkMutexParallel(b *testing.B) {
	for _, cs := range cases {
		b.Run(fmt.Sprintf("W%v-R%v", cs.write, cs.read), func(b *testing.B) {
			//b.SetParallelism(1000)
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					for i := 0; i < cs.write; i++ {
						box.SetValueM(i%100, 0.123)
					}
					for i := 0; i < cs.read; i++ {
						_ = box.GetValueM(i % 100)
					}
				}
			})
		})
	}
}

func BenchmarkRWMutexParallel(b *testing.B) {
	for _, cs := range cases {
		b.Run(fmt.Sprintf("W%v-R%v", cs.write, cs.read), func(b *testing.B) {
			//b.SetParallelism(1000)
			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					for i := 0; i < cs.write; i++ {
						box.SetValueRW(i%100, 0.123)
					}
					for i := 0; i < cs.read; i++ {
						_ = box.GetValueRW(i % 100)
					}
				}
			})
		})
	}
}
