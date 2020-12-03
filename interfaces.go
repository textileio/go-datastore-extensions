package dsextensions

import "github.com/ipfs/go-datastore/query"

type QueryExt struct {
	query.Query
	SeekPrefix string
}

type QueryExtensions interface {
	QueryExtended(q QueryExt) (query.Results, error)
}
