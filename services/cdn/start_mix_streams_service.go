package cdn

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

// StartMixStreamsService invokes the cdn.StartMixStreamsService API synchronously
// api document: https://help.aliyun.com/api/cdn/startmixstreamsservice.html
func (client *Client) StartMixStreamsService(request *StartMixStreamsServiceRequest) (response *StartMixStreamsServiceResponse, err error) {
	response = CreateStartMixStreamsServiceResponse()
	err = client.DoAction(request, response)
	return
}

// StartMixStreamsServiceWithChan invokes the cdn.StartMixStreamsService API asynchronously
// api document: https://help.aliyun.com/api/cdn/startmixstreamsservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartMixStreamsServiceWithChan(request *StartMixStreamsServiceRequest) (<-chan *StartMixStreamsServiceResponse, <-chan error) {
	responseChan := make(chan *StartMixStreamsServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartMixStreamsService(request)
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

// StartMixStreamsServiceWithCallback invokes the cdn.StartMixStreamsService API asynchronously
// api document: https://help.aliyun.com/api/cdn/startmixstreamsservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartMixStreamsServiceWithCallback(request *StartMixStreamsServiceRequest, callback func(response *StartMixStreamsServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartMixStreamsServiceResponse
		var err error
		defer close(result)
		response, err = client.StartMixStreamsService(request)
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

// StartMixStreamsServiceRequest is the request struct for api StartMixStreamsService
type StartMixStreamsServiceRequest struct {
	*requests.RpcRequest
	MixStreamName  string           `position:"Query" name:"MixStreamName"`
	MixAppName     string           `position:"Query" name:"MixAppName"`
	MainStreamName string           `position:"Query" name:"MainStreamName"`
	MixType        string           `position:"Query" name:"MixType"`
	SecurityToken  string           `position:"Query" name:"SecurityToken"`
	MainDomainName string           `position:"Query" name:"MainDomainName"`
	MixTemplate    string           `position:"Query" name:"MixTemplate"`
	MixDomainName  string           `position:"Query" name:"MixDomainName"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	MainAppName    string           `position:"Query" name:"MainAppName"`
}

// StartMixStreamsServiceResponse is the response struct for api StartMixStreamsService
type StartMixStreamsServiceResponse struct {
	*responses.BaseResponse
	RequestId          string                                     `json:"RequestId" xml:"RequestId"`
	MixStreamsInfoList MixStreamsInfoListInStartMixStreamsService `json:"MixStreamsInfoList" xml:"MixStreamsInfoList"`
}

// CreateStartMixStreamsServiceRequest creates a request to invoke StartMixStreamsService API
func CreateStartMixStreamsServiceRequest() (request *StartMixStreamsServiceRequest) {
	request = &StartMixStreamsServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "StartMixStreamsService", "", "")
	return
}

// CreateStartMixStreamsServiceResponse creates a response to parse from StartMixStreamsService response
func CreateStartMixStreamsServiceResponse() (response *StartMixStreamsServiceResponse) {
	response = &StartMixStreamsServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
