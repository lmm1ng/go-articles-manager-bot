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
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(ReadArticle).
				WithCallbackData(fmt.Sprintf("%s %d", ReadArticle, articleId)),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(DeleteArticle).
				WithCallbackData(fmt.Sprintf("%s %d", DeleteArticle, articleId)),
		),
	)
}
