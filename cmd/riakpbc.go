package main

import (
	"github.com/kr/pretty"
	"log"
	"mrb/riakpbc"
)

func main() {
	riak, err := riakpbc.Dial("127.0.0.1:8081")

	if err != nil {
		log.Print(err)
		return
	}

	obj, _ := riak.FetchObject("bucket", "key")
	log.Printf("%s", pretty.Formatter(obj))

	bux, _ := riak.ListBuckets()
	log.Printf("%s", pretty.Formatter(bux))

	info, _ := riak.GetServerInfo()
	log.Printf("%s", pretty.Formatter(info))

	storeresp, _ := riak.StoreObject("bucket", "keyzles", "{'keyzle':'deyzle'}")
	log.Printf("%s", pretty.Formatter(storeresp))

	nobj, _ := riak.FetchObject("bucket", "keyzles")
	log.Printf("%s", pretty.Formatter(nobj))

	nval := uint32(1)
	allowmult := false

	nobj, _ = riak.SetBucket("squadrons", &nval, &allowmult)
	log.Printf("%s", pretty.Formatter(nobj))

	storeresp, _ = riak.StoreObject("squadrons", "nymets", "{'players':['deyzle','freyzle','chezyle']}")
	log.Printf("%s", pretty.Formatter(storeresp))

	obj, _ = riak.FetchObject("squadrons", "nymets")
	log.Printf("%s", pretty.Formatter(obj))

	riak.Close()
}
