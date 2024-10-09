package Storage

import (
	error2 "Telegrambot/Library/error"
	"encoding/gob"
	"errors"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Storage struct {
	basePath string // Info in which folder, we save files
}

const (
	defaultPerm = 0774
)

var ErrorsNew = errors.New("no saved page")

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *Page) (err error) {
	//Определяем способы обработки ошибок
	defer func() { err = error2.Wrap("cant save page", err) }()

	// Определяем куда мы будем сохранять файлы
	filePath := filepath.Join(s.basePath, page.userName)

	//Создаем все нужный директории пути filePath
	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}

	// Получаем имя файла
	fname, err := fileName(page)
	if err != nil {
		return err
	}

	//Теперь добавим к пути до файла само имя файла
	filePath = filepath.Join(filePath, fname)

	//Создаем файл
	file, err := os.Create(filePath)

	//Обработка ошибки метода Close()
	defer func() { _ = file.Close() }()

	//Теперь нужно стерилизовать страницу, т.е привести страницу
	//к формату, который мы можем записать в файл и восстановить исходную структура
	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

}

func (s Storage) PickRandom(username string) (page *Page, err error) {
	//Определяем способы обработки ошибок
	defer func() { err = error2.Wrap("cant pick random page", err) }()

	// Определяем куда мы будем сохранять файлы
	filePath := filepath.Join(s.basePath, username)

	//Получаем список файлов
	files, err := os.ReadDir(filePath)
	if err != nil {
		return nil, err
	}

	//Проверка на наличие файлов
	if len(files) == 0 {
		return nil, ErrorsNew
	}

	//Генерируем случайное число,
	rand.Seed(time.Now().UnixNano())

	//Получаем само число
	n := rand.Intn(len(files))

	//Получаем файл с номером, который мы сгенерировали
	file := files[n]

	//Декодируем файл(open file ---> decode)

}

func (s Storage) decodePage(filepath string) (*Page, error) {
	//Открываем файл
	f, err := os.Open(filepath)
	if err != nil {
		return nil, error2.Wrap("cant decode page", err)
	}
	//Закрываем
	defer func() { _ = f.Close() }()

	//FINISH IT...
}

// Определяем имя файла
func fileName(p *Page) (string, error) {
	return p.Hash()
}
