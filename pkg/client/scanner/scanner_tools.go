// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package scanner

// Composer represents a container which produces a Composer result artifact.
var Composer = Tool{
	Name:      "Composer",
	Version:   "dummy",
	DockerImg: "docker.pkg.github.com/osrgroup/product-model-toolkit/scanner-composer:dummy",
	Cmd:       `/bin/sh -c "cp example.json result/example.json"`,
	Results:   []string{"example.json"},
}
