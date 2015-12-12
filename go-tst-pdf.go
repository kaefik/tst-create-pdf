package main
  import (
    "fmt"
    "github.com/signintech/gopdf"
  )

  func main() {

   fmt.Println("start program")

    pdf := gopdf.GoPdf{}
    pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
    pdf.AddPage()
//    err := pdf.AddTTFFont("HDZB_5", "ttf/wts11.ttf")
	err := pdf.AddTTFFont("HDZB_5", "ttf/Roboto-Regular.ttf")
    if err != nil {
        fmt.Print(err.Error())
        return
    }

    err = pdf.SetFont("HDZB_5", "", 4)
    if err != nil {
        fmt.Print(err.Error())
        return
    }
	pdf.Cell(nil, "Hello")
	pdf.Cell(nil, " Привет как дела?")
    pdf.WritePdf("hello.pdf")
	fmt.Println("end program")
  }
