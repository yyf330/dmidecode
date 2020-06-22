package smbios_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yyf330/dmidecode/smbios"
)

func TestRead(t *testing.T) {
	ss, err := smbios.ReadStructures()
	if assert.NoError(t, err) {
		t.Log(ss)
	}
}
