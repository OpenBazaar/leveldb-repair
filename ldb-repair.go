package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	ldbopt "github.com/syndtr/goleveldb/leveldb/opt"
)

type options struct {
	LevelDBPath string `short:"d" long:"datastore-path" description:"full or relative path to the root of the levelDB datastore"`
}

func main() {
	var (
		opts      = &options{}
		argParser = flags.NewParser(opts, flags.Default)
	)
	if _, err := argParser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			fmt.Printf("parsing arguments: %s\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("opening %s...", opts.LevelDBPath)
	db, err := leveldb.OpenFile(opts.LevelDBPath, &ldbopt.Options{})
	if err != nil && errors.IsCorrupted(err) {
		db, recoverErr := leveldb.RecoverFile(opts.LevelDBPath, &ldbopt.Options{})
		if recoverErr != nil {
			fmt.Printf("recovering: %s\n", recoverErr)
			os.Exit(2)
		}
		defer db.Close()
		fmt.Println("recovery complete")
		os.Exit(0)
	}
	defer db.Close()
	fmt.Println("opened but not corrupted: recovery skipped")
	os.Exit(0)
}
