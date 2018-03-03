// Copyright 2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type SecProto int

const (
	NoEnc SecProto = iota
	WpaPsk
	WpaEap
	NotSupportedProto
)

type WifiOption struct {
	Essid     string
	AuthSuite SecProto
}

type UserInputMessage struct {
	args []string
}

type State struct {
	nearbyWifis []WifiOption
	curEssid    string
}

var (
	StatusRequestChannel = make(chan bool)
	UserInputChannel     = make(chan UserInputMessage)
	StateChannel         = make(chan State)
)