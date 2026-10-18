// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package merger

import (
	"github.com/imdario/mergo"
)

func MergeStructs[T comparable](withOverride bool, structs ...T) (T, error) {
	dst := new(T)

	if !withOverride {
		for _, s := range structs {
			if err := mergo.Merge(dst, s); err != nil {
				return *dst, err
			}
		}

		return *dst, nil
	}

	for _, s := range structs {
		if err := mergo.Merge(dst, s, mergo.WithOverride); err != nil {
			return *dst, err
		}
	}

	return *dst, nil
}
