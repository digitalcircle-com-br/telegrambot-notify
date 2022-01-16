package telegrambotnotify

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var url string
var apikey string
var cli http.Client

func Init(nurl string, napikey string) {
	url = nurl
	apikey = napikey

}

func Notify(ch string, msg string) error {
	msgo := fmt.Sprintf(`{
		"ch":"%s",
		"msg":"%s"
	}`, ch, msg)
	sr := strings.NewReader(msgo)

	req, err := http.NewRequest(http.MethodPost, url, sr)
	if err != nil {
		return err
	}
	req.Header = http.Header{}
	req.Header.Add("X-API-KEY", apikey)
	res, err := cli.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode > 399 {
		return errors.New(res.Status)
	}
	return nil
}
