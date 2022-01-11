/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetNotificationParametersChangeOfState struct {
	*BACnetNotificationParameters
	InnerOpeningTag *BACnetOpeningTag
	ChangeOfState   *BACnetPropertyStates
	StatusFlags     *BACnetStatusFlags
	InnerClosingTag *BACnetClosingTag
}

// The corresponding interface
type IBACnetNotificationParametersChangeOfState interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetNotificationParametersChangeOfState) PeekedTagNumber() uint8 {
	return uint8(1)
}

func (m *BACnetNotificationParametersChangeOfState) InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, peekedTagNumber uint8) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func NewBACnetNotificationParametersChangeOfState(innerOpeningTag *BACnetOpeningTag, changeOfState *BACnetPropertyStates, statusFlags *BACnetStatusFlags, innerClosingTag *BACnetClosingTag, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, peekedTagNumber uint8) *BACnetNotificationParameters {
	child := &BACnetNotificationParametersChangeOfState{
		InnerOpeningTag:              innerOpeningTag,
		ChangeOfState:                changeOfState,
		StatusFlags:                  statusFlags,
		InnerClosingTag:              innerClosingTag,
		BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, peekedTagNumber),
	}
	child.Child = child
	return child.BACnetNotificationParameters
}

func CastBACnetNotificationParametersChangeOfState(structType interface{}) *BACnetNotificationParametersChangeOfState {
	castFunc := func(typ interface{}) *BACnetNotificationParametersChangeOfState {
		if casted, ok := typ.(BACnetNotificationParametersChangeOfState); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetNotificationParametersChangeOfState); ok {
			return casted
		}
		if casted, ok := typ.(BACnetNotificationParameters); ok {
			return CastBACnetNotificationParametersChangeOfState(casted.Child)
		}
		if casted, ok := typ.(*BACnetNotificationParameters); ok {
			return CastBACnetNotificationParametersChangeOfState(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetNotificationParametersChangeOfState) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfState"
}

func (m *BACnetNotificationParametersChangeOfState) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfState) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.LengthInBits()

	// Simple field (changeOfState)
	lengthInBits += m.ChangeOfState.LengthInBits()

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.LengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.LengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfState) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetNotificationParametersChangeOfStateParse(readBuffer utils.ReadBuffer, tagNumber uint8, peekedTagNumber uint8) (*BACnetNotificationParameters, error) {
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfState"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetContextTagParse(readBuffer, peekedTagNumber, BACnetDataType_OPENING_TAG)
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := CastBACnetOpeningTag(_innerOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (changeOfState)
	if pullErr := readBuffer.PullContext("changeOfState"); pullErr != nil {
		return nil, pullErr
	}
	_changeOfState, _changeOfStateErr := BACnetPropertyStatesParse(readBuffer, uint8(0))
	if _changeOfStateErr != nil {
		return nil, errors.Wrap(_changeOfStateErr, "Error parsing 'changeOfState' field")
	}
	changeOfState := CastBACnetPropertyStates(_changeOfState)
	if closeErr := readBuffer.CloseContext("changeOfState"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, pullErr
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsParse(readBuffer, uint8(1))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field")
	}
	statusFlags := CastBACnetStatusFlags(_statusFlags)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerClosingTag, _innerClosingTagErr := BACnetContextTagParse(readBuffer, peekedTagNumber, BACnetDataType_CLOSING_TAG)
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := CastBACnetClosingTag(_innerClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfState"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfState{
		InnerOpeningTag:              CastBACnetOpeningTag(innerOpeningTag),
		ChangeOfState:                CastBACnetPropertyStates(changeOfState),
		StatusFlags:                  CastBACnetStatusFlags(statusFlags),
		InnerClosingTag:              CastBACnetClosingTag(innerClosingTag),
		BACnetNotificationParameters: &BACnetNotificationParameters{},
	}
	_child.BACnetNotificationParameters.Child = _child
	return _child.BACnetNotificationParameters, nil
}

func (m *BACnetNotificationParametersChangeOfState) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfState"); pushErr != nil {
			return pushErr
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return pushErr
		}
		_innerOpeningTagErr := m.InnerOpeningTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return popErr
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (changeOfState)
		if pushErr := writeBuffer.PushContext("changeOfState"); pushErr != nil {
			return pushErr
		}
		_changeOfStateErr := m.ChangeOfState.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("changeOfState"); popErr != nil {
			return popErr
		}
		if _changeOfStateErr != nil {
			return errors.Wrap(_changeOfStateErr, "Error serializing 'changeOfState' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return pushErr
		}
		_statusFlagsErr := m.StatusFlags.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return popErr
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return pushErr
		}
		_innerClosingTagErr := m.InnerClosingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return popErr
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfState"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfState) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
