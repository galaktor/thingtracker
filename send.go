package main

import(
	"github.com/jordan-wright/email"
	"github.com/russross/blackfriday"
)

func main() {
	e := email.NewEmail()
	e.From = "Foo <owner-email@foo.bar>"
	e.To = []string{"p1-email@foo.bar", "p2-email@foo.bar"}
	e.Subject = "Hello, world!"
	// send both plain text AND html to support different browsers
	e.Text = []byte("# Header\r\n\r\nThis is a [TEST](http://www.github.com)")
	e.HTML = []byte(blackfriday.MarkdownBasic(e.Text))
//	e.Send("MAILHOST:25", nil)
}
