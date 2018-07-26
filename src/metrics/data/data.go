package data

func TelegrafBackgroundScript() string {
	return "telegraf --config ./telegraf.conf >&1 2>&1 &"
}
