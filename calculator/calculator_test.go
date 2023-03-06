package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {

	eval := "3 + 2"

	tokens := tokenize(eval)
	assert.Equal(t, 3, len(tokens))

	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "+", tokens[1].String())
	assert.Equal(t, "2", tokens[2].String())

	eval = "3 + 2 + 1"

	tokens = tokenize(eval)
	assert.Equal(t, 5, len(tokens))

	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "+", tokens[1].String())
	assert.Equal(t, "2", tokens[2].String())
	assert.Equal(t, "+", tokens[3].String())
	assert.Equal(t, "1", tokens[4].String())

	eval = "3 - 1"

	tokens = tokenize(eval)
	assert.Equal(t, 3, len(tokens))

	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "-", tokens[1].String())
	assert.Equal(t, "1", tokens[2].String())

	eval = "3 * 2"

	tokens = tokenize(eval)
	assert.Equal(t, 3, len(tokens))

	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "*", tokens[1].String())
	assert.Equal(t, "2", tokens[2].String())

	eval = "3 / 2"

	tokens = tokenize(eval)
	assert.Equal(t, 3, len(tokens))

	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "/", tokens[1].String())
	assert.Equal(t, "2", tokens[2].String())
}

func TestPostfix(t *testing.T) {

	eval := "3 + 2"

	tokens := postfix(eval)

	// 3 2 +
	assert.Equal(t, 3, len(tokens))
	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "2", tokens[1].String())
	assert.Equal(t, "+", tokens[2].String())

	eval = "3 + 2 + 1"

	tokens = postfix(eval)

	// 3 2 + 1 +
	assert.Equal(t, 5, len(tokens))
	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "2", tokens[1].String())
	assert.Equal(t, "+", tokens[2].String())
	assert.Equal(t, "1", tokens[3].String())
	assert.Equal(t, "+", tokens[4].String())

	eval = "3 + 2 - 1"

	tokens = postfix(eval)

	// 3 2 + 1 -
	assert.Equal(t, 5, len(tokens))
	assert.Equal(t, "3", tokens[0].String())
	assert.Equal(t, "2", tokens[1].String())
	assert.Equal(t, "+", tokens[2].String())
	assert.Equal(t, "1", tokens[3].String())
	assert.Equal(t, "-", tokens[4].String())
}

func TestEvaluate(t *testing.T) {
	eval := "3 + 2"

	// 3 2 +

	rst, success := Evaluate(eval)

	assert.True(t, success)
	assert.Equal(t, 5, rst)

	eval = "124 + 200"

	rst, success = Evaluate(eval)

	assert.True(t, success)
	assert.Equal(t, 324, rst)

	eval = "200 - 100"

	rst, success = Evaluate(eval)

	assert.True(t, success)
	assert.Equal(t, 100, rst)

	eval = "20 * 10"

	rst, success = Evaluate(eval)

	assert.True(t, success)
	assert.Equal(t, 200, rst)

	eval = "20 / 10"

	rst, success = Evaluate(eval)

	assert.True(t, success)
	assert.Equal(t, 2, rst)

	eval = "2 + 3 + 4 - 3"

	rst, success = Evaluate(eval)

	assert.True(t, success)
	assert.Equal(t, 6, rst)
}

func TestPriority(t *testing.T) {
	eval := "3 + 2 * 4"

	rst, success := Evaluate(eval)
	assert.True(t, success)
	assert.Equal(t, 11, rst)
}
