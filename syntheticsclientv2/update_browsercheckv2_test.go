//go:build unit_tests
// +build unit_tests

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

package syntheticsclientv2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var (
	updateBrowserCheckV2Body  = `{"test":{"name":"browser-beep-test","business_transactions":[{"name":"Synthetic transaction 1","steps":[{"name":"Go to URL","type":"go_to_url","url":"https://www.splunk.com","action":"go_to_url","options":{"url":"https://www.splunk.com"},"wait_for_nav":true},{"name":"Nexter step","type":"click_element","selector_type":"id","selector":"free-splunk-click-desktop","wait_for_nav":false}]}],"urlProtocol":"https://","startUrl":"www.splunk.com","location_ids":["aws-us-east-1"],"device_id":2,"frequency":5,"scheduling_strategy":"round_robin","active":true,"advanced_settings":{"authentication":{"username":"boopuser","password":"{{env.beep-var}}"},"cookies":[{"key":"super","value":"duper","domain":"www.batmansagent.com","path":"/boom/goes/beep"}],"headers":[{"name":"batman","value":"Agentoz","domain":"www.batmansagent.com"}],"host_overrides":[],"user_agent":"Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X; Splunk Synthetics) AppleWebKit/604.1.25 (KHTML, like Gecko) Version/11.0 Mobile/15A5304j Safari/604.1","verify_certificates":true}}}`
	inputBrowserCheckV2Update = BrowserCheckV2Input{}
)

func TestUpdateBrowserCheckV2(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/tests/browser/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		_, err := w.Write([]byte(updateBrowserCheckV2Body))
		if err != nil {
			t.Fatal(err)
		}
	})

	err := json.Unmarshal([]byte(updateBrowserCheckV2Body), &inputBrowserCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	resp, _, err := testClient.UpdateBrowserCheckV2(10, &inputBrowserCheckV2Update)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)

	if !reflect.DeepEqual(resp.Test.Name, inputBrowserCheckV2Update.Test.Name) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.Test.Name, inputBrowserCheckV2Update.Test.Name)
	}

}
