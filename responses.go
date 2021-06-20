package main

type DeviceCodeResponse struct {
	DeviceCode              string `json:"deviceCode"`
	ExpiresIn               int64  `json:"expiresIn"`
	Interval                int64  `json:"interval"`
	UserCode                string `json:"userCode"`
	VerificationURI         string `json:"verificationUri"`
	VerificationURIComplete string `json:"verificationUriComplete"`
}
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	User         struct {
		UserId       int64         `json:"userId"`
		Email        interface{} `json:"email"`
		CountryCode  string      `json:"countryCode"`
		FullName     interface{} `json:"fullName"`
		FirstName    interface{} `json:"firstName"`
		LastName     interface{} `json:"lastName"`
		Nickname     interface{} `json:"nickname"`
		Username     string      `json:"username"`
		Address      interface{} `json:"address"`
		City         interface{} `json:"city"`
		Postalcode   interface{} `json:"postalcode"`
		UsState      interface{} `json:"usState"`
		PhoneNumber  interface{} `json:"phoneNumber"`
		Birthday     interface{} `json:"birthday"`
		Gender       interface{} `json:"gender"`
		ImageId      interface{} `json:"imageId"`
		ChannelId    int         `json:"channelId"`
		ParentId     int         `json:"parentId"`
		AcceptedEULA bool        `json:"acceptedEULA"`
		Created      int64       `json:"created"`
		Updated      int64       `json:"updated"`
		FacebookUid  int         `json:"facebookUid"`
		AppleUid     interface{} `json:"appleUid"`
		NewUser      bool        `json:"newUser"`
	} `json:"user"`
}
