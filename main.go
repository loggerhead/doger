package main

import (
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/queue"
	"github.com/loggerhead/doger/api"
	"github.com/loggerhead/doger/crawler"
	"github.com/zolamk/colly-mongo-storage/colly/mongo"
)

func main() {
	if err := initAll(); err != nil {
		panic(err)
	}
}

func initAll() (err error) {
	rand.Seed(time.Now().UTC().UnixNano())
	initLogger()

	c := colly.NewCollector(
		colly.Async(true),
		colly.Debugger(&debug.LogDebugger{}),
		colly.URLFilters(
			regexp.MustCompile(api.GenURL("/user")+"/.*"),
			regexp.MustCompile(api.GenURL("/dynamic")+"/.*"),
		),
	)

	err = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Delay:       1000 * time.Millisecond,
		RandomDelay: 500 * time.Millisecond,
		Parallelism: 1,
	})

	if err != nil {
		return
	}

	storage := &mongo.Storage{
		Database: "doger",
		URI:      "mongodb://127.0.0.1:27017",
	}

	if err = crawler.InitDB("users", storage.URI); err != nil {
		return
	}

	if err = c.SetStorage(storage); err != nil {
		return
	}

	if err = crawler.Init(c); err != nil {
		return
	}

	if err = initAndRunQueue(c); err != nil {
		return
	}
	log.Printf("finished Queue\n")

	c.Wait()
	log.Printf("finished all\n")
	return
}

func initLogger() {
	logFile, err := os.OpenFile("doger.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(logFile)
}

func initAndRunQueue(c *colly.Collector) error {
	maxDID := <-crawler.GetMaxDID()
	log.Printf("begin initAndRunQueue: %v\n", maxDID)

	q, _ := queue.New(10, &queue.InMemoryQueueStorage{MaxSize: maxDID})

	go func() {
		for i := maxDID - 1; i > 0; i-- {
			if err := q.AddURL(api.ReqUtil.GenURLDynamicDetail(i)); err != nil {
				log.Printf("ERROR q.AddURL failed: %v %v\n", i, err)
			}
			if i%100000 == 0 {
				log.Printf("q.AddURL succ: %v\n", i)
			}
		}
	}()

	log.Printf("q.run: %v\n", maxDID)
	return q.Run(c)
}
