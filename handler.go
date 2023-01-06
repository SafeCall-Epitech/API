package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

type Credentials struct {
	Uri string `json:"uri"`
}

func genUUID() string {
	id := uuid.New()
	return id.String()
}

// This function will get the uri in the json file to id to the db
func getCredentials() string {
	fileContent, err := os.Open("config.json")

	if err != nil {
		log.Fatal(err)
		return ""
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	res := Credentials{}
	json.Unmarshal([]byte(byteResult), &res)

	return res.Uri
}

func LoginHandler(id, psw string) string {
	uri := getCredentials()
	users := GetUsers(uri)

	for _, info := range users {
		if info["login"] == id && info["psw"] == psw {
			userID := info["login"].(string)
			return userID
		}
	}

	return "failed"
}

func userToProto(username, psw string) UserMessage {
	user := UserMessage{
		Id:       1,
		Username: username,
		Password: psw,
		Settings: &SettingsMessage{
			DoNotDisturb: true,
			Language:     "eng",
		},
	}
	return user
}

func RegisterHandler(id, psw string) string {
	uri := getCredentials()
	users := GetUsers(uri)

	if len(id) < 5 {
		return "id too short" // Id too short
	}

	if len(psw) < 7 {
		return "password too short" // password too short
	}

	for _, info := range users {
		if info["login"] == id {
			return "Id already taken" // Id already taken
		}
	}

	protoUser := userToProto(id, psw)
	binary, _ := proto.Marshal(&protoUser)
	if AddUser(uri, id, psw, string(binary)) != true {
		return "Unknown error while registration"
	}
	if CreateProfile(uri, id) != true {
		return "Unknown error while profile creation"
	}
	return "200"
}

func postProfileHandler(endpoint, userID, data string) string {
	uri := getCredentials()
	if len(userID) > 45 {
		return "User ID too long"
	}

	if endpoint == "FullName" {
		if len(data) > 30 {
			return "Full Name too long"
		}
	}

	if endpoint == "PhoneNB" {
		if len(data) > 15 {
			return "Phone NB too long"
		}
	}

	UpdateProfile(uri, endpoint, userID, data)
	return "success"
}

func getProfileHandler(userID string) Profile {
	resp := getProfile(userID)

	return StringToProfile(resp)
}

func searchName(userID string) string {
	resp := searchNameQuery(userID)
	tmp := 1
	dest := strings.Split(resp, "\"")
	result := ""

	a := len(dest) - 2
	a = a / 4

	for a != 0 {
		tmp = tmp + 4
		result = result + "," + dest[tmp]
		a = a - 1
	}

	return result[1:]
}
