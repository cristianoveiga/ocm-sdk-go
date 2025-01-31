/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"io"
	"net/http"
)

func readHTPasswdUserDeleteRequest(request *HTPasswdUserDeleteServerRequest, r *http.Request) error {
	return nil
}
func writeHTPasswdUserDeleteRequest(request *HTPasswdUserDeleteRequest, writer io.Writer) error {
	return nil
}
func readHTPasswdUserDeleteResponse(response *HTPasswdUserDeleteResponse, reader io.Reader) error {
	return nil
}
func writeHTPasswdUserDeleteResponse(response *HTPasswdUserDeleteServerResponse, w http.ResponseWriter) error {
	return nil
}
func readHTPasswdUserGetRequest(request *HTPasswdUserGetServerRequest, r *http.Request) error {
	return nil
}
func writeHTPasswdUserGetRequest(request *HTPasswdUserGetRequest, writer io.Writer) error {
	return nil
}
func readHTPasswdUserGetResponse(response *HTPasswdUserGetResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalHTPasswdUser(reader)
	return err
}
func writeHTPasswdUserGetResponse(response *HTPasswdUserGetServerResponse, w http.ResponseWriter) error {
	return MarshalHTPasswdUser(response.body, w)
}
func readHTPasswdUserUpdateRequest(request *HTPasswdUserUpdateServerRequest, r *http.Request) error {
	var err error
	request.body, err = UnmarshalHTPasswdUser(r.Body)
	return err
}
func writeHTPasswdUserUpdateRequest(request *HTPasswdUserUpdateRequest, writer io.Writer) error {
	return MarshalHTPasswdUser(request.body, writer)
}
func readHTPasswdUserUpdateResponse(response *HTPasswdUserUpdateResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalHTPasswdUser(reader)
	return err
}
func writeHTPasswdUserUpdateResponse(response *HTPasswdUserUpdateServerResponse, w http.ResponseWriter) error {
	return MarshalHTPasswdUser(response.body, w)
}
