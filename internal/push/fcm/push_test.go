package fcm

import (
	"fmt"
	"testing"

	"github.com/OpenIMSDK/Open-IM-Server/internal/push"

	"github.com/stretchr/testify/assert"
)

func Test_Push(t *testing.T) {
	t.SkipNow()

	offlinePusher := NewFcm()
	resp, err := offlinePusher.Push([]string{"test_uid"}, "test", "test", "12321", push.PushOpts{})
	assert.Nil(t, err)
	fmt.Println(resp)
}
