package entropy

import (
	"image"
	"image/draw"
	"io/ioutil"
	"math/rand"
	"time"

	"code.google.com/p/freetype-go/freetype"
	"code.google.com/p/freetype-go/freetype/truetype"
)

const (
	WORDPOOL = "ABCDEFGHJKMNPQRSTUVWXYZ23456789"
)

type Captcha struct {
	Image *image.RGBA
	Text  string
}

/*
生成验证码
bg:背景色
fg:前景色
length:字符长度
width:宽度
height:高度
size:字体大小
fontPath:字体文件路径
*/
func GenerateCaptcha(bg, fg *image.Uniform, length int, width int, height int, size float64, fontPath string) *Captcha {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	font, err := truetype.Parse(fontBytes)

	if err != nil {
		panic(err)
	}
	cap := &Captcha{}
	cap.Text = randString(length)
	cap.Image = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(cap.Image, cap.Image.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetFont(font)
	c.SetFontSize(size)
	c.SetClip(cap.Image.Bounds())
	c.SetDst(cap.Image)
	c.SetSrc(fg)
	pt := freetype.Pt(0, int(c.PointToFix32(size)>>8))
	for _, s := range cap.Text {
		_, err = c.DrawString(string(s), pt)
		if err != nil {
			panic(err)
			return nil
		}
		pt.X += c.PointToFix32(size * 0.5)
	}
	return cap
}

func randString(length int) string {
	ret := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		ret += string(WORDPOOL[rand.Intn(len(WORDPOOL))])
	}
	return ret
}
