package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
    "fmt"
	"os"
)

func main() {
	// set up the session with the right region and load the credentials.
	session, err_session := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	})
	
	if err_session != nil {
		fmt.Printf("%s\n", err_session)
	}	
	
	// the bucket name to which I am uploading the files.
	bucket_name := "test-bucket-cindy"
	
	f, err_f := os.Open("cat.jpg")
	if err_f != nil {
		fmt.Printf("%s\n", err_f)
		os.Exit(1)
	}
	defer f.Close()
	
	// create an new uploader and pass it the session.
	uploader := s3manager.NewUploader(session)	
	
	res, res_err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket_name),
		Key: aws.String("cat.jpg"),	
		Body: f,
	})
	
	if res_err != nil {
		fmt.Printf("%s\n", res_err)
	}
	
	fmt.Printf("%v\n", res)
}
