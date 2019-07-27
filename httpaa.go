// Copyright 2019 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package httpaa is like net/http but
// provides a http.Handle function with
// all others a's.
//
// This package is inspired by
// https://twitter.com/chimeracoder/status/1154156377301929984.
package httpaa

import "net/http"

// HaandleFunc is like http.HandleFunc but with aa.
func HaandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, handler)
}

// HåndleFunc is like http.HandleFunc but with å.
func HåndleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, handler)
}

// HändleFunc is like http.HandleFunc but with ä.
func HändleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, handler)
}

// HândleFunc is like http.HandleFunc but with â.
func HândleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, handler)
}

// HàndleFunc is like http.HandleFunc but with à.
func HàndleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, handler)
}

// HándleFunc is like http.HandleFunc but with á.
func HándleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, handler)
}

// HædnleFunc is like http.HandleFunc but with æ.
func HændleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, handler)
}

// HāndleFunc is like http.HandleFunc but with ā.
func HāndleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, handler)
}

