/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   18-Mar-2019
 * @Project: Vylm.io
 * @Filename: acm.go
 * @Last modified by:   inaki
 * @Last modified time: 11-Aug-2019
 * @Copyright: Iñaki Rodriguez
 */
package awsservices

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DoAction() {
	fmt.Println("Checking ACM Certificates...")

	fmt.Println("0")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewSharedCredentials("", "default"),
	})
	if err != nil {
		log.Fatal(err)
	}
	// credens, err := sess.Config.Credentials.Get()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(credens)

	fmt.Println("1")
	svc := s3.New(sess)
	fmt.Println("2")
	input := &s3.ListBucketsInput{}
	fmt.Println("3")
	result, err := svc.ListBuckets(input)
	fmt.Println("4")
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}
