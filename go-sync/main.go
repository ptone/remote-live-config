package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

var projectID = "ptone-serverless"

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	p := "/tmp/demo.json"
	k := strings.Replace(p, "/", ":", -1)
	docref := client.Collection("file-configs").Doc(k)
	_, err = client.Collection("file-configs").Doc(":::writetest").Set(ctx, map[string]interface{}{
		"message": "Warning - DB should be read only",
	})
	if err == nil {
		log.Fatal("config sync should not have write permission")
	}
	dociter := docref.Snapshots(ctx)
	for {
		dsnap, err := dociter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(dsnap.Data())
		confData, err := json.MarshalIndent(dsnap.Data(), "", "  ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// jsonStr := string(empData)
		err = ioutil.WriteFile(p, confData, 0644)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// fmt.Println("The JSON data is:")
		// fmt.Println(jsonStr)
	}

}
