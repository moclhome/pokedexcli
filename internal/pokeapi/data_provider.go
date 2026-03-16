package pokeapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (c *Client) GetDataFromCacheOrInternet(url string, flagDebug bool) ([]byte, error) {
	cache := c.cache
	entry, found := cache.Get(url)
	if found {
		return entry, nil
	} else {
		if flagDebug {
			fmt.Printf("Not found. Starting request for %s\n", url)
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}

		defer closing(res.Body, flagDebug)

		if res.StatusCode != http.StatusOK {
			err := errors.New("Status code: " + strconv.Itoa(res.StatusCode))
			return nil, err
		}
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		cache.Add(url, data)
		return data, nil
	}

}

func closing(body io.ReadCloser, flagDebug bool) {
	if flagDebug {
		fmt.Println("closing result body")
	}
	body.Close()
}
