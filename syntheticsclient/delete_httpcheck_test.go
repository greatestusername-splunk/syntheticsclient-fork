// Copyright 2021 Splunk, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package syntheticsclient

import (
	"net/http"
	"testing"
)

var (
	deleteHttpRespBody = `{"result":"success","message":"testcase successfully deleted","errors":[]}`
)

func TestDeleteHttpCheck(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/checks/http/19", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.Write([]byte(deleteHttpRespBody))
	})

	resp, err := testClient.DeleteHttpCheck(19)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Message != "testcase successfully deleted" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Message, "testcase successfully deleted")
	}
	if resp.Result != "success" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Result, "success")
	}
}
