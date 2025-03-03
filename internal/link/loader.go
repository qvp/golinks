package link

import (
	"io"
	"log"
	"net/http"
)

func LoadHtml(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %s\n", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Ошибка при чтении тела ответа: %s\n", err)
		return "", err
	}

	return string(body), nil
}
