package main

// #include <stdlib.h>
import "C"

import (
	"errors"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"

	"encoding/json"

	"unsafe"

	"firebase.google.com/go/auth"
	"gitlab.com/squarealfa/dart_framework/firebase-auth-admin/go/dart_api_dl"
)

var client *auth.Client

const initError = "please call initialize before any other authorization method"

//export initialize
func initialize(api unsafe.Pointer, credentialsFile *C.char) *C.char {
	dart_api_dl.Init(api)

	gCredentialsFile := C.GoString(credentialsFile)

	opt := option.WithCredentialsFile(gCredentialsFile)
	ctx := context.Background()
	var app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		C.CString(err.Error())
	}

	cli, err := app.Auth(ctx)
	if err != nil {
		return C.CString(err.Error())
	}

	client = cli

	return nil
}

//export initializeWithJson
func initializeWithJson(api unsafe.Pointer, credentialsJson *C.char) *C.char {
	dart_api_dl.Init(api)

	gCredentialsJson := C.GoString(credentialsJson)

	var credentialsMap map[string]interface{}
	bCredentials := []byte(gCredentialsJson)
	err := json.Unmarshal(bCredentials, &credentialsMap)
	if err != nil {
		return C.CString(err.Error())
	}

	opt := option.WithCredentialsJSON(bCredentials)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return C.CString(err.Error())
	}

	cli, err := app.Auth(ctx)
	if err != nil {
		return C.CString(err.Error())
	}

	client = cli

	return nil
}

//export verifyToken
func verifyToken(port int64, idToken *C.char) {
	if !_checkClient(port) {
		return
	}
	gIdToken := C.GoString(idToken)

	go doVerifyToken(port, gIdToken)
}

func doVerifyToken(port int64, idToken string) {
	ctx := context.Background()
	sToken, err := client.VerifyIDToken(ctx, idToken)
	_handleResult(port, sToken, err)
}

//export verifyIDTokenAndCheckRevoked
func verifyIDTokenAndCheckRevoked(port int64, idToken *C.char) {
	if !_checkClient(port) {
		return
	}
	gIdToken := C.GoString(idToken)
	go doVerifyIDTokenAndCheckRevoked(port, gIdToken)
}

func doVerifyIDTokenAndCheckRevoked(port int64, idToken string) {
	ctx := context.Background()
	sToken, err := client.VerifyIDTokenAndCheckRevoked(ctx, idToken)
	_handleResult(port, sToken, err)
}

//export revokeRefreshTokens
func revokeRefreshTokens(port int64, idToken *C.char) {
	if !_checkClient(port) {
		return
	}
	gIdToken := C.GoString(idToken)
	go doRevokeRefreshTokens(port, gIdToken)
}

func doRevokeRefreshTokens(port int64, idToken string) {
	ctx := context.Background()
	err := client.RevokeRefreshTokens(ctx, idToken)
	_handleEmptyResult(port, err)

}

//export setCustomUserClaims
func setCustomUserClaims(port int64, uid *C.char, claims *C.char) {
	if !_checkClient(port) {
		return
	}

	gClaims := C.GoString(claims)
	gUid := C.GoString(uid)

	go doSetCustomClaims(port, gUid, gClaims)

}

func doSetCustomClaims(port int64, uid string, claims string) {
	var claimsMap map[string]interface{}
	bClaims := []byte(claims)
	err := json.Unmarshal(bClaims, &claimsMap)

	if err != nil {
		dart_api_dl.SendErrorToPort(port, err)
		return
	}

	ctx := context.Background()
	err = client.SetCustomUserClaims(ctx, uid, claimsMap)
	_handleEmptyResult(port, err)

}

//export createUser
func createUser(port int64, userJson *C.char) {
	if !_checkClient(port) {
		return
	}

	gUserJson := []byte(C.GoString(userJson))

	var userMap map[string]interface{}
	json.Unmarshal(gUserJson, &userMap)

	params := (&auth.UserToCreate{})

	if userMap["email"] != nil {
		params.Email(userMap["email"].(string))
	}
	if userMap["emailVerified"] != nil {
		params.EmailVerified(userMap["emailVerified"].(bool))
	}
	if userMap["disabled"] != nil {
		params.Disabled(userMap["disabled"].(bool))
	}
	if userMap["displayName"] != nil {
		params.DisplayName(userMap["displayName"].(string))
	}
	if userMap["password"] != nil {
		params.Password(userMap["password"].(string))
	}
	if userMap["phoneNumber"] != nil {
		params.PhoneNumber(userMap["phoneNumber"].(string))
	}
	if userMap["photoUrl"] != nil {
		params.PhotoURL(userMap["photoUrl"].(string))
	}
	if userMap["uid"] != nil {
		params.UID(userMap["uid"].(string))
	}

	ctx := context.Background()

	go doCreateUser(port, ctx, params)
}

func doCreateUser(port int64, ctx context.Context, params *auth.UserToCreate) {
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		dart_api_dl.SendErrorToPort(port, err)
		return
	}

	dart_api_dl.SendDataToPort(port, u.UID)
}

