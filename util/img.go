/*
公共图片处理工具

BetaX Blog
Copyright © 2024 SkyeZhang <skai-zhang@hotmail.com>
*/

package util

import (
	"bytes"
	"image"
	"image/jpeg"
	_ "image/png"
)

func Compression(data []byte) ([]byte, error) {
	// 解码图片
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// 创建一个新的缓冲区来存储 JPEG 图像数据
	var buf bytes.Buffer
	// 转换为 JPG 并进行压缩
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100}); err != nil {
		return nil, err
	}
	// 从缓冲区获取 JPEG 图像数据
	return buf.Bytes(), nil

}
