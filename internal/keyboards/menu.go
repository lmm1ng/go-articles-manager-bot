package keyboards

import (
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
			tu.KeyboardButton(AddArticle),
			tu.KeyboardButton(GetVibe),
		),
		tu.KeyboardRow(
			tu.KeyboardButton(ShowArticles),
			tu.KeyboardButton(RandomArticle),
		)).WithResizeKeyboard()
}
