package main

import (
	"log"
	"os"
	"net/mail"
	"strconv"

	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-imap"
)

const SpamScoreThreshold = 15.0 // Could be lower?

func main() {
	c, err := client.DialTLS(os.Getenv("HOSTNAME") + ":993", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	if err := c.Login(os.Getenv("USERNAME"), os.Getenv("PASSWORD")); err != nil {
		log.Println("Failed to login:", err)
		return
	}

	mbox, err := c.Select("Junk", false)
	if err != nil {
		log.Println("Failed to select Junk folder:", err)
		return
	}

	seqset := new(imap.SeqSet)
	seqset.AddRange(1, mbox.Messages)

	messages := make(chan *imap.Message, mbox.Messages)

	section := &imap.BodySectionName{
		BodyPartName: imap.BodyPartName{
			Specifier: imap.HeaderSpecifier,
			Fields: []string{"X-Rspamd-Score"},
		},
	}


	if err := c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}, messages); err != nil {
		log.Println("Failed to fetch:", err)
		return
	}

	seqsetToDelete := new(imap.SeqSet)

	for msg := range messages {
		//log.Println("*", msg.Envelope.Subject)
		for name, value := range msg.Body {
			if name.BodyPartName.Specifier == imap.HeaderSpecifier {
				//log.Println(value)
				//r := strings.NewReader(value)
				m, err := mail.ReadMessage(value)
				if err != nil {
					log.Fatal(err)
				}

				score, err := strconv.ParseFloat(m.Header.Get("X-Rspamd-Score"), 64)
				if err != nil {
					continue
				}

				if score > SpamScoreThreshold {
					seqsetToDelete.AddNum(msg.SeqNum)
					log.Printf("DELETING %.2f %s", score, msg.Envelope.Subject)
				}
			}
		}
	}

	if !seqsetToDelete.Empty() {
		item := imap.FormatFlagsOp(imap.AddFlags, true)
		value := []interface{}{imap.DeletedFlag}
		if err := c.Store(seqsetToDelete, item, value, nil); err != nil {
			log.Fatal(err)
		}

		if err := c.Expunge(nil); err != nil {
			log.Println("Failed to Expunge:", err)
		}
	}
}
