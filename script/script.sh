#!/bin/sh

ssh -l pi proxy73.rt3.io -p 35745 ./sketchbook/tobalobs/script-remote.sh $1 &
# ssh pi@10.42.0.230 ls
# cd ..
# cd Documents/tobalobs/
# python monitor.py $1 &
# cd ..
