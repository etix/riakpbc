package riakpbc

import (
	"encoding/json"
)

var commandToNum = map[string]byte{
	"RpbErrorResp":         0,
	"RpbPingReq":           1,
	"RpbPingResp":          2,
	"RpbGetClientIdReq":    3,
	"RpbGetClientIdResp":   4,
	"RpbSetClientIdReq":    5,
	"RpbSetClientIdResp":   6,
	"RpbGetServerInfoReq":  7,
	"RpbGetServerInfoResp": 8,
	"RpbGetReq":            9,
	"RpbGetResp":           10,
	"RpbPutReq":            11,
	"RpbPutResp":           12,
	"RpbDelReq":            13,
	"RpbDelResp":           14,
	"RpbListBucketsReq":    15,
	"RpbListBucketsResp":   16,
	"RpbListKeysReq":       17,
	"RpbListKeysResp":      18,
	"RpbGetBucketReq":      19,
	"RpbGetBucketResp":     20,
	"RpbSetBucketReq":      21,
	"RpbSetBucketResp":     22,
	"RpbMapRedReq":         23,
	"RpbMapRedResp":        24,
}

// Store an object in riak
func StoreObject(c *Conn, bucket string, key string, content string) (b []byte, err error) {
	jval, err := json.Marshal(content)

	reqstruct := &RpbPutReq{
		Bucket: []byte(bucket),
		Key:    []byte(key),
		Content: &RpbContent{
			Value:       []byte(jval),
			ContentType: []byte("application/json"),
		},
	}

	marshaledRequest, err := marshalRequest(reqstruct)
	if err != nil {
		return nil, err
	}

	formattedRequest, err := formatMarshaledData("RpbPutReq", marshaledRequest)
	if err != nil {
		return
	}

	err = writeRequest(c, formattedRequest)
	if err != nil {
		return nil, err
	}

	marshaledResponse, err := readResponse(c)
	if err != nil {
		return nil, err
	}

	rawresp, err := unmarshalResponse(marshaledResponse)
	if err != nil {
		return nil, err
	}
	b = rawresp.([]byte)

	return b, nil
}

// Fetch an object from a bucket
func FetchObject(c *Conn, bucket string, key string) (b []byte, err error) {
	reqstruct := &RpbGetReq{
		Bucket: []byte(bucket),
		Key:    []byte(key),
	}

	marshaledRequest, err := marshalRequest(reqstruct)
	if err != nil {
		return nil, err
	}

	formattedRequest, err := formatMarshaledData("RpbGetReq", marshaledRequest)
	if err != nil {
		return nil, err
	}

	err = writeRequest(c, formattedRequest)
	if err != nil {
		return nil, err
	}

	marshaledResponse, err := readResponse(c)
	if err != nil {
		return nil, err
	}

	rawresp, err := unmarshalResponse(marshaledResponse)
	if err != nil {
		return nil, err
	}
	b = rawresp.([]byte)

	return b, nil
}

// List all buckets
func ListBuckets(c *Conn) (b [][]byte, err error) {
	reqdata := []byte{0, 0, 0, 1, 15}

	err = writeRequest(c, reqdata)
	if err != nil {
		return nil, err
	}

	marshaledResponse, err := readResponse(c)
	if err != nil {
		return nil, err
	}

	rawresp, err := unmarshalResponse(marshaledResponse)

	if err != nil {
		return nil, err
	}

	b = rawresp.([][]byte)

	return b, nil
}

// List all keys from bucket
func ListKeys(c *Conn, bucket string) (b [][]byte, err error) {
	reqstruct := &RpbListKeysReq{
		Bucket: []byte(bucket),
	}

	marshaledRequest, err := marshalRequest(reqstruct)
	if err != nil {
		return nil, err
	}

	formattedRequest, err := formatMarshaledData("RpbListKeysReq", marshaledRequest)
	if err != nil {
		return nil, err
	}

	err = writeRequest(c, formattedRequest)
	if err != nil {
		return nil, err
	}

	marshaledResponse, err := readResponse(c)
	if err != nil {
		return nil, err
	}

	rawresp, err := unmarshalResponse(marshaledResponse)

	if err != nil {
		return nil, err
	}
	b = rawresp.([][]byte)

	return b, nil
}

// Get server info
func GetServerInfo(c *Conn) (b []byte, err error) {
	reqdata := []byte{0, 0, 0, 1, 7}

	err = writeRequest(c, reqdata)
	if err != nil {
		return nil, err
	}

	marshaledResponse, err := readResponse(c)
	if err != nil {
		return nil, err
	}

	rawresp, err := unmarshalResponse(marshaledResponse)
	if err != nil {
		return nil, err
	}
	b = rawresp.([]byte)

	return b, nil
}

// Get bucket info
func GetBucket(c *Conn, bucket string) (b []byte, err error) {
	return b, nil
}

// Create bucket
func SetBucket(c *Conn, bucket string, nval *uint32, allowmult *bool) (b []byte, err error) {
	propstruct := &RpbBucketProps{
		NVal:      nval,
		AllowMult: allowmult,
	}

	reqstruct := &RpbSetBucketReq{
		Bucket: []byte(bucket),
		Props:  propstruct,
	}

	marshaledRequest, err := marshalRequest(reqstruct)
	if err != nil {
		return nil, err
	}

	formattedRequest, err := formatMarshaledData("RpbSetBucketReq", marshaledRequest)
	if err != nil {
		return
	}

	err = writeRequest(c, formattedRequest)
	if err != nil {
		return nil, err
	}

	marshaledResponse, err := readResponse(c)
	if err != nil {
		return nil, err
	}

	rawresp, err := unmarshalResponse(marshaledResponse)
	if err != nil {
		return nil, err
	}
	b = rawresp.([]byte)

	return b, nil
}
