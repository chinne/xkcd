package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Create a folder to store the comic images
	os.Mkdir("Pictures/comics", os.ModePerm)

	// Send a GET request to the homepage of the website
	res, err := http.Get("https://xkcd.com/")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	// Parse the HTML document using goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Find the URL of the comic image
	imgSrc, exists := doc.Find("#comic img").Attr("src")
	if !exists {
		fmt.Println("Error: Comic image not found")
		return
	}

	// Find the title of the comic
	title := doc.Find("#comic img").AttrOr("title", "")
	if title == "" {
		fmt.Println("Error: Title not found")
		return
	}

	// Download the comic image
	res, err = http.Get("https:" + imgSrc)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	// Create a file to save the image
	fileName := filepath.Join("Pictures/comics", strings.ReplaceAll(title, "/", "_")+filepath.Ext(imgSrc))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Copy the image from the response body to the file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Comic image saved successfully to", fileName)

}
