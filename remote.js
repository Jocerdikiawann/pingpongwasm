const btnLeft = document.getElementById("left");
const btnRight = document.getElementById("right");
const btnUp = document.getElementById("up");
const btnDown = document.getElementById("down");

const options = {
  clean: true,
  connectTimeout: 4000,
  clientId: "remotePingPong2",
};
const connectUrl = "wss://broker.hivemq.com:8884/mqtt";
const client = mqtt.connect(connectUrl, options);

client.on("reconnect", (error) => {
  console.log("reconnecting:", error);
});

client.on("error", (error) => {
  console.log("Connection failed:", error);
});

btnLeft.addEventListener("click", function () {
  client.publish("pingpong/remote", "ArrowLeft");
});

btnRight.addEventListener("click", function () {
  client.publish("pingpong/remote", "ArrowRight");
});

btnUp.addEventListener("click", function () {
  client.publish("pingpong/remote", "ArrowUp");
});

btnDown.addEventListener("click", function () {
  client.publish("pingpong/remote", "ArrowDown");
});
