// +build appengine

package service

import (
	"io/ioutil"

	"appengine"
	"appengine/urlfetch"
)

func Download(url string, c appengine.Context) ([]byte, error) {
	client := urlfetch.Client(c)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	return data, err
}
