# SMS
Four different utilities to send SMS text message from the __macos__ command line:

* `sms    ` : Send text messages via __macos__ Messages application (AppleScript).
* `sms.sh ` : Send text messages via https://textbelt.com service (Bash).
* `smscli ` : Send text messages via https://textbelt.com service (Go).
* `smsclic` : Send text messages via https://textbelt.com service (C).

These utilities were more of a quick exercise on how to write HTTP calls using different languages.

## Install
You can install any of these 4 utilities using my Homebrew Tap. For example:

`brew install lencap/tools/sms`

Once the tap is registered, the other utilities can easily be installed with just `brew install smsclic`, etc. 

Each `Makefile` within the respective language directory also allows an option to install to `/usr/local/bin/`, if you rather do that:

`make install`

CORRECTION: Didn't bother creating a brew formula for the Bash one.

## Compilation
For the ones needing compilation, read the respective `Makefile` for more info.

## Configuration
Run each utility without any arguments to see the __usage__, which provides details for what to configure. The simpler `sms` one just requires that you're logged on to your Messages app. The other 3 need a KEY (from the https://textbelt.com service) to be in a text file as:

`svckey = KEY`

Checkout the code for the name of the respective file.

## Run
Run with the 2 obvious arguments (cellphone number & the actual message):

`sms 2015554444 "Hello world"`

## Notes
These have only been tested on __macos__ 10.15.2
