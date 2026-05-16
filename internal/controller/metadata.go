/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"maps"
	"strings"
)

func mergeManagedLabels(desired, current map[string]string) map[string]string {
	return mergeMetadataFunc(desired, current, isManagedLabel)
}

func mergeMetadata(desired, current map[string]string) map[string]string {
	return mergeMetadataFunc(desired, current, func(string) bool { return false })
}

func mergeMetadataFunc(desired, current map[string]string, isManaged func(string) bool) map[string]string {
	if len(desired) == 0 && len(current) == 0 {
		return nil
	}

	merged := make(map[string]string, len(desired)+len(current))
	maps.Copy(merged, current)
	for key := range current {
		if isManaged(key) {
			delete(merged, key)
		}
	}
	maps.Copy(merged, desired)
	if len(merged) == 0 {
		return nil
	}
	return merged
}

func isManagedLabel(key string) bool {
	return key == managedByLabel || strings.HasPrefix(key, "kudeploy.com/")
}
