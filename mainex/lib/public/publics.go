package public

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"
)

type PublicsModel struct{}

var Publics = PublicsModel{}

func (this PublicsModel) IsEmail(email string) bool {
	regEmail := regexp.MustCompile("^\\w+@\\w+\\.\\w{2,4}$")
	return regEmail.MatchString(email)
}

//去除空字符串
func (this PublicsModel) Trim(str string) string {
	// 去除空格
	str = strings.Replace(str, " ", "", -1)

	return str
}

//加密md5
func (this PublicsModel) Md5(mds string) string {

	md5h := md5.New()
	md5h.Write([]byte(mds))
	cipherStr := md5h.Sum(nil)
	//加密输出md5值
	return hex.EncodeToString(cipherStr)
}

// if your img's are properly formed with doublequotes then use this, it's more efficient.
// var imgRE = regexp.MustCompile(`<img[^>]+\bsrc="([^"]+)"`)
func (this PublicsModel) FindImages(htm string) []string {
	var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	imgs := imgRE.FindAllStringSubmatch(htm, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
	}
	return out
}

//匹配html或markdown中的图片,http开头,jpg/png在中间,后面跟一段参数,注意去掉^%首尾匹配，不然文章匹配不到
//在线验证正则表达式网站https://regex101.com/r/eY1yW5/2
func (this PublicsModel) CoverGirl(htm string, count int) []string {
	//var imgRE = regexp.MustCompile(`http.*(png|gif|jpg)[a-zA-Z0-9?\%-/]*`) //中间会包含空格
	//var imgRE = regexp.MustCompile(`http.[^ ]*(png|gif|jpg)[a-zA-Z0-9?\%-/]*`) //中间不包含空格
	//var imgRE = regexp.MustCompile(`http.[^ ]*(png|gif|jpg|jpeg)[a-zA-Z0-9?\%\-/]*`) //转义-，否则会匹配到')'
	//var imgRE = regexp.MustCompile(`http.[^ ].*(png|gif|jpg|jpeg)[a-zA-Z0-9?\%\-\/]*`) //排除2个中间的httpxx.jpg sadsd httpyy.png，这样会识别成一个
	var imgRE = regexp.MustCompile(`http.[^ ]*(png|gif|jpg|jpeg)[a-zA-Z0-9?\%\-\/]*`)
	imgs := imgRE.FindAllString(htm, -1) //得到所有匹配项

	//fmt.Println(imgRE.FindAllStringSubmatch(htm, -1))
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i]
	}
	//截取前3
	//count := 3
	if len(out) > count {
		s := out[0:count]
		return s
	} else {
		s := out[0:len(out)]
		return s
	}

}
