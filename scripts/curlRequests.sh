#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/friends/1
curl localhost:9090/invites/2
curl localhost:9090/relationships -XPOST -d '{"user1": {"userid":2, "relationshiptype":3}, "user2": {"userid":6, "relationshiptype":4}}'
curl localhost:9090/relationships -XPUT -d '{"id":1, "user1": {"userid":1, "relationshiptype":1}, "user2": {"userid":2, "relationshiptype":1}}'
curl localhost:9090/relationships/1 -XDELETE
