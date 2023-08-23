import grpc from 'k6/net/grpc';
import { check } from 'k6';

const conf = {
  baseURL: __ENV.BASE_URL || "grpcbin.test.k6.io:9001"
}

export let options = {
  stages: [
    { target: 10, duration: "30s" },
  ]
};

const client = new grpc.Client();
client.load(['definitions'], 'addsvc.proto');

export default () => {
  console.log('connecting: '+ conf.baseURL);
  client.connect(conf.baseURL, {
    // plaintext: false
  });

  const response = client.invoke('addsvc.Add/Sum', {
    a: 1,
    b: 2,
  });
  console.log(response.message.v); // should print 3


  check(response, {
    'status is OK': (r) => r && r.status === grpc.StatusOK,
    'response is correct': (r) => r && r.message && r.message.v === 3
  });


  client.close();
};
