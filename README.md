# AWS-Tools-Extra
Different utils to make simple tasks easily:
- S3 UPLOADER
- INSTANCE SCHEDULER PANEL
- INSTANCE WEEK REPORT
- SSL CERTIFICATE VALIDITY

### INSTRUCTIONS

##### Download go client first:
*https://golang.org/dl/*
##### Install brew
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
##### Install external dependencies
- *go get -v -u github.com/schollz/progressbar*
- *go get -v -u github.com/gen2brain/raylib-go/raylib*
- *go get fyne.io/fyne*
- *go get -u github.com/aws/aws-sdk-go/...*

##### Run the create client
*go run main.go*

##### Create binary (within /src)
*go build*
