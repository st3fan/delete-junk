#!/usr/bin/env python3.7


#
# delete-junk.py - delete junk email with a X-Rspamd-Score higher than 17.5 to
# the trash.
#


import email, os, sys
from imapclient import IMAPClient


def delete_junk(hostname, username, password, threshold):

    with IMAPClient(hostname, use_uid=True) as client:

        client.login(username, password)
        client.select_folder('Junk')

        for msgid, data in client.fetch(client.search(), ['ENVELOPE','RFC822']).items():
            envelope = data[b'ENVELOPE']

            message = email.message_from_bytes(data[b'RFC822'])
            spam_score = float(message.get("X-Rspamd-Score", 0.0))

            if spam_score > threshold:
                print('MOVING #%d: %2.2f %s' % (msgid, spam_score, envelope.subject.decode()))
                client.move(msgid, 'Trash')


if __name__ == "__main__":

    hostname = os.getenv("JUNK_HOSTNAME")
    username = os.getenv("JUNK_USERNAME")
    password = os.getenv("JUNK_PASSWORD")

    if hostname is None or username is None or password is None:
        print("You have to set the JUNK_{HOSTNAME,USERNAME,PASSWORD} environment variables.")
        sys.exit(1)

    delete_junk(hostname, username, password, 17.5)