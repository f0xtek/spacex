# spacex

Go client to get various details about SpaceX.

At the moment, it only retrieves the following details abut the latest launch:

- Mission Name
- Flight Number
- Launch Site
- Launch Date
- YouTube Video Link

## Usage

```bash
$ spacex
The latest successful SpaceX mission was Starlink 1 (v0.9).
The flight number was 79 and launched from Cape Canaveral Air Force Station Space Launch Complex 40 on 24 May 19 03:30 BST.
The rocker used was a Falcon 9.

Here is a link to the launch video stream: https://www.youtube.com/watch?v=riBaVeDTEWI
```

## Building from source

```bash
git clone https://github.com/drmarconi/spacex.git
cd spacex
go build
```
