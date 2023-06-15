package utils_test

import (
	"errors"
	"testing"

	"github.com/Timwood0x10/sei-chain/utils"
	"github.com/stretchr/testify/require"
)

func TestHardFail(t *testing.T) {
	hardFailer := func() {
		panic(utils.DecorateHardFailError(errors.New("some error")))
	}
	panicHandlingFn := func() {
		defer utils.PanicHandler(func(_ any) {})()
		hardFailer()
	}
	require.Panics(t, panicHandlingFn)
}
