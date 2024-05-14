// Copyright 2024 Jelly Terra
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package util

import (
	"github.com/BurntSushi/toml"
	"os"
)

func UnmarshalTOMLFile(path string, v any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return toml.Unmarshal(b, v)
}
