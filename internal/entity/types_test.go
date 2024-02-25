package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotification(t *testing.T) {
	n := NewNotification("Hello test", "news", "test@gamil.com")
	assert.Equal(t, "news", n.Type)
	assert.Equal(t, "test@gamil.com", n.Recipient)
}
