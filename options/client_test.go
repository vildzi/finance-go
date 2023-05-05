package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
	tests "github.com/vildzi/finance-go/testing"
)

func TestGetStraddle(t *testing.T) {

	iter := GetStraddle(tests.TestStraddleSymbol)
	success := iter.Next()
	assert.True(t, success)
	assert.Nil(t, iter.Err())
	assert.Equal(t, iter.Meta().UnderlyingSymbol, tests.TestStraddleSymbol)
}
