package utils

import (
	"fmt"
	"gin-demo/common"
	"gin-demo/model"
	"gin-demo/response"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("adwddwfklwklwwewklklklolwfojwofowoadwpdwopfw")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn((len(letters)))]
	}
	return string(result)
}

func HttpGet(url string) (result io.ReadCloser, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	//defer resp.Body.Close()

	return resp.Body, nil
}

func Reptile(c *gin.Context) {
	start := 1
	end := 10
	DB := common.GetDb()
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E6%9D%8E%E6%AF%85&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("err:", err)
			continue
		}

		doc, err := goquery.NewDocumentFromReader(result)
		if err != nil {
			log.Fatal(err)
			return
		}

		doc.Find("#content li").Each(func(i int, s *goquery.Selection) {
			item := s.Find("a.j_th_tit")
			text := item.Text()
			href, _ := item.Attr("href")
			if text != "" {
				var title model.Pachong
				title.Name = text
				title.Href = href
				fmt.Println(text, href)
				DB.Create(&title)
			}

		})
	}
	response.Success(c, nil, "成功")
}
