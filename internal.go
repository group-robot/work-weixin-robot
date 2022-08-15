package work_weixin_robot

// Article 	图文消息
type Article struct {
	// Title 标题
	Title string
	// Description 标题，不超过128个字节，超过会自动截断
	Description string
	// Url 点击后跳转的链接
	Url string
	// PicUrl 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150
	PicUrl string
}

// NewArticle create Article
func NewArticle(title, url string) *Article {
	return &Article{
		Title: title,
		Url:   url,
	}
}

// SetDesc set Article.Description
func (article *Article) SetDesc(desc string) *Article {
	article.Description = desc
	return article
}

// SetPicUrl set Article.PicUrl
func (article *Article) SetPicUrl(picUrl string) *Article {
	article.PicUrl = picUrl
	return article
}

// ToMap to map
func (article *Article) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title":       article.Title,
		"description": article.Description,
		"url":         article.Url,
		"picurl":      article.PicUrl,
	}
}

// DescColor 来源文字
type DescColor int

const (
	// GreyDescColor 灰色
	GreyDescColor DescColor = iota
	// BlackDescColor 黑色
	BlackDescColor
	// RedDescColor 红色
	RedDescColor
	// GreenDescColor 绿色
	GreenDescColor
)

// CardSource 卡片来源样式信息
type CardSource struct {
	// IconUrl 	来源图片的url
	IconUrl string
	// Desc 来源图片的描述，建议不超过13个字
	Desc string
	// DescColor 来源文字的颜色
	DescColor DescColor
}

// NewCardSource create CardSource
func NewCardSource() *CardSource {
	return &CardSource{
		DescColor: GreyDescColor,
	}
}

// SetIconUrl set CardSource.IconUrl
func (card *CardSource) SetIconUrl(iconUrl string) *CardSource {
	card.IconUrl = iconUrl
	return card
}

// SetDesc set CardSource.Desc
func (card *CardSource) SetDesc(desc string) *CardSource {
	card.Desc = desc
	return card
}

// SetDescColor set CardSource.DescColor
func (card *CardSource) SetDescColor(descColor DescColor) *CardSource {
	card.DescColor = descColor
	return card
}

// ToMap to map
func (card *CardSource) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"icon_url":   card.IconUrl,
		"desc":       card.Desc,
		"desc_color": card.DescColor,
	}
}

// CardMainTitle 模版卡片的主要内容
type CardMainTitle struct {
	// Title 一级标题，建议不超过26个字
	Title string
	// Desc 标题辅助信息，建议不超过30个字
	Desc string
}

// NewCardMainTitle create CardMainTitle
func NewCardMainTitle() *CardMainTitle {
	return &CardMainTitle{}
}

// SetTitle set CardMainTitle.Title
func (card *CardMainTitle) SetTitle(title string) *CardMainTitle {
	card.Title = title
	return card
}

// SetDesc set CardMainTitle.Desc
func (card *CardMainTitle) SetDesc(desc string) *CardMainTitle {
	card.Desc = desc
	return card
}

// ToMap to map
func (card *CardMainTitle) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title": card.Title,
		"desc":  card.Desc,
	}
}

// CardEmphasisContent 关键数据样式
type CardEmphasisContent struct {
	// Title 关键数据样式的数据内容，建议不超过10个字
	Title string
	// Desc 关键数据样式的数据描述内容，建议不超过15个字
	Desc string
}

// NewCardEmphasisContent create CardEmphasisContent
func NewCardEmphasisContent() *CardEmphasisContent {
	return &CardEmphasisContent{}
}

// SetTitle set CardEmphasisContent.Title
func (card *CardEmphasisContent) SetTitle(title string) *CardEmphasisContent {
	card.Title = title
	return card
}

// SetDesc set CardEmphasisContent.Desc
func (card *CardEmphasisContent) SetDesc(desc string) *CardEmphasisContent {
	card.Desc = desc
	return card
}

// ToMap to map
func (card *CardEmphasisContent) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title": card.Title,
		"desc":  card.Desc,
	}
}

