package awshelper

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func S3Client() (*s3.S3, error) {
	awsAccessKey := os.Getenv("ACCESS_KEY")
	awsSecretKey := os.Getenv("SECRET")

	// Create a new session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	})

	if err != nil {
		fmt.Println("Error creating session,", err)
		return nil, err
	}
	// Create a new S3 service client
	svc := s3.New(sess)

	return svc, nil
}

func ConvertS3Compatible(name string) string {
	// Convert to lowercase
	name = strings.ToLower(name)

	// Replace spaces with hyphens
	name = strings.ReplaceAll(name, " ", "-")

	// Remove any characters that are not alphanumeric, hyphen, or period
	reg, _ := regexp.Compile("[^a-z0-9-.]+")
	name = reg.ReplaceAllString(name, "")

	// Ensure the name is between 3 and 63 characters long
	if len(name) < 3 {
		name = "s3-" + name
	}
	if len(name) > 63 {
		name = name[:63]
	}

	return name
}

func SendFileToS3(svc *s3.S3, bucket string, key string, file string) error {
	// Open the file
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// Create a new S3 uploader
	uploader := s3manager.NewUploaderWithClient(svc)

	// Upload the file
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   f,
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return err
	}

	return nil
}

func CreateBucketWithPublicAccess(reponame string, prefix string) (*string, error) {
	svc, err := S3Client()
	s3CompatibleName := ConvertS3Compatible(prefix + "-" + reponame)

	if err != nil {
		fmt.Println("Error creating S3 client service,", err)
		return nil, nil
	}

	output, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(s3CompatibleName),
	})

	if err != nil {
		fmt.Println("Error creating bucket,", err)
		return nil, nil
	}

	_, err = svc.DeletePublicAccessBlock(&s3.DeletePublicAccessBlockInput{
		Bucket: aws.String(s3CompatibleName),
	})

	if err != nil {
		fmt.Println("Error removing public access block,", err)
		return nil, nil
	}

	policy := `{
		"Version":"2012-10-17",
		"Statement":[{
			"Sid":"PublicReadGetObject",
			"Effect":"Allow",
			"Principal": "*",
			"Action":["s3:GetObject"],
			"Resource":["arn:aws:s3:::` + s3CompatibleName + `/*"]
		}]
	}`

	_, err = svc.PutBucketPolicy(&s3.PutBucketPolicyInput{
		Bucket: aws.String(s3CompatibleName),
		Policy: aws.String(policy),
	})

	if err != nil {
		fmt.Println("Error setting bucket policy", err)
		return nil, nil
	}

	return output.Location, nil
}
