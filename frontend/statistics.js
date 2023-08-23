const fs = require('fs');
const pwd = __dirname;
let results = [];

results.push(...loadFramework('gws'));
results.push(...loadFramework('gorilla'));
results.push(...loadFramework('nhooyr'));
results.push(...loadFramework('gobwas'));
console.log(results);
fs.writeFileSync(pwd + '/public/final.json', JSON.stringify(results));

function loadFramework(framework) {
    let content = fs.readFileSync(pwd + '/logs/' + framework + ".log", 'utf8');
    let arr = content.split('\n');
    let m = {};

    arr.forEach((s) => {
        if (s === '') {
            return
        }
        let item = JSON.parse(s);
        let child = m[item.payload];
        if (child === undefined) {
            m[item.payload] = [];
        }
        m[item.payload].push({
            payload: item.payload,
            framework: framework,
            iops: item.iops,
            p50: item.p50,
            p90: item.p90,
            p99: item.p99,
        });
    })

    for (let k in m) {
        m[k].sort((a, b) => {
            if (a.iops > b.iops) { return 1; }
            else if (a.iops < b.iops) { return -1; }
            else { return 0; }
        });
    }

    let results = [];
    for (let k in m) {
        let v = m[k];
        let index = Math.floor(v.length / 2);
        results.push(v[index]);
    }
    return results;
}