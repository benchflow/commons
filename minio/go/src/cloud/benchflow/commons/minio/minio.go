package minio

import (
	"github.com/minio/minio-go"
	"os/exec"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

//Number of hash characters
const numOfHashCharacters int = 4

/*
// We keep thos function in case the API fails because of a Minio update
func callMinioClient(fileName string, minioHost string, minioKey string) {
		//TODO: change, we are using sudo to elevate the priviledge in the container, but it is not nice
		//NOTE: it seems that the commands that are not in PATH, should be launched using sh -c
		log.Printf("sh -c sudo /app/mc --quiet cp " + fileName + " " + minioHost + "/runs/" + minioKey)
		cmd := exec.Command("sh", "-c", "sudo /app/mc --quiet cp " + fileName + " " + minioHost + "/runs/" + minioKey)
    	var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
		    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		    return
		}
		fmt.Println("Result: " + out.String())
}
*/

//Function to send a gzip file to Minio
func SendGzipToMinio(fileName string, minioHost string, minioPort string, minioKey string, accessKey string, secretAccessKey string) {
	sendToMinio(fileName, minioHost, minioPort, minioKey, accessKey, secretAccessKey, "application/gzip")
	}

//Function to send a file to Minio using the Minio client for Golang
func sendToMinio(fileName string, minioHost string, minioPort string, minioKey string, accessKey string, secretAccessKey string, fileType string) error {
	minioClient, err := minio.New(minioHost+":"+minioPort, accessKey, secretAccessKey, true)
	if err != nil {
		fmt.Println(err)
	    return err
	    }
	_, err = minioClient.FPutObject("runs", minioKey, fileName, fileType)
	if err != nil {
	    fmt.Println(err)
	    return err
		}
	return nil
	}

//Function to hash a key with MD5
func hashKey(key string) string {
	hasher := md5.New()
    hasher.Write([]byte(key))
    hashString := hex.EncodeToString(hasher.Sum(nil))
	return (hashString[:numOfHashCharacters]+"/"+key)
	}

//Function to generate the Minio key including the hash
func GenerateKey(fileName string, trialID string, experimentID string, containerName string, collectorName string, dataName string) string {
	key := experimentID+"/"+trialID+"/"+containerName+"/"+collectorName+"/"+dataName+"/"+fileName
	return hashKey(key)
}

//Function to compress a file with Gzip
func GzipFile(fileName string) {
	cmd := exec.Command("gzip", fileName)
	err := cmd.Start()
	cmd.Wait()
	if err != nil {
		panic(err)
		}
}