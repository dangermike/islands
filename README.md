# Islands

As an interview question, I've asked people to solve this problem, so I decided to try it out.

The goal is to count islands in a grid. The grid is defined as a contiguous land mass where contiguous land is adjacent vertically or hoizontally, but not diagonally. It is tempting to try to scan from the top-left, but that misses curls in the land.

Instead, we can treat this as a flood-fill problem. Every time we find land, we travel to every adjacent coordinate looking for more land. Every time we find land, we mark it as "visited." This could be done recursively, but if there is a _gigantic_ land mass we might blow the stack. So instead we manage our own stack. Yes, this is silly. Yes, I did it anyway.

## Use in interviewing

The solution presented here in the `countIslands` and `visit` methods is not unreasonable to expect in an interview context, though I would want to include the test data and harness.

I'm unsure of the value of these kinds of questions in interviewing, but this seems like a reasonable problem of this type.

* There isn't any reliance on fancy algorithms
* There is an obvious path that is wrong in an obvious way
* Only a few lines of code are required

## Running it

You can run through all of the test data islands using `go test` or benchmark the solution using `go test -bench .`

Additionally, you can pass in your own islands either via `stdin` or by executing with a filename. Mark land as `1` and water as `0` and the application should do the rest.