//export updateUser
func updateUser(port int64, userJson *C.char) {
	if !_checkClient(port) {
		return
	}

	gUserJson := []byte(C.GoString(userJson))

	var userMap map[string]interface{}
	json.Unmarshal(gUserJson, &userMap)
	gUid := userMap["uid"].(string)

	params := (&auth.UserToUpdate{})

	if userMap["email"] != nil {
		params.Email(userMap["email"].(string))
	}
	if userMap["emailVerified"] != nil {
		params.EmailVerified(userMap["emailVerified"].(bool))
	}
	if userMap["disabled"] != nil {
		params.Disabled(userMap["disabled"].(bool))
	}
	if userMap["displayName"] != nil {
		params.DisplayName(userMap["displayName"].(string))
	}
	if userMap["password"] != nil {
		params.Password(userMap["password"].(string))
	}
	if userMap["phoneNumber"] != nil {
		params.PhoneNumber(userMap["phoneNumber"].(string))
	}
	if userMap["photoUrl"] != nil {
		params.PhotoURL(userMap["photoUrl"].(string))
	}

	go doUpdateUser(port, gUid, params)
}

func doUpdateUser(port int64, gUid string, params *auth.UserToUpdate) {
	ctx := context.Background()
	_, err := client.UpdateUser(ctx, gUid, params)
	_handleEmptyResult(port, err)
}

//export deleteUser
func deleteUser(port int64, uid *C.char) {
	if !_checkClient(port) {
		return
	}

	gUid := C.GoString(uid)

	go doDeleteUser(port, gUid)
}

func doDeleteUser(port int64, gUid string) {
	ctx := context.Background()
	err := client.DeleteUser(ctx, gUid)
	_handleEmptyResult(port, err)
}

//export getCustomClaims
func getCustomClaims(port int64, uid *C.char) {
	if !_checkClient(port) {
		return
	}
	gUid := C.GoString(uid)

	go doGetCustomClaims(port, gUid)
}

func doGetCustomClaims(port int64, gUid string) {
	ctx := context.Background()
	usr, err := client.GetUser(ctx, gUid)
	if err != nil {
		dart_api_dl.SendErrorToPort(port, err)
		return
	}

	cclaims := usr.CustomClaims
	_handleResult(port, cclaims, err)
}

//export createCustomTokenWithClaims
func createCustomTokenWithClaims(port int64, uid *C.char, claims *C.char) {
	if !_checkClient(port) {
		return
	}

	gUid := C.GoString(uid)
	gClaims := C.GoString(claims)

	var claimsMap map[string]interface{}
	bClaims := []byte(gClaims)
	json.Unmarshal(bClaims, &claimsMap)

	go doCreateCustomTokenWithClaims(port, gUid, claimsMap)

}

func doCreateCustomTokenWithClaims(port int64, gUid string, claimsMap map[string]interface{}) {
	ctx := context.Background()
	idToken, err := client.CustomTokenWithClaims(ctx, gUid, claimsMap)

	if err != nil {
		dart_api_dl.SendErrorToPort(port, err)
		return
	}

	dart_api_dl.SendDataToPort(port, idToken)
}

//export createCustomToken
func createCustomToken(port int64, uid *C.char) {
	if !_checkClient(port) {
		return
	}

	gUid := C.GoString(uid)
	go doCreateCustomToken(port, gUid)
}

func doCreateCustomToken(port int64, gUid string) {
	ctx := context.Background()
	idToken, err := client.CustomToken(ctx, gUid)

	if err != nil {
		dart_api_dl.SendErrorToPort(port, err)
		return
	}

	dart_api_dl.SendDataToPort(port, idToken)
}

//export getUser
func getUser(port int64, uid *C.char) {
	if !_checkClient(port) {
		return
	}

	gUid := C.GoString(uid)

	go doGetUser(port, gUid)
}

func doGetUser(port int64, gUid string) {
	ctx := context.Background()
	usr, err := client.GetUser(ctx, gUid)
	_handleResult(port, usr, err)
}

//export getUserByEmail
func getUserByEmail(port int64, email *C.char) {
	if !_checkClient(port) {
		return
	}

	gEmail := C.GoString(email)

	go doGetUserByEmail(port, gEmail)
}

func doGetUserByEmail(port int64, email string) {
	ctx := context.Background()
	usr, err := client.GetUserByEmail(ctx, email)
	_handleResult(port, usr, err)
}

//export getUserByPhoneNumber
func getUserByPhoneNumber(port int64, phoneNumber *C.char) {
	if !_checkClient(port) {
		return
	}

	gPhoneNumber := C.GoString(phoneNumber)

	go doGetUserByPhoneNumber(port, gPhoneNumber)
}

func doGetUserByPhoneNumber(port int64, phoneNumber string) {
	ctx := context.Background()
	usr, err := client.GetUserByPhoneNumber(ctx, phoneNumber)
	_handleResult(port, usr, err)
}

func _handleResult(port int64, v interface{}, err error) {
	if err != nil {
		dart_api_dl.SendErrorToPort(port, err)
		return
	}

	var jToken []byte
	jToken, err = json.Marshal(v)
	if err != nil {
		dart_api_dl.SendErrorToPort(port, err)
		return
	}

	gToken := string(jToken)
	dart_api_dl.SendDataToPort(port, gToken)

}

func _handleEmptyResult(port int64, err error) {
	if err != nil {
		dart_api_dl.SendErrorToPort(port, err)
		return
	}

	dart_api_dl.SendDataToPort(port, "")
}

func _checkClient(port int64) bool {
	if client == nil {
		dart_api_dl.SendErrorToPort(port, errors.New(initError))
		return false
	}
	return true

}

func main() {
}
