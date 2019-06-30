package manager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvailable(t *testing.T) {
	manager := Create()
	_, err := manager.Create([]MemberName{"alydnh", "realgang", "tracy"}, nil)
	assert.Nil(t, err)
	_, err = manager.Create([]MemberName{"alydnh", "realgang"}, nil)
	assert.Nil(t, err)

	rooms, err := manager.Available("alydnh")
	assert.Nil(t, err)
	assert.Len(t, rooms, 2)

	rooms, err = manager.Available("tracy")
	assert.Nil(t, err)
	assert.Len(t, rooms, 1)

}
