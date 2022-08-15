package work_weixin_robot

// MsgType  消息类型
type MsgType string

const (
	// TextMsgType 文本类型
	TextMsgType MsgType = "text"
	// MarkdownMsgType markdown类型
	MarkdownMsgType MsgType = "markdown"
	// ImageMsgType 图片类型
	ImageMsgType MsgType = "image"
	// NewsMsgTye 图文类型
	NewsMsgTye MsgType = "news"
	// FileMsgType 文件类型
	FileMsgType MsgType = "file"
	// TemplateCardMsgType 模版卡片类型
	TemplateCardMsgType MsgType = "template_card"
)

// Message base message struct
type Message interface {
	// ToMessageMap for JSON  serialization
	ToMessageMap() map[string]interface{}
}

// BaseMessage base message interface
type BaseMessage interface {
	// GetMsgType  get MsgType
	GetMsgType() MsgType
}

// TextMessage text message
type TextMessage struct {
	// Content 	文本内容，最长不超过2048个字节，必须是utf8编码
	Content string
	// UserIds userid的列表，提醒群中的指定成员(@某个成员)，
	UserIds []string
	// Mobiles 手机号列表，提醒手机号对应的群成员(@某个成员)
	Mobiles []string
}

// NewTextMessage create TextMessage
func NewTextMessage(content string) *TextMessage {
	return &TextMessage{
		Content: content,
		UserIds: []string{},
		Mobiles: []string{},
	}
}

// NewTextMessageAtAll create @all TextMessage
func NewTextMessageAtAll(content string) *TextMessage {
	return &TextMessage{
		Content: content,
		UserIds: []string{"@all"},
		Mobiles: []string{},
	}
}

// SetUserIds set TextMessage.UserIds
func (message *TextMessage) SetUserIds(userIds ...string) *TextMessage {
	message.UserIds = userIds
	return message
}

// AddUserIds add TextMessage.UserIds
func (message *TextMessage) AddUserIds(userIds ...string) *TextMessage {
	message.UserIds = append(message.UserIds, userIds...)
	return message
}

// SetMobiles set TextMessage.Mobiles
func (message *TextMessage) SetMobiles(mobiles ...string) *TextMessage {
	message.Mobiles = mobiles
	return message
}

// AddMobiles add TextMessage.Mobiles
func (message *TextMessage) AddMobiles(mobiles ...string) *TextMessage {
	message.Mobiles = append(message.Mobiles, mobiles...)
	return message
}

func (message *TextMessage) GetMsgType() MsgType {
	return TextMsgType
}
func (message *TextMessage) ToMessageMap() map[string]interface{} {
	content := map[string]interface{}{}
	content["content"] = message.Content
	content["mentioned_list"] = message.UserIds
	content["mentioned_mobile_list"] = message.Mobiles
	return map[string]interface{}{
		"msgtype": message.GetMsgType(),
		"text":    content,
	}
}

// MarkdownMessage markdown类型
type MarkdownMessage struct {
	// Content markdown内容，最长不超过4096个字节
	Content string
}

// NewMarkdownMessage create MarkdownMessage
func NewMarkdownMessage(content string) *MarkdownMessage {
	return &MarkdownMessage{
		Content: content,
	}
}

func (message *MarkdownMessage) GetMsgType() MsgType {
	return MarkdownMsgType
}

func (message *MarkdownMessage) ToMessageMap() map[string]interface{} {
	content := map[string]interface{}{
		"content": message.Content,
	}
	return map[string]interface{}{
		"msgtype":  message.GetMsgType(),
		"markdown": content,
	}
}

// ImageMessage 图片类型
type ImageMessage struct {
	// Base64 图片内容的base64编码
	Base64 string
	// Md5	图片内容（base64编码前）的md5值
	Md5 string
}

// NewImageMessage create ImageMessage
func NewImageMessage(base64, md5 string) *ImageMessage {
	return &ImageMessage{
		Base64: base64,
		Md5:    md5,
	}
}

