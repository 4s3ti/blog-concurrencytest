#!/bin/python

import json
import uuid
import random
import sys

total = int(sys.argv[1])
file=(f'systems_{total}.json')
regions = ["eu-north-1", "eu-west-1", "eu-west-2", "us-west-1"]
states = ["READY", "PLAN_FAILED", "RESUME_FAILED", "PAUSED", "STORAGE_CREATION_FAILED", "FOO_FAILED"]
auxStates = ["PLANING", "RESUMING", "PAUSING", "TO_BE_PLANED", "TO_BE_PAUSED"]
types = ["team", "lab"]
versions = ["1.2.3-gec414e3c24", "2.3.2-bde194c123a", "3.4.0-foobarbaz"]
items = []

for i in range(total):
    item = {
        "id" : str(uuid.uuid4()),
        "name" : (f'system{i}'),
        "region" : random.choice(regions),
        "state" : random.choice(states),
        "auxState" : random.choice(auxStates),
        "apiEndpoint" : (f'https://system{i}.mydomain.com'),
        "vidispineType" : random.choice(types),
        "vidispineVersion" : random.choice(versions)
    }
    items.append(item)

data = {"items": items}

with open(file, "w") as f:
    json.dump(data, f, indent=2)
