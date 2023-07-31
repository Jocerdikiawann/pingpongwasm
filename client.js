const options = {
  clean: true,
  connectTimeout: 4000,
  clientId: "remotePingPong1",
};
const connectUrl = "wss://broker.hivemq.com:8884/mqtt";
const client = mqtt.connect(connectUrl, options);

client.subscribe("pingpong/remote", { qos: 2 });

client.on("reconnect", (error) => {
  console.log("reconnecting:", error);
});

client.on("error", (error) => {
  console.log("Connection failed:", error);
});

client.on("message", (_, message) => {
  Box(message.toString());
});
