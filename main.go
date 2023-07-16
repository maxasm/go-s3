package main

// The following are the packages you need.
// aws
// aws/session -> Create an AWS Session using a given region
// service/s3 -> Working with Amazon S3

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "fmt"
)

func main() {
	// set up the session with the right region
	session, err_session := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	})
	
	if err_session != nil {
		fmt.Printf("%s\n", err_session)
	}	

	// create an S3 service client
	svc := s3.New(session)
	
	// list buckets
	bkts, err_bkts := svc.ListBuckets(nil)
	
	if err_bkts != nil {
		fmt.Printf("%s\n", err_bkts)
	}
	
	fmt.Printf("%v\n", bkts)
}
