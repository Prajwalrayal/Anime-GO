package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func downloadFile(url, filename string) error {
	// fmt.Println("Hello")
	filename += ".mp4"
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	return err
}

func download_episode(url string) {

	println(url)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	htmlContent := string(body)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".dowload a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")

		fmt.Println(exists)
		if exists {
			fmt.Printf("Downloading: %s\n", link)
			fileName := filepath.Base(link)
			err := downloadFile(link, fileName)
			if err != nil {
				log.Println("Error downloading file:", err)
			} else {
				fmt.Printf("Downloaded: %s\n", fileName)
			}
		}
	})
}

func main() {

	var name string
	fmt.Print("Enter Name in format Dr-Stone: ")
	fmt.Scanln(&name)

	var episode string
	fmt.Print("Enter Episode number: ")
	fmt.Scanln(&episode)
	res, err := http.Get("https://ww1.gogoanimes.fi/" + name + "-episode-" + episode)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	htmlContent := string(body)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}
	link := doc.Find("li.dowloads a").AttrOr("href", "")

	download_episode(link)
}
