package keyboards

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func NewArticleInlineKeyboard(articleId uint32, read bool, deletable bool) *telego.InlineKeyboardMarkup {
	var buttons = [][]telego.InlineKeyboardButton{}

	readBtn := ReadArticle
	if read {
		readBtn = UnreadArticle
	}

	buttons = append(buttons, tu.InlineKeyboardRow(
		tu.InlineKeyboardButton(readBtn).
			WithCallbackData(fmt.Sprintf("%s %d %t", readBtn, articleId, read)),
	))

	if deletable {
		buttons = append(buttons, tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(DeleteArticle).
				WithCallbackData(fmt.Sprintf("%s %d", DeleteArticle, articleId)),
		))
	}

	return tu.InlineKeyboard(buttons...)
}
