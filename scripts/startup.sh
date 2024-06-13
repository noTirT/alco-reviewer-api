#!/bin/bash
docker compose up -d
echo "Start ngrok in background on port 8080"
ngrok start alcotracker > /dev/null &
ngrok_url=$(grep "domain" "/home/tom/.config/ngrok/ngrok.yml"| sed "s/.*domain: //")

#if [ -n "$TMUX" ]; then
#    echo -n "\033]8;;$ngrok_url\033\\$ngrok_url\033]8;;\033\\"
#else
#    echo -e "ngrok public URL: \e]8;;$ngrok_url\a$ngrok_url\e]8;;\a"
#fi


