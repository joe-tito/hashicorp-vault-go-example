# Go Example

This repository is meant to demonstrate some of the basic functionality of HashiCorp Vault. Specifically, it demonstrates storing and accessing secrets stored in HashiCorp's KV2 (key/value pair) engine. 

This repository demonstrates the following Vault functionality:

- How to connect a Java application to a running instance of Vault
- How to write a key/value secret to Vault
- How to read a key/value secret from Vault

---

## Go Installation

If you do not have Go installed on your machine already, please follow the instructions below to install them.

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

## Add SpringVault to an existing project

To install the SpringVault into an existing project, add one of the following to your project:

- Maven
    ```bash
    <dependency>
        <groupId>org.springframework.vault</groupId>
        <artifactId>spring-vault-core</artifactId>
        <version>${version}</version>
    </dependency>
    ```

To view additional documentation / how-to's for VaultSpring, check out their github repository located here: https://github.com/spring-projects/spring-vault

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

You can now navigate to http://localhost:8200 to launch the Vault Spring demo.

If the web app ran successfull, you should see the following output in your browser:

```bash
Vault is AWESOME! :)
```

Then in your terminal, you should see the following output:

```bash
Using Vault: http://localhost:8200
Writing secret: test/data/demo
Writing data:
item1: adp is awesome
item2: spring is awesome
item3: protect your secrets
Reading secret: test/data/demo
item1: adp is awesome
item2: spring is awesome
item3: protect your secrets
Writing secret: test/data/demo
Writing data:
item1: adp is awesome
item2: updated secrets are neat
item3: this is neat
Reading secret: test/data/demo
item1: adp is awesome
item2: updated secrets are neat
item3: this is neat
```

To validate the KV was written successfully, you can also:

- Login to the Vault UI: http://127.0.0.1:8200/ui
- Verify via the command line `vault kv get test/demo`