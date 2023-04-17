package push

import "github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"

var PushTerminal = []int{constant.IOSPlatformID, constant.AndroidPlatformID, constant.WebPlatformID}

type OfflinePusher interface {
	Push(userIDList []string, title, detailContent, operationID string, opts PushOpts) (resp string, err error)
}

type PushOpts struct {
	Signal        Signal
	IOSPushSound  string
	IOSBadgeCount bool
	Data          string
}

type Signal struct {
	ClientMsgID string
}
