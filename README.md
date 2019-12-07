Have two modules:
* `executable`: Executable file which is compiled
* `libEmbed`: Library where pkger is used to embed a file

My issue:
* Test cases in `/libEmbed` open file and pass without issue. Run `go test ./...` in this module and all tests will pass.
* Test cases in `/executable` cannot find the file on the local drive. Run `go test ./...` in this module and all tests will *FAIL*.
* Executable file works just fine. File will be found in executable. Can go into `/build` and run the `build.sh` script, which will run `pkger` and create the `pkged.go`, and build an executable called `main` in the same directory.

Suspected Cause:
* When loading items in dev environment, pkger is using the `/executabe` module name in the development
environment for some reason, instead of the module in which the function executing lives.