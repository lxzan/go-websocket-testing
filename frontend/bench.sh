#!/usr/bin/env bash

C=1000
SLEEP=5
Concurrency=8

rm ./logs/*

for loop in 1 2 3 4 5
do
    wsbench echo -c ${C} -n 2000000 -p 500 -u ws://127.0.0.1:8000/connect -o ./logs/gws.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 2000000 -p 500 -u ws://127.0.0.1:8001/connect -o ./logs/gorilla.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 2000000 -p 500 -u ws://127.0.0.1:8002/connect -o ./logs/nhooyr.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 2000000 -p 500 -u ws://127.0.0.1:8003/connect -o ./logs/gobwas.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 2000000 -p 500 -u ws://127.0.0.1:8004/connect -o ./logs/nbio.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 2000000 -p 500 -u ws://127.0.0.1:8005/connect -o ./logs/x-net.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
done

for loop in 1 2 3 4 5
do
    wsbench echo -c ${C} -n 1000000 -p 2000 -u ws://127.0.0.1:8000/connect -o ./logs/gws.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 1000000 -p 2000 -u ws://127.0.0.1:8001/connect -o ./logs/gorilla.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 1000000 -p 2000 -u ws://127.0.0.1:8002/connect -o ./logs/nhooyr.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 1000000 -p 2000 -u ws://127.0.0.1:8003/connect -o ./logs/gobwas.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 1000000 -p 2000 -u ws://127.0.0.1:8004/connect -o ./logs/nbio.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 1000000 -p 2000 -u ws://127.0.0.1:8005/connect -o ./logs/x-net.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
done

for loop in 1 2 3 4 5
do
    wsbench echo -c ${C} -n 500000 -p 8000 -u ws://127.0.0.1:8000/connect -o ./logs/gws.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 500000 -p 8000 -u ws://127.0.0.1:8001/connect -o ./logs/gorilla.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 500000 -p 8000 -u ws://127.0.0.1:8002/connect -o ./logs/nhooyr.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 500000 -p 8000 -u ws://127.0.0.1:8003/connect -o ./logs/gobwas.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 500000 -p 8000 -u ws://127.0.0.1:8004/connect -o ./logs/nbio.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
    wsbench echo -c ${C} -n 500000 -p 8000 -u ws://127.0.0.1:8005/connect -o ./logs/x-net.log --concurrency=${Concurrency} --latency
    sleep ${SLEEP}
done

node statistics.js