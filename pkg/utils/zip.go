package utils

import (
	"compress/zlib"

	"bytes"
	"io/ioutil"
)

// ZIP 流压缩
func Zip(data []byte) *bytes.Buffer {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	defer w.Close()

	w.Write([]byte(data))

	return &buf
}

// zip 解压
func Unzip(b *bytes.Buffer) ([]byte, error) {
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}
