#!/bin/bash
# sms.sh
# v1.1.0
# Use textbelt.com services to send SMS

[[ "$#" -ne 2 ]] && printf "Usage: $0 TEL 'message'\n" && exit
Tel="$1"
Msg="$2"
KeyName=svckey
Cfg=~/.smsclirc
Key="$(grep $KeyName $Cfg | awk '{print $3}')"

if [[ $Key ]]; then 
  printf "[$Tel]\n"
  printf "[$Msg]\n"
  curl -X POST "https://textbelt.com/text" \
    --data-urlencode key=$Key \
    --data-urlencode phone="$Tel" \
    --data-urlencode message="$Msg"
  printf "\n"
else
  printf "Error. Unable to locate '$KeyName = KEY' in $Cfg\n"
fi

exit 0
