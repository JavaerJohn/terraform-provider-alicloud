package rds

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

func (client *Client) DescribeBackups(request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
	response = CreateDescribeBackupsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeBackupsWithChan(request *DescribeBackupsRequest) (<-chan *DescribeBackupsResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackups(request)
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

func (client *Client) DescribeBackupsWithCallback(request *DescribeBackupsRequest, callback func(response *DescribeBackupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackups(request)
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

type DescribeBackupsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	BackupId             string           `position:"Query" name:"BackupId"`
	BackupLocation       string           `position:"Query" name:"BackupLocation"`
	BackupStatus         string           `position:"Query" name:"BackupStatus"`
	BackupMode           string           `position:"Query" name:"BackupMode"`
	StartTime            string           `position:"Query" name:"StartTime"`
	EndTime              string           `position:"Query" name:"EndTime"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeBackupsResponse struct {
	*responses.BaseResponse
	RequestId        string                 `json:"RequestId" xml:"RequestId"`
	TotalRecordCount string                 `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       string                 `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  string                 `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalBackupSize  int                    `json:"TotalBackupSize" xml:"TotalBackupSize"`
	Items            ItemsInDescribeBackups `json:"Items" xml:"Items"`
}

func CreateDescribeBackupsRequest() (request *DescribeBackupsRequest) {
	request = &DescribeBackupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeBackups", "rds", "openAPI")
	return
}

func CreateDescribeBackupsResponse() (response *DescribeBackupsResponse) {
	response = &DescribeBackupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
