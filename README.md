# scheduler-example


The Purpose for this is to demostrate how you could leverage k8s jobs/crons to achieve the task of "scheduling jobs" that run in parallel. 

Project structure

- Root

`main.go` is the 'main' cron that runs as a Golang script which kicks off other "jobs" as goroutines or synchrnous functions.

Spins up a job in as a container in k8s that pulls in the python image under `jobs/trigger-endpont`


- Jobs

    - Trigger Endpont
    
        This python file is just a placeholder that represents how easy it is to put a quick script that kicks off a downstream method.

## Run

```bash
Docker build . -t scheduler-example:1.0
```

```bash
Docker build jobs/trigger-endpoint -t trigger-job:latest
```

```bash
helm upgrade scheduler-example --install helm/ -f helm/values.yaml --namespace scheduler-example
```