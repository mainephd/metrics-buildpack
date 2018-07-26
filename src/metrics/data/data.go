package data

func TelegrafBackgroundScript() string {
	return "[[ -f ./telegraf.env ]] && source ./telegraf.env; telegraf --config ./telegraf.conf >&1 2>&1 &"
}
