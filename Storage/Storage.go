package Storage

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
