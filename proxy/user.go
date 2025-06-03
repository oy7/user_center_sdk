package proxy

import (
	"fmt"

	"github.com/google/uuid"
	pbUser "github.com/oy7/user_center_sdk/proto/user"
)

type User struct {
	Url       string
	Source    string
	Token     string
	RequestId string
	hook      func(logContext string)
}

func Init(url, source, requestId string) User {
	if requestId == "" {
		requestId = uuid.New().String()
	}
	requestId = fmt.Sprintf("%s_sdk_%s", source, requestId)
	return User{
		Url:       url,
		Source:    source,
		RequestId: requestId,
		hook: func(logContext string) {

		},
	}
}

func (u *User) SetLogHook(f func(logContext string)) {
	u.hook = f
}

// ServiceSmsSendLogin 发送验证码
func (u User) ServiceSmsSendLogin(phoneNumber string, smsCodeType pbUser.E_SMS_CODE_TYPE) (*pbUser.SMSSendLoginResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	smsSendLoginReq := &pbUser.SMSSendLoginReq{
		SmsCodeType: smsCodeType,
		PhoneNumber: phoneNumber,
	}
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] SmsSendLogin, req:%+v", u.RequestId, smsSendLoginReq))
	resp, err := client.SmsSendLogin(ctx, smsSendLoginReq)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] SmsSendLogin, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// UserChangeMobile 修改手机号
func (u User) UserChangeMobile(userId int64, phoneNumber, code string) (*pbUser.UpdateUserInfoResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	updateUserInfoReq := &pbUser.UpdateUserInfoReq{
		Uid:        uint64(userId),
		ModifyType: pbUser.UserModifyType_REBIND_PHONE,
		Phone:      phoneNumber,
		VerifyCode: code,
	}
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] UserChangeMobile-UpdateUserInfo, req:%+v", u.RequestId, updateUserInfoReq))
	resp, err := client.UpdateUserInfo(ctx, updateUserInfoReq)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] UserChangeMobile-UpdateUserInfo, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// ApiUserLogin 用户登录
func (u User) ApiUserLogin(req *pbUser.UserLoginReq) (*pbUser.UserLoginResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] UserLogin, req:%+v", u.RequestId, req))
	resp, err := client.UserLogin(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] UserLogin, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// GetBaseInfo 获取用户基本信息
