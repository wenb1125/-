package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
)

func MD5HashString(data string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	psswordBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(psswordBytes)
}

/**
 * 对一个io操作的reader（通常为文件)进行数据读取，并计算hash，返回md5哈希值
 */
func MD5HashReader(reader io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	md5Hash := md5.New()
	md5Hash.Write(bytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}
