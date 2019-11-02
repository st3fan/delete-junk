# Delete Junk
*Stefan Arentz, November 2019*

This is a short Go program that goes into your IMAP Junk folder and deletes all spam with a score higher than *SpamScoreThreshold*. It takes three environment variables to determine where your IMAP account lives: `USERNAME`, `PASSWORD` and `HOSTNAME`.

I get an insane amount of spam, which is all properly filtered as spam, but not deleted. So I run this every five minutes to keep my spam folder manageable. Once in a while I go in there and move stuff out that is not spam.

May be rough since it was written for personal use. If you find this useful or have a request for an enhancement, feel free to file an issue or a pull request.


```
2019/11/02 08:42:12 DELETING 19.82 "Miracle Shake" Treats Root Cause of Diabetes
2019/11/02 08:42:12 DELETING 22.87 At home teeth whitening - snow white teeth in just 2 weeks
2019/11/02 08:42:12 DELETING 44.36 could you meet me at the weekend?
2019/11/02 08:42:12 DELETING 41.88 I can not find
2019/11/02 08:42:12 DELETING 44.08 With our brand solutions it will!
2019/11/02 08:42:12 DELETING 43.98 Try it on lower prices!
2019/11/02 08:42:12 DELETING 16.16 Best Smartphone on the market with all the latest technology
2019/11/02 08:42:12 DELETING 22.83 Drink this, Drop 10 Sizes From your waist
2019/11/02 08:42:12 DELETING 42.67 will we schedule an appointment for tomorrow?
2019/11/02 08:42:12 DELETING 38.30 Bet you'll never find better offer for Viagra.
```
