package msg

import (
	"context"
	"strings"

	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/config"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/db"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/log"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/grpc-etcdv3/getcdv3"
	pbChat "github.com/OpenIMSDK/Open-IM-Server/pkg/proto/msg"
	pbCommon "github.com/OpenIMSDK/Open-IM-Server/pkg/proto/sdk_ws"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/utils"
)

func TagSendMessage(operationID string, user *db.User, recvID, content string, senderPlatformID int32) {
	log.NewInfo(operationID, utils.GetSelfFuncName(), "args: ", user.UserID, recvID, content)
	var req pbChat.SendMsgReq
	var msgData pbCommon.MsgData
	msgData.SendID = user.UserID
	msgData.RecvID = recvID
	msgData.ContentType = constant.Custom
	msgData.SessionType = constant.SingleChatType
	msgData.MsgFrom = constant.UserMsgType
	msgData.Content = []byte(content)
	msgData.SenderFaceURL = user.FaceURL
	msgData.SenderNickname = user.Nickname
	msgData.Options = map[string]bool{}
	msgData.Options[constant.IsSenderConversationUpdate] = false
	msgData.Options[constant.IsSenderNotificationPush] = false
	msgData.CreateTime = utils.GetCurrentTimestampByMill()
	msgData.ClientMsgID = utils.GetMsgID(user.UserID)
	msgData.SenderPlatformID = senderPlatformID
	req.MsgData = &msgData
	req.OperationID = operationID
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImMsgName, operationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		return
	}

	client := pbChat.NewMsgClient(etcdConn)
	respPb, err := client.SendMsg(context.Background(), &req)
	if err != nil {
		log.NewError(operationID, utils.GetSelfFuncName(), "send msg failed", err.Error())
		return
	}
	if respPb.ErrCode != 0 {
		log.NewError(operationID, utils.GetSelfFuncName(), "send tag msg failed ", respPb)
	}
}
