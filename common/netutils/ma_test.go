package netutils

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMovingAverage(t *testing.T) {

	ma := NewMovingAverage(time.Second, 5)

	// avg-value is zero
	assert.Equal(t, 0.0, ma.Avg())

	// async add 50 values
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ma.Add(1.)
		}()
	}
	wg.Wait()

	// check avg-value
	assert.InDelta(t, 50.0, ma.Avg(), 0.000001)

	// wait 1 second
	time.Sleep(time.Second)

	// async add 100 values
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ma.Add(1.)
		}()
	}
	wg.Wait()

	// check avg-value
	assert.InDelta(t, (50.+100.)/2, ma.Avg(), 0.000001)
}
