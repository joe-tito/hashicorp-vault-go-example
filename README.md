# Go Example

This repository is meant to demonstrate some of the basic functionality of HashiCorp Vault. Specifically, it demonstrates storing and accessing secrets stored in HashiCorp's KV2 (key/value pair) engine. 

This repository demonstrates the following Vault functionality:

- How to connect a Java application to a running instance of Vault
- How to write a key/value secret to Vault
- How to read a key/value secret from Vault

---

## Go Installation

If you do not have Go installed on your machine already, please follow the instructions below to install it.

1) Connect to your Linux instance 

2) Install Go

    - Amazon Linux
    ```bash
    sudo yum install golang -y
    ```

3) Verify Go was installed properly

    ```bash
    go version

    # You should see something similar to:
    # go version go1.15.8 linux/amd64
    ```

---

## Clone this repository

Clone this repository to your local machine.

```bash
git clone git@gitlab.com:joetito1/vault-spring-sample.git
```

---

## Configure your environment to connect to Vault

Set the following environment variables to configure which Vault server this application connects to. These variables tell your app where your instance of Vault is running, which token to use for authentication and which namespace to use.

```bash
export VAULT_ADDR='http://127.0.0.1:8200'
export VAULT_TOKEN='[TOKEN]'
export VAULT_NAMESPACE=''
```

---

## Create a test secrets engine

Run the following to create a test secrets engine we can read from and write to:

```
vault secrets enable --path=test kv-v2
```

---

## Run the Go app

Build and install the Go program

```bash
go install example.com/vault/demo
```

Execute the binary generated above

```bash
~/go/bin/demo
```

If the app ran successfull, you should see the following output:

```bash
Using Vault: http://127.0.0.1:8200
Writing secret: test/data/demo
Reading secret: test/data/demo
item1: secrets are cool
item2: adp rocks
item3: golang is neat
Writing secret: test/data/demo
Reading secret: test/data/demo
item1: protect your secrets
item2: adp is amazing
item3: golang is fun
```

To validate the KV was written successfully, you can also:

- Login to the Vault UI: http://127.0.0.1:8200/ui
- Verify via the command line `vault kv get test/demo`