package main

import "testing"
import "strings"
import "errors"
import "github.com/docker/go-plugins-helpers/volume"
import "github.com/Azure/azure-sdk-for-go/storage"
import "github.com/bouk/monkey"

func TestCreateWithoutShareCallFunc(t *testing.T) {
	v := volumeDriver{}
	req := volume.Request{}
	createResult := v.Create(req)
	if !(strings.Contains(createResult.Err, "missing") && strings.Contains(createResult.Err, "share")) {
		t.Error("'share' is a required argument and if it's not passed, Create func should return an error!")
	}
}

func TestCreateWithShareCallFunc(t *testing.T) {
	v := volumeDriver{}
	req := volume.Request{
		Options: map[string]string{
			"share": "testshare",
		},
	}

	monkey.Patch(storage.FileServiceClient.CreateShareIfNotExists, func(fsc storage.FileServiceClient, name string) (bool, error) {
		return false, errors.New("Something went wrong with share " + name)
	})

	createResult := v.Create(req)
	if !(strings.Contains(createResult.Err, "wrong") && strings.Contains(createResult.Err, "testshare")) {
		t.Error("Create should return possible error from CreateShareIfNotExists method!")
	}
}

//TODO: cover the behaviour with 'remotepath' parameter in Create function
//      TestCreateWithRemotepathParamCallFunc, patch CreateDirectory
