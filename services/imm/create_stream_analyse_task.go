package imm

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

// CreateStreamAnalyseTask invokes the imm.CreateStreamAnalyseTask API synchronously
// api document: https://help.aliyun.com/api/imm/createstreamanalysetask.html
func (client *Client) CreateStreamAnalyseTask(request *CreateStreamAnalyseTaskRequest) (response *CreateStreamAnalyseTaskResponse, err error) {
	response = CreateCreateStreamAnalyseTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateStreamAnalyseTaskWithChan invokes the imm.CreateStreamAnalyseTask API asynchronously
// api document: https://help.aliyun.com/api/imm/createstreamanalysetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateStreamAnalyseTaskWithChan(request *CreateStreamAnalyseTaskRequest) (<-chan *CreateStreamAnalyseTaskResponse, <-chan error) {
	responseChan := make(chan *CreateStreamAnalyseTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateStreamAnalyseTask(request)
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

// CreateStreamAnalyseTaskWithCallback invokes the imm.CreateStreamAnalyseTask API asynchronously
// api document: https://help.aliyun.com/api/imm/createstreamanalysetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateStreamAnalyseTaskWithCallback(request *CreateStreamAnalyseTaskRequest, callback func(response *CreateStreamAnalyseTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateStreamAnalyseTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateStreamAnalyseTask(request)
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

// CreateStreamAnalyseTaskRequest is the request struct for api CreateStreamAnalyseTask
type CreateStreamAnalyseTaskRequest struct {
	*requests.RpcRequest
	GrabType        string           `position:"Query" name:"GrabType"`
	Project         string           `position:"Query" name:"Project"`
	StartTime       string           `position:"Query" name:"StartTime"`
	NotifyEndpoint  string           `position:"Query" name:"NotifyEndpoint"`
	StreamUri       string           `position:"Query" name:"StreamUri"`
	NotifyTopicName string           `position:"Query" name:"NotifyTopicName"`
	EndTime         string           `position:"Query" name:"EndTime"`
	SaveType        requests.Boolean `position:"Query" name:"SaveType"`
	Interval        string           `position:"Query" name:"Interval"`
	TgtUri          string           `position:"Query" name:"TgtUri"`
}

// CreateStreamAnalyseTaskResponse is the response struct for api CreateStreamAnalyseTask
type CreateStreamAnalyseTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	TaskType  string `json:"TaskType" xml:"TaskType"`
}

// CreateCreateStreamAnalyseTaskRequest creates a request to invoke CreateStreamAnalyseTask API
func CreateCreateStreamAnalyseTaskRequest() (request *CreateStreamAnalyseTaskRequest) {
	request = &CreateStreamAnalyseTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateStreamAnalyseTask", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateStreamAnalyseTaskResponse creates a response to parse from CreateStreamAnalyseTask response
func CreateCreateStreamAnalyseTaskResponse() (response *CreateStreamAnalyseTaskResponse) {
	response = &CreateStreamAnalyseTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
