package Storage

import (
	error2 "Telegrambot/Library/error"
	"encoding/gob"
	"errors"
	"fmt"
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

func (s *Storage) Save(page *Page) (err error) {
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

func (s *Storage) PickRandom(username string) (page *Page, err error) {
	//Определяем обработку ошибки
	defer func() { err = error2.Wrap("cant pick random page", err) }()

	//Получаем путь к директокрии с файлами
	path := filepath.Join(s.basePath, username)

	//Получаем список файлов
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	//Проверка на наличие файлов
	if len(files) == 0 {
		return nil, ErrorsNew
	}

	//Генерируем случайное число
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	//Берем файл под номером n
	file := files[n]

	//open decode
	return s.decodePage(filepath.Join(path, file.Name()))

}

func (s *Storage) Remove(p *Page) (err error) {

	fileName, err := fileName(p)
	if err != nil {
		return error2.Wrap("cant remove page", err)
	}

	path := filepath.Join(s.basePath, p.userName, fileName)

	if err := os.Remove(path); err != nil {
		return error2.Wrap(fmt.Sprintf("cant remove page %s", path), err)
	}

	return nil
}

func (s *Storage) IsExist(p *Page) (bool, error) {
	fileName, err := fileName(p)
	if err != nil {
		return false, error2.Wrap("cant remove page", err)
	}

	path := filepath.Join(s.basePath, p.userName, fileName)

	switch _, err := os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		return false, error2.Wrap(fmt.Sprintf("cant check if file %s exist", path), err)
	}
	return true, nil
}

func (s Storage) decodePage(filepath string) (*Page, error) {
	//Open the file
	f, err := os.Open(filepath)
	if err != nil {
		return nil, error2.Wrap("cant open file per path", err)
	}

	//Close the file
	defer func() { _ = f.Close() }()

	//Создаем переменную, в которой файл будет декодирован
	var p Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, error2.Wrap("cant decode page", err)
	}

	return &p, nil
}

// Определяем имя файла
func fileName(p *Page) (string, error) {
	return p.Hash()
}
