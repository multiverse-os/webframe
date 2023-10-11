package session

import (
	"fmt"
	"net/http"
	"time"
)

type Session struct {
	//Id       record.Id
	//Data     record.Data
	Id       string
	Data     string
	ExpireAt time.Time
}

func New(expireAt time.Time) *Session {
	return &Session{
		ExpireAt: expireAt,
	}
}

func NewSession(w http.ResponseWriter, expireAt time.Time) string {
	// NOTE: This id holds order information, and other useful information
	// which is why we use it in addition to our scramblekeys
	fmt.Println("creating a new session")

	// TODO: Create the record in the session collection, so the cookie only
	// stores the ID we use to look up all the data. Making the cookie small,
	// and data auto-saved essentially.

	//newSession := New(expireAt)

	cookie := http.Cookie{
		Name: "sid",
		//Value:   Record.Id.String(),
		Value:   "create-session-id",
		Expires: expireAt,
	}
	http.SetCookie(w, &cookie)
	return ""
}

// TODO: Not sure about naming, need to differentiate the concept of a
// non-registered session from a registered user session, they share a lot but
// are fundamentally different as well
func NewRegisteredSession(sid, uid, password string) string {
	//TODO: Validate session can make this request
	//TODO: Add flash messages depending on outcome
	fmt.Println("A user is attempting to login with: ")
	fmt.Println("  UID     : ", uid)
	fmt.Println("  PASSWORD: ", password)
	//user, err := .session.Collections["session"].Get(uid)
	//TODO: Take database data and convert to user
	//if err != nil {
	//	session := self.Session[sid]
	//	if session.FailedLoginAttempts >= 6 {
	//		fmt.Println("session has failed too many login attempts, should just block it for x amount of time")
	//	} else {
	//		//TODO: Attempt login
	//		session.FailedLoginAttempts += 1
	//		fmt.Println(" USER JSON: " + user)
	//	}
	//} else {
	//	fmt.Println("[uid not found]")
	//}
	return ""
}