func (message *ImageMessage) GetMsgType() MsgType {
	return ImageMsgType
}
func (message *ImageMessage) ToMessageMap() map[string]interface{} {
	image := map[string]interface{}{
		"base64": message.Base64,
		"md5":    message.Md5,
	}
	return map[string]interface{}{
		"msgtype": message.GetMsgType(),
		"image":   image,
	}
}

// NewsMessage 图文类型
type NewsMessage struct {
	// Articles 图文消息，一个图文消息支持1到8条图文
	Articles []*Article
}

// NewNewsMessage create NewsMessage
func NewNewsMessage(articles ...*Article) *NewsMessage {
	return &NewsMessage{
		Articles: articles,
	}
}

// AddArticles add NewsMessage.Articles
func (message *NewsMessage) AddArticles(articles ...*Article) *NewsMessage {
	message.Articles = append(message.Articles, articles...)
	return message
}
func (message *NewsMessage) GetMsgType() MsgType {
	return NewsMsgTye
}
func (message *NewsMessage) ToMessageMap() map[string]interface{} {
	var articles []map[string]interface{}
	for _, article := range message.Articles {
		articles = append(articles, article.ToMap())
	}
	articleMap := map[string]interface{}{
		"articles": articles,
	}
	return map[string]interface{}{
		"msgtype": message.GetMsgType(),
		"news":    articleMap,
	}
}

// FileMessage 文件类型
type FileMessage struct {
	// MediaId 文件id
	MediaId string
}

// NewFileMessage create FileMessage
func NewFileMessage(mediaId string) *FileMessage {
	return &FileMessage{
		mediaId,
	}
}

func (message *FileMessage) GetMsgType() MsgType {
	return FileMsgType
}
func (message *FileMessage) ToMessageMap() map[string]interface{} {
	file := map[string]interface{}{
		"media_id": message.MediaId,
	}
	return map[string]interface{}{
		"msgtype": message.GetMsgType(),
		"file":    file,
	}
}

type CardType string

// CardBaseMessage 卡片类型
type CardBaseMessage interface {
	Message
	BaseMessage
	// CardMessageMap 卡片message map
	CardMessageMap() map[string]interface{}
}

// CardTextNoticeMessage 模版卡片的模版类型,类型为: text_notice
type CardTextNoticeMessage struct {
	// Source 卡片来源样式信息，不需要来源样式可不填写
	Source *CardSource
	// MainTitle 模版卡片的主要内容，包括一级标题和标题辅助信息
	MainTitle *CardMainTitle
	// EmphasisContent 关键数据样式
	EmphasisContent *CardEmphasisContent
	// QuoteArea 引用文献样式，建议不与关键数据共用
	QuoteArea *CardQuoteArea
	// SubTitleText 二级普通文本，建议不超过112个字
	SubTitleText string
	// HorizontalContents 二级标题+文本列表，列表长度不超过6
	HorizontalContents []*CardHorizontalContent
	// Jumps 跳转指引样式的列表
	Jumps []*CardJump
	// Action 整体卡片的点击跳转事件
	Action *CardAction
}

// NewCardTextNoticeMessage create CardTextNoticeMessage
func NewCardTextNoticeMessage(mainTitle *CardMainTitle, action *CardAction) *CardTextNoticeMessage {
	return &CardTextNoticeMessage{
		MainTitle:          mainTitle,
		HorizontalContents: []*CardHorizontalContent{},
		Jumps:              []*CardJump{},
		Action:             action,
	}
}

// SetSource set CardTextNoticeMessage.Source
func (message *CardTextNoticeMessage) SetSource(source *CardSource) *CardTextNoticeMessage {
	message.Source = source
	return message
}

// SetEmphasisContent set CardTextNoticeMessage.EmphasisContent
func (message *CardTextNoticeMessage) SetEmphasisContent(emphasisContent *CardEmphasisContent) *CardTextNoticeMessage {
	message.EmphasisContent = emphasisContent
	return message
}

// SetQuoteArea set CardTextNoticeMessage.QuoteArea
func (message *CardTextNoticeMessage) SetQuoteArea(quoteArea *CardQuoteArea) *CardTextNoticeMessage {
	message.QuoteArea = quoteArea
	return message
}

