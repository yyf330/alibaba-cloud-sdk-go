package vod

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SubmitAIJob invokes the vod.SubmitAIJob API synchronously
// api document: https://help.aliyun.com/api/vod/submitaijob.html
func (client *Client) SubmitAIJob(request *SubmitAIJobRequest) (response *SubmitAIJobResponse, err error) {
	response = CreateSubmitAIJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitAIJobWithChan invokes the vod.SubmitAIJob API asynchronously
// api document: https://help.aliyun.com/api/vod/submitaijob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitAIJobWithChan(request *SubmitAIJobRequest) (<-chan *SubmitAIJobResponse, <-chan error) {
	responseChan := make(chan *SubmitAIJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitAIJob(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SubmitAIJobWithCallback invokes the vod.SubmitAIJob API asynchronously
// api document: https://help.aliyun.com/api/vod/submitaijob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitAIJobWithCallback(request *SubmitAIJobRequest, callback func(response *SubmitAIJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitAIJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitAIJob(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SubmitAIJobRequest is the request struct for api SubmitAIJob
type SubmitAIJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	UserData             string `position:"Query" name:"UserData"`
	Types                string `position:"Query" name:"Types"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	Input                string `position:"Query" name:"Input"`
	Config               string `position:"Query" name:"Config"`
}

// SubmitAIJobResponse is the response struct for api SubmitAIJob
type SubmitAIJobResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	AIJobList AIJobListInSubmitAIJob `json:"AIJobList" xml:"AIJobList"`
}

// CreateSubmitAIJobRequest creates a request to invoke SubmitAIJob API
func CreateSubmitAIJobRequest() (request *SubmitAIJobRequest) {
	request = &SubmitAIJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SubmitAIJob", "vod", "openAPI")
	return
}

// CreateSubmitAIJobResponse creates a response to parse from SubmitAIJob response
func CreateSubmitAIJobResponse() (response *SubmitAIJobResponse) {
	response = &SubmitAIJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
