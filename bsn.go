/*
 * http://www.apache.org/licenses/LICENSE-2.0.txt
 *
 * Copyright 2017 OpsVision Solutions
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package bsn

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	//"io"
	"net"
	"net/http"
	//"os"
	"time"
)

type Client struct {
	Controller    string `json:"controller"`
	SessionCookie string `json:"session_cookie"`
}

func New(controller string, port string) *Client {
	client := &Client{}
	client.Controller = fmt.Sprintf("https://%s:%s", controller, port)

	return client
}

// GetSwitches retrieves a list of switches from the BSN controller
func (c *Client) GetSwitches() Switches {
	var endpoint = fmt.Sprintf("%s/api/v1/data/controller/applications/bcf/info/fabric/switch", c.Controller)
	var switches Switches

	// Create the request
	req, err := http.NewRequest("GET", endpoint, nil)

	// Set headers
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Cookie", c.SessionCookie)

	// Send the payload
	resp, err := c.getClient().Do(req)
	if err != nil {
		fmt.Printf("Error: %s\n" + err.Error())
	}
	defer resp.Body.Close()

	// Decode the response
	json.NewDecoder(resp.Body).Decode(&switches)
	//io.Copy(os.Stdout, resp.Body)
	return switches
}

// Authenticate with the BSN controller
func (c *Client) Authenticate(creds *Credentials) {
	var endpoint = fmt.Sprintf("%s/api/v1/auth/login", c.Controller)
	var payload bytes.Buffer
	var authresp AuthResponse

	// Encode the payload; username and password
	json.NewEncoder(&payload).Encode(creds)

	// Create the request
	req, err := http.NewRequest("POST", endpoint, &payload)

	// Set headers
	req.Header.Set("content-type", "application/json")

	// Send the payload
	resp, err := c.getClient().Do(req)
	if err != nil {
		fmt.Printf("Error: %s\n" + err.Error())
	}
	defer resp.Body.Close()

	// Decode the response
	json.NewDecoder(resp.Body).Decode(&authresp)

	// Store the session cookie
	c.SessionCookie = "session_cookie=" + authresp.SessionCookie
}

// getClient sets up our http client
func (c *Client) getClient() *http.Client {
	// Setup transport settings
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	// Create a client
	client := &http.Client{
		Timeout:   time.Second * 10, // 10 second timeout
		Transport: tr,
	}

	return client
}
