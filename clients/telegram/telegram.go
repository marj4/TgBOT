package main

import (
	error2 "Telegrambot/Library/error"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string // host.com/bot <token>
	http     http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendUpdatesMethod = "sendUpdates"
)

// Create client
func New(host, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		http:     http.Client{},
	}

}

// Create basePath
func newBasePath(token string) string {
	return "bot" + token
}

// Implement methood Update()
func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}

	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res updateResponse

	//Распарсим json()
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, err

}

func (c *Client) SendUpdates(chatId int, text string) error {
	q := url.Values{}

	q.Add("chatId", strconv.Itoa(chatId))
	q.Add("text", text)

	_, err := c.doRequest(sendUpdatesMethod, q)
	if err != nil {
		return error2.Wrap("can`t send message", err)
	}

	return nil

}

func (c *Client) doRequest(methood string, query url.Values) ([]byte, error) {
	const errMsg = "can`n do request"
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, methood),
	}

	//Prepare request(create object)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, error2.Wrap(errMsg, err)
	}

	//Передаем объекту req параметры запроса query
	req.URL.RawQuery = query.Encode()

	//Отправляем запрос
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, error2.Wrap(errMsg, err)
	}

	//Получим содержимое
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, error2.Wrap(errMsg, err)
	}

	//Возвращаем результат
	return body, nil

}
