package dataworks_public

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

// CreateNodeComplement invokes the dataworks_public.CreateNodeComplement API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/createnodecomplement.html
func (client *Client) CreateNodeComplement(request *CreateNodeComplementRequest) (response *CreateNodeComplementResponse, err error) {
	response = CreateCreateNodeComplementResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNodeComplementWithChan invokes the dataworks_public.CreateNodeComplement API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/createnodecomplement.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateNodeComplementWithChan(request *CreateNodeComplementRequest) (<-chan *CreateNodeComplementResponse, <-chan error) {
	responseChan := make(chan *CreateNodeComplementResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNodeComplement(request)
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

// CreateNodeComplementWithCallback invokes the dataworks_public.CreateNodeComplement API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/createnodecomplement.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateNodeComplementWithCallback(request *CreateNodeComplementRequest, callback func(response *CreateNodeComplementResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNodeComplementResponse
		var err error
		defer close(result)
		response, err = client.CreateNodeComplement(request)
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

// CreateNodeComplementRequest is the request struct for api CreateNodeComplement
type CreateNodeComplementRequest struct {
	*requests.RpcRequest
	ProjectEnv     string           `position:"Body" name:"ProjectEnv"`
	StartBizDate   string           `position:"Body" name:"StartBizDate"`
	Parallelism    requests.Boolean `position:"Body" name:"Parallelism"`
	BizBeginTime   string           `position:"Body" name:"BizBeginTime"`
	EndBizDate     string           `position:"Body" name:"EndBizDate"`
	IncludeNodeIds string           `position:"Body" name:"IncludeNodeIds"`
	BizEndTime     string           `position:"Body" name:"BizEndTime"`
	Name           string           `position:"Body" name:"Name"`
	ExcludeNodeIds string           `position:"Body" name:"ExcludeNodeIds"`
	NodeId         requests.Integer `position:"Body" name:"NodeId"`
	NodeProjectId  requests.Integer `position:"Body" name:"NodeProjectId"`
}

// CreateNodeComplementResponse is the response struct for api CreateNodeComplement
type CreateNodeComplementResponse struct {
	*responses.BaseResponse
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           int64  `json:"Data" xml:"Data"`
}

// CreateCreateNodeComplementRequest creates a request to invoke CreateNodeComplement API
func CreateCreateNodeComplementRequest() (request *CreateNodeComplementRequest) {
	request = &CreateNodeComplementRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateNodeComplement", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateNodeComplementResponse creates a response to parse from CreateNodeComplement response
func CreateCreateNodeComplementResponse() (response *CreateNodeComplementResponse) {
	response = &CreateNodeComplementResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
