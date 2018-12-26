// Copyright 2018 singularitynet foundation.
// All rights reserved.
// <<add licence terms for code reuse>>

// package for monitoring and reporting the daemon metrics
package metrics

import (
	"reflect"
	"testing"
)

func Test_callgRPCServiceHeartbeat(t *testing.T) {
	type args struct {
		grpcAddress string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := callgRPCServiceHeartbeat(tt.args.grpcAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("callgRPCServiceHeartbeat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("callgRPCServiceHeartbeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_callHTTPServiceHeartbeat(t *testing.T) {
	type args struct {
		serviceURL string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := callHTTPServiceHeartbeat(tt.args.serviceURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("callHTTPServiceHeartbeat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("callHTTPServiceHeartbeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_callRegisterService(t *testing.T) {
	type args struct {
		daemonID   string
		serviceURL string
	}
	tests := []struct {
		name       string
		args       args
		wantStatus bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStatus := callRegisterService(tt.args.daemonID, tt.args.serviceURL); gotStatus != tt.wantStatus {
				t.Errorf("callRegisterService() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
