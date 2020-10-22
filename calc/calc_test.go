package calc_test

import (
	"testing"

	"boostcamp-git.th-service.co.in/root/demogopipeline/calc"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 3, calc.Add(1, 2))
}

func TestDel(t *testing.T) {
	assert.Equal(t, 1, calc.Del(3, 2))
}
