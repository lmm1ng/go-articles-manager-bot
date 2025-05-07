package keyboards

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func NewMainMenuKeyboard() *telego.ReplyKeyboardMarkup {
	return tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton(Profile),
			tu.KeyboardButton(Statistics),
		),
		tu.KeyboardRow(
			tu.KeyboardButton(AddArticle).WithText(AddArticle),
		),
		tu.KeyboardRow(
			tu.KeyboardButton(ShowArticles),
			tu.KeyboardButton(RandomArticle),
		)).WithResizeKeyboard()
}

func NewArticleInlineKeyboard(articleId uint32, deletable bool) *telego.InlineKeyboardMarkup {
	var buttons = [][]telego.InlineKeyboardButton{}

	buttons = append(buttons, tu.InlineKeyboardRow(
		tu.InlineKeyboardButton(ReadArticle).
			WithCallbackData(fmt.Sprintf("%s %d", ReadArticle, articleId)),
	))

	if deletable {
		buttons = append(buttons, tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(DeleteArticle).
				WithCallbackData(fmt.Sprintf("%s %d", DeleteArticle, articleId)),
		))
	}

	return tu.InlineKeyboard(buttons...)
}

func NewArticlesListInlineKeyboard(page uint16, read bool, articlesCount uint16, perPage uint16) *telego.InlineKeyboardMarkup {
	var visibilityButton string

	if read {
		visibilityButton = HideRead
	} else {
		visibilityButton = ShowRead
	}

	aButtons := []telego.InlineKeyboardButton{}

	for pos := range articlesCount {
		absPos := page*perPage - perPage + pos + 1
		aButtons = append(
			aButtons,
			tu.InlineKeyboardButton(fmt.Sprintf("%s %d", SelectArticle, page)).
				WithCallbackData(fmt.Sprintf("%s %d %t %d", SelectArticle, page, read, absPos)),
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
