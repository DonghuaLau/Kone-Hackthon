#!/bin/sh

client_id="ac442003-88f2-4393-90a8-688ec41aac78"
client_secret="K2yE3dS5bX8uI6wC0hP2yU8nK5wG5dU8eI5eA6yF0qG6xD1eA0"

building_id=9990000508
from_area="area:9990000508:13000"
to_area="area:9990000508:5000"

curl --request GET \
	--url https://api.kone.com/api/building/${building_id}/area/7000/arealink/ \
	--header "accept: application/vnd.api+json" \
	--header "x-ibm-client-id: ${client_id}" \
	--header "x-ibm-client-secret: ${client_secret}"


