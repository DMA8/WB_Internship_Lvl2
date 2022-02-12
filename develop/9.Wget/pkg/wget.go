package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

//MainDownload downloads given URL, css, js, png...
func MainDownload(inpURL string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go wGet(&wg, inpURL)
	resp, err := http.Get(inpURL)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(resp.Body)
	domain, _ := url.Parse(inpURL)
	for line, err := r.ReadString('\n'); err != io.EOF; {
		if err != nil {
			log.Fatal(err)
		}
		a := getAllSrcInHTMLLine(line, domain.Host)
		for _, j := range a{
			wg.Add(1)
			go wGet(&wg, j)
		}
		line, err = r.ReadString('\n')
	}
	wg.Wait()
}

func wGet(wg *sync.WaitGroup, inpURL string){
	defer wg.Done() 
	if !checkAvailabilyty(inpURL){
		fmt.Println(inpURL, "is not available")
		return
	}
	fileName := getFileName(inpURL)
	createFile(fileName)
	download(inpURL, fileName)
}

func createFile(fileName string){
	if !checkFileNameValid(fileName){
		fmt.Println(fileName)
		log.Fatal(errors.New("wrong file name"))
	}
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func checkFileNameValid(name string) bool{//!!! Переписать! Валится!
	pattern := `^[/a-zA-Z0-9._-]{1,255}`
	if ok, _ := regexp.Match(pattern, []byte(name)); ok{
		return true
	}
	return false
}

func getFileName(URLRaw string) string{ 
	prefix := "/mnt/c/Users/DB/OneDrive/html/" // путь, куда надо сохранить файл
	parsedURL, err := url.Parse(URLRaw)
	if err != nil{
		log.Fatal(err)
	}
	path := parsedURL.Path
	splittedPath := strings.Split(path, "/")
	if ext := filepath.Ext(URLRaw); ext == ""{
		name := splittedPath[len(splittedPath) - 1]
		if name == ""{
			name = "index"
		}
		return prefix + name + ".html"
	}
	return prefix + splittedPath[len(splittedPath) - 1]
}

func download(url string, fileName string){
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend) // os.Open() не сработал
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(file, resp.Body) // MAGIC Чтобы не хранить ответ в памяти, сразу пишем его в файл
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	defer file.Close()
}

func checkAvailabilyty(rawURL string) bool{
	resp, err := http.Get(rawURL)
	if err != nil {
		return false
	}
	return resp.StatusCode == 200
}

func getAllSrcInHTMLLine(htmLine, relativePrefix string) []string{ //TODO не работает с закр тегом 
	urlsFoud := make([]string, 0)
	srcRegEx, err := regexp.Compile(`src=\"(.*?)"`)
	if err != nil {
		log.Fatal(err)
	}
	urls := srcRegEx.FindAllSubmatch([]byte(htmLine), 50)
	for _, i := range urls{
		for _, j := range i[1:]{
			splitted := strings.Split(string(j), "/n")
			for _, h := range splitted{
				h := strings.TrimPrefix(string(h), "//")
				if strings.HasPrefix(h, "/"){
					h = fmt.Sprintf("%s%s",relativePrefix, h )
				}
				if !strings.HasPrefix(h, "https://") && !strings.HasPrefix(h, "http://"){
					h = "https://" + h
				}
				h = strings.TrimSuffix(h, ">")
				h = strings.TrimSuffix(h, "/")
				urlsFoud = append(urlsFoud, h)
			}
		}
	}
	return urlsFoud
}
