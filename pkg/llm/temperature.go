// Copyright 2025 venslabs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package llm

import "strings"

// temperatureRefusingModels lists Anthropic model generations known to reject
// an explicit temperature parameter outright ("`temperature` is deprecated for
// this model."). Anthropic's capabilities object carries no temperature key,
// and adjacent generations (sonnet-4-6 vs sonnet-5) report identical
// capabilities while differing on temperature, so a best-effort list is the
// only way to fail fast. A model missing from the list still fails with a
// clear error after its first refused call instead of retrying.
var temperatureRefusingModels = []string{
	// Current generation; e.g. claude-sonnet-5. Extend as new generations
	// confirm they refuse an explicit temperature.
	"sonnet-5",
	"opus-4-7",
	"opus-4-8",
	"opus-5",
	"opus-5-5",
	"fable",
}

// ModelRefusesTemperature reports whether model is a known temperature
// refuser. It lives in the shared llm package so every provider applies the
// same check.
func ModelRefusesTemperature(model string) bool {
	m := strings.ToLower(model)
	for _, known := range temperatureRefusingModels {
		if strings.Contains(m, known) {
			return true
		}
	}
	return false
}
