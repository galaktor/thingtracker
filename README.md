# ThingTracker
A stupid simple cooperative web application to track collaboration on a *thing*.

# IMPORTANT DISCLAIMER
Intended *only* for local use on your work machine, in your own private network, with a few colleagues, MAX.

It's a minimal simple app for personal usage. It is **NOT** designed to handle random user input, prevent abuse for spam, handle large amounts of incoming traffic and all the other things a public, scalable, public-facing web server would have to. This way this can stay stupic simple. But it would be crazy to use it with public access, so let me stress this point again:

**NOT FOR USE AS PUBLIC SERVER**. Do not do it. Don't even think about it. Just don't.

Clear enough?!

That said, it's super nice as a personal tool because it's so simple!

# Kudos
Blogs and libraries which I used to get this working, in no particular order:

* The http router used is [gorilla/mux](https://github.com/gorilla/mux).
* Initial early test was based on tutorial found at [TheNewStack](http://thenewstack.io/make-a-restful-json-api-go/).
* Email sending via [jordan-wright/email](https://github.com/jordan-wright/email)
* Markdown-to-HTML in emails handled with [russross/blackfriday](https://github.com/russross/blackfriday)

# TODOS
* Close things when all done?
* Highlight overdue dates
* Create timers for all open things
* Set remind interval on every thing
* ~~Send email via SMTP~~
* When timers expire, send emails to all un-done participants
* Prettify HTML
* ~~Make ThingLink <a href>~~

# License
Copyright 2015 Raphael Estrada.
The source code is licensed under the GNU General Public License. See `LICENSE` file.
