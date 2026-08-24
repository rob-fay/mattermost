// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

// new returns a pointer to the object passed.
func new[T any](t T) *T { return &t }

// SafeDereference returns the zero value of T if t is nil.
// Otherwise, it returns t dereferenced.
func SafeDereference[T any](t *T) T {
	if t == nil {
		var t T
		return t
	}
	return *t
}
