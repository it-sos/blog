#!/bin/bash

trap 'stopMonitor; exit' SIGINT SIGTERM SIGHUP SIGQUIT SIGKILL

function checkMonitor()
{
    pgrep -f 'studynotes' > /dev/null
}

function stopMonitor()
{
    pkill -f 'studynotes'
}

function startMonitor()
{
    checkMonitor
    if test $? == 0; then
        return 0
    fi
    basedir=$(dirname $0)
    nohup $basedir/../cmd/studynotes/studynotes -w $basedir/../ >> studynotes.log 2>&1 &
}

startMonitor

while :
do
    checkMonitor
    if test $? != 0; then
        break
    fi
    sleep 1
done
