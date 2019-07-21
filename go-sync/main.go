// Copyright 2019 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

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
		confData, err := json.MarshalIndent(dsnap.Data(), "", "  ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = ioutil.WriteFile(p, confData, 0644)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}
