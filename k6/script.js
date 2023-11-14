import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    scenarios: {
        smokeTest: {
            executor: 'constant-vus',
            exec: 'testFunction',
            vus: 1,
            duration: '1m',
            startTime: '0m',
            tags: { test_type: 'smoke' },
        },
        averageLoadTest: {
            executor: 'ramping-vus',
            exec: 'testFunction',
            startVUs: 50,
            stages: [
                { duration: '5m', target: 100 }, // Normal load
            ],
            startTime: '2m',
            tags: { test_type: 'average' },
        },
        stressTest: {
            executor: 'ramping-vus',
            exec: 'testFunction',
            startVUs: 100,
            stages: [
                { duration: '2m', target: 500 }, // Ramp up
                { duration: '3m', target: 5000}, // Stay at peak
                { duration: '2m', target: 0 },  // Ramp down
            ],
            startTime: '8m',
            tags: { test_type: 'stress' },
        },
        spikeTest: {
            executor: 'ramping-vus',
            exec: 'testFunction',
            startVUs: 1,
            stages: [
                { duration: '1m', target: 10000 }, // Spike
                { duration: '1m', target: 10000 }, // Stay at spike
                { duration: '1m', target: 1 },  // Back to normal
            ],
            startTime: '15m',
            tags: { test_type: 'spike' },
        },
        // breakpointTest: {
        //     executor: 'ramping-vus',
        //     exec: 'testFunction',
        //     startVUs: 1,
        //     stages: [
        //         { duration: '10m', target: 20000 }, // Gradually ramp up
        //     ],
        //     startTime: '18m',
        //     tags: { test_type: 'breakpoint' },
        // },
        // soakTest: {
        //    executor: 'constant-vus',
        //    exec: 'testFunction',
        //    vus: 100,
        //   duration: '60m', // Long duration
        //    startTime: '0m',
        //    tags: { test_type: 'soak' },
        // }
    },
};

export function testFunction () {
    const url = 'https://client.itsag2t1.com/api/hosted/login';
    const payload = JSON.stringify({
        email: 'auyeongweibin@gmail.com',
        password: 'password',
        company: 'ascenda',
    });

    const params = {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
    };

    const response = http.post(url, payload, params);
    check(response, {
        'status is 200': (r) => r.status === 200,
    });

    sleep(1);
}
