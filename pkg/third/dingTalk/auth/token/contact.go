package dingTalk

// https://open.dingtalk.com/document/isvapp-server/dingtalk-retrieve-user-information

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dingtalkcontact_1_0 "github.com/alibabacloud-go/dingtalk/contact_1_0"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用 Token 初始化账号Client
 * @return Client
 * @throws Exception
 */
func dingTalkContact10Client() (_result *dingtalkcontact_1_0.Client, _err error) {
	config := &openapi.Config{}
	config.Protocol = tea.String("https")
	config.RegionId = tea.String("central")
	_result = &dingtalkcontact_1_0.Client{}
	_result, _err = dingtalkcontact_1_0.NewClient(config)
	return _result, _err
}

func getUserResponse(accessToken string, options string) (info *dingtalkcontact_1_0.GetUserResponse, _err error) {
	client, _err := dingTalkContact10Client()
	if _err != nil {
		return nil, _err
	}

	getUserHeaders := &dingtalkcontact_1_0.GetUserHeaders{}
	getUserHeaders.XAcsDingtalkAccessToken = tea.String(accessToken)
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		info, _err = client.GetUserWithOptions(tea.String(options), getUserHeaders, &util.RuntimeOptions{})
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		err := &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			err = _t
		} else {
			err.Message = tea.String(tryErr.Error())
		}
		if !tea.BoolValue(util.Empty(err.Code)) && !tea.BoolValue(util.Empty(err.Message)) {
			// err 中含有 code 和 message 属性，可帮助开发定位问题
		}

	}
	return info, _err
}

func GetUserContact(accessToken string) (*dingtalkcontact_1_0.GetUserResponse, error) {
	return getUserResponse(accessToken, "me")
}

func GetFriendContact(accessToken string, unionId string) (*dingtalkcontact_1_0.GetUserResponse, error) {
	return getUserResponse(accessToken, unionId)
}
