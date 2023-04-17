package msg

import (
	"context"
	"time"

	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/db"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/log"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/token_verify"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/proto/msg"
	commonPb "github.com/OpenIMSDK/Open-IM-Server/pkg/proto/sdk_ws"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/utils"
)

func (rpc *rpcChat) DelMsgList(_ context.Context, req *commonPb.DelMsgListReq) (*commonPb.DelMsgListResp, error) {
	log.NewInfo(req.OperationID, utils.GetSelfFuncName(), "req: ", req.String())
	resp := &commonPb.DelMsgListResp{}
	select {
	case rpc.delMsgCh <- deleteMsg{
		UserID:      req.UserID,
		OpUserID:    req.OpUserID,
		SeqList:     req.SeqList,
		OperationID: req.OperationID,
	}:
	case <-time.After(1 * time.Second):
		resp.ErrCode = constant.ErrSendLimit.ErrCode
		resp.ErrMsg = constant.ErrSendLimit.ErrMsg
		return resp, nil
	}
	log.NewInfo(req.OperationID, utils.GetSelfFuncName(), "resp: ", resp.String())
	return resp, nil
}
func (rpc *rpcChat) DelSuperGroupMsg(_ context.Context, req *msg.DelSuperGroupMsgReq) (*msg.DelSuperGroupMsgResp, error) {
	log.NewInfo(req.OperationID, utils.GetSelfFuncName(), "req: ", req.String())
	if !token_verify.CheckAccess(req.OpUserID, req.UserID) {
		log.NewError(req.OperationID, "CheckAccess false ", req.OpUserID, req.UserID)
		return &msg.DelSuperGroupMsgResp{ErrCode: constant.ErrAccess.ErrCode, ErrMsg: constant.ErrAccess.ErrMsg}, nil
	}
	resp := &msg.DelSuperGroupMsgResp{}
	groupMaxSeq, err := db.DB.GetGroupMaxSeq(req.GroupID)
	if err != nil {
		log.NewError(req.OperationID, "GetGroupMaxSeq false ", req.OpUserID, req.UserID, req.GroupID)
		resp.ErrCode = constant.ErrDB.ErrCode
		resp.ErrMsg = err.Error()
		return resp, nil
	}
	err = db.DB.SetGroupUserMinSeq(req.GroupID, req.UserID, groupMaxSeq)
	if err != nil {
		log.NewError(req.OperationID, "SetGroupUserMinSeq false ", req.OpUserID, req.UserID, req.GroupID)
		resp.ErrCode = constant.ErrDB.ErrCode
		resp.ErrMsg = err.Error()
		return resp, nil
	}
	log.NewInfo(req.OperationID, utils.GetSelfFuncName(), "resp: ", resp.String())
	return resp, nil
}
