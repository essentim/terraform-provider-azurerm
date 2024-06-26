// Package iotcentral implements the Azure ARM Iotcentral service API version 2022-10-31-preview.
//
// Azure IoT Central is a service that makes it easy to connect, monitor, and manage your IoT devices at scale.
package iotcentral

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseDomain is the default value for base domain
	DefaultBaseDomain = "azureiotcentral.com"
)

// BaseClient is the base client for Iotcentral.
type BaseClient struct {
	autorest.Client
	BaseDomain string
	Subdomain  string
}

// New creates an instance of the BaseClient client.
func New(subdomain string) BaseClient {
	return NewWithoutDefaults(subdomain, DefaultBaseDomain)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(subdomain string, baseDomain string) BaseClient {
	return BaseClient{
		Client:     autorest.NewClientWithUserAgent(UserAgent()),
		BaseDomain: baseDomain,
		Subdomain:  subdomain,
	}
}
