package message

// Text 文本消息
type Link struct {
	CommonToken
	Content     CDATA `xml:"Content"`
	Title       CDATA `xml:"Title"`
	Description CDATA `xml:"Description"`
	Url         CDATA `xml:"Url"`
	MsgId       int64 `xml:"MsgId"`
}

// NewLink 初始化链接文本
func NewLink(MsgId int64, Title string, Description string, Url string) *Link {
	link := new(Link)
	link.MsgId = MsgId
	link.Title = CDATA(Title)
	link.Description = CDATA(Description)
	link.Url = CDATA(Url)
	return link
}
