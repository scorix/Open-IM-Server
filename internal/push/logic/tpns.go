package logic

import (
	tpns "github.com/OpenIMSDK/Open-IM-Server/internal/push/sdk/tpns-server-sdk-go/go"
	"github.com/OpenIMSDK/Open-IM-Server/internal/push/sdk/tpns-server-sdk-go/go/auth"
	"github.com/OpenIMSDK/Open-IM-Server/internal/push/sdk/tpns-server-sdk-go/go/common"
	"github.com/OpenIMSDK/Open-IM-Server/internal/push/sdk/tpns-server-sdk-go/go/req"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/config"
)

var badgeType = -2
var iosAcceptId = auth.Auther{AccessID: config.Config.Push.Tpns.Ios.AccessID, SecretKey: config.Config.Push.Tpns.Ios.SecretKey}

func IOSAccountListPush(accounts []string, title, content, jsonCustomContent string) {
	var iosMessage = tpns.Message{
		Title:   title,
		Content: content,
		IOS: &tpns.IOSParams{
			Aps: &tpns.Aps{
				BadgeType: &badgeType,
				Sound:     "default",
				Category:  "INVITE_CATEGORY",
			},
			CustomContent: jsonCustomContent,
			//CustomContent: `"{"key\":\"value\"}"`,
		},
	}
	pushReq, reqBody, err := req.NewListAccountPush(accounts, iosMessage)
	if err != nil {
		return
	}
	iosAcceptId.Auth(pushReq, auth.UseSignAuthored, iosAcceptId, reqBody)
	common.PushAndGetResult(pushReq)
}
