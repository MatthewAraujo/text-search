package main

import (
	"flag"
	"log"
	"time"

	"github.com/MatthewAraujo/text-search/utils"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "dump", "wiki.gz", "Path to the dump file")
	flag.StringVar(&query, "query", "Small wild cat", "Query to search for")
	flag.Parse()
	log.Println("Full text search in Wikipedia dump")

	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatalf("Error loading documents: %v", err)
	}

	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))
	start = time.Now()

	idx := make(utils.Index)
	idx.Add(docs)

	log.Printf("Indexing  %d took %v", len(docs), time.Since(start))
	start = time.Now()

	matchedIds := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIds), time.Since(start))
	for _, id := range matchedIds {
		log.Printf("Matched document: %v", docs[id])
	}
}
