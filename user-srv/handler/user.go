package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/v2/util/log"
	"io/ioutil"
	"net/http"
	"user-srv/lib"
	proto "user-srv/submodules/common/protob"
)

type UserService struct {
}

type userCenterRes struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (s *UserService) ValidateToken(ctx context.Context, req *proto.ReqValidateToken, res *proto.Response) error {
	token := req.Token

	//------------------请求用户中心验证token-------------------
	client := &http.Client{}
	url := lib.Config.Url.Uims.Scheme + "://" + lib.Config.Url.Uims.Host + "/saasuims/uims/uimsmanage/AppUserInfo"

	request, _ := http.NewRequest("GET", url, nil)
	//请求头
	request.Header.Set("Authorization", token)
	//请求参数
	params := request.URL.Query()
	params.Set("c_id", req.CId)
	request.URL.RawQuery = params.Encode()
	//发起请求
	response, err := client.Do(request)
	if err != nil {
		log.Info("ValidateToken请求用户中心失败:", err, "|token: ", token)
		lib.NewResponse(1000).WithMsg(err.Error()).ToProtoRes(res)
		return nil
	}
	defer response.Body.Close()

	//------------------解析响应结果-------------------
	if response.StatusCode == 401 {
		lib.NewResponse(1006).ToProtoRes(res)
		return nil
	}
	body, _ := ioutil.ReadAll(response.Body)
	var userRes userCenterRes
	if jsonErr := json.Unmarshal(body, &userRes); jsonErr != nil {
		fmt.Println("json.Unmarshal:", jsonErr, "userRes:", userRes)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}

	//------------------处理响应结果-------------------
	if response.StatusCode == 401 {
		lib.NewResponse(1006).ToProtoRes(res)
		return nil
	}
	if userRes.Code == 6003 {
		lib.NewResponse(30103).ToProtoRes(res)
		return nil
	}
	if userRes.Code == 6021 {
		lib.NewResponse(30104).ToProtoRes(res)
		return nil
	}
	if userRes.Code == 1001 {
		lib.NewResponse(1001).WithMsg(userRes.Message).ToProtoRes(res)
		return nil
	}
	if userRes.Code != 200 {
		lib.NewResponse(30100).ToProtoRes(res)
		return nil
	}
	lib.Success().ToProtoRes(res)
	return nil
}

func (s *UserService) BaseInfoEmployeeListNoPage(ctx context.Context, req *proto.ReqBaseInfoEmployeeListNoPage, res *proto.Response) error {
	token := req.Token

	log.Info("req:", req.GetIds())
	//------------------请求用户中心验证token-------------------
	client := &http.Client{}
	url := lib.Config.Url.Uims.Scheme + "://" + lib.Config.Url.Uims.Host + "/saasuims/uims/backend/employee/BaseInfoEmployeeListNoPage"

	request, _ := http.NewRequest("GET", url, nil)
	//请求头
	request.Header.Set("Authorization", token)
	//请求参数
	params := request.URL.Query()
	params.Set("c_id", req.CId)
	params.Set("ids", req.Ids)
	params.Set("real_name", req.RealName)
	params.Set("mobile", req.Mobile)
	params.Set("email", req.Email)
	//params.Set("status", req.GetStatus())
	request.URL.RawQuery = params.Encode()
	//发起请求
	response, err := client.Do(request)
	if err != nil {
		log.Info("请求用户中心失败:", err)
		lib.NewResponse(1000).WithMsg(err.Error()).ToProtoRes(res)
		return nil
	}
	defer response.Body.Close()

	//------------------解析响应结果-------------------
	if response.StatusCode == 401 {
		lib.NewResponse(1006).ToProtoRes(res)
		return nil
	}
	body, _ := ioutil.ReadAll(response.Body)
	var userRes userCenterRes
	if jsonErr := json.Unmarshal(body, &userRes); jsonErr != nil {
		fmt.Println("json.Unmarshal:", jsonErr, "userRes:", userRes)
		lib.NewResponse(1000).ToProtoRes(res)
		return nil
	}
	fmt.Println("params:", params, "userRes:", userRes)
	lib.Success().WithData(userRes.Data).ToProtoRes(res)
	return nil
}
