package main

import (
	"database/sql"
	"log"
	"mailinglist/jsonapi"
	"mailinglist/mdb"
	"sync"

	"github.com/alexflint/go-arg"
)

var args struct {
	Dbpath   string `arg:"env:MAILINGLIST_DB`
	BindJson string `arg:"env:MAILINGLIST_BIND_JSON`
}

func main() {
	arg.MustParse(&args)

	if args.Dbpath == "" {
		args.Dbpath = "list.db"
	}

	if args.BindJson == "" {
		args.BindJson = ":8000"
	}

	log.Printf("using database '%v'\n", args.Dbpath)
	db, err := sql.Open("sqlite3", args.Dbpath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mdb.TryCreate(db)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		log.Printf("starting JSON API server...\n")
		jsonapi.Serve(db, args.BindJson)
		wg.Done()

	}()
	wg.Wait()
}
