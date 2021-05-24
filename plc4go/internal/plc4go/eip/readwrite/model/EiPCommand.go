//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type EiPCommand uint16

type IEiPCommand interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	EiPCommand_RegisterSession   EiPCommand = 0x0065
	EiPCommand_UnregisterSession EiPCommand = 0x0066
	EiPCommand_SendRRData        EiPCommand = 0x006F
)

var EiPCommandValues []EiPCommand

func init() {
	_ = errors.New
	EiPCommandValues = []EiPCommand{
		EiPCommand_RegisterSession,
		EiPCommand_UnregisterSession,
		EiPCommand_SendRRData,
	}
}

func EiPCommandByValue(value uint16) EiPCommand {
	switch value {
	case 0x0065:
		return EiPCommand_RegisterSession
	case 0x0066:
		return EiPCommand_UnregisterSession
	case 0x006F:
		return EiPCommand_SendRRData
	}
	return 0
}

func EiPCommandByName(value string) EiPCommand {
	switch value {
	case "RegisterSession":
		return EiPCommand_RegisterSession
	case "UnregisterSession":
		return EiPCommand_UnregisterSession
	case "SendRRData":
		return EiPCommand_SendRRData
	}
	return 0
}

func CastEiPCommand(structType interface{}) EiPCommand {
	castFunc := func(typ interface{}) EiPCommand {
		if sEiPCommand, ok := typ.(EiPCommand); ok {
			return sEiPCommand
		}
		return 0
	}
	return castFunc(structType)
}

func (m EiPCommand) LengthInBits() uint16 {
	return 16
}

func (m EiPCommand) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func EiPCommandParse(readBuffer utils.ReadBuffer) (EiPCommand, error) {
	val, err := readBuffer.ReadUint16("EiPCommand", 16)
	if err != nil {
		return 0, nil
	}
	return EiPCommandByValue(val), nil
}

func (e EiPCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("EiPCommand", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e EiPCommand) name() string {
	switch e {
	case EiPCommand_RegisterSession:
		return "RegisterSession"
	case EiPCommand_UnregisterSession:
		return "UnregisterSession"
	case EiPCommand_SendRRData:
		return "SendRRData"
	}
	return ""
}

func (e EiPCommand) String() string {
	return e.name()
}
