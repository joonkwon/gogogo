package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bbalet/stopwords"
	"golang.org/x/net/html"
)

var url = "http://shakespeare.mit.edu/"
var dataFolder = "./data"

var rexSpecial = regexp.MustCompile(`[\?\.,"_\(\):;!]`)
var rexSigleQuote = regexp.MustCompile(`([a-zA-z]+)'[ s]`)
var wordsCount = make(map[string]int)

type wordfreq struct {
	word string
	freq int
}

func main() {

	res, _ := http.Get(url)

	z := html.NewTokenizer(res.Body)
	defer res.Body.Close()

	links := []string{}
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "a" {
				path := t.Attr[0].Val
				if strings.HasPrefix(path, "Poetry") {
					// links = append(links, fmt.Sprintf(url+"%s", path))
					continue
				}
				links = append(links, fmt.Sprintf(url+"%s/full.html", strings.Split(path, "/")[0]))
			}
		}
	}
	for i, link := range links {
		if i == 0 || i == 1 || i == 39 || i == 40 {
			continue
		}
		saveHTMLToTxt(dataFolder, link)
	}

	err := filepath.Walk(dataFolder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fp, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("failed to read file: %s: caused by: %s", path, err.Error())
			}
			reader := bufio.NewReader(fp)

			bytes, err := ioutil.ReadAll(reader)
			if err != nil {
				return fmt.Errorf("failed in reading file: %s: error: %s", path, err.Error())
			}
			// fmt.Println(string(bytes))
			countWords(cleanStr(string(bytes)), wordsCount)

		}
		return nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// build wordfrequency
	wordfrequency := []wordfreq{}
	for k, v := range wordsCount {
		wf := wordfreq{word: k, freq: v}
		wordfrequency = append(wordfrequency, wf)
	}

	sort.Slice(wordfrequency, func(i, j int) bool {
		return wordfrequency[i].freq < wordfrequency[j].freq
	})

	ln := len(wordfrequency)
	// print top 20
	for i := ln - 1; i >= ln-20; i-- {
		fmt.Printf("%s: %d\n", wordfrequency[i].word, wordfrequency[i].freq)
	}

	// print top 20 - 40
	for i := ln - 21; i >= ln-40; i-- {
		fmt.Printf("%s: %d\n", wordfrequency[i].word, wordfrequency[i].freq)
	}

	// print bottom 20
	for i := 0; i < 20; i++ {
		fmt.Printf("%s: %d\n", wordfrequency[i].word, wordfrequency[i].freq)
	}
}

func saveHTMLToTxt(dataFolder string, link string) {
	res, err := http.Get(link)
	if err != nil || res.StatusCode != 200 {
		fmt.Printf("error: status code: %d, link: %s\n", res.StatusCode, link)
		res.Body.Close()
		return
	}
	lc := strings.Split(link, "/")
	ln := len(lc)

	fp := dataFolder + "/" + lc[ln-2] + ".txt"
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("failed to open file: %s\n", fp)
	}
	defer f.Close()
	wr := bufio.NewWriter(f)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Printf("failed to create goquery document: %s", err.Error())
		return
	}
	wr.WriteString(doc.Find("tbody").Children().First().Find("td").Text())
	doc.Find("table").Siblings().Each(func(in int, el *goquery.Selection) {
		wr.WriteString(el.Text())
	})
}

func cleanStr(in string) string {
	ret := removeSpecialChars(in)
	ret = stopwords.CleanString(ret, "en", false)

	return ret
}

func removeSpecialChars(in string) string {
	ret := rexSpecial.ReplaceAllString(in, " ")
	ret = rexSigleQuote.ReplaceAllString(ret, "$1")
	return ret
}

func countWords(in string, count map[string]int) {
	words := strings.Fields(in)
	for _, w := range words {
		word := strings.ToLower(w)
		count[word]++
	}
}
