import http from "k6/http";
import { check, sleep } from "k6";

// Configuration for soak test
export let options = {
  vus: 20, // 20 Virtual Users
  duration: "30m", // Run for 30 minutes
};

export default function () {
  let res = http.get("http://localhost:8080");
  check(res, {
    "is status 200": (r) => r.status === 200,
  });
  sleep(1);
}
