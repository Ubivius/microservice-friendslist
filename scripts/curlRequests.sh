#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/friends/1
curl localhost:9090/invites/2
curl localhost:9090/relationships -XPOST -d '{"user_1": {"user_id":2, "relationship_type":3}, "user_2": {"user_id":6, "relationship_type":4}}'
curl localhost:9090/relationships -XPUT -d '{"id":1, "user_1": {"user_id":1, "relationship_type":1}, "user_2": {"user_id":2, "relationship_type":1}}'
curl localhost:9090/relationships/1 -XDELETE
