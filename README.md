riakpbc
=======

A Riak Protocol Buffer Client in Go.

A simple `riakpbc` program:

```go
package main

import (
	"github.com/kr/pretty"
	"log"
	"mrb/riakpbc"
)

func main() {
	// connect to the riak cluster
	conn, err := riakpbc.Dial("127.0.0.1:8081")

	if err != nil {
		log.Print(err)
		return
	}

	// get the value of 'key' in the 'bucket' bucket and print it
	obj, _ := riakpbc.FetchObject(conn, "bucket", "key")
	log.Printf("%s", pretty.Formatter(obj))
}
```

The rest of the API:

```
func Dial(addr string) (*Conn, error)
    Dial connects to a single riak server.

FetchObject(c *Conn, bucket string, key string) (b []byte, err error)
    Fetch an object from a bucket

func GetBucket(c *Conn, bucket string) (b []byte, err error)
    Get bucket info

func GetServerInfo(c *Conn) (b []byte, err error)
    Get server info

func ListBuckets(c *Conn) (b [][]byte, err error)
    List all buckets

func ListKeys(c *Conn, bucket string) (b [][]byte, err error)
    List all keys from bucket

func SetBucket(c *Conn, bucket string, nval *uint32, allowmult *bool) (b []byte, err error)
    Create bucket

func StoreObject(c *Conn, bucket string, key string, content string) (b []byte, err error)
    Store an object in riak
```

### Credits

riakpbc is (c) Michael R. Bernstein, 2012

### Licensing

riakpbc is distributed under the MIT License, see `LICENSE` file for details.
