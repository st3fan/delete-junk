// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

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
		for name, value := range msg.Body {
			if name.BodyPartName.Specifier == imap.HeaderSpecifier {
				m, err := mail.ReadMessage(value)
				if err != nil {
					continue
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
			log.Println("Failed to mark messages to be deleted:", err)
			return
		}

		if err := c.Expunge(nil); err != nil {
			log.Println("Failed to expunge messages:", err)
			return
		}
	}
}
