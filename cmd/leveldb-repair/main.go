package main

import (
	"fmt"
	"os"

	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/errors"
	ldbopts "github.com/btcsuite/goleveldb/leveldb/opts"
	flags "github.com/jessevdk/go-flags"
)

var options struct {
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

	db, err := leveldb.OpenFile(opts.LevelDBPath, &ldbopts.Options{})
	if err != nil && errors.IsCorrupted(err) && !nopts.GetReadOnly() {
		db, recoverErr = leveldb.RecoverFile(opts.LevelDBPath, &ldbopts.Options{})
		if recoverErr != nil {
			fmt.Printf("failed during recovery: %s\n", recoverErr)
			os.Exit(2)
		}
		defer db.Close()
		fmt.Println("recovery completed")
		os.Exit(0)
	}
	fmt.Println("datastore opened successfully, recovery skipped")
	os.Exit(0)
}
