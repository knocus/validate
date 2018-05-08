
# Validate
(WORK-IN-PROGRESS) A golang package for simple validation of email addresses. This package is to be used along with sending a confirmation mail.

## Installation 
```go
    go get -u github.com/knocus/validate
```   
 ## Usage
 
```go
    import (
	    "github.com/knocus/validate"
	)
	func main(){
		err := validate.Email("some.email@gmail.com");
		if(err == nil){
			/* send confirmation link / other logic here */
 		} else {
			/* prompt the user about error / ask email again. */
		}
    }
```