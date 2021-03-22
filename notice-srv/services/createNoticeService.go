package services

import (
	"errors"
	"github.com/micro/go-micro/v2/util/log"
	"notice-srv/validator"
)

//消息创建行为接口
type BehaviorNoticeCreateInterface interface {
	//执行业务抽象方法
	Do(repo *RepoNoticeCreate) error
}

//--------------------------------------------
//请求对象
type Request struct {
	UserId   int64
	UserName string
	OpenId   string
	Content  string
	Template interface{}
}

//--------------------------------------------

//参数校验行为
type ParamsBehaviour struct{}

//参数校验
func (behaviour *ParamsBehaviour) Do(repo *RepoNoticeCreate) error {
	errs := validator.Builder.Struct(repo.InfoParams)
	if vErr := validator.TransError(errs); vErr != "" {
		return errors.New("1001")
	}
	return nil
}

//--------------------------------------------
//加载订阅配置
//--------------------------------------------
//消息行为
type NoticeBehaviour struct{}

//创建并加载消息实体信息
func (behavior *NoticeBehaviour) Do(repo *RepoNoticeCreate) error {
	return nil
}

//--------------------------------------------
//消息发送人行为
type NoticeUserBehaviour struct{}

//创建并加载消息发送人信息
func (behavior *NoticeUserBehaviour) Do(repo *RepoNoticeCreate) error {
	return nil
}

//--------------------------------------------
//执行发送
type SendBehaviour struct{}

//加载不同渠道执行发送
func (behavior *SendBehaviour) Do(repo *RepoNoticeCreate) error {
	//if repo.InfoNotice[""] {
	//
	//}
	return nil
}

//--------------------------------------------
//--------------------------------------------

//用于创建消息的仓库
type RepoNoticeCreate struct {
	//请求参数
	InfoParams interface{}
	//验证器信息
	InfoValidator interface{}
	//消息订阅配置信息
	//InfoSubscription interface{}
	//消息实体信息
	InfoNotice interface{}
	//发送人信息
	InfoUser interface{}
	//行为列表
	behaviourList []BehaviorNoticeCreateInterface
}

//初始化仓库
func (repo *RepoNoticeCreate) Init(r *Request) *RepoNoticeCreate {
	//根据Request初始化请求信息和验证器信息
	repo.InfoParams = r
	return repo
}

//注册行为
func (repo *RepoNoticeCreate) RegisterBehaviour(behaviour ...BehaviorNoticeCreateInterface) *RepoNoticeCreate {
	repo.behaviourList = append(repo.behaviourList, behaviour...)
	return repo
}

//执行行为列表
func (repo *RepoNoticeCreate) Run() error {
	for _, behaviour := range repo.behaviourList {
		if err := behaviour.Do(repo); err != nil {
			log.Error("创建提醒失败：", err)
			return err
		}
	}
	return nil
}
