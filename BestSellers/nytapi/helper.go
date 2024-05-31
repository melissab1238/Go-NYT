package nytapi

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/melissab1238/GO-NYT/BestSellers/config"
)

// helper function
func GetJsonFromUrl(url string) ([]byte, error) {
	url = fmt.Sprintf("%s?api-key=%s", url, config.APIKEY)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
