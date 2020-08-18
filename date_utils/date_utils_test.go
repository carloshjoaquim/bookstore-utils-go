package date_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNow(t *testing.T) {
	now := GetNow()
	assert.NotNil(t, now)
}

func TestGetNowDBFormat(t *testing.T) {
	nowDbFormat := GetNowDBFormat()
	assert.NotNil(t, nowDbFormat)
}

func TestGetNowString(t *testing.T) {
	nowString := GetNowDBFormat()
	assert.NotNil(t, nowString)
}
