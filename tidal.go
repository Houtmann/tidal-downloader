package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type login struct {
	DeviceCode        string `json:"device_code"`
	UserCode          string `json:"user_code"`
	VerificationURL   string `json:"verification_url"`
	AuthCheckTimeout  int64  `json:"auth_check_timeout"`
	AuthCheckInterval int64  `json:"auth_check_interval"`
	UserID            int64  `json:"user_id"`
	CountryCode       string `json:"country_code"`
	AccessToken       string `json:"access_token"`
	RefreshToken      string `json:"refresh_token"`
	ExpiresIn         int64  `json:"expires_in"`
}

func (l login) save() error {
	file, _ := json.MarshalIndent(l, "", " ")
	err := ioutil.WriteFile("credentials.json", file, 0644)
	if err != nil {
		return fmt.Errorf("can't save credentials on disk %w", err)
	}
	return nil
}

type Tidal struct {
	login *login
}

func NewTidal() (*Tidal, error) {
	tidal := &Tidal{login: &login{}}
	err := tidal.login.GetDeviceCode()
	if err != nil {
		return nil, err
	}
	return tidal, nil

}
func (t *Tidal) Auth() error {
	fmt.Printf("go to %s/%s", t.login.VerificationURL, t.login.UserCode)
	for time.Now().Before(time.Now().Add(time.Duration(t.login.AuthCheckTimeout) * time.Second)) {
		auth, err := t.auth()
		if auth == true {
			break
		}
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}

func (t *Tidal) auth() (bool, error) {
	data := url.Values{}
	data.Set("client_id", ClientID)
	data.Set("grant_type", "urn:ietf:params:oauth:grant-type:device_code")
	data.Set("device_code", t.login.DeviceCode)
	data.Set("scope", "r_usr+w_usr+w_sub")

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/token", AuthURL), strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	r.SetBasicAuth(ClientID, ClientSecret)

	resp, _ := client.Do(r)
	if resp.StatusCode != 200 {
		return false, nil
	}
	var response AuthResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}
	t.login.AccessToken = response.AccessToken
	t.login.CountryCode = response.User.CountryCode
	t.login.ExpiresIn = response.ExpiresIn
	t.login.RefreshToken = response.RefreshToken
	t.login.UserID = response.User.UserId
	err = t.login.save()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (l login) AccessTokenValid() (bool, error) {
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/sessions", APIURL), nil)
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", l.AccessToken))
	r.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		return false, err
	}
	if resp.StatusCode != 200 {
		return false, err
	}
	return true, nil
}

func (t *Tidal) Configure() error {
	err := t.login.openCredentials()
	if err != nil {
		return  err
	}
	valid, err := t.login.AccessTokenValid()
	if err != nil {
		return err
	}
	if !valid {
		err := t.Auth()
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *login) openCredentials() error {
	file, err := ioutil.ReadFile("credentials.json")

	if err != nil {
		return err
	}
	data := login{}

	if err := json.Unmarshal([]byte(file), &data); err != nil{
		return err
	}
	l.AccessToken = data.AccessToken
	l.AuthCheckInterval = data.AuthCheckInterval
	l.AuthCheckTimeout = data.AuthCheckTimeout
	l.CountryCode = data.CountryCode
	l.DeviceCode = data.DeviceCode
	l.ExpiresIn = data.ExpiresIn
	l.RefreshToken = data.RefreshToken
	l.UserID = data.UserID
	return nil
}

func (l *login) GetDeviceCode() error {
	data := url.Values{}
	data.Set("client_id", ClientID)
	data.Set("score", "r_usr+w_usr+w_sub")

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/device_authorization", AuthURL), strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	var response DeviceCodeResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	l.DeviceCode = response.DeviceCode
	l.UserCode = response.UserCode
	l.VerificationURL = response.VerificationURI
	l.AuthCheckTimeout = response.ExpiresIn
	l.AuthCheckInterval = response.Interval
	return nil

}

func (t Tidal) DownloadsTrack() error  {
 return nil
}