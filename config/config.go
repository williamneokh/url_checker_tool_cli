package config

type AppConfig struct {
	Repeat      bool
	RepeatTime  int
	IsNotifyBot bool
	IsGoRoutine bool
	CheckCount int
	BotToken string
	TelegramChannel string
}
