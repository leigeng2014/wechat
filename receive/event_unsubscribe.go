package receive

//取消关注
/*
<xml>
	<ToUserName><![CDATA[gh_ba43c0866a80]]></ToUserName>
	<FromUserName><![CDATA[oIawIjx961XLlZrGQStq3QR39-Ps]]></FromUserName>
	<CreateTime>1524729089</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[unsubscribe]]></Event>
	<EventKey><![CDATA[]]></EventKey>
</xml>
*/
type EventUnSubscribe struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int64  `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	EventKey     string `json:"EventKey"`
}
