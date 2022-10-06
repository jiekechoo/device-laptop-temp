// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a simple example of a device service.
package main

import (
	"github.com/edgexfoundry/device-laptop-temp"
	"github.com/edgexfoundry/device-laptop-temp/driver"
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/startup"
)

const (
	serviceName string = "device-laptop-temp"
)

func main() {
	sd := driver.SimpleDriver{}
	startup.Bootstrap(serviceName, device.Version, &sd)
}
