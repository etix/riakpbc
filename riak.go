package riakpbc

import (
	"encoding/json"
)

// Store an object in riak
func (c *Conn) StoreObject(bucket string, key string, content string) (b []byte, err error) {
	jval, err := json.Marshal(content)

	reqstruct := &RpbPutReq{
		Bucket: []byte(bucket),
		Key:    []byte(key),
		Content: &RpbContent{
			Value:       []byte(jval),
			ContentType: []byte("application/json"),
		},
	}

}

// Fetch an object from a bucket
func (c *Conn) FetchObject(bucket string, key string) (b []byte, err error) {
	reqstruct := &RpbGetReq{
		Bucket: []byte(bucket),
		Key:    []byte(key),
	}

}

// List all buckets
func (c *Conn) ListBuckets() (b [][]byte, err error) {
	reqdata := []byte{0, 0, 0, 1, 15}

	err = writeRequest(c, reqdata)
	if err != nil {
		return nil, err
	}
}

// List all keys from bucket
func (c *Conn) ListKeys(bucket string) (b [][]byte, err error) {
	reqstruct := &RpbListKeysReq{
		Bucket: []byte(bucket),
	}

}

// Get server info
func (c *Conn) GetServerInfo() (b []byte, err error) {
	reqdata := []byte{0, 0, 0, 1, 7}

	err = writeRequest(c, reqdata)
	if err != nil {
		return nil, err
	}
}

// Get bucket info
func (c *Conn) GetBucket(bucket string) (b []byte, err error) {
	return b, nil
}

// Create bucket
func (c *Conn) SetBucket(bucket string, nval *uint32, allowmult *bool) (response []byte, err error) {
	propstruct := &RpbBucketProps{
		NVal:      nval,
		AllowMult: allowmult,
	}

	reqstruct := &RpbSetBucketReq{
		Bucket: []byte(bucket),
		Props:  propstruct,
	}

	err = makeRequest(reqstruct, "RpbSetBucketReq")

	response = getResp(bar, baz)

	return response, nil
}

/*

	marshaledRequest, err := marshalRequest(reqstruct)
	if err != nil {
		return nil, err
	}

	formattedRequest, err := prependRequestHeader("RpbPutReq", marshaledRequest)
	if err != nil {
		return
	}

	err = writeRequest(c, formattedRequest)
	if err != nil {
		return nil, err
	}
*/
