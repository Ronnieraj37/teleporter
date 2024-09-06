// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// The unpackgen binary generates Solidity code to invert `abi.encodePacked()`
// for known, unambiguous output sizes. See the [unpack] package for details.
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/pflag"

	"github.com/ava-labs/teleporter/utils/unpack"
)

const dynFlagSentinel = "dyn"

func main() {
	var cfg config

	pflag.StringVar(&cfg.implPath, "out", "./Unpack.sol", "Output path to which the Solidity implementations are written")
	pflag.StringVar(&cfg.testPath, "test_out", "./Unpack.t.sol", "Output path to which the Solidity tests are written")
	pflag.StringVar(&cfg.solcVersion, "solc_version", "", "If non-empty, used verbatim as the version of the `pragma solidity` directive")

	bytesMsg := fmt.Sprintf("Semicolon-delimited list of groups of comma-delimited byte sizes. Use literal %q for (at most one) dynamically sized byte array per group.", dynFlagSentinel) //nolint:lll
	pflag.Var(&cfg.bytes, "byte_sizes", bytesMsg)

	pflag.Parse()

	if err := cfg.run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type config struct {
	implPath, testPath, solcVersion string
	bytes                           bytesFlag
}

// run creates files at both paths and then calls `unpack.Generate()`.
func (cfg *config) run() error {
	impl, err := os.Create(cfg.implPath)
	if err != nil {
		return err
	}
	test, err := os.Create(cfg.testPath)
	if err != nil {
		return err
	}

	rel, err := filepath.Rel(filepath.Dir(cfg.testPath), cfg.implPath)
	if err != nil {
		return fmt.Errorf("determine path from test to implementation: %v", err)
	}

	if err := unpack.Generate(impl, test, cfg.solcVersion, rel, cfg.bytes.val...); err != nil {
		return err
	}
	return errors.Join(impl.Close(), test.Close())
}

type bytesFlag struct {
	val [][]uint
}

var _ pflag.Value = (*bytesFlag)(nil)

// Set splits `s` by semicolons then splits each result by commas, parsing all
// resulting values as uints. It trims whitespace. If a value is "dyn" it
// is replaced with unpack.Dynamic.
func (f *bytesFlag) Set(s string) error {
	groups := strings.Split(s, ";")
	var val [][]uint

	for _, g := range groups {
		var bytes []uint
		for _, bStr := range strings.Split(g, ",") {
			if strings.ToLower(bStr) == dynFlagSentinel {
				bytes = append(bytes, unpack.Dynamic)
				continue
			}
			b, err := strconv.ParseUint(strings.TrimSpace(bStr), 10, 32)
			if err != nil {
				return err
			}
			bytes = append(bytes, uint(b))
		}
		val = append(val, bytes)
	}

	f.val = val
	return nil
}

// String is the inverse of `bytesFlag.Set()`.
func (f *bytesFlag) String() string {
	var parts []string
	for _, grp := range f.val {
		var vs []string
		for _, v := range grp {
			var vStr string
			if v == unpack.Dynamic {
				vStr = dynFlagSentinel
			} else {
				vStr = fmt.Sprintf("%d", v)
			}
			vs = append(vs, vStr)
		}
		parts = append(parts, strings.Join(vs, ","))
	}
	return strings.Join(parts, ";")
}

func (f *bytesFlag) Type() string {
	return fmt.Sprintf("%T", *f)
}