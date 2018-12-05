#leveldb-repair

Repairs a corrupted leveldb instance.

## Purpose

Sometimes levelDB gets corrupted. This repo builds into a binary that fixes it without any external depedencies.

## Build

```
> dep ensure
> go build ./ldb-repair.go
```

## Usage

Help
`./ldb-repair -h  # shows help and options`

Repair
`./ldb-repair -d ./path/to/datastore/root`

