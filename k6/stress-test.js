import http from "k6/http";
import { check, sleep } from "k6";

// Configuration for stress test
export let options = {
  stages: [
    { duration: "1m", target: 10 }, // Start with 10 users
    { duration: "1m", target: 500 }, // Gradually increase to 50 users
    { duration: "1m", target: 1000 }, // Gradually increase to 100 users
    { duration: "1m", target: 2000 }, // Gradually increase to 200 users
    { duration: "1m", target: 300 }, // Peak load: 300 users
    { duration: "1m", target: 0 }, // Ramp-down to 0 users
  ],
};

export default function () {
  let res = http.get("http://server:8080/get");
  check(res, {
    "is status 200": (r) => r.status === 200,
  });
  sleep(1);
}