// CardImage  图片样式
type CardImage struct {
	// Url 图片的url
	Url string
	// AspectRatio 图片的宽高比，宽高比要小于2.25，大于1.3，不填该参数默认1.3
	AspectRatio float32
}

// NewCardImage create CardImage
func NewCardImage(url string) *CardImage {
	return &CardImage{
		Url:         url,
		AspectRatio: 1.3,
	}
}

// SetAspectRation set CardImage.AspectRatio
func (card *CardImage) SetAspectRation(aspectRation float32) *CardImage {
	card.AspectRatio = aspectRation
	return card
}

// ToMap to map
func (card *CardImage) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"url":          card.Url,
		"aspect_ratio": card.AspectRatio,
	}
}

// ClickType  区域点击事件
type ClickType int

const (
	// ClickNone 没有点击事件
	ClickNone ClickType = iota
	// ClickUrl 跳转url
	ClickUrl
	// ClickMiniApp 跳转小程序
	ClickMiniApp
)

// CardImageTextArea 左图右文样式
type CardImageTextArea struct {
	// ClickType 左图右文样式区域点击事件
	ClickType ClickType
	// Url 点击跳转的url, CardImageTextArea.ClickType 为 ClickUrl
	Url string
	// Appid 点击跳转的小程序的appid, CardImageTextArea.ClickType 为 ClickMiniApp
	Appid string
	// PagePath 点击跳转的小程序的pagePath
	PagePath string
	// Title 左图右文样式的标题
	Title string
	// Desc 左图右文样式的描述
	Desc string
	// ImageUrl 	左图右文样式的图片url
	ImageUrl string
}

// NewCardImageTextArea create CardImageTextArea
func NewCardImageTextArea(imageUrl string) *CardImageTextArea {
	return &CardImageTextArea{
		ImageUrl:  imageUrl,
		ClickType: ClickNone,
	}
}

// SetType set CardImageTextArea.ClickType
func (card *CardImageTextArea) SetType(clickType ClickType) *CardImageTextArea {
	card.ClickType = clickType
	return card
}

// SetUrl set CardImageTextArea.Url
func (card *CardImageTextArea) SetUrl(url string) *CardImageTextArea {
	card.Url = url
	return card
}

// SetAppid set CardImageTextArea.Appid
func (card *CardImageTextArea) SetAppid(appid string) *CardImageTextArea {
	card.Appid = appid
	return card
}

// SetPagePath set CardImageTextArea.PagePath
func (card *CardImageTextArea) SetPagePath(pagePath string) *CardImageTextArea {
	card.PagePath = pagePath
	return card
}

// SetTitle set CardImageTextArea.Title
func (card *CardImageTextArea) SetTitle(title string) *CardImageTextArea {
	card.Title = title
	return card
}

// SetDesc set CardImageTextArea.Desc
func (card *CardImageTextArea) SetDesc(desc string) *CardImageTextArea {
	card.Desc = desc
	return card
}

// ToMap to map
func (card *CardImageTextArea) ToMap() map[string]interface{} {
	image := map[string]interface{}{}
	image["image_url"] = card.ImageUrl
	image["desc"] = card.Desc
	image["title"] = card.Title
	image["type"] = card.ClickType
	image["url"] = card.Url
	image["appid"] = card.Appid
	image["pagepath"] = card.PagePath
	return image
}

// CardQuoteArea 引用文献样式，建议不与关键数据共用
type CardQuoteArea struct {
	// ClickType 引用文献样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
	ClickType ClickType
	// Url 点击跳转的url，quote_area.type是1时必填
	Url string
	// Appid 点击跳转的小程序的appid，quote_area.type是2时必填
	Appid string
	// PagePath 点击跳转的小程序的pagepath，quote_area.type是2时选填
	PagePath string
	// Title 引用文献样式的标题
	Title string
	// QuoteText 引用文献样式的引用文案
	QuoteText string
}

