#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/friends/a2181017-5c53-422b-b6bc-036b27c04fc8
curl localhost:9090/invites/e2382ea2-b5fa-4506-aa9d-d338aa52af44
curl localhost:9090/relationships -XPOST -d '{"user_1": {"user_id":"a2181017-5c53-422b-b6bc-036b27c04fc8", "relationship_type":"PendingOutgoing"}, "user_2": {"user_id":"e2382ea2-b5fa-4506-aa9d-d338aa52af44", "relationship_type":"PendingIncoming"}}'
curl localhost:9090/relationships -XPUT -d '{"id":"eb9aff9f-8c4e-47c3-9f6d-bd9aac3d9f31", "user_1": {"user_id":"a2181017-5c53-422b-b6bc-036b27c04fc8", "relationship_type":"Friend"}, "user_2": {"user_id":"e2382ea2-b5fa-4506-aa9d-d338aa52af44", "relationship_type":"Friend"}}'
curl localhost:9090/relationships/a2181017-5c53-422b-b6bc-036b27c04fc8 -XDELETE
