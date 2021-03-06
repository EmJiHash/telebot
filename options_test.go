package telebot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBtn(t *testing.T) {
	r := &ReplyMarkup{}

	assert.Equal(t, &ReplyButton{Text: "T"}, r.Text("T").Reply())
	assert.Equal(t, &ReplyButton{Text: "T", Contact: true}, r.Contact("T").Reply())
	assert.Equal(t, &ReplyButton{Text: "T", Location: true}, r.Location("T").Reply())
	assert.Equal(t, &ReplyButton{Text: "T", Poll: PollAny}, r.Poll("T", PollAny).Reply())

	assert.Nil(t, r.Data("T", "u").Reply())
	assert.Equal(t, &InlineButton{Unique: "u", Text: "T"}, r.Data("T", "u").Inline())
	assert.Equal(t, &InlineButton{Unique: "u", Text: "T", Data: "1|2"}, r.Data("T", "u", "1", "2").Inline())
	assert.Equal(t, &InlineButton{Text: "T", URL: "url"}, r.URL("T", "url").Inline())
	assert.Equal(t, &InlineButton{Text: "T", InlineQuery: "q"}, r.Query("T", "q").Inline())
	assert.Equal(t, &InlineButton{Text: "T", InlineQueryChat: "q"}, r.QueryChat("T", "q").Inline())
	assert.Equal(t, &InlineButton{Text: "T", Login: &Login{Text: "T"}}, r.Login("T", &Login{Text: "T"}).Inline())
}

func TestReplyInline(t *testing.T) {
	r := &ReplyMarkup{}
	r.Reply(
		r.Row(r.Text("Menu")),
		r.Row(r.Text("Settings")),
	)

	assert.Equal(t, [][]ReplyButton{
		{{Text: "Menu"}},
		{{Text: "Settings"}},
	}, r.ReplyKeyboard)

	i := &ReplyMarkup{}
	i.Inline(r.Row(
		r.Data("Previous", "prev"),
		r.Data("Next", "next"),
	))

	assert.Equal(t, [][]InlineButton{{
		{Unique: "prev", Text: "Previous"},
		{Unique: "next", Text: "Next"},
	}}, i.InlineKeyboard)
}
