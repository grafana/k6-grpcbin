import grpc from 'k6/net/grpc';
import { check, sleep } from 'k6';

const conf = {
  baseURL: __ENV.BASE_URL || "grpcbin.test.k6.io:9001"
}

export let options = {
  stages: [
    { target: 5000, duration: "5m" },
  ],
  ext: {
    loadimpact: {
      name: "grpcbin_demo.js",
      distribution: {
        scenarioLabel2: { loadZone: "amazon:ie:dublin", percent: 100 }
      }
    }
  }
};


const client = new grpc.Client();
client.load(['definitions'], 'hello.proto');

export default () => {
  client.connect(conf.baseURL, {
    // plaintext: false
  });

  const data = { greeting: 'Bert' };
  const response = client.invoke('hello.HelloService/SayHello', data);

  check(response, {
    'status is OK': (r) => r && r.status === grpc.StatusOK,
    'response is correct': (r) => r && r.message && r.message.reply === "hello Bert"
  });

  // console.log(JSON.stringify(response.message));

  client.close();
  sleep(1);
};
