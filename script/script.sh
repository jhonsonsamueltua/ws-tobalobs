#!/bin/sh

ssh pi@$1 -p $2 -y ./sketchbook/tobalobs/script-remote.sh $3 &
# ssh pi@10.42.0.230 ls
# cd ..
# cd Documents/tobalobs/
# python monitor.py $1 &
# cd ..
