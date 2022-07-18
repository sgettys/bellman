# bellman
Town crier utility for sending messages. This utility will read contents from a file or from stdin and send that content to all of the configured "criers". Criers are configured via the bellman.yaml config file.

# Usage
By default the stdout crier will be configured if no config file is specified. This will simply echo the data to stdout. This can be useful for debugging when using bellman to construct complex messages.
ex: 
```
bellman -d 'some data'
```

Bellman can also read file contents and forward that to its criers using the `-f` flag.
ex:
```
bellman -f /some/file/path -c my_bellman_config.yaml
```

A specific configured crier can be selected by specifying the name using the `-n` flag. If no crier is specified then the message is sent to all of the configured criers. No configured crier will result in the message being sent to stdout.

# Criers
Criers can be specified in the config file under the `receivers` field.

Example config file:
```
loglevel: trace
receivers:
  - name: logger
    stdout: {}
```
### Stdout
```
  - name: logger
    stdout: {}
```
### Azsb
Azure Service Bus will attempt to connect to the service bus using the connection string provided under `endpoint`. If it is unable to authenticate with the parameters found in the connection string it will then attempt to connect using the environment.
It uses the following logic to attempt to connect:
  * 1. Client Credentials: attempt to authenticate with a Service Principal via "AZURE_TENANT_ID", "AZURE_CLIENT_ID" and "AZURE_CLIENT_SECRET"
  * 2. Client Certificate: attempt to authenticate with a Service Principal via "AZURE_TENANT_ID", "AZURE_CLIENT_ID", "AZURE_CERTIFICATE_PATH" and "AZURE_CERTIFICATE_PASSWORD"
  * 3. Managed Identity (MI): attempt to authenticate via the MI assigned to the Azure resource


The Azure Environment used can be specified using the name of the Azure Environment set in "AZURE_ENVIRONMENT" var.

```
  - name: exazsb
    azsb:
        endpoint: <azure service bus endpoint>
        topic: <topic to publish to>

```
### Nats
```
  - name: exnats
    nats:
      endpoint: <nats endpoint>
      topic: <topic to publish to>

```

## Adding a new Crier
To add a new crier add code that implements the [crier interface](./pkg/criers/crier.go). Add the name of the crier to the [ReceiverConfig](./pkg/criers/receiver.go) and the [GetCrier](./pkg/criers/receiver.go) factory method.