package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

func validateFunctionName(functionName string) error {
	if strings.Count(functionName, "*") != 1 {
		return fmt.Errorf("functionName should have only one placeholder (*)")
	}

	return nil
}

func validateTarget(target string) error {
	if !strings.HasPrefix(target, "0x") {
		return fmt.Errorf("target should start with 0x: %s", target)
	}

	if _, err := strconv.ParseUint(strings.ReplaceAll(target[2:], "*", ""), 16, 64); err != nil {
		return fmt.Errorf("invalid target: %s", target)
	}

	if len(target) != 10 {
		return fmt.Errorf("invalid target length, must be 10 symbols")
	}

	return nil
}
