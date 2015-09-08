package main

import(
	"fmt"
	
	"github.com/jordan-wright/email"
	"github.com/russross/blackfriday"
)

// TODO: use text template api
const texttempl = `### Hi there!
%v asked you recently if you could have a look at A Thing. It would be much appreciated if you could find a few minutes to have a look.

[%v](%v)

Once you're finished, you can mark yourself 'done' in the 'edit' form.

Thank you very much for taking the time.
-Thingtracker`

func (t *Thing) EmailParticipants() error {
	if len(t.Participants) == 0 {
		return nil
	}

	e := email.NewEmail()
	e.From = t.Owner.Email
	e.To = []string{}
	for _,p := range t.Participants {
		e.To = append(e.To, p.Email)
	}
	e.Subject = "A friendly reminder..."
	e.Text = []byte(fmt.Sprintf(texttempl, t.Owner.Email, t.ThingName, t.ThingLink))
	e.HTML = []byte(blackfriday.MarkdownBasic(e.Text))
	fmt.Printf("%v\n", e.To)
	println(string(e.Text))
	
	return e.Send("MAILHOST:25", nil)
//	return nil

}
