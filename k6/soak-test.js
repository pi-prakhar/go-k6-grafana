import http from "k6/http";
import { check, sleep } from "k6";

// Configuration for soak test
export let options = {
  vus: 1000, // 20 Virtual Users
  duration: "30m", // Run for 30 minutes
};

export default function () {
  let res = http.get("http://server:8080/get");
  check(res, {
    "is status 200": (r) => r.status === 200,
  });
  sleep(1);
}
