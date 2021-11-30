# A simple webhook caller #

This image is primarily used to ping/call a URL on regular intervals using Kubernetes (k8s) CronJob.

A sample job would look something like this:
```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: hello
spec:
  schedule: "*/5 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: hello
            image: goavega-software/invoke-webhook
            imagePullPolicy: IfNotPresent
            env:
            - name: URL
              value: "https://reqres.in/api/users/2"
            - name: METHOD
              value: "GET"
            - name: PAYLOAD
              value: ""
          restartPolicy: OnFailure
```

#golang #kubernetes