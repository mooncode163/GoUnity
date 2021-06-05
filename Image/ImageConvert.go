package main

import (
	"fmt"
	"os"
	"strings"

	// "strings"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"path/filepath"

	"github.com/nfnt/resize"
)

func main() {
	fmt.Println("ConverImage")

	filesrc := "1.jpg"
	filedst := "dst.png"
	if len(os.Args) > 1 {
		// who = strings.Join(os.Args[1:], " ")
		filesrc = os.Args[1]
		filedst = os.Args[2]
	}
	fmt.Println("image=", filesrc)
	ConverImage(filesrc, filedst, 512, 1024)
	// ConverImage(filesrc, "dst.png", 512, 1024)
}

// imagetype 0 jpg  1 png
func ConverImage(filesrc string, filedst string, width uint, height uint) {
	file1, _ := os.Open(filesrc) //打开图片1
	defer file1.Close()

	// image.Decode 图片
	var (
		img1 image.Image
		err  error
	)
	if img1, _, err = image.Decode(file1); err != nil {
		log.Fatal(err)
		return
	}
	// 调用resize库进行图片缩放(高度填0，resize.Resize函数中会自动计算缩放图片的宽高比)
	m1 := resize.Resize(width, height, img1, resize.Lanczos3)

	// 将两个图片合成一张
	newWidth := m1.Bounds().Max.X                                      //新宽度 = 随意一张图片的宽度
	newHeight := m1.Bounds().Max.Y                                     // 新图片的高度为两张图片高度的和
	newImg := image.NewNRGBA(image.Rect(0, 0, newWidth, newHeight))    //创建一个新RGBA图像
	draw.Draw(newImg, newImg.Bounds(), m1, m1.Bounds().Min, draw.Over) //画上第一张缩放后的图片

	// 保存文件
	ext := filepath.Ext(filesrc)
	imgfile, _ := os.Create(filedst)
	defer imgfile.Close()
	// {
	// 	// jpeg.Encode(imgfile, newImg, &jpeg.Options{100})
	// }
	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {
		err = jpeg.Encode(imgfile, newImg, &jpeg.Options{Quality: 100})
	} else if strings.EqualFold(ext, ".png") {
		err = png.Encode(imgfile, newImg)
	} else if strings.EqualFold(ext, ".gif") {
		err = gif.Encode(imgfile, newImg, &gif.Options{NumColors: 256})
	}

}