// SetSubTitle set CardTextNoticeMessage.SubTitleText
func (message *CardTextNoticeMessage) SetSubTitle(subTitle string) *CardTextNoticeMessage {
	message.SubTitleText = subTitle
	return message
}

// SetHorizontalContents set CardTextNoticeMessage.HorizontalContents
func (message *CardTextNoticeMessage) SetHorizontalContents(horizontalContents ...*CardHorizontalContent) *CardTextNoticeMessage {
	message.HorizontalContents = horizontalContents
	return message
}

// AddHorizontalContents add CardTextNoticeMessage.HorizontalContents
func (message *CardTextNoticeMessage) AddHorizontalContents(horizontalContents ...*CardHorizontalContent) *CardTextNoticeMessage {
	message.HorizontalContents = append(message.HorizontalContents, horizontalContents...)
	return message
}

// SetJumps set CardTextNoticeMessage.Jumps
func (message *CardTextNoticeMessage) SetJumps(jumps ...*CardJump) *CardTextNoticeMessage {
	message.Jumps = jumps
	return message
}

// AddJumps add CardTextNoticeMessage.Jumps
func (message *CardTextNoticeMessage) AddJumps(jumps ...*CardJump) *CardTextNoticeMessage {
	message.Jumps = append(message.Jumps, jumps...)
	return message
}
func (message *CardTextNoticeMessage) CardMessageMap() map[string]interface{} {
	cardMessage := map[string]interface{}{}
	cardMessage["card_type"] = "text_notice"
	if message.Source != nil {
		cardMessage["source"] = message.Source.ToMap()
	}
	cardMessage["main_title"] = message.MainTitle.ToMap()
	if message.EmphasisContent != nil {
		cardMessage["emphasis_content"] = message.EmphasisContent.ToMap()
	}
	if message.QuoteArea != nil {
		cardMessage["quote_area"] = message.QuoteArea.ToMap()
	}
	cardMessage["sub_title_text"] = message.SubTitleText
	if len(message.HorizontalContents) > 0 {
		var contents []map[string]interface{}
		for _, content := range message.HorizontalContents {
			contents = append(contents, content.ToMap())
		}
		cardMessage["horizontal_content_list"] = contents
	}
	if len(message.Jumps) > 0 {
		var jumps []map[string]interface{}
		for _, jump := range message.Jumps {
			jumps = append(jumps, jump.ToMap())
		}
		cardMessage["jump_list"] = jumps
	}
	cardMessage["card_action"] = message.Action.ToMap()
	return cardMessage
}
func (message *CardTextNoticeMessage) GetMsgType() MsgType {
	return TemplateCardMsgType
}
func (message *CardTextNoticeMessage) ToMessageMap() map[string]interface{} {
	return map[string]interface{}{
		"msgtype":       message.GetMsgType(),
		"template_card": message.CardMessageMap(),
	}
}

// CardNewsNoticeMessage 模版卡片的模版类型,图文展示模版卡片的类型为news_notice
type CardNewsNoticeMessage struct {
	Source             *CardSource
	MainTitle          *CardMainTitle
	Image              *CardImage
	ImageTextArea      *CardImageTextArea
	QuoteArea          *CardQuoteArea
	VerticalContents   []*CardVerticalContent
	HorizontalContents []*CardHorizontalContent
	Jumps              []*CardJump
	Action             *CardAction
}

// NewCardNewsNoticeMessage create CardNewsNoticeMessage
func NewCardNewsNoticeMessage(mainTitle *CardMainTitle, image *CardImage, action *CardAction) *CardNewsNoticeMessage {
	return &CardNewsNoticeMessage{
		MainTitle:          mainTitle,
		Image:              image,
		VerticalContents:   []*CardVerticalContent{},
		HorizontalContents: []*CardHorizontalContent{},
		Jumps:              []*CardJump{},
		Action:             action,
	}
}

// SetSource set CardNewsNoticeMessage.Source
func (message *CardNewsNoticeMessage) SetSource(source *CardSource) *CardNewsNoticeMessage {
	message.Source = source
	return message
}

