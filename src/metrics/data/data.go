package data

// TelegrafBackgroundScript Returns a one line bash command to source in a custom environment and start telegraf in the background
func TelegrafBackgroundScript() string {
	return "[[ -f ./telegraf.env ]] && chmod +x ./telegraf.env && source ./telegraf.env; telegraf --config ./telegraf.conf >&1 2>&1 &\n"
}
