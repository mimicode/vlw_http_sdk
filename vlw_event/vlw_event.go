package vlw_event

// 解析vlw_event 消息

type Common struct {
	Event string `json:"event"`
}

// Enable 启动插件
type Enable struct {
	Common `json:",inline"`
	Data   string `json:"data"`
}

// StopUsing 停止插件
type StopUsing struct {
	Common `json:",inline"`
	Data   string `json:"data"`
}

// Uninstall 卸载插件
type Uninstall struct {
	Common `json:",inline"`
	Data   string `json:"data"`
}

// Login 新的账号登录成功/下线
type Login struct {
	Common `json:",inline"`
	Data   LoginData `json:"data"`
}

type LoginData struct {
	LoginType int    `json:"login_type"` // 0 登录成功 / 1 即将离线
	Wxid      string `json:"Wxid"`       //登录/离线 的Wxid
	RobotType int    `json:"robot_type"` // 来源微信类型 0 正常微信 / 1 企业微信
}

// EventGroupChat 群消息事件
type EventGroupChat struct {
	Common `json:",inline"`
	Data   EventGroupChatData `json:"data"`
}

type EventGroupChatData struct {
	RobotWxID     string `json:"robot_wxid"`      //机器人Wxid
	MsgType       int    `json:"type"`            //1/文本消息 3/图片消息 34/语音消息  42/名片消息  43/视频 47/动态表情 48/地理位置  49/分享链接  2001/红包  2002/小程序  2003/群邀请 更多请参考sdk模块常量值
	FromGroup     string `json:"from_group"`      //来源群号
	FromGroupName string `json:"from_group_name"` //来源群名称
	FromWxID      string `json:"from_wxid"`       //具体发消息的群成员id
	FromName      string `json:"from_name"`       //具体发消息的群成员昵称
	Msg           string `json:"msg"`             //消息内容
	MsgSource     string `json:"msg_source"`      //可选，附带属性（群消息有艾特人员时，返回被艾特数组）
	ClientID      string `json:"clientid"`        //企业微信可用
	RobotType     string `json:"robot_type"`      //来源微信类型 0 正常微信 / 1 企业微信
}

// EventPrivateChat 私聊消息事件
type EventPrivateChat struct {
	Common `json:",inline"`
	Data   EventPrivateChatData `json:"data"`
}

type EventPrivateChatData struct {
	RobotWxID string `json:"robot_wxid"` //机器人Wxid
	MsgType   int    `json:"type"`       //1/文本消息 3/图片消息 34/语音消息  42/名片消息  43/视频 47/动态表情 48/地理位置  49/分享链接  2001/红包  2002/小程序  2003/群邀请 更多请参考sdk模块常量值
	FromWxID  string `json:"from_wxid"`  //来源用户ID
	FromName  string `json:"from_name"`  //来源用户昵称
	Msg       string `json:"msg"`        //消息内容
	ClientID  string `json:"clientid"`   //企业微信可用
	RobotType string `json:"robot_type"` //来源微信类型 0 正常微信 / 1 企业微信
}

// EventDeviceCallback 设备回调事件
type EventDeviceCallback struct {
	Common `json:",inline"`
	Data   EventDeviceCallbackData `json:"data"`
}
type EventDeviceCallbackData struct {
	RobotWxID string `json:"robot_wxid"` //机器人Wxid
	MsgType   int    `json:"type"`       //1/文本消息 3/图片消息 34/语音消息  42/名片消息  43/视频 47/动态表情 48/地理位置  49/分享链接  2001/红包  2002/小程序  2003/群邀请 更多请参考sdk模块常量值
	Msg       string `json:"msg"`        //消息内容
	ToWxID    string `json:"to_wxid"`    //接收用户ID
	ToName    string `json:"to_name"`    //接收用户昵称
	ClientID  string `json:"clientid"`   //企业微信可用
	RobotType string `json:"robot_type"` //来源微信类型 0 正常微信 / 1 企业微信
}

// EventFrieneVerify 好友请求事件
type EventFrieneVerify struct {
	Common `json:",inline"`
	Data   EventFrieneVerifyData `json:"data"`
}
type EventFrieneVerifyData struct {
	RobotWxID string `json:"robot_wxid"` //机器人Wxid
	MsgType   int    `json:"type"`       //3/微信号,手机搜索添加 12/QQ号 13/通讯录 14/群内添加 17/名片推荐 18/附近 29/摇一摇 30/扫一扫
	FromWxID  string `json:"from_wxid"`  //请求者wxid
	FromName  string `json:"from_name"`  //请求者昵称
	V1        string `json:"v1"`
	V2        string `json:"v2"`
	JsonMsg   string `json:"json_msg"`   //文本型, , 好友验证信息JSON(群内添加时，包含群id) (名片推荐添加时，包含推荐人id及昵称) (微信号、手机号搜索等添加时,具体JSON结构请查看日志）
	RobotType string `json:"robot_type"` //来源微信类型 0 正常微信 / 1 企业微信
}

