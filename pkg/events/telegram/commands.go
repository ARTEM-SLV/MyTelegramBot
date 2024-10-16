package telegram

import "strings"

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)
}
