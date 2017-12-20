const { events, Job , Group} = require("brigadier");
const dest = "$GOPATH/src/github.com/technosophos/ulid";

events.on("push", (e, p) => {
  console.log(e.payload)
  var gh = JSON.parse(e.payload)
  // var test = new Job("test", "golang:1.9")
  // test.storage.enabled = false
  // test.tasks = [
  //   "mkdir -p " + dest,
  //   "cp -a /src/* " + dest,
  //   "cd " + dest,
  //   "go get -u github.com/golang/dep/cmd/dep",
  //   "dep ensure",
  //   "make test"
  // ];
  // test.run()
  var job = new Job("do-nothing", "alpine:3.4")
  job.run()
});