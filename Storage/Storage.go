package Storage

import (
	error2 "Telegrambot/Library/error"
	"crypto/sha1"
	"fmt"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExist(p *Page) (bool, error)
}

// Страница(ссылка),которую мы скинули боту
type Page struct {
	url      string
	userName string
}

func (p *Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.url); err != nil {
		return "", error2.Wrap("cant calculate hash", err)
	}

	if _, err := io.WriteString(h, p.userName); err != nil {
		return "", error2.Wrap("cant calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil // получаем итоговый хеш

}
