package main

import (
	"github.com/kr/pretty"
	"log"
	"mrb/riakpbc"
)

func main() {
	conn, err := riakpbc.Dial("127.0.0.1:8081")

	if err != nil {
		log.Print(err)
		return
	}

	obj, _ := riakpbc.FetchObject(conn, "bucket", "key")
	log.Printf("%s", pretty.Formatter(obj))

	bux, _ := riakpbc.ListBuckets(conn)
	log.Printf("%s", pretty.Formatter(bux))

	info, _ := riakpbc.GetServerInfo(conn)
	log.Printf("%s", pretty.Formatter(info))

	storeresp, _ := riakpbc.StoreObject(conn, "bucket", "keyzles", "{'keyzle':'deyzle'}")
	log.Printf("%s", pretty.Formatter(storeresp))

	nobj, _ := riakpbc.FetchObject(conn, "bucket", "keyzles")
	log.Printf("%s", pretty.Formatter(nobj))

	nval := uint32(1)
	allowmult := false

	nobj, _ = riakpbc.SetBucket(conn, "bucketier", &nval, &allowmult)
	log.Printf("%s", pretty.Formatter(nobj))

	storeresp, _ = riakpbc.StoreObject(conn, "bucketier", "keyzles", "{'keyzle':'deyzle'}")
	log.Printf("%s", pretty.Formatter(storeresp))

	obj, _ = riakpbc.FetchObject(conn, "bucketier", "keyzles")
	log.Printf("%s", pretty.Formatter(obj))

	conn.Close()
}
