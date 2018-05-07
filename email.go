package validate

import (
	"errors"
	"net"
	"strings"
)



var (
	ErrInvalidEmail = errors.New("Invalid email address")
)

func Email(email string) (error) {

	host, err := ParseHost(email) 
	if err != nil {
		return ErrInvalidEmail
	}
	
	_, err = net.LookupMX(host)
	if err != nil {
		return ErrInvalidEmail
 	} 
	return nil
}


func ParseHost(email string) (host string, err error){
	
	at := strings.LastIndex(email, "@")

	if at >= 0 {
		host = email[at+1:]
		err = nil
	} else {
		host = ""
		err = ErrInvalid
	}
	
	return
}
