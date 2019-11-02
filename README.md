# Delete Junk
*Stefan Arentz, November 2019*

This is a short Go program that goes into your IMAP Junk folder and deletes all spam with a score higher than *SpamScoreThreshold*. It takes three environment variables to determine where your IMAP account lives: `USERNAME`, `PASSWORD` and `HOSTNAME`.

I get an insane amount of spam, which is all properly filtered as spam, but not deleted. So I run this every five minutes to keep my spam folder manageable. Once in a while I go in there and move stuff out that is not spam.

May be rough since it was written for personal use. If you find this useful or have a request for an enhancement, feel free to file an issue or a pull request.

