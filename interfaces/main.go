package main

import (
	"fmt"
	"net/http"
	"sync"

)

// ============================================
// worker Pool
// ============================================

func fetchURL(workerID int, url string){
	resp, err := http.Get(url)
	if err!=nil{
		fmt.Println("Worker: erro ao acessar", workerID, url, err )
	}
	defer resp.Body.Close()
	fmt.Println("Worker: concluido para ", workerID, " com status ", url, resp.StatusCode)

}

func worker(id int, urls <- chan string, wg *sync.WaitGroup){
	defer wg.Done()

	for url := range urls{
		fetchURL(id, url)
	}
}

func main() {

	urls := []string{//array de strings
		"https://www.google.com", 
		"https://www.github.com",
		"https://www.golang.com", 
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
	}


	const numWorks =3 //numeros de works no pool
	urlChannel :=make(chan string, len(urls))
	var wg sync.WaitGroup

	for i:=0; i<= numWorks; i++{
		wg.Add(1)
		go worker(i, urlChannel, &wg)
	}

	for _,url := range urls{
		urlChannel <- url
	}

	close(urlChannel)

	wg.Wait()
	fmt.Println("....FIM.")
}


