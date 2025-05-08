package keyboards

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func NewProfileInlineKeyboard(public bool) *telego.InlineKeyboardMarkup {
	var publicButton string

	if public {
		publicButton = SetPublic
	} else {
		publicButton = HideUser
	}

	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(EditDesc).WithCallbackData(EditDesc),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(publicButton).
				WithCallbackData(fmt.Sprintf("%s %t", publicButton, !public))),
	)
}
