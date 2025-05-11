# http-from-tcp

Boot.dev Golang HTTP from TCP Project.

## TCP & UDP

Transmission Control Protocol is the primary communication protocol for the internet. Allows *ordered data* to be safely sent across the internet.

Data is split and sent in *packets*. Packets are sent and arrive potentially out of order and are reassembled on the receiver side. Without TCP you cannot guarantee the order.

High level differences between TCP and User Datagramn Protocol (UDP):
|              | TCP | UDP |
| ---          | --- | --- |
Connection     | Yes | No  |
Handshake      | Yes | No  |
In Order       | Yes | No  |
Blazingly Fast | No  | Yes |

TCP establishes a connection between sender and receiver with a handshake and ensures that all data is sent in order. UDP just sends the data, and the receiver will make sense of it, *but UDP does not care if the receiver is listening or not*.

## Files vs Network

Files and network connections behave similarly, they are both streams of bytes you can reand from and write to.

When reading from a file, you're in control of the reading process, you *pull* data from the file:
- when to read;
- how much to read;
- when to stop reading.

When reading form a network connection, the data is *pushed* to us by the remote sender. We don't have control over when the data arrives, how much arrives or when it stops arriving.

## netcat

- Send data
```shell
printf "toto" | nc -w 1 127.0.0.1 42069
```
- Listen for UDP packets without initiating a connection
```shell
nc -u -l [port]
```
