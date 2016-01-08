package config

import (
	"testing"

	v2testing "github.com/v2ray/v2ray-core/testing"
	"github.com/v2ray/v2ray-core/testing/assert"
)

func TestRegisterInboundConfig(t *testing.T) {
	v2testing.Current(t)
	initializeConfigCache()

	protocol := "test_protocol"
	creator := func([]byte) (interface{}, error) {
		return true, nil
	}

	err := RegisterInboundConnectionConfig(protocol, creator)
	assert.Error(err).IsNil()

	configObj, err := CreateInboundConnectionConfig(protocol, nil)
	assert.Bool(configObj.(bool)).IsTrue()
	assert.Error(err).IsNil()

	configObj, err = CreateOutboundConnectionConfig(protocol, nil)
	assert.Pointer(configObj).IsNil()
}

func TestRegisterOutboundConfig(t *testing.T) {
	v2testing.Current(t)
	initializeConfigCache()

	protocol := "test_protocol"
	creator := func([]byte) (interface{}, error) {
		return true, nil
	}

	err := RegisterOutboundConnectionConfig(protocol, creator)
	assert.Error(err).IsNil()

	configObj, err := CreateOutboundConnectionConfig(protocol, nil)
	assert.Bool(configObj.(bool)).IsTrue()
	assert.Error(err).IsNil()

	configObj, err = CreateInboundConnectionConfig(protocol, nil)
	assert.Pointer(configObj).IsNil()
}