package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	finance "github.com/vildzi/finance-go"
	tests "github.com/vildzi/finance-go/testing"
)

func TestGetCryptoPair(t *testing.T) {
	tests.SetMarket(finance.MarketStateRegular)

	q, err := Get(tests.TestCryptoPairSymbol)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, finance.MarketStateRegular, q.MarketState)
	assert.Equal(t, tests.TestCryptoPairSymbol, q.Symbol)
}
