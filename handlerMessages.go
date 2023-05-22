package main

type Messages struct {
	Sender  string `bson:"Sender"`
	Message string `bson:"Message"`
}

func GetConversation(userID string) []string {
	url := "http://facteur:3000/get_all_conv/" + userID
	resp := getAllConvQuery(url)
	return resp
}

func GetMessagesHandler(userID, friendID string) []Messages {
	url := "http://facteur:3000/conv/" + userID + "/" + friendID
	resp := getFromMessage(url)
	var messages []Messages
	for _, v := range resp {
		tmp := Messages{
			Sender:  v[0],
			Message: v[1],
		}
		messages = append(messages, tmp)
	}

	return messages
}

func PostMessageHandler(userID, friendID, message string) {
	url := "http://facteur:3000/send_message"
	postFacteur(url, userID, friendID, message)
}
