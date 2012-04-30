package riakpbc

import (
	"code.google.com/p/goprotobuf/proto"
	"encoding/json"
	"errors"
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

var numToCommand = map[int]string{
	0:  "RpbErrorResp",
	1:  "RpbPingReq",
	2:  "RpbPingResp",
	3:  "RpbGetClientIdReq",
	4:  "RpbGetClientIdResp",
	5:  "RpbSetClientIdReq",
	6:  "RpbSetClientIdResp",
	7:  "RpbGetServerInfoReq",
	8:  "RpbGetServerInfoResp",
	9:  "RpbGetReq",
	10: "RpbGetResp",
	11: "RpbPutReq",
	12: "RpbPutResp",
	13: "RpbDelReq",
	14: "RpbDelResp",
	15: "RpbListBucketsReq",
	16: "RpbListBucketsResp",
	17: "RpbListKeysReq",
	18: "RpbListKeysResp",
	19: "RpbGetBucketReq",
	20: "RpbGetBucketResp",
	21: "RpbSetBucketReq",
	22: "RpbSetBucketResp",
	23: "RpbMapRedReq",
	24: "RpbMapRedResp",
}

var (
	ErrLengthZero    = errors.New("length response 0")
	ErrCorruptHeader = errors.New("corrupt header")
)

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

func marshalRequest(reqstruct interface{}) (marshaledRequest []byte, err error) {

	marshaledRequest, err = proto.Marshal(reqstruct)

	return marshaledRequest, err
}

func formatMarshaledData(commandName string, marshaledReqData []byte) (formattedData []byte, e error) {
	msgbuf := []byte{}
	formattedData = []byte{}

	mn := []byte{0, 0, 0}
	comm := []byte{commandToNum[commandName]}

	msgbuf = append(msgbuf, comm...)
	msgbuf = append(msgbuf, marshaledReqData...)

	length := []byte{byte(len(msgbuf))}

	formattedData = append(formattedData, mn...)
	formattedData = append(formattedData, length...)
	formattedData = append(formattedData, msgbuf...)

	return formattedData, nil
}

func writeRequest(c *Conn, formattedRequest []byte) (err error) {
	_, err = c.conn.Write(formattedRequest)
	return err
}

func readResponse(c *Conn) (respraw []byte, err error) {
	respraw = make([]byte, 512)

	c.conn.Read(respraw)

	_ = respraw[3]

	return respraw, nil
}

func unmarshalResponse(respraw []byte) (respbuf interface{}, err error) {
	if len(respraw) < 5 {
		err = ErrCorruptHeader
		return respbuf, err
	}

	resplength := respraw[3]

	if resplength == byte(1) {
		err = ErrLengthZero
		return respbuf, err
	}

	resptype := respraw[4]

	structname := numToCommand[int(resptype)]

	respbuf = respraw[5 : resplength+3]

	if structname == "RpbGetResp" {
		respstruct := &RpbGetResp{}
		err = proto.Unmarshal(respbuf.([]byte), respstruct)
		respbuf = respstruct.Content[0].Value
	}

	if structname == "RpbGetServerInfoResp" {
		respstruct := &RpbGetServerInfoResp{}
		err = proto.Unmarshal(respbuf.([]byte), respstruct)
		respbuf = respstruct.Node
	}

	if structname == "RpbListBucketsResp" {
		respstruct := &RpbListBucketsResp{}
		err = proto.Unmarshal(respbuf.([]byte), respstruct)
		respbuf = respstruct.Buckets
	}

	if structname == "RpbListKeysResp" {
		respstruct := &RpbListKeysResp{}
		err = proto.Unmarshal(respbuf.([]byte), respstruct)
		respbuf = respstruct.Keys
	}

	return respbuf, nil
}
