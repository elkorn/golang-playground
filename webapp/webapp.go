package main

import (
	"bytes"
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"code.google.com/p/freetype-go/freetype/raster"
	"code.google.com/p/goauth2/oauth"
)

var config = &oauth.Config{
	ClientId:     "test",
	ClientSecret: "secret",
	Scope:        "https://www.googleapis.com/auth/buzz",
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://accounts.google.com/o/oauth2/token",
	RedirectURL:  "http://localhost:8080/post",
}

func errorHandler(fn http.HandlerFunc) http.HandlerFunc {
	errorTpl, _ := template.ParseFiles("error.html")
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				fmt.Println("Sending error response.")
				w.WriteHeader(500)
				errorTpl.Execute(w, e)
			}

			fmt.Println("All ok.")
		}()
		fn(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	uploadTpl, _ := template.ParseFiles("upload.html")
	uploadTpl.Execute(w, "")
}

func rgba(m image.Image) *image.RGBA {
	if r, ok := m.(*image.RGBA); ok {
		return r
	}

	b := m.Bounds()
	r := image.NewRGBA(b)
	draw.Draw(r, b, m, image.ZP, draw.Src)
	return r
}

func pt(x, y int) raster.Point {
	return raster.Point{X: raster.Fix32(x << 8), Y: raster.Fix32(y << 8)}
}

func moustache(m image.Image, x, y, size, droopFactor int) image.Image {
	mrgba := rgba(m)
	p := raster.NewRGBAPainter(mrgba)
	p.SetColor(color.RGBA{0, 0, 0, 255})

	w, h := m.Bounds().Dx(), m.Bounds().Dy()
	r := raster.NewRasterizer(w, h)

	var (
		mag   = raster.Fix32((10 + size) << 8)
		width = pt(20, 0).Mul(mag)
		mid   = pt(x, y)
		droop = pt(0, droopFactor).Mul(mag)
		left  = mid.Sub(width).Add(droop)
		right = mid.Add(width).Add(droop)
		bow   = pt(0, 5).Mul(mag).Sub(droop)
		curlx = pt(10, 0).Mul(mag)
		curly = pt(0, 2).Mul(mag)
		risex = pt(2, 0).Mul(mag)
		risey = pt(0, 5).Mul(mag)
	)

	r.Start(left)
	r.Add3(
		mid.Sub(curlx).Add(curly),
		mid.Sub(risex).Sub(risey),
		mid)
	r.Add3(
		mid.Add(risex).Sub(risey),
		mid.Add(curlx).Add(curly),
		right)
	r.Add2(
		mid.Add(bow),
		left)
	r.Rasterize(p)
	return mrgba
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		handler(w, r)
		return
	}

	fmt.Println("Uploading the file.")
	f, _, _ := r.FormFile("image")

	// At this point we are sure that the file exists so we can order for it to be closed.
	defer f.Close()
	t, _ := ioutil.TempFile(".", "image-")

	defer t.Close()
	io.Copy(t, f)
	fmt.Printf("Uploading image %v\n", t.Name())
	http.Redirect(w, r, "/edit?id="+t.Name()[6:], 302)
}

func view(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, "image-"+r.FormValue("id"))
}

func edit(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("edit.html")
	tpl.Execute(w, r.FormValue("id"))
}

func img(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Imaging.")
	f, _ := os.Open("image-" + r.FormValue("id"))
	m, _, _ := image.Decode(f)

	v := func(n string) int { // helper closure
		i, _ := strconv.Atoi(r.FormValue(n))
		return i
	}

	m = moustache(m, v("x"), v("y"), v("s"), v("d"))

	b := &bytes.Buffer{}
	w.Header().Set("Content-Type", "image/jpeg")
	jpeg.Encode(w, m, nil)
	b.WriteTo(w)
}

func main() {
	fmt.Println("Setting up...")
	http.HandleFunc("/", errorHandler(upload))
	http.HandleFunc("/view", errorHandler(view))
	http.HandleFunc("/edit", errorHandler(edit))
	http.HandleFunc("/img", errorHandler(img))
	fmt.Println("Ready to handle.")
	http.ListenAndServe(":8080", nil)
}
