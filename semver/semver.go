/*
 *       Copyright 2019, 2020 Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package semver is Semantic Versioning as defined on <http://semver.org>.
package semver

// Version is the interface to access version numbers.
type Version interface {
	Major() uint
	Minor() uint
	Patch() uint
	Label() string
	Metadata() string
}

type tVersion struct {
	major    uint
	minor    uint
	patch    uint
	label    string
	metadata string
	str      string
}

// New returns an object implementing the Version interface.
func New(major, minor, patch uint, label, metadata string) Version {
	version := new(tVersion)
	version.major = major
	version.minor = minor
	version.patch = patch
	version.label = label
	version.metadata = metadata
	version.str = String(major, minor, patch, label, metadata)
	return version
}

// Major returns the major number of version.
func (version *tVersion) Major() uint {
	return version.major
}

// Minor returns the minor number of version.
func (version *tVersion) Minor() uint {
	return version.minor
}

// Patch returns the patch number of version.
func (version *tVersion) Patch() uint {
	return version.patch
}

// String returns version as string formatted "<major>.<minor>.<patch>".
func (version *tVersion) String() string {
	return version.str
}

// Label returns the additional label of pre-release version.
func (version *tVersion) Label() string {
	return version.label
}

// Metadata returns the metadata of version.
func (version *tVersion) Metadata() string {
	return version.metadata
}

// String returns major, minor and patch number with additional label
// and metadata formatted "<major>.<minor>.<patch>-<label>+<metadata>".
func String(major, minor, patch uint, label, metadata string) string {
	var versionStr string
	if major < 10 && minor < 10 {
		if patch < 10 {
			versionArr := make([]byte, 5)
			versionArr[0] = byte(major + 48)
			versionArr[1] = 46
			versionArr[2] = byte(minor + 48)
			versionArr[3] = 46
			versionArr[4] = byte(patch + 48)
			versionStr = string(versionArr)

		} else if patch < 100 {
			versionArr := make([]byte, 6)
			patchTenth := patch / 10
			versionArr[0] = byte(major + 48)
			versionArr[1] = 46
			versionArr[2] = byte(minor + 48)
			versionArr[3] = 46
			versionArr[4] = byte(patchTenth + 48)
			versionArr[5] = byte(patch - patchTenth*10 + 48)
			versionStr = string(versionArr)

		} else {
			majorByte := byte(major + 48)
			minorByte := byte(minor + 48)
			versionStr = string([]byte{majorByte, 46, minorByte, 46}) + uintToString(patch)
		}
	} else {
		versionStr = uintToString(major) + "." + uintToString(minor) + "." + uintToString(patch)
	}
	if len(label) > 0 {
		versionStr = versionStr + "-" + label
	}
	if len(metadata) > 0 {
		versionStr = versionStr + "+" + metadata
	}
	return versionStr
}

func uintToString(value uint) string {
	var byteArr [20]byte
	var decimalPlaces int
	tenth := value / 10
	byteArr[19] = byte(value - tenth*10 + 48)
	value = tenth

	for decimalPlaces = 1; value > 0 && decimalPlaces < 20; decimalPlaces += 1 {
		tenth := value / 10
		byteArr[19-decimalPlaces] = byte(value - tenth*10 + 48)
		value = tenth
	}
	return string(byteArr[20-decimalPlaces:])
}
