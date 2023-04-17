package callback

import (
	"Open_IM/pkg/common/constant"
	sdk_ws "Open_IM/pkg/proto/sdk_ws"

	"github.com/golang/protobuf/proto"
)

func GetContent(msg *sdk_ws.MsgData) string {
	if msg.ContentType >= constant.NotificationBegin && msg.ContentType <= constant.NotificationEnd {
		var tips sdk_ws.TipsComm
		_ = proto.Unmarshal(msg.Content, &tips)
		//marshaler := jsonpb.Marshaler{
		//	OrigName:     true,
		//	EnumsAsInts:  false,
		//	EmitDefaults: false,
		//}
		content := tips.JsonDetail
		return content
	} else {
		return string(msg.Content)
	}
}
