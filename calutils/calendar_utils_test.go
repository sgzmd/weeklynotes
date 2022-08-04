package calutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFirstMonday(t *testing.T) {
	dt := FirstMondayOfYear(2022)
	expected, _ := time.Parse("2006-01-02", "2022-01-03")
	assert.Equal(t, expected, dt)

	d2 := FirstMondayOfYear(2020)
	expected2, _ := time.Parse("2006-01-02", "2020-01-06")
	assert.Equal(t, expected2, d2)
}
