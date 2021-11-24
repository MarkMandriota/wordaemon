# wordaemon
*simple helper for word game*

Replaces a word in the clipboard with a random word in the dictionary starting with the last letter of the original word.
* Text changes every of 1/3 of second. Default phase can be changed copying `~<time.Duration>` to the clipboard.

* To exit from daemon, copy to the clipboard `-`.
* If you want only stop daemon, temporary, copy `]`, you can resume it copying `[`.


### Commands list
* `~<time.Duration>` - set phase
* `-` - send kill signal
* `]` - stop
* `[` - resume
