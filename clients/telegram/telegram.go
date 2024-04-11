package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"never_read_list/lib/e"
)

type Client struct {
	host string
	basePath string
	client http.Client
}

const (
	getUpdatesMeethod = "getUpdates"
	sendMessageMeethod = "sendMessage"
)

func New(host string, token string) Client {
	return Client{
		host: host,
		basePath: basePath(token),
		client: http.Client{},
	}
}

func (c *Client) Updates(offset int, limit int ) (updates []Update, err error) {
	defer func() { err = e.WrapIfErr("Cannot perform update", err) } ()

	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMeethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) SendMessages(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMeethod, q)
	if err != nil {
		return e.Wrap("Cannot send message", err)
	}

	return nil
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error ) {
	defer func() { err = e.WrapIfErr("Cannot make request", err) } ()

	u := url.URL{
		Scheme: "https",
		Host: c.host,
		Path: path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func basePath(token string) string {
	return "bot" + token
}