# http-from-tcp

Boot.dev Golang HTTP from TCP Project.

## TCP & UDP

Transmission Control Protocol is the primary communication protocol for the internet. Allows *ordered data* to be safely sent across the internet.

Data is split and sent in *packets*. Packets are sent and arrive potentially out of order and are reassembled on the receiver side. Without TCP you cannot guarantee the order, TCP guarantees that the data arrive in order.

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

## HTTP

HTTP/1.1 is a text protocol that works over TCP. If the HTTP request or response is too big to fit into a single TCP packet it can be broken up into many packets and reconstructed in the correct order on the receiver side.

`HTTP-message` is the format that the text in the HTTP request or response must follow (CRLF is `\r\n`):
```
start-line CRLF
*( field-line CRLF )
CRLF
[ message-body ]
```

| Part | Example | Description |
| --- | --- | --- |
| start-line CRLF | POST /users/primeagen HTTP/1.1 | The request (for a request) or status (for a response) line |
| *( field-line CRLF ) | Host: google.com | Zero or more lines of HTTP headers. These are key-value pairs. |
| CRLF | | A blank line that separates the headers from the body. |
| [ message-body ] | {"key": "value"} | The body of the message. This is optional. |

Check [RFC 9112](https://datatracker.ietf.org/doc/html/rfc9112) and [RFC 9110](https://datatracker.ietf.org/doc/html/rfc9110) for detailed info about HTTP semantics.

Example of HTTP `GET` request:
```
GET /goodies HTTP/1.1
Host: localhost:42069
User-Agent: curl/8.6.0
Accept: */*
```
- Request: `GET /goodies HTTP/1.1`
- Headers:
    ```
    Host: localhost:42069
    User-Agent: curl/8.6.0
    Accept: */*
    ```
- Body: empty

Example of HTTP `POST` request:
```
POST /coffee HTTP/1.1            # start-line CRLF
Host: localhost:42069            # *( field-line CRLF )
User-Agent: curl/8.6.0           # *( field-line CRLF )
Accept: */*                      # *( field-line CRLF )
Content-Type: application/json   # *( field-line CRLF )
Content-Length: 22               # *( field-line CRLF )
                                 # CRLF
{"flavor":"dark mode"}           # [ message-body ]
```
Similar to `GET` but with a body (`{"flavor":"dark mode"}`) and content related headers (`content-type`, `conten-length`).

## netcat

- Send data
```shell
printf "toto" | nc -w 1 127.0.0.1 42069
```
- Listen for UDP packets without initiating a connection
```shell
nc -u -l [port]
```
