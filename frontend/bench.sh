#!/usr/bin/env bash

C=1000
N=200
SLEEP=3

rm ./logs/*

for loop in 1 2 3 4 5
do
    wsbench echo -c ${C} -n 1000 -p 100 -u ws://localhost:8000/connect -o ./logs/gws.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 1000 -p 100 -u ws://localhost:8001/connect -o ./logs/gorilla.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 1000 -p 100 -u ws://localhost:8002/connect -o ./logs/nhooyr.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 1000 -p 100 -u ws://localhost:8003/connect -o ./logs/gobwas.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 1000 -p 100 -u ws://localhost:8004/connect -o ./logs/nbio.log
    sleep ${SLEEP}
done

for loop in 1 2 3 4 5
do
    wsbench echo -c ${C} -n 500 -p 1000 -u ws://localhost:8000/connect -o ./logs/gws.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 500 -p 1000 -u ws://localhost:8001/connect -o ./logs/gorilla.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 500 -p 1000 -u ws://localhost:8002/connect -o ./logs/nhooyr.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 500 -p 1000 -u ws://localhost:8003/connect -o ./logs/gobwas.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 500 -p 1000 -u ws://localhost:8004/connect -o ./logs/nbio.log
    sleep ${SLEEP}
done

for loop in 1 2 3 4 5
do
    wsbench echo -c ${C} -n 200 -p 4000 -u ws://localhost:8000/connect -o ./logs/gws.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 200 -p 4000 -u ws://localhost:8001/connect -o ./logs/gorilla.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 200 -p 4000 -u ws://localhost:8002/connect -o ./logs/nhooyr.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 200 -p 4000 -u ws://localhost:8003/connect -o ./logs/gobwas.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 200 -p 4000 -u ws://localhost:8004/connect -o ./logs/nbio.log
    sleep ${SLEEP}
done

for loop in 1 2 3 4 5
do
    wsbench echo -c ${C} -n 100 -p 8000 -u ws://localhost:8000/connect -o ./logs/gws.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 100 -p 8000 -u ws://localhost:8001/connect -o ./logs/gorilla.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 100 -p 8000 -u ws://localhost:8002/connect -o ./logs/nhooyr.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 100 -p 8000 -u ws://localhost:8003/connect -o ./logs/gobwas.log
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 100 -p 8000 -u ws://localhost:8004/connect -o ./logs/nbio.log
    sleep ${SLEEP}
done

node statistics.js