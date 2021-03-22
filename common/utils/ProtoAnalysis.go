package utils

import (
	"encoding/json"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
)

//需要获取 proto.any格式中的data,需要进行解码
//ps:需要保证存进any中的数据是进行过jsonMarshal的
func UnmarshalProtoAny(pb *anypb.Any) (interface{}, error) {
	//UnmarshalAny
	pData := &anypb.Any{}
	if anyErr := ptypes.UnmarshalAny(pb, pData); anyErr != nil {
		return nil, anyErr
	}

	//json decode
	var data interface{}
	if jsonErr := json.Unmarshal(pData.Value, &data); jsonErr != nil {
		return nil, jsonErr
	}
	return data, nil
}

//需要获取 proto.any格式中的data,解码为指定结构
func UnmarshalProtoAnyTo(pb *anypb.Any, data interface{}) error {
	//UnmarshalAny
	pData := &anypb.Any{}
	if anyErr := ptypes.UnmarshalAny(pb, pData); anyErr != nil {
		return anyErr
	}

	//json decode
	if jsonErr := json.Unmarshal(pData.Value, &data); jsonErr != nil {
		return jsonErr
	}
	return nil
}
