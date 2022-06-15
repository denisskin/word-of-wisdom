package netutils

import (
	"sync"
	"time"
)

type MovingAverage struct {
	step int64 // duration in nano sec
	size int64 // buffer size

	buf []float64  // buffer
	t   int64      // last push time (in steps)
	n   int64      //
	sum float64    //
	mx  sync.Mutex //
}

// NewMovingAverage makes moving average aggregator
func NewMovingAverage(step time.Duration, size int) *MovingAverage {
	return &MovingAverage{
		step: int64(step),
		size: int64(size),
		buf:  make([]float64, size),
		t:    timestamp(int64(step)),
		n:    1,
	}
}

// Avg returns actual moving average value
func (m *MovingAverage) Avg() float64 {
	return m.Add(0)
}

// Add aggregates new value and returns actual moving average value
func (m *MovingAverage) Add(v float64) float64 {
	t := timestamp(m.step)

	m.mx.Lock()
	defer m.mx.Unlock()

	if t-m.t > m.size {
		for i := range m.buf { // clear buf
			m.buf[i] = 0
		}
		m.t, m.n, m.sum = t, m.size, 0
	}
	for ; m.t < t; m.t++ {
		i := int((m.t + 1) % m.size)
		m.sum -= m.buf[i]
		m.buf[i] = 0
		if m.n < m.size {
			m.n++
		}
	}
	if v != 0 {
		m.sum += v
		m.buf[int(t%m.size)] += v
	}
	return m.sum / float64(m.n) // avg-value
}

func timestamp(step int64) int64 {
	return time.Now().UnixNano() / step
}