// NewCardQuoteArea create CardQuoteArea
func NewCardQuoteArea(clickType ClickType) *CardQuoteArea {
	return &CardQuoteArea{
		ClickType: clickType,
	}
}

// SetUrl set CardQuoteArea.Url
func (quote *CardQuoteArea) SetUrl(url string) *CardQuoteArea {
	quote.Url = url
	return quote
}

// SetAppid set CardQuoteArea.Appid
func (quote *CardQuoteArea) SetAppid(appid string) *CardQuoteArea {
	quote.Appid = appid
	return quote
}

// SetPagePath set CardQuoteArea.PagePath
func (quote *CardQuoteArea) SetPagePath(pagePath string) *CardQuoteArea {
	quote.PagePath = pagePath
	return quote
}

// SetTitle set CardQuoteArea.Title
func (quote *CardQuoteArea) SetTitle(title string) *CardQuoteArea {
	quote.Title = title
	return quote
}

// SetQuoteText set CardQuoteArea.QuoteText
func (quote *CardQuoteArea) SetQuoteText(quoteText string) *CardQuoteArea {
	quote.QuoteText = quoteText
	return quote
}

// ToMap to map
func (quote *CardQuoteArea) ToMap() map[string]interface{} {
	image := map[string]interface{}{}
	image["quote_text"] = quote.QuoteText
	image["title"] = quote.Title
	image["type"] = quote.ClickType
	image["url"] = quote.Url
	image["appid"] = quote.Appid
	image["pagepath"] = quote.PagePath
	return image
}

// CardVerticalContent 卡片二级垂直内容
type CardVerticalContent struct {
	// Title 卡片二级标题，建议不超过26个字
	Title string
	// Desc	二级普通文本，建议不超过112个字
	Desc string
}

// NewCardVerticalContent create CardVerticalContent
func NewCardVerticalContent(title string) *CardVerticalContent {
	return &CardVerticalContent{
		Title: title,
	}
}

// SetDesc set CardVerticalContent.Desc
func (vertical *CardVerticalContent) SetDesc(desc string) *CardVerticalContent {
	vertical.Desc = desc
	return vertical
}

// ToMap to map
func (vertical *CardVerticalContent) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title": vertical.Title,
		"desc":  vertical.Desc,
	}
}

// HorizontalType 	链接类型
type HorizontalType int

const (
	// TextHorizontalType 普通文本
	TextHorizontalType HorizontalType = iota
	// UrlHorizontalType 跳转url
	UrlHorizontalType
	// FileHorizontalType 下载附件
	FileHorizontalType
	// AtHorizontalType @员工
	AtHorizontalType
)

// CardHorizontalContent 二级标题+文本列表
type CardHorizontalContent struct {
	// HorizontalType 模版卡片的二级标题信息内容支持的类型
	HorizontalType HorizontalType
	// KeyName 二级标题，建议不超过5个字
	KeyName string
	// Value 二级文本,如果horizontal_content_list.type是2，该字段代表文件名称（要包含文件类型），建议不超过26个字
	Value string
	// Url 链接跳转的url,horizontal_content_list.type是1时必填
	Url string
	// MedialId 附件的media_id，horizontal_content_list.type是2时必填
	MedialId string
	// UserId 被@的成员的userid，horizontal_content_list.type是3时必填
	UserId string
}

// NewCardHorizontalContent create CardHorizontalContent
func NewCardHorizontalContent(keyname string) *CardHorizontalContent {
	return &CardHorizontalContent{
		KeyName:        keyname,
		HorizontalType: TextHorizontalType,
	}
}

// SetType set CardHorizontalContent.HorizontalType
func (horizontal *CardHorizontalContent) SetType(horizontalType HorizontalType) *CardHorizontalContent {
	horizontal.HorizontalType = horizontalType
	return horizontal
}

// SetValue set CardHorizontalContent.Value
func (horizontal *CardHorizontalContent) SetValue(value string) *CardHorizontalContent {
	horizontal.Value = value
	return horizontal
}

