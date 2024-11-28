package stats

import (
	"github.com/stretchr/testify/assert"
	"github.com/venerasf/go-lua-libs/tests"
	"testing"
)

func TestApi(t *testing.T) {
	assert.NotZero(t, tests.RunLuaTestFile(t, Preload, "./test/test_api.lua"))
}
