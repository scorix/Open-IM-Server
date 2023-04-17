package token_verify

import (
	"testing"

	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParseToken(t *testing.T) {
	token, _, err := CreateToken("", constant.IOSPlatformID)
	require.NoError(t, err)

	_, err = GetClaimFromToken(token)
	assert.NoError(t, err)
}
