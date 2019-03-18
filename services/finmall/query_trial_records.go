package finmall

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

// QueryTrialRecords invokes the finmall.QueryTrialRecords API synchronously
// api document: https://help.aliyun.com/api/finmall/querytrialrecords.html
func (client *Client) QueryTrialRecords(request *QueryTrialRecordsRequest) (response *QueryTrialRecordsResponse, err error) {
	response = CreateQueryTrialRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTrialRecordsWithChan invokes the finmall.QueryTrialRecords API asynchronously
// api document: https://help.aliyun.com/api/finmall/querytrialrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTrialRecordsWithChan(request *QueryTrialRecordsRequest) (<-chan *QueryTrialRecordsResponse, <-chan error) {
	responseChan := make(chan *QueryTrialRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTrialRecords(request)
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

// QueryTrialRecordsWithCallback invokes the finmall.QueryTrialRecords API asynchronously
// api document: https://help.aliyun.com/api/finmall/querytrialrecords.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryTrialRecordsWithCallback(request *QueryTrialRecordsRequest, callback func(response *QueryTrialRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTrialRecordsResponse
		var err error
		defer close(result)
		response, err = client.QueryTrialRecords(request)
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

// QueryTrialRecordsRequest is the request struct for api QueryTrialRecords
type QueryTrialRecordsRequest struct {
	*requests.RpcRequest
	UserId string `position:"Query" name:"UserId"`
}

// QueryTrialRecordsResponse is the response struct for api QueryTrialRecords
type QueryTrialRecordsResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateQueryTrialRecordsRequest creates a request to invoke QueryTrialRecords API
func CreateQueryTrialRecordsRequest() (request *QueryTrialRecordsRequest) {
	request = &QueryTrialRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("finmall", "2018-07-23", "QueryTrialRecords", "finmall", "openAPI")
	return
}

// CreateQueryTrialRecordsResponse creates a response to parse from QueryTrialRecords response
func CreateQueryTrialRecordsResponse() (response *QueryTrialRecordsResponse) {
	response = &QueryTrialRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}