import http from "k6/http";
import { check, sleep } from "k6";

// Configuration for a steady load test
export let options = {
  vus: 1000, // 1000 Virtual Users
  duration: "1m", // Test duration of 1 minute
};

export default function () {
  let res = http.get("http://server:8080/get"); // Update with your server's URL
  check(res, {
    "is status 200": (r) => r.status === 200,
    "response time < 200ms": (r) => r.timings.duration < 200,
  });
  sleep(1); // Simulate user think time
}