// SetUrl set CardHorizontalContent.Url
func (horizontal *CardHorizontalContent) SetUrl(url string) *CardHorizontalContent {
	horizontal.Url = url
	return horizontal
}

// SetUrl set CardHorizontalContent.MedialId
func (horizontal *CardHorizontalContent) setMediaId(mediaId string) *CardHorizontalContent {
	horizontal.MedialId = mediaId
	return horizontal
}

// SetUserId set CardHorizontalContent.UserId
func (horizontal *CardHorizontalContent) SetUserId(userId string) *CardHorizontalContent {
	horizontal.UserId = userId
	return horizontal
}

// ToMap to map
func (horizontal *CardHorizontalContent) ToMap() map[string]interface{} {
	content := map[string]interface{}{}
	content["keyname"] = horizontal.KeyName
	content["value"] = horizontal.Value
	content["type"] = horizontal.HorizontalType
	content["url"] = horizontal.Url
	content["media_id"] = horizontal.MedialId
	content["userid"] = horizontal.UserId
	return content
}

// CardJump 跳转指引样式的列表
type CardJump struct {
	// JumpType 跳转链接类型，0或不填代表不是链接，1 代表跳转url，2 代表跳转小程序
	JumpType ClickType
	// Title 跳转链接样式的文案内容，建议不超过13个字
	Title string
	// Url 跳转链接的url，jump_list.type是1时必填
	Url string
	// Appid 跳转链接的小程序的appid，jump_list.type是2时必填
	Appid string
	// PagePath 跳转链接的小程序的pagepath，jump_list.type是2时选填
	PagePath string
}

// NewCardJump create CardJump
func NewCardJump(title string) *CardJump {
	return &CardJump{
		JumpType: ClickNone,
		Title:    title,
	}
}

// SetType set CardJump.JumpType
func (jump *CardJump) SetType(jumpType ClickType) *CardJump {
	jump.JumpType = jumpType
	return jump
}

// SetUrl set CardJump.Url
func (jump *CardJump) SetUrl(url string) *CardJump {
	jump.Url = url
	return jump
}

// SetAppId set CardJump.Appid
func (jump *CardJump) SetAppId(appId string) *CardJump {
	jump.Appid = appId
	return jump
}

// SetPagePath set CardJump.PagePath
func (jump *CardJump) SetPagePath(pagePath string) *CardJump {
	jump.PagePath = pagePath
	return jump
}

// ToMap to map
func (jump *CardJump) ToMap() map[string]interface{} {
	image := map[string]interface{}{}
	image["title"] = jump.Title
	image["type"] = jump.JumpType
	image["url"] = jump.Url
	image["appid"] = jump.Appid
	image["pagepath"] = jump.PagePath
	return image
}

// CardAction 整体卡片的点击跳转事件
type CardAction struct {
	// ClickType 	卡片跳转类型
	ClickType ClickType
	// Url	跳转事件的url，card_action.type是1时必填
	Url string
	// Appid 跳转事件的小程序的appid，card_action.type是2时必填
	Appid string
	// PagePath 跳转事件的小程序的pagepath，card_action.type是2时选填
	PagePath string
}

// NewCardAction create CardAction
func NewCardAction(clickType ClickType) *CardAction {
	return &CardAction{
		ClickType: clickType,
	}
}

// SetUrl set CardAction.Url
func (card *CardAction) SetUrl(url string) *CardAction {
	card.Url = url
	return card
}

// SetAppId set CardAction.Appid
func (card *CardAction) SetAppId(appId string) *CardAction {
	card.Appid = appId
	return card
}

// SetPagePath set CardAction.PagePath
func (card *CardAction) SetPagePath(pagePath string) *CardAction {
	card.PagePath = pagePath
	return card
}

// ToMap to map
func (card *CardAction) ToMap() map[string]interface{} {
	image := map[string]interface{}{}
	image["type"] = card.ClickType
	image["url"] = card.Url
	image["appid"] = card.Appid
	image["pagepath"] = card.PagePath
	return image
}
