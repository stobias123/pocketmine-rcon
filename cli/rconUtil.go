package main

const (
	consoleDefault      = "\033[39m"
	consoleBlack        = "\033[30m"
	consoleRed          = "\033[31m"
	consoleGreen        = "\033[32m"
	consoleYellow       = "\033[33m"
	consoleBlue         = "\033[34m"
	consoleMagenta      = "\033[35m"
	consoleCyan         = "\033[36m"
	consoleLightGray    = "\033[37m"
	consoleDarkdGray    = "\033[90m"
	consoleLightRed     = "\033[91m"
	consoleLightGreen   = "\033[92m"
	consoleLightYellow  = "\033[93m"
	consoleLightBlue    = "\033[94m"
	consoleLightMagenta = "\033[95m"
	consoleLightCyan    = "\033[96m"
	consoleWhite        = "\033[97m"

	minecraftDarkRed     = "§4"
	minecraftRed         = "§c"
	minecraftGold        = "§6"
	minecraftYellow      = "§e"
	minecraftDarkGreen   = "§2"
	minecraftGreen       = "§a"
	minecraftAqua        = "§b"
	minecraftDarkAqua    = "§3"
	minecraftDarkBlue    = "§1"
	minecraftBlue        = "§9"
	minecraftLightPurple = "§d"
	minecraftDarkPurple  = "§5"
	minecraftWhite       = "§f"
	minecraftGray        = "§7"
	minecraftDarkGray    = "§8"
	minecraftBlack       = "§0"
)

func getConsoleColor(minecraftColor string) string {
	colorMap := map[string]string{
		minecraftDarkRed:     consoleRed,
		minecraftRed:         consoleLightRed,
		minecraftGold:        consoleYellow,
		minecraftYellow:      consoleYellow,
		minecraftDarkGreen:   consoleGreen,
		minecraftGreen:       consoleLightGreen,
		minecraftAqua:        consoleLightCyan,
		minecraftDarkAqua:    consoleCyan,
		minecraftDarkBlue:    consoleBlue,
		minecraftBlue:        consoleLightBlue,
		minecraftLightPurple: consoleMagenta,
		minecraftDarkPurple:  consoleMagenta,
		minecraftWhite:       consoleWhite,
		minecraftGray:        consoleDarkdGray,
		minecraftDarkGray:    consoleLightGray,
		minecraftBlack:       consoleDefault,
	}
	return colorMap[minecraftColor]
}

func getMinecraftColors() []string {
	return []string{
		minecraftDarkRed,
		minecraftRed,
		minecraftGold,
		minecraftYellow,
		minecraftDarkGreen,
		minecraftGreen,
		minecraftAqua,
		minecraftDarkAqua,
		minecraftDarkBlue,
		minecraftBlue,
		minecraftLightPurple,
		minecraftDarkPurple,
		minecraftWhite,
		minecraftGray,
		minecraftDarkGray,
		minecraftBlack}
}

func getConsoleColors() []string {
	return []string{
		consoleDefault,
		consoleBlack,
		consoleRed,
		consoleGreen,
		consoleYellow,
		consoleBlue,
		consoleMagenta,
		consoleCyan,
		consoleLightGray,
		consoleDarkdGray,
		consoleLightRed,
		consoleLightGreen,
		consoleLightYellow,
		consoleLightBlue,
		consoleLightMagenta,
		consoleLightCyan,
		consoleWhite}
}
