![blaster](assets/blaster-128.png)

# Blaster
> CLI based message pump

[![Build Status](https://travis-ci.org/buddyspike/blaster.svg?branch=master)](https://travis-ci.org/buddyspike/blaster) [![codecov](https://codecov.io/gh/buddyspike/blaster/branch/master/graph/badge.svg)](https://codecov.io/gh/buddyspike/blaster) [![Go Report Card](https://goreportcard.com/badge/github.com/buddyspike/blaster)](https://goreportcard.com/report/github.com/buddyspike/blaster)

Blaster is a CLI utility to pump messages from various message queue services. Users can write custom message processing logic in their language of choice and rely on Blaster for optimal work scheduling and fault tolerance.

## Table of Contents

- [Introduction](#Introduction) 
- [Example](#Example)
- [Usage](#Usage)
   * [Global Options](#Global-Options)
   * [Broker Options](#Broker-Options)
     - [SQS](#SQS)
     - [Kafka](#Kafka)
- [Message Schema](#Message-Schema)
- [Deploy with Docker](#Deploy-With-Docker)
- [Contributing](#Contributing)
- [Credits](#Credits)

## Introduction
A typical workflow to consume messages in a message queue is; fetch one message, process, remove and repeat. This seemingly straightforward process however is often convoluted by the logic required to deal with subtleties in message processing. Following list summarises some of the common complexities without being too exhaustive. 

- Read messages in batches to reduce network round-trips.
- Enhance the work distribution by intelligently filling read ahead buffers. 
- Retry handling the messages when there are intermittent errors.
- Reduce the stress on recovering downstream services with circuit breakers.
- Process multiple messages simultaneously.
- Prevent exhausting the host resources by throttling the maximum number of messages processed at any given time.

Blaster simplifies the message handling code by providing a flexible message processing pipeline with built-in features to deal with the well-known complexities. 

It's built with minimum overhead (i.e. cpu/memory) to ensure that the handlers are cost effective when operated in pay-as-you-go infrastructures (e.g. AWS Fargate).

## How it works
Blaster has an extensible pipeline to bind to a target queue and efficiently deliver messages to user's message handler. Message handler is launched and managed as a child process of Blaster and communication between the two processes occur via a well defined interface over HTTP.

## Example

**Step 1: Write a handler**

Blaster message handler is any executable exposing the message handling logic as an HTTP endpoint. In this example, it is a script written in Javascript.

```javascript
#!/usr/bin/env node

const express = require('express');

const app = express();
app.use(express.json());

// Messages are submitted to the handler endpoint as HTTP POST requests
app.post('/', (req, res) => {
    console.log(req.body);
    res.send('ok');
});

// By default blaster forwards messages to http://localhost:8312/
app.listen(8312, () => { console.log('listening'); });
```

**Step 2: Launch it with blaster**

Now that we have a message handler, we can launch blaster to handle messages stored in a supported broker. For instance, to process messages in an AWS SQS queue called test with the script created in step 1, launch blaster with following command (this should be executed in the directory containing node script):

```
chmod +x ./handler.js
AWS_REGION=ap-southeast-2 blaster sqs --queue-name "test" --handler-command ./handler.js
```

## FAQ

**Why doesn't it support forwarding messages to an external HTTP endpoint?**
Blaster abstracts the complexities of scheduling and fault tolerance in a message consumer. In order to make smart decisions about how work is scheduled, Blaster requires visibility to resource utilisation of user's code handling the message. Blaster does not support delivering messages to an external HTTP point because this level of transparency is not easily achievable with externally managed processes (e.g. They could be executing in a remote host).

## Usage

```
blaster <broker> [options]

```

### Global Options
`--handler-command`

Command to launch the handler.

`--handler-args`

Comma separated list of arguments to the handler command.

`--max-handlers`

Maximum number of concurrent handlers to execute. By default this is calculated by multiplying the number of CPUs available to the process by 256.

`--startup-delay-seconds`

Number of seconds to wait on start before delivering messages to the handler. Default setting is five seconds. Turning startup delay off by setting it to zero will notify blaster that handler's readiness endpoint should be probed instead of a static delay. Readiness endpoint must listen for HTTP GET requests at the handler URL. When handler is ready to accept message, readiness endpoint must respond with an HTTP 200.

`--handler-url`

Endpoint that handler is listening on. Default value is http://localhost:8312/

`--retry-count`

When the handler does not respond with an HTTP 200, blaster retries the delivery for the number of times indicated by this option.

`--retry-delay-seconds`

Number of seconds to wait before retrying the delivery of a message.

`--version`

Show blaster version

### Broker Options

#### SQS

Use `AWS_REGION`, `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` env variables to specify the environment of the SQS queue.

`--queue-name`

Name of the queue. Blaster will resolve the queue URL using its name and the region set in `AWS_REGION`.

`--max-number-of-messages`

Maximum number of messages to receive in a single poll from SQS. Default setting is 1 and the maximum value supported is 10.

`--wait-time-seconds`

Blaster uses long polling when receiving messages from SQS. Use this option to control the delay between polls. Default setting is 1.

#### Kafka

Blaster creates a consumer group with the specified name to receive messages from a Kafka topic. An instance of handler executable is launched for each partition assigned to the current blaster instance. Since the handler process is isolated in its own address space, it alleviates the need to synchronise access to shared memory in handler code. As a result of this multi-process design, Kafka message handlers should listen on the designated port advertised via `BLASTER_HANDLER_PORT` environment variable (as shown in the sample code snippet below).

Kafka binding in blaster is also aware of partition re-balances that may occur due to new members (i.e. new blaster instances) joining the consumer group. During a re-balance event, blaster gracefully brings the current handler processes down and launches new ones as per new partition assignment. This is a useful feature to auto scale the message processing nodes based on their resource consumption.

```
#!/usr/bin/env node

const express = require('express');

const app = express();
app.use(express.json());

app.post('/', (req, res) => {
    console.log(`pid: ${process.pid} partion: ${req.body.properties.partitionId} offset: ${req.body.properties.offset} messageId: ${req.body.messageId}: ${req.body.body}`);
    return res.send('ok');
});

// Bind to the port assigned by blaster or default port. Using default
// port would only work if the topic has a single partition.
const port = process.env['BLASTER_HANDLER_PORT'] || 8312;
app.listen(port, () => {
    console.log('listening on port ', port);
});

```
Complete example can be found [here](samples/kafaka-node)

`--brokers`

Comma separated list of broker addresses.

`--topic`

Name of the topic to read messages from.

`--group`

Name of the consumer group. Blaster creates a handler instance for each partition assigned to a member of the consumer group. Each message is sequentially delivered to the handler in the order they are received.

`--start-from-oldest`

Force blaster to start reading the partition from oldest available offset.

`--buffer-size`

Number of messages to be read into the local buffer.

## Message Schema

Since Blaster is designed to work with many different message brokers, it converts the message to a  general purpose format before forwarding it to the handler.

```
{
    "$schema": "http://json-schema.org/schema#",
    "$id": "https://github.com/buddyspike/blaster/message-schema.json",
    "title": "Message",
    "type": "object",
    "properties": {
        "messageId": {
            "type": "string",
            "description": "Unique message id that is generally assigned by the broker"
        },
        "body": {
            "type": "string",
            "description": "Message body with the content"
        },
        "properties": {
            "type": "object",
            "description": "Additional information available in the message such as headers"
        }
    }
}
```

## Deploy with Docker

To deploy blaster handler in a docker container, copy the linux binary from [Releases](https://github.com/buddyspike/blaster/releases) to the path and set the entry point with desired options.

```
from node:10.15.3-alpine

RUN mkdir /usr/local/handler
WORKDIR /usr/local/handler
COPY .tmp/blaster /usr/local/bin/
COPY *.js *.json /usr/local/handler/

RUN npm install

ENTRYPOINT ["blaster", "sqs", "--handler-command", "./index.js", "--startup-delay-seconds", "0"]
```

Full example can be found [here](https://github.com/buddyspike/blaster/tree/master/samples/docker).

## Contributing

```
git clone https://github.com/buddyspike/blaster
cd blaster

# Run tests
make test

# Build binary
make build

# Build binary and copy it to path
make install

# Build cross compiled binaries
./build.sh
```

## Credits

- Icons made by <a href="https://www.flaticon.com/authors/nikita-golubev" title="Nikita Golubev">Nikita Golubev</a> from <a href="https://www.flaticon.com/" title="Flaticon"> www.flaticon.com</a>

<sub><sup>Made in Australia with ❤ <sub><sup>
