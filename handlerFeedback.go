package main

type Feedback struct {
	Username string `bson:"Username"`
	Date     string `bson:"Date"`
	Message  string `bson:"Message"`
	Status   string `bson:"Status"`
}

func NewFeedbackHandler(user, date, message string) {
	url := getCredentials()
	message = checkForBannedWords(message)
	feedback := Feedback{user, date, message, "NEW"}
	AddFeedback(url.Uri, feedback)
}

func EditFeedbackHandler(user, date, state string) bool {
	url := getCredentials()
	UpdateFeedback(url.Uri, user, date, state)
	return true
}

func GetFeedbackHandler() []Feedback {
	url := getCredentials()
	resp, _ := GetFeedbacks(url.Uri)
	return resp
}

func DelFeedbackHandler(user, date string) string {
	url := getCredentials()
	resp := DeleteFeedback(url.Uri, user, date)

	if !resp { // if resp == false {
		return "No feedback found"
	}
	return "Feedback correctly deleted"
}
