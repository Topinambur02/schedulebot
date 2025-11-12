package states

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"topinambur02.com/m/v2/model"
)

type StateContext struct {
	states map[string]State
}

func NewStateContext(bot *tgbotapi.BotAPI) *StateContext {
	sc := &StateContext{
		states: make(map[string]State),
	}
	
	sc.states["/start"] = &StartState{}
	sc.states["Показать текущее расписание"] = &CurrentScheduleState{}
	sc.states["Удалить кэш для Redis"] = &DeleteCacheState{}
	sc.states["Посмотреть расписание на завтра"] = &TomorrowScheduleState{}
	sc.states["default"] = &DefaultState{}
	
	return sc
}

func (sc *StateContext) HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, user model.User) {
	state, exists := sc.states[message.Text]
	if !exists {
		state = sc.states["default"]
	}
	
	state.Handle(bot, message, user)
}
