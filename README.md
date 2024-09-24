# Grid of Greed

## Introduction

A game server for a iOS game that allows collaborates to build designs with a simple grid and color system.

## Upcoming Updates

### Game Features

- Allow grid color selection options.
- Adding a player list and highlights to show current players and
  last game grid updates to improve interactivity.
- Adding to login / signup for user persistance.
- Adding account info api for stat tracking.

### Bugfixes, Performance Updates and Tech Improvements

- Update message protocol to reduce payload size for reduced latency.
- Handle race conditions between multiple concurrent users accessing the same
  resources.

## The Server

Built with websockets, the server serves as a central for all the game logic and state, communicating closley with the iOS frontend.

The server primarily uses two concurrent goroutines that communicate with one another.

### Handle Players Connections

Player connections is handled in its own concurrent goroutine, serving each unique connection and reading their incoming messages. The messages are then sent along a channel to the be handled in the _Message Hub_.

### Message Hub

The message hub is a concurrent goroutine that handles all the incoming websocket messages, distrubuting the received game channel messages to the correct game logic and websocket response.
