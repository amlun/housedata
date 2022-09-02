package service

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
)

var (
	client *http.Client
)

func init() {
	jar, _ := cookiejar.New(nil)
	client = &http.Client{
		Jar: jar,
	}
}

func GetNum(url, selector string) (num int, err error) {
	log.WithField("url", url).WithField("selector", selector).Info("抓取数据")
	res, err := client.Get(url)
	if err != nil {
		log.Error(err)
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Error(err)
		return
	}
	total := doc.Find(selector).Text()
	num, err = strconv.Atoi(strings.TrimSpace(total))
	return
}
