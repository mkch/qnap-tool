package qnap

import (
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/mkch/wol"
)

// ErrAuthFailed is the error returned by Shutdown or
// other functions  to indicate a authentication error.
var ErrAuthFailed = errors.New("authentication failed")

func login(baseUrl, user, password string) (sid string, err error) {
	var reqUrl = baseUrl + "/cgi-bin/authLogin.cgi"
	var formData = url.Values{}
	formData.Add("user", user)
	formData.Add("pwd", encodePwd(password))
	resp, err := http.Post(reqUrl, "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	var xmlResp = struct {
		AuthPassed bool   `xml:"authPassed"`
		AuthSID    string `xml:"authSid"`
	}{}
	if err = xml.NewDecoder(resp.Body).Decode(&xmlResp); err != nil {
		return "", err
	}
	if !xmlResp.AuthPassed {
		return "", ErrAuthFailed
	}
	return xmlResp.AuthSID, nil
}

// Shutdown shuts down the QNAP NAS.
// baseUrl is the base url of QNAP admin page, in format of "http(s)://IP:port" (http://192.168.0.111:5000 for example).
// user and password is the user name and password used to login.
func Shutdown(baseUrl, user, password string) error {
	sid, err := login(baseUrl, user, password)
	if err != nil {
		return err
	}
	var reqUrl = baseUrl + "/cgi-bin/sys/sysRequest.cgi?sid=" + url.QueryEscape(sid) + "&subfunc=power_mgmt&apply=shutdown"
	resp, err := http.Post(reqUrl, "", nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	var xmlResp = struct {
		AuthPassed bool `xml:"authPassed"`
	}{}
	if err = xml.NewDecoder(resp.Body).Decode(&xmlResp); err != nil {
		return err
	}
	if !xmlResp.AuthPassed {
		return ErrAuthFailed
	}
	return nil
}

// Wake sends a Wake-on-LAN request to the NIC.
// If Wake-on-LAN is enabled, the plugged but powered off NAS should be turned on.
func Wake(macAddr string) error {
	return wol.Wake(macAddr)
}
