[![CircleCI](https://circleci.com/gh/knocus/validate.svg?style=svg)](https://circleci.com/gh/knocus/validate)
[![Coverage Status](https://coveralls.io/repos/github/knocus/validate/badge.svg?branch=master)](https://coveralls.io/github/knocus/validate?branch=master)
# Validate
A golang package for simple validation of email addresses. This package is to be used along with sending a confirmation mail.

## Installation 

    go get -u github.com/knocus/validate
   
 ## Usage
 

    import (
	    "github.com/knocus/validate"
	)
	func main(){
		err := 
		if(err == nil){
			/* send confirmation link / other logic here */
 		} else {
			/* prompt the user about error / ask email again. */
		}
	}