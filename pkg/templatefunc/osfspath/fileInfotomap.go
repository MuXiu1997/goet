/*
The MIT License (MIT)

Copyright (c) 2018-2022 Tom Payne

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Code sourced from: Tom Payne, https://github.com/twpayne/chezmoi/blob/4cbdbd1820930a33914805d980979cf74ce0a75b/pkg/cmd/templatefuncs.go#L404
// Copyright belongs to Tom Payne
// Modified by MuXiu1997

package osfspath

import "io/fs"

func fileInfoToMap(fileInfo fs.FileInfo) map[string]any {
	return map[string]any{
		"name":    fileInfo.Name(),
		"size":    fileInfo.Size(),
		"mode":    int(fileInfo.Mode()),
		"perm":    int(fileInfo.Mode().Perm()),
		"modTime": fileInfo.ModTime().Unix(),
		"isDir":   fileInfo.IsDir(),
		"type":    fileModeTypeNames[fileInfo.Mode()&fs.ModeType],
	}
}

var fileModeTypeNames = map[fs.FileMode]string{
	0:                 "file",
	fs.ModeDir:        "dir",
	fs.ModeSymlink:    "symlink",
	fs.ModeNamedPipe:  "named pipe",
	fs.ModeSocket:     "socket",
	fs.ModeDevice:     "device",
	fs.ModeCharDevice: "char device",
}
