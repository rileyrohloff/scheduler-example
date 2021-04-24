import requests
import json
import sys



if __name__ == '__main__':
    # obviously this isn't triggering anything, but this an example of hitting an API endpoint
    # that would kick off a downstream process. Think publishing to kafka, or kicking off another sequence.
    args = sys.argv
    print(args)
    try:
        res = requests.get("https://swapi.dev/api/people/1") if len(args) <= 2 else requests.get(f"https://swapi.dev/api/{args[1]}/{args[2]}")
        print(json.dumps(res.json(), indent=4))
    except Exception as err:
        sys.exit(err)
