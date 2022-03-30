package dingTalk

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dingtalkoauth2_1_0 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

// DingTalk AccessToken 文档
// https://open.dingtalk.com/document/orgapp-server/obtain-user-token

/**
 * 使用 Token 初始化账号Client
 * @return Client
 * @throws Exception
 */

func dingTalkOAuth210Client() (_result *dingtalkoauth2_1_0.Client, _err error) {
	config := &openapi.Config{}
	config.Protocol = tea.String("https")
	config.RegionId = tea.String("central")
	_result = &dingtalkoauth2_1_0.Client{}
	_result, _err = dingtalkoauth2_1_0.NewClient(config)
	return _result, _err
}

func accessToken(clientId string, clientSecret string, code string, refreshToken string, grantType string) (token *dingtalkoauth2_1_0.GetUserTokenResponse, err error) {
	client, err := dingTalkOAuth210Client()
	if err != nil {
		return nil, err
	}

	getUserTokenRequest := &dingtalkoauth2_1_0.GetUserTokenRequest{
		ClientId:     tea.String(clientId),
		ClientSecret: tea.String(clientSecret),
		Code:         tea.String(code),
		RefreshToken: tea.String(refreshToken),
		GrantType:    tea.String(grantType),
	}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		token, err = client.GetUserToken(getUserTokenRequest)
		if err != nil {
			return err
		}

		return nil
	}()

	if tryErr != nil {
		var _err = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			_err = _t
		} else {
			_err.Message = tea.String(tryErr.Error())
		}
		if !tea.BoolValue(util.Empty(_err.Code)) && !tea.BoolValue(util.Empty(_err.Message)) {
			// err 中含有 code 和 message 属性，可帮助开发定位问题
			fmt.Println(_err)
		}

	}
	return token, err
}

// GetAccessToken 第一次获取AccessToken调用，返回AccessToken，ExpireTimer和RefreshToken.
func GetAccessToken(clientId string, clientSecret string, code string) (*dingtalkoauth2_1_0.GetUserTokenResponse, error) {
	return accessToken(clientId, clientSecret, code, "", "authorization_code")
}

// RefreshAccessToken 超时7200s，需要刷新AccessToken.
func RefreshAccessToken(clientId string, clientSecret string, code string, refreshToken string) (*dingtalkoauth2_1_0.GetUserTokenResponse, error) {
	return accessToken(clientId, clientSecret, code, refreshToken, "refresh_token")
}