// SetImageTextArea set CardNewsNoticeMessage.ImageTextArea
func (message *CardNewsNoticeMessage) SetImageTextArea(imageTextArea *CardImageTextArea) *CardNewsNoticeMessage {
	message.ImageTextArea = imageTextArea
	return message
}

// SetQuoteArea set CardNewsNoticeMessage.QuoteArea
func (message *CardNewsNoticeMessage) SetQuoteArea(quoteArea *CardQuoteArea) *CardNewsNoticeMessage {
	message.QuoteArea = quoteArea
	return message
}

// SetVerticalContents set CardNewsNoticeMessage.VerticalContents
func (message *CardNewsNoticeMessage) SetVerticalContents(verticalContents ...*CardVerticalContent) *CardNewsNoticeMessage {
	message.VerticalContents = verticalContents
	return message
}

// AddVerticalContents add CardNewsNoticeMessage.VerticalContents
func (message *CardNewsNoticeMessage) AddVerticalContents(verticalContents ...*CardVerticalContent) *CardNewsNoticeMessage {
	message.VerticalContents = append(message.VerticalContents, verticalContents...)
	return message
}

// SetHorizontalContents set CardNewsNoticeMessage.HorizontalContents
func (message *CardNewsNoticeMessage) SetHorizontalContents(horizontalContents ...*CardHorizontalContent) *CardNewsNoticeMessage {
	message.HorizontalContents = horizontalContents
	return message
}

// AddHorizontalContents add CardNewsNoticeMessage.HorizontalContents
func (message *CardNewsNoticeMessage) AddHorizontalContents(horizontalContents ...*CardHorizontalContent) *CardNewsNoticeMessage {
	message.HorizontalContents = append(message.HorizontalContents, horizontalContents...)
	return message
}

// SetJumps set CardNewsNoticeMessage.Jumps
func (message *CardNewsNoticeMessage) SetJumps(jumps ...*CardJump) *CardNewsNoticeMessage {
	message.Jumps = jumps
	return message
}

// AddJumps add CardNewsNoticeMessage.Jumps
func (message *CardNewsNoticeMessage) AddJumps(jumps ...*CardJump) *CardNewsNoticeMessage {
	message.Jumps = append(message.Jumps, jumps...)
	return message
}

func (message *CardNewsNoticeMessage) CardMessageMap() map[string]interface{} {
	cardMessage := map[string]interface{}{}
	cardMessage["card_type"] = "news_notice"
	if message.Source != nil {
		cardMessage["source"] = message.Source.ToMap()
	}
	cardMessage["main_title"] = message.MainTitle.ToMap()
	if message.Image != nil {
		cardMessage["card_image"] = message.Image.ToMap()
	}
	if message.ImageTextArea != nil {
		cardMessage["image_text_area"] = message.ImageTextArea.ToMap()
	}
	if message.QuoteArea != nil {
		cardMessage["quote_area"] = message.QuoteArea.ToMap()
	}
	if len(message.VerticalContents) > 0 {
		var contents []map[string]interface{}
		for _, content := range message.VerticalContents {
			contents = append(contents, content.ToMap())
		}
		cardMessage["vertical_content_list"] = contents
	}
	if len(message.HorizontalContents) > 0 {
		var contents []map[string]interface{}
		for _, content := range message.HorizontalContents {
			contents = append(contents, content.ToMap())
		}
		cardMessage["horizontal_content_list"] = contents
	}
	if len(message.Jumps) > 0 {
		var jumps []map[string]interface{}
		for _, jump := range message.Jumps {
			jumps = append(jumps, jump.ToMap())
		}
		cardMessage["jump_list"] = jumps
	}
	cardMessage["card_action"] = message.Action.ToMap()
	return cardMessage
}
func (message *CardNewsNoticeMessage) GetMsgType() MsgType {
	return TemplateCardMsgType
}
func (message *CardNewsNoticeMessage) ToMessageMap() map[string]interface{} {
	return map[string]interface{}{
		"msgtype":       message.GetMsgType(),
		"template_card": message.CardMessageMap(),
	}
}
