// Tideland Go Library - Cell Behaviors - Unit Tests - Configurator
//
// Copyright (C) 2015 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package behaviors_test

//--------------------
// IMPORTS
//--------------------

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/tideland/golib/audit"
	"github.com/tideland/golib/cells"
	"github.com/tideland/golib/cells/behaviors"
	"github.com/tideland/golib/configuration"
)

//--------------------
// TESTS
//--------------------

// TestConfigurationLoad tests the successful loading of a configuration.
func TestConfigurationLoad(t *testing.T) {
	assert := audit.NewTestingAssertion(t, true)
	env := cells.NewEnvironment("configuration-load")
	defer env.Stop()
	tempDir, filename := createConfigurationFile(assert, "{config {foo 42}}")
	defer tempDir.Restore()

	sigc := make(chan bool, 1)
	spf := func(ctx cells.Context, event cells.Event) error {
		if event.Topic() == behaviors.ConfigurationTopic {
			value, ok := event.Payload().Get(behaviors.ConfigurationPayload)
			assert.True(ok)
			config, ok := value.(configuration.Configuration)
			assert.True(ok)
			v, err := config.At("foo").Value()
			assert.Nil(err)
			assert.Equal(v, "42")

			sigc <- true
		}
		return nil
	}

	env.StartCell("configurator", behaviors.NewConfiguratorBehavior(nil))
	env.StartCell("simple", behaviors.NewSimpleProcessorBehavior(spf))
	env.Subscribe("configurator", "simple")

	pvs := cells.PayloadValues{
		behaviors.ConfigurationFilenamePayload: filename,
	}
	env.EmitNew("configurator", behaviors.ReadConfigurationTopic, pvs, nil)
	assert.Wait(sigc, 100*time.Millisecond)
}

//--------------------
// HELPER
//--------------------

// createConfigurationFile creates a temporary configuration file.
func createConfigurationFile(assert audit.Assertion, content string) (*audit.TempDir, string) {
	tempDir := audit.NewTempDir(assert)
	configFile, err := ioutil.TempFile(tempDir.String(), "configuration")
	assert.Nil(err)
	configFileName := configFile.Name()
	_, err = configFile.WriteString(content)
	assert.Nil(err)
	configFile.Close()

	return tempDir, configFileName
}

// EOF
