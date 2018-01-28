package tests

import (
	"testing"
	"github.com/fronbasal/vertretungsplan/helpers"
	"github.com/stretchr/testify/assert"
)

func TestIservLogin(t *testing.T) {
	err, success := helpers.IServLogin("daniel.malik", "notmyactualpassword")
	assert.Nil(t, err)
	assert.True(t, success)
}
