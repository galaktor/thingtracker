package main

import(
	"fmt"
	
	"github.com/jordan-wright/email"
	"github.com/russross/blackfriday"
)

// TODO: use text template api
const texttempl = `### Hi there!
%v asked you recently if you could have a look at [A Thing](%v). It would be much appreciated if you could find a few minutes to do so.

[%v](%v)

Once you're finished, you can mark yourself 'done' in the '[edit](%v)' form.

Thank you very much for taking the time.

*Thingtracker*`

func (t *Thing) EmailParticipants() error {
	if len(t.Participants) == 0 {
		return nil
	}

	showurl := fmt.Sprint(URL_ROOT, "/show/", t.Id)
	editurl := fmt.Sprint(URL_ROOT, "/edit/", t.Id)
	
	e := email.NewEmail()
	e.From = t.Owner.Email
	e.To = []string{}
	for _,p := range t.Participants {
		if !p.Done {
			e.To = append(e.To, p.Email)
		}
	}
	e.Subject = "A friendly reminder..."
	e.Text = []byte(fmt.Sprintf(texttempl, t.Owner.Email, showurl, t.ThingName, t.ThingLink, editurl))
	e.HTML = []byte(blackfriday.MarkdownBasic(e.Text))
	fmt.Printf("%v\n", e.To)
	println(string(e.Text))
	
	return e.Send("MAILHOST:25", nil)
//	return nil

}
