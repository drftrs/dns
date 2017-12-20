package drftrs

import (
	"bufio"
	"net/http"
	"strings"
)

func CompileBlacklist(url string) ([]string, error) {
	var blacklist []string
	res, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		blacklist = append(blacklist, domain)
	}
	return blacklist, nil
}
