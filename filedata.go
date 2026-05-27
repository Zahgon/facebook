// A facebook graph api client in go.
// https://github.com/huandu/facebook/
//
// Copyright 2012, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/facebook/blob/master/LICENSE

package facebook

import (
	"io"
)

// BinaryData represents binary data from a given source.
type BinaryData struct {
	Filename    string    // filename used in multipart form writer.
	Source      io.Reader // file data source.
	ContentType string    // content type of the data.
}

// BinaryFile represents a file on disk.
type BinaryFile struct {
	Filename    string // filename used in multipart form writer.
	Path        string // path to file. must be readable.
	ContentType string // content type of the file.
}

// Data creates new binary data holder.
func Data(filename string, source io.Reader) *BinaryData { _ = "STUB: not implemented"; return nil }

// DataWithContentType creates new binary data holder with arbitrary content type.
func DataWithContentType(filename string, source io.Reader, contentType string) *BinaryData {
	_ = "STUB: not implemented"
	return nil
}

// File creates a binary file holder.
func File(filename string) *BinaryFile { _ = "STUB: not implemented"; return nil }

// FileAlias creates a binary file holder and specific a different path for reading.
func FileAlias(filename, path string) *BinaryFile { _ = "STUB: not implemented"; return nil }

// FileAliasWithContentType creates a new binary file holder with arbitrary content type.
func FileAliasWithContentType(filename, path, contentType string) *BinaryFile {
	_ = "STUB: not implemented"
	return nil
}
