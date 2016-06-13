// Copyright 2015-2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-agent/agent/ec2 (interfaces: EC2MetadataClient,HttpClient)

package mock_ec2

import (
	http "net/http"

	ec2 "github.com/aws/amazon-ecs-agent/agent/ec2"
	gomock "github.com/golang/mock/gomock"
)

// Mock of EC2MetadataClient interface
type MockEC2MetadataClient struct {
	ctrl     *gomock.Controller
	recorder *_MockEC2MetadataClientRecorder
}

// Recorder for MockEC2MetadataClient (not exported)
type _MockEC2MetadataClientRecorder struct {
	mock *MockEC2MetadataClient
}

func NewMockEC2MetadataClient(ctrl *gomock.Controller) *MockEC2MetadataClient {
	mock := &MockEC2MetadataClient{ctrl: ctrl}
	mock.recorder = &_MockEC2MetadataClientRecorder{mock}
	return mock
}

func (_m *MockEC2MetadataClient) EXPECT() *_MockEC2MetadataClientRecorder {
	return _m.recorder
}

func (_m *MockEC2MetadataClient) DefaultCredentials() (*ec2.RoleCredentials, error) {
	ret := _m.ctrl.Call(_m, "DefaultCredentials")
	ret0, _ := ret[0].(*ec2.RoleCredentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEC2MetadataClientRecorder) DefaultCredentials() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DefaultCredentials")
}

func (_m *MockEC2MetadataClient) InstanceIdentityDocument() (*ec2.InstanceIdentityDocument, error) {
	ret := _m.ctrl.Call(_m, "InstanceIdentityDocument")
	ret0, _ := ret[0].(*ec2.InstanceIdentityDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEC2MetadataClientRecorder) InstanceIdentityDocument() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstanceIdentityDocument")
}

func (_m *MockEC2MetadataClient) ReadResource(_param0 string) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "ReadResource", _param0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockEC2MetadataClientRecorder) ReadResource(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadResource", arg0)
}

// Mock of HttpClient interface
type MockHttpClient struct {
	ctrl     *gomock.Controller
	recorder *_MockHttpClientRecorder
}

// Recorder for MockHttpClient (not exported)
type _MockHttpClientRecorder struct {
	mock *MockHttpClient
}

func NewMockHttpClient(ctrl *gomock.Controller) *MockHttpClient {
	mock := &MockHttpClient{ctrl: ctrl}
	mock.recorder = &_MockHttpClientRecorder{mock}
	return mock
}

func (_m *MockHttpClient) EXPECT() *_MockHttpClientRecorder {
	return _m.recorder
}

func (_m *MockHttpClient) Get(_param0 string) (*http.Response, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHttpClientRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}
