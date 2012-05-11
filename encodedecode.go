package riakpbc

import (
	"code.google.com/p/goprotobuf/proto"
)

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

func marshalRequest(reqstruct interface{}) (marshaledRequest []byte, err error) {

	marshaledRequest, err = proto.Marshal(reqstruct)

	return marshaledRequest, err
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
