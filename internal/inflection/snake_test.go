package inflection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnake(t *testing.T) {
	tests := map[string]string{
		"test":        "test",
		"Test":        "test",
		"Test123":     "test123",
		"TestTest":    "test_test",
		"TestTest123": "test_test123",
	}

	for camel, snake := range tests {
		assert.Equalf(t, snake, Snake(camel), "camel: %s snake: %s", camel, snake)
	}
}
