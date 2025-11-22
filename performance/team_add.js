import http from "k6/http";
import { check, sleep } from "k6";

// Test configuration
export const options = {
  thresholds: {
    http_req_duration: ["p(99) < 50"],
  },
  stages: [
    { duration: "10s", target: 10 },
  ],
};


function generateRandomData() {
  const timestamp = Date.now();
  const randomSuffix = Math.floor(Math.random() * 1000000);
  const vuId = __VU;
  const iterId = __ITER;

  return {
    name: `TestTeam_VU${vuId}_Iter${iterId}_${timestamp}_${randomSuffix}`,
    users: Array.from({length: 6}, (_, i) => ({
        id: `user_id_VU${vuId}_Iter${iterId}_${i+1}_${timestamp}`,
        username: `user_VU${vuId}_Iter${iterId}_${i+1}_${timestamp}`,
        is_active: Math.random() > 0.1,
    }))
  };
}

export default function () {
  let res = http.post(
    "http://0.0.0.0:8080/api/v1/team/add/",
    JSON.stringify(generateRandomData())
);
  check(res, { "status was 200": (r) => r.status < 300 });
  sleep(2);
}
