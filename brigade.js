const { events, Job , Group} = require("brigadier");
const dest = "$GOPATH/src/github.com/technosophos/ulid";

events.on("push", (e, p) => {
  console.log(e.payload)
  var gh = JSON.parse(e.payload)
  var test = new Job("test", "golang:1.9")
  test.tasks = [
    "mkdir -p " + dest,
    "cp -a /src/* " + dest,
    "cd " + dest,
    "go get -u github.com/golang/dep/cmd/dep",
    "dep ensure",
    "make test"
  ];
  test.run()
});