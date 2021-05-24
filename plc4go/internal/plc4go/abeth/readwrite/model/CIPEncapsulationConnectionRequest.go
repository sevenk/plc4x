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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type CIPEncapsulationConnectionRequest struct {
	Parent *CIPEncapsulationPacket
}

// The corresponding interface
type ICIPEncapsulationConnectionRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *CIPEncapsulationConnectionRequest) CommandType() uint16 {
	return 0x0101
}

func (m *CIPEncapsulationConnectionRequest) InitializeParent(parent *CIPEncapsulationPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.Parent.SessionHandle = sessionHandle
	m.Parent.Status = status
	m.Parent.SenderContext = senderContext
	m.Parent.Options = options
}

func NewCIPEncapsulationConnectionRequest(sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *CIPEncapsulationPacket {
	child := &CIPEncapsulationConnectionRequest{
		Parent: NewCIPEncapsulationPacket(sessionHandle, status, senderContext, options),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastCIPEncapsulationConnectionRequest(structType interface{}) *CIPEncapsulationConnectionRequest {
	castFunc := func(typ interface{}) *CIPEncapsulationConnectionRequest {
		if casted, ok := typ.(CIPEncapsulationConnectionRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*CIPEncapsulationConnectionRequest); ok {
			return casted
		}
		if casted, ok := typ.(CIPEncapsulationPacket); ok {
			return CastCIPEncapsulationConnectionRequest(casted.Child)
		}
		if casted, ok := typ.(*CIPEncapsulationPacket); ok {
			return CastCIPEncapsulationConnectionRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CIPEncapsulationConnectionRequest) GetTypeName() string {
	return "CIPEncapsulationConnectionRequest"
}

func (m *CIPEncapsulationConnectionRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *CIPEncapsulationConnectionRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *CIPEncapsulationConnectionRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func CIPEncapsulationConnectionRequestParse(readBuffer utils.ReadBuffer) (*CIPEncapsulationPacket, error) {
	if pullErr := readBuffer.PullContext("CIPEncapsulationConnectionRequest"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("CIPEncapsulationConnectionRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CIPEncapsulationConnectionRequest{
		Parent: &CIPEncapsulationPacket{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *CIPEncapsulationConnectionRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CIPEncapsulationConnectionRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("CIPEncapsulationConnectionRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *CIPEncapsulationConnectionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
