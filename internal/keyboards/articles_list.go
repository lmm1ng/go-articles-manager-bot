package keyboards

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type ArticleListEl struct {
	DbId   uint32
	ListId uint16
	Text   string
	Read   bool
}

func NewArticlesListInlineKeyboard(page uint16, read bool, articleIds []*ArticleListEl) *telego.InlineKeyboardMarkup {
	var visibilityButton string

	if read {
		visibilityButton = HideRead
	} else {
		visibilityButton = ShowRead
	}

	aButtons := []telego.InlineKeyboardButton{}

	for _, aId := range articleIds {
		aButtons = append(
			aButtons,
			tu.InlineKeyboardButton(fmt.Sprintf("%s %d", SelectArticle, aId.ListId)).
				WithCallbackData(fmt.Sprintf("%s %d", SelectArticle, aId.DbId)),
		)
	}

	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(PrevPage).
				WithCallbackData(fmt.Sprintf("%s %d %t", PrevPage, page-1, read)),
			tu.InlineKeyboardButton(visibilityButton).
				WithCallbackData(fmt.Sprintf("%s %t", visibilityButton, !read)),
			tu.InlineKeyboardButton(NextPage).
				WithCallbackData(fmt.Sprintf("%s %d %t", NextPage, page+1, read)),
		),
		tu.InlineKeyboardRow(aButtons...),
	)
}
