import http from "k6/http";
import { check, sleep } from "k6";

import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";
import { textSummary } from "https://jslib.k6.io/k6-summary/0.0.2/index.js";

// Test configuration
export const options = {
  thresholds: {
    http_req_duration: ["p(99) < 50"],
  },
  stages: [
    { duration: "10s", target: 50 },
  ],
};


export function setup() {
  const timestamp = Date.now();
  const randomSuffix = Math.floor(Math.random() * 1000000);

  let payload = {
    name: `TestTeam_${timestamp}_${randomSuffix}`,
    users: Array.from({length: 3}, (_, i) => ({
        id: `user_id_${i+1}_${timestamp}`,
        username: `user_${i+1}_${timestamp}`,
        is_active: true,
    }))
  };

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const response = http.post("http://0.0.0.0:8080/api/v1/team/add/", JSON.stringify(payload), params);

  check(response, {
    'setup completed successfully': (r) => r.status === 201 || r.status === 200,
  });

  if (response.status !== 201 && response.status !== 200) {
    console.error('Setup failed:', response.body);
    throw new Error('Setup failed');
  }

  const responseData = response.json();

  return responseData.data || responseData;
}

export default function(data) {


  let payload = {
    created_by_id: data.members[0].id,
    pull_request_id: `pr_id_for_${data.members[0].id}_${Date.now()}`,
    pull_request_name: `pr_name_for_${data.members[0].id}__${Date.now()}`,
  };

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const response = http.post("http://0.0.0.0:8080/api/v1/pr/open/", JSON.stringify(payload), params);
  console.log(response.status)
  check(response, {
    'status is 201': (r) => r.status === 201,
  });

  sleep(2);
}
