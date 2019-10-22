package dnsknocker

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

// DeleteResourceRecord invokes the dnsknocker.DeleteResourceRecord API synchronously
// api document: https://help.aliyun.com/api/dnsknocker/deleteresourcerecord.html
func (client *Client) DeleteResourceRecord(request *DeleteResourceRecordRequest) (response *DeleteResourceRecordResponse, err error) {
	response = CreateDeleteResourceRecordResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteResourceRecordWithChan invokes the dnsknocker.DeleteResourceRecord API asynchronously
// api document: https://help.aliyun.com/api/dnsknocker/deleteresourcerecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteResourceRecordWithChan(request *DeleteResourceRecordRequest) (<-chan *DeleteResourceRecordResponse, <-chan error) {
	responseChan := make(chan *DeleteResourceRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteResourceRecord(request)
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

// DeleteResourceRecordWithCallback invokes the dnsknocker.DeleteResourceRecord API asynchronously
// api document: https://help.aliyun.com/api/dnsknocker/deleteresourcerecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteResourceRecordWithCallback(request *DeleteResourceRecordRequest, callback func(response *DeleteResourceRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteResourceRecordResponse
		var err error
		defer close(result)
		response, err = client.DeleteResourceRecord(request)
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

// DeleteResourceRecordRequest is the request struct for api DeleteResourceRecord
type DeleteResourceRecordRequest struct {
	*requests.RpcRequest
	AccessID      string `position:"Body" name:"AccessID"`
	RRTTL         string `position:"Body" name:"RRTTL"`
	AccessSecret  string `position:"Body" name:"AccessSecret"`
	RRLine        string `position:"Body" name:"RRLine"`
	DomainName    string `position:"Body" name:"DomainName"`
	RRValue       string `position:"Body" name:"RRValue"`
	ZoneName      string `position:"Body" name:"ZoneName"`
	TransactionId string `position:"Body" name:"TransactionId"`
	Group         string `position:"Body" name:"Group"`
	RRType        string `position:"Body" name:"RRType"`
}

// DeleteResourceRecordResponse is the response struct for api DeleteResourceRecord
type DeleteResourceRecordResponse struct {
	*responses.BaseResponse
	ResultCode    string `json:"ResultCode" xml:"ResultCode"`
	ResultMessage string `json:"ResultMessage" xml:"ResultMessage"`
	Success       string `json:"Success" xml:"Success"`
	TransactionId string `json:"TransactionId" xml:"TransactionId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteResourceRecordRequest creates a request to invoke DeleteResourceRecord API
func CreateDeleteResourceRecordRequest() (request *DeleteResourceRecordRequest) {
	request = &DeleteResourceRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DnsKnocker", "2019-09-10", "DeleteResourceRecord", "dns_knocker", "openAPI")
	return
}

// CreateDeleteResourceRecordResponse creates a response to parse from DeleteResourceRecord response
func CreateDeleteResourceRecordResponse() (response *DeleteResourceRecordResponse) {
	response = &DeleteResourceRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}