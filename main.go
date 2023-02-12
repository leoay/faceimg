package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/nfnt/resize"
)

func filelist(originpath string) []string {
	root := originpath
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

// 从本地文件夹，生成封面
func gengzhface(inputdir string, outpath string) {
	img := image.NewNRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{2350, 1000},
	})

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, color.White)
		}
	}

	imglist := filelist(inputdir)[1:]
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(imglist), func(i, j int) {
		imglist[i], imglist[j] = imglist[j], imglist[i]
	})
	imglist = imglist[0:5]
	i := 0
	for _, file := range imglist {
		img1, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer img1.Close()
		jpgimg1, _ := jpeg.Decode(img1)
		img1_1 := resize.Resize(0, uint(img.Bounds().Dy()), jpgimg1, resize.Lanczos3)
		fromW := i * (10 + img1_1.Bounds().Dx())
		fromH := img1_1.Bounds().Dy()
		draw.Draw(img,
			image.Rectangle{
				Min: image.Point{X: fromW, Y: 0},
				Max: image.Point{X: fromW + img1_1.Bounds().Dx(), Y: fromH}},
			img1_1,
			img1_1.Bounds().Min,
			draw.Over,
		)
		i++
	}

	// imgxl := image.NewNRGBA(image.Rect(0, 0, 500, 260))
	// for x := 0; x < imgxl.Bounds().Dx(); x++ {
	// 	for y := 0; y < imgxl.Bounds().Dy(); y++ {
	// 		imgxl.SetNRGBA(x, y, color.NRGBA{
	// 			R: 0, G: 0, B: 0, A: 100})
	// 	}
	// }

	// draw.Draw(img,
	// 	image.Rectangle{
	// 		Min: image.Point{X: img.Bounds().Dx() - 540, Y: img.Bounds().Dy() - 310},
	// 		Max: image.Point{X: img.Bounds().Dx(), Y: img.Bounds().Dy()}},
	// 	imgxl,
	// 	imgxl.Bounds().Min,
	// 	draw.Over,
	// )

	//写字
	// fontBytes, err := ioutil.ReadFile("fonts/PangMenZhengDaoBiaoTiTi-1.ttf")
	// if err != nil {
	// 	log.Println(err)
	// }

	// font, err := freetype.ParseFont(fontBytes)
	// if err != nil {
	// 	log.Println(err)
	// }

	// f := freetype.NewContext()
	// f.SetDPI(72)
	// f.SetFont(font)
	// f.SetFontSize(115)
	// f.SetClip(img.Bounds())
	// f.SetDst(img)
	// f.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 255, B: 255, A: 255}))

	//pt := freetype.Pt(img.Bounds().Dx()-500, img.Bounds().Dy()-160)

	// _, err = f.DrawString("臻选壁纸", pt)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//写字2
	// fontBytes, err = ioutil.ReadFile("fonts/PangMenZhengDaoBiaoTiTi-1.ttf")
	// if err != nil {
	// 	log.Println(err)
	// }

	// font, err = freetype.ParseFont(fontBytes)
	// if err != nil {
	// 	log.Println(err)
	// }

	// f = freetype.NewContext()
	// f.SetDPI(72)
	// f.SetFont(font)
	// f.SetFontSize(58)
	// f.SetClip(img.Bounds())
	// f.SetDst(img)
	// f.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 255, B: 255, A: 255}))

	//pt = freetype.Pt(img.Bounds().Dx()-500, img.Bounds().Dy()-90)

	// _, err = f.DrawString("zhenxuanpaper.com", pt)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	newfile, _ := os.Create(outpath)
	defer newfile.Close()
	jpeg.Encode(newfile, img, &jpeg.Options{100})
}

func main() {
	gengzhface("./imglist", "./test222.jpg")
}