// EventGroupNameChange 群名称变动事件
type EventGroupNameChange struct {
	Common `json:",inline"`
	Data   EventGroupNameChangeData `json:"data"`
}
type EventGroupNameChangeData struct {
	RobotWxID string `json:"robot_wxid"` //机器人Wxid
	FromGroup string `json:"from_group"` // 群号
	NewName   string `json:"new_name"`   // 新群名
	OldName   string `json:"old_name"`   //旧群名
	FromName  string `json:"from_name"`  //操作者昵称
	ClientID  string `json:"clientid"`   //企业微信可用
	RobotType string `json:"robot_type"` //来源微信类型 0 正常微信 / 1 企业微信
}

// EventGroupMemberAdd 群成员增加事件（新人进群）
type EventGroupMemberAdd struct {
	Common `json:",inline"`
	Data   EventGroupMemberAddData `json:"data"`
}

type EventGroupMemberAddData struct {
	RobotWxID     string `json:"robot_wxid"`      //机器人Wxid
	FromGroup     string `json:"from_group"`      // 群号
	FromGroupName string `json:"from_group_name"` // 群名
	Guest         string `json:"guest"`           //新人
	Inviter       string `json:"inviter"`         // 邀请者
	ClientID      string `json:"clientid"`        //企业微信可用
	RobotType     string `json:"robot_type"`      //来源微信类型 0 正常微信 / 1 企业微信
}

//EventGroupMemberDecrease 群成员减少
type EventGroupMemberDecrease struct {
	Common `json:",inline"`
	Data   EventGroupMemberDecreaseData `json:"data"`
}
type EventGroupMemberDecreaseData struct {
	RobotWxID     string `json:"robot_wxid"`      //机器人Wxid
	FromGroup     string `json:"from_group"`      // 群号
	FromGroupName string `json:"from_group_name"` // 群名
	ToWxID        string `json:"to_wxid"`         // 被操作者wxid
	ToName        string `json:"to_name"`         //被操作者昵称
	Time          string `json:"time"`            //操作时间
	ClientID      string `json:"clientid"`        //企业微信可用
	RobotType     string `json:"robot_type"`      //来源微信类型 0 正常微信 / 1 企业微信
}

// EventInvitedInGroup 被邀请入群事件，运行这里（企业微信不传递此事件）
type EventInvitedInGroup struct {
	Common `json:",inline"`
	Data   EventInvitedInGroupData `json:"data"`
}

type EventInvitedInGroupData struct {
	RobotWxID string `json:"robot_wxid"` //机器人Wxid
	FromWxid  string `json:"from_wxid"`  // 邀请机器人入群的id
	Msg       string `json:"msg"`        //消息内容
	RobotType string `json:"robot_type"` //来源微信类型 0 正常微信 / 1 企业微信
}

// EventQRcodePayment 面对面收款事件（二维码收款时，运行这里）（企业微信不传递此事件）
type EventQRcodePayment struct {
	Common `json:",inline"`
	Data   EventQRcodePaymentData `json:"data"`
}

type EventQRcodePaymentData struct {
	RobotWxID          string `json:"robot_wxid"`           //机器人Wxid
	ToWxID             string `json:"to_wxid"`              // 收款者wxid
	PayerWxID          string `json:"payer_wxid"`           //  付款者wxid
	PayerNickname      string `json:"payer_nickname"`       //  付款者昵称
	PayerPayID         string `json:"payer_pay_id"`         //
	ReceivedMoneyIndex string `json:"received_money_index"` //
	Money              string `json:"money"`                //
	TotalMoney         string `json:"total_money"`          //
	Remark             string `json:"remark"`               //
	SceneDesc          string `json:"scene_desc"`           //
	Scene              string `json:"scene"`                //  -1 扫码后退出 / 1 已扫码，未付款 / 2 付款完成 / 3 付款完成后的日志 / 4 付款完成后的日志（商家版），这里重点说明一下，如要做收款播报，只需要判断等于2或者等3的时候就可以了
	Time string `json:"time"` //时间
}
