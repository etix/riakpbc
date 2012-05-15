package riakpbc

import (
	"code.google.com/p/goprotobuf/proto"
)

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

/*
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


	if resplength == byte(1) {

		if resptype == 10 {
			err = ErrObjectNotFound
		}
		return nil, err
	}

  rawmessage = respraw[5 : resplength+3]
*/

func unmarshalResponse(respraw []byte) (respbuf interface{}, err error) {
	// should accept message and struct
	structname := numToCommand[int(resptype)]

	respbuf = respraw[5 : resplength+3]

	//respstruct := getResponseStruct(structname

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

	if structname == "RpbSetBucketResp" {
		respstruct := &RpbSetBucketResp{}
		err = proto.Unmarshal(respbuf.([]byte), respstruct)
		respbuf = []byte("Success")
	}

	return respbuf, nil
}

func validateResponseHeader(respraw []byte) (err error) {
	if len(respraw) < 5 {
		err = ErrCorruptHeader
		return nil, err
	}

	resplength := int(respraw[3])

	if resplength < 0 {
		err = ErrLengthZero
		return nil, err
	}

	resptype := respraw[4]

	if resptype < 0 || resptype > 24 {
		err = ErrNoSuchCommand
		return nil, err
	}

	return nil
}
