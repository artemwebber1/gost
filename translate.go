package gost

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Translate(sourceLang, destinationtLang, text string) (string, error) {
	text = strings.Replace(text, " ", "+", -1)

	url := fmt.Sprintf("https://ftapi.pythonanywhere.com/translate?sl=%s&dl=%s&text=%s", sourceLang, destinationtLang, text)

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resDataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var resData translationApiResponse
	json.Unmarshal(resDataBytes, &resData)

	return resData.DestinationText, nil
}
