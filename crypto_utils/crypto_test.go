package crypto_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMd5(t *testing.T) {
	md5 := GetMd5("TESTING")

	assert.NotNil(t, md5)
}

func TestSameStringGenerateSameMD5(t *testing.T) {
	firstMd5 := GetMd5("TESTING")
	secondM5 := GetMd5("TESTING")

	assert.NotNil(t, firstMd5)
	assert.NotNil(t, secondM5)
	assert.EqualValues(t, firstMd5, secondM5)
}
