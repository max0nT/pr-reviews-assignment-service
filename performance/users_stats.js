import http from "k6/http";
import { check, sleep } from "k6";
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";


export const options = {
  thresholds: {
    http_req_duration: ["p(99) < 100"],
  },
  stages: [
    { duration: "10s", target: 50 },
  ],
};


export function handleSummary(data) {
  return {
    "summary.html": htmlReport(data),
  };
}


export function setup() {
  const timestamp = Date.now();
  const totalUsers = 600;
  const usersPerTeam = 10;
  const totalTeams = totalUsers / usersPerTeam;
  const prsPerUser = 5;

  for (let teamIndex = 0; teamIndex < totalTeams; teamIndex++) {
    const teamSuffix = `${timestamp}_${teamIndex}`;

    let payload = {
      name: `TestTeam_${teamSuffix}`,
      users: Array.from({length: usersPerTeam}, (_, i) => {
        const userNumber = teamIndex * usersPerTeam + i + 1;
        return {
          id: `user_id_${userNumber}_${timestamp}`,
          username: `user_${userNumber}_${timestamp}`,
          is_active: true,
        };
      })
    };

    let params = {
      headers: {
        'Content-Type': 'application/json',
      },
    };

    let response = http.post("http://0.0.0.0:8080/api/v1/team/add/", JSON.stringify(payload), params);

    check(response, {
      [`team ${teamIndex + 1} created successfully`]: (r) => r.status === 201 || r.status === 200,
    });

    if (response.status !== 201 && response.status !== 200) {
      console.error(`Team ${teamIndex + 1} creation failed:`, response.body);
      continue;
    }

    const responseData = response.json();

    for (let userIndex = 0; userIndex < usersPerTeam; userIndex++) {
      const user = responseData.members[userIndex];
      const userNumber = teamIndex * usersPerTeam + userIndex + 1;

      for (let prIndex = 0; prIndex < prsPerUser; prIndex++) {
        let prPayload = {
          created_by_id: user.id,
          pull_request_id: `pr_id_${userNumber}_${prIndex + 1}_${timestamp}`,
          pull_request_name: `pr_name_${userNumber}_${prIndex + 1}_${timestamp}`,
        };

        let prParams = {
          headers: {
            'Content-Type': 'application/json',
          },
        };

        let prResponse = http.post("http://0.0.0.0:8080/api/v1/pr/open/", JSON.stringify(prPayload), prParams);

        check(prResponse, {
          [`PR ${prIndex + 1} for user ${userNumber} created successfully`]: (r) => r.status === 201 || r.status === 200,
        });

        if (prResponse.status !== 201 && prResponse.status !== 200) {
          console.error(`PR ${prIndex + 1} creation failed for user ${userNumber}:`, prResponse.body);
        }
      }
    }
  }
}


export default function() {
    let res = http.get(
        "http://0.0.0.0:8080/api/v1/user/stats/",
    );
    check(res, { "status was 200": (r) => r.status < 300 });
}
