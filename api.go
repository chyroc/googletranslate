package googletranslate

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Language string

const (
	Zh Language = "zh"
	En Language = "en"
)

func Translate(query string, from, to Language) (string, error) {
	apiURL, body, headers := generateParam(query, string(from), string(to))
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(body))
	if err != nil {
		return "", err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	res := []string{}
	err = json.Unmarshal(bs, &res)
	if err != nil {
		return "", err
	}
	if len(res) == 0 {
		return "", nil
	}
	return res[0], nil
}

var httpClient = &http.Client{Timeout: time.Second * 3}

func generateParam(q, from, to string) (string, string, map[string]string) {
	query := url.Values{}
	for k, v := range map[string]string{
		"anno":   "3",
		"client": "te_lib",
		"format": "html",
		"v":      "1.0",
		"key":    "AIzaSyBOti4mM-6x9WDnZIjIeyEU21OpBXqWBgw",
		"logld":  "vTE_20200210_00",
		"sl":     from,
		"tl":     to,
		"sp":     "nmt",
		"tc":     "1",
		"sr":     "1",
		"tk":     generateTk(q, generateTkk()),
		"mode":   "1",
	} {
		query.Add(k, v)
	}

	apiURL := "https://translate.googleapis.com/translate_a/t?" + query.Encode()

	u := url.Values{}
	u.Add("q", q)
	body := u.Encode()

	headers := map[string]string{
		"Content-Type":   "application/x-www-form-urlencoded",
		"User-Agent":     "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36",
		"Content-Length": strconv.Itoa(len(body)),
	}

	return apiURL, body, headers
}

func generateTkk() string {
	// return "458341.3287556325"
	return ""
}

// https://gist.github.com/vielhuber/b7739bf50b2edcf636c43a8f8910def9
func generateTk(query, ckk string) string {
	return ""
}