func (u User) GetBaseInfo(userId uint64) (*pbUser.GetUserInfoResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	req := &pbUser.GetUserInfoReq{
		Uid: userId,
	}
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetUserInfo, req:%+v", u.RequestId, req))
	resp, err := client.GetUserInfo(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetUserInfo, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// GetBaseInfoDecode 获取用户基本信息解密
func (u User) GetBaseInfoDecode(userId int64) (*pbUser.GetUserInfoResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	req := &pbUser.GetUserInfoReq{
		Uid:        uint64(userId),
		IsRealAuth: true,
	}
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetUserInfo, req:%+v", u.RequestId, req))
	resp, err := client.GetUserInfo(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetUserInfo, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// GetUserPhonesByUidList 批量获取手机号
func (u User) GetUserPhonesByUidList(userIds []uint64) (*pbUser.GetUserPhonesResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	req := &pbUser.GetUserPhonesReq{
		UidList: userIds,
	}
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetUserPhonesByUidList, req:%+v", u.RequestId, req))
	resp, err := client.GetUserPhonesByUidList(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetUserPhonesByUidList, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// GetUserIdByPhone 手机号获取用户uid
func (u User) GetUserIdByPhone(phone string) (*pbUser.GetUserIdByPhoneResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	req := &pbUser.GetUserIdByPhoneReq{
		Phone: phone,
	}
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetUserIdByPhone, req:%+v", u.RequestId, req))
	resp, err := client.GetUserIdByPhone(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetUserIdByPhone, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// SetBaseInfo 设置用户信息
func (u User) SetBaseInfo(req *pbUser.UpdateUserInfoReq) (*pbUser.UpdateUserInfoResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] UpdateUserInfo, req:%+v", u.RequestId, req))
	resp, err := client.UpdateUserInfo(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] UpdateUserInfo, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// BindWeChat 绑定微信
func (u User) BindWeChat(req *pbUser.OpenIDBindReq) (*pbUser.OpenIDBindResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] OpenIDBind, req:%+v", u.RequestId, req))
	resp, err := client.OpenIDBind(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] OpenIDBind, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// func (u User) SyncWeChatData(body ApiBindWeChatReq) (ResponseUserBaseInfo, error) {
// 	res := ResponseUserBaseInfo{}
// 	request := ghttp.FromValues{}
// 	request.Add("request_id", gtype.UniqueId())
// 	request.Add("source", u.Source)
// 	request.Add("body", body)
// 	uri := "/api/service_wechat/sync_nick_profile"
// 	u.hook(fmt.Sprintf("Request Uri:%s, body:%s", uri, request.EncodeJson()))
// 	gr, err := ghttp.PostJsonRetry(u.Url+"/api/service_wechat/sync_nick_profile", request, nil, time.Second*3, 3)
// 	u.hook(fmt.Sprintf("Response Uri:%s, body:%s", uri, gr.Body))
// 	if err != nil {
// 		return res, err
// 	}
// 	if gr.StatusCode != 200 {
// 		return res, errors.New(fmt.Sprintf("请求失败, http.status.code: %d", gr.StatusCode))
// 	}
// 	err = json.Unmarshal([]byte(gr.Body), &res)
// 	if err != nil {
// 		return res, errors.New(fmt.Sprintf("解析失败. %s", err))
// 	}
// 	if res.Code != 0 {
// 		return res, errors.New(gtype.ToString(res.Message))
// 	}
// 	return res, nil
// }

// RealName 实名认证
func (u User) RealName(userId uint64, userName, userIdNumber string) (*pbUser.UserCertificationResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	req := &pbUser.UserCertificationReq{
		Uid:          userId,
		UserName:     userName,
		UserIdNumber: userIdNumber,
	}
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] UserCertification, req:%+v", u.RequestId, req))
	resp, err := client.UserCertification(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] UserCertification, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// GetTreeUser 获取用户所属组织详情包含权限信息
func (u User) GetTreeUser(userId uint64) (*pbUser.GetOrgTreeUserResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	req := &pbUser.GetOrgTreeUserReq{
		UserId: userId,
	}
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetOrgTreeUser, req:%+v", u.RequestId, req))
	resp, err := client.GetOrgTreeUser(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetOrgTreeUser, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// BindTreeUser 绑定用户组织关系
func (u User) BindTreeUser(userId uint64, orgId uint32) (*pbUser.BindUserToOrganizationResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	req := &pbUser.BindUserToOrganizationReq{
		UserId: userId,
		OrgId:  orgId,
	}
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] BindOrgTreeUser, req:%+v", u.RequestId, req))
	resp, err := client.BindOrgTreeUser(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] BindOrgTreeUser, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// UnBindTreeUser 解绑用户组织关系
func (u User) UnBindTreeUser(userId uint64, orgId uint32) (*pbUser.UnbindUserToOrganizationResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	req := &pbUser.UnbindUserToOrganizationReq{
		UserId: userId,
		OrgId:  orgId,
	}
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] UnBindOrgTreeUser, req:%+v", u.RequestId, req))
	resp, err := client.UnBindOrgTreeUser(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] UnBindOrgTreeUser, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// GetTreeUserChildren 根据ID查询组织下的组织
func (u User) GetTreeUserChildren(orgId uint32) (*pbUser.GetOrganizationChildrenResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	req := &pbUser.GetOrganizationReq{
		Id: orgId,
	}
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetOrgTreeChildren, req:%+v", u.RequestId, req))
	resp, err := client.GetOrgTreeChildren(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetOrgTreeChildren, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// 获取用户实名信息
func (u User) GetUserSensitiveInfo(userId uint64) (*pbUser.GetUserSensitiveInfoResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	req := &pbUser.GetUserSensitiveInfoReq{}
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetUserSensitiveInfo, req:%+v", u.RequestId, req))
	resp, err := client.GetUserSensitiveInfo(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetUserSensitiveInfo, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// GetUserBankInfo 获取用户银行卡信息
func (u User) GetUserBankInfo(req *pbUser.GetUserBankInfoReq) (*pbUser.GetUserBankInfoResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetUserBankInfo, req:%+v", u.RequestId, req))
	resp, err := client.GetUserBankInfo(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] GetUserBankInfo, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// AddUpUserBankInfo 新增/更新用户银行卡信息
func (u User) AddUpUserBankInfo(req *pbUser.AddUpUserBankInfoReq) (*pbUser.AddUpUserBankInfoResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] AddUserBankInfo, req:%+v", u.RequestId, req))
	resp, err := client.AddUpUserBankInfo(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] AddUserBankInfo, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}

// DeleteUserBankInfo 删除用户银行卡信息
func (u User) DeleteUserBankInfo(req *pbUser.DeleteUserBankInfoReq) (*pbUser.DeleteUserBankInfoResp, error) {
	conn, err := GetConnect(u.Url)
	if err != nil {
		u.hook(fmt.Sprintf("GetConnect [RequestId:%s] err:%v", u.RequestId, err))
		return nil, err
	}
	defer conn.Close()
	client := pbUser.NewUserServerClient(conn.Value())
	ctx := GetMetadataCtx(u.RequestId, u.Source, u.Token)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] DeleteUserBankInfo, req:%+v", u.RequestId, req))
	resp, err := client.DeleteUserBankInfo(ctx, req)
	u.hook(fmt.Sprintf("grpcRequest [RequestId:%s] DeleteUserBankInfo, resp:%+v; err:%v", u.RequestId, resp, err))

	return resp, err
}
