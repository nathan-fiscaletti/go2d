# Go2D

A Canvas based 2D Game Engine written in Go built on SDL.

## Running a Go2D Game

If you are only trying to **run** a game that was created using Go2D, you will need to install SDL.

See [Installing SDL](./docs/sdl.md) for more information on installing SDL for your particular system.

```sh
$ cd pong
$ CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -tags static -ldflags "-s -w" pong.go
$ ./pong
```

## Demos

There are two Demos packaged with Go2D which demonstrate the use of some of it's data structures.

* [Pong](./pong): A basic example of the classic game of Pong with a mild twist where the paddles shrink in size each time the ball collides with them.

* [Shooter](./shooter): A simple Shooter game where balls will spawn at the top and you can shoot them by pressing the space bar.

