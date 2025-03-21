package main

import "fmt"

//Struct
type messagesToSend struct {
	message string
	sender user
	recepient user
}

//Embedded struct
type sender struct {
	rateLimit int
	user
}

//Nested Struct
type user struct {
	name string
	number int
	Membership
}

type authenticationInfo struct {
	username string
	password string
}

type Membership struct {
	Type string
	MessageCharLimit int
}

func canSendMessages(mToSend messagesToSend) bool {
	if(mToSend.sender.number == 0) {
		return false
	}

	if(mToSend.sender.name == "") {
		return false
	}

	if(mToSend.recepient.number == 0) {
		return false
	}

	if(mToSend.recepient.name == "") {
		return false
	}

	return true
}

// Example of struct method
func (authInfo authenticationInfo) getBasicAuth () string {
	authUsername := authInfo.username
	authPasswrd := authInfo.password

	return fmt.Sprintf("Authorization: Basic %s:%s", authUsername, authPasswrd)
}

func newUser (name string, membershipType string) user {
	var messageCharLimit int

	if membershipType == "premium" {
		messageCharLimit = 1000
	} else {
		messageCharLimit = 100
	}

	return user {
		name: name,
		number: 9089098908,
		Membership: Membership{
			Type: membershipType,
			MessageCharLimit: messageCharLimit,
		},
	}
}

func (usr user) sendMessage (message string, messageLength int) (string, bool) {
	if messageLength > usr.MessageCharLimit {
		return "", false
	}

	return message, true
}

func main() {
	fmt.Println("Learning Go")

	sender1 := user {
		name: "Nitin",
		number: 1234567890,
	}

	//Example of an embedded struct
	embeddedSender := sender {
		rateLimit: 100,
		user: user{
			name: "Nitin",
			number: 9807654321,
		},
	}

	fmt.Println(embeddedSender.name, embeddedSender.number, embeddedSender.rateLimit, embeddedSender.user)

	recepient := user {
		name: "System",
		number: 0,
	}

	mToSend := messagesToSend {
		message: "Hey there i am learning go",
		sender: sender1,
		recepient: recepient,
	}

	canSendMsg := canSendMessages(mToSend)

	fmt.Printf("Can send message %v\n", canSendMsg);

	authInfo := authenticationInfo {
		username: "nitin.test@example.com",
		password: "testPassword",
	}

	fmt.Println(authInfo.getBasicAuth())

	createdUser := newUser("Shubham", "premium")
	msg := "Hey there i am learning golang"

	sentMsg, isMessageSent := createdUser.sendMessage(msg, len(msg))

	fmt.Println(sentMsg, isMessageSent)
}