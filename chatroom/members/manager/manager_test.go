package manager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	manager := Create()
	m, err := manager.Login(MemberName("alydnh"), nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Token())
}
