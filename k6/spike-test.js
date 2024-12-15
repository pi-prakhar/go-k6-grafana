import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 10 },  // Ramp-up to 10 users over 30 seconds
    { duration: '10s', target: 5000 }, // Spike: Jump to 100 users
    { duration: '10s', target: 10 },  // Quick ramp-down to 10 users
    { duration: '20s', target: 0 },   // Recovery: Drop back to 0 users
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'], // 95% of requests should be below 500ms
  },
};

export default function () {
  const url = 'http://server:8080/get'; // Replace with your server's endpoint

  // Simulate a basic GET request
  const res = http.get(url);

  // Optional: Log or check the response
  if (res.status !== 200) {
    console.error(`Request failed: ${res.status} - ${res.body}`);
  }

  sleep(1); // Simulate user "think time" between requests
}