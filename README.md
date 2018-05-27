## openfaas-bcrypt-hash-verifier
[![Go Report Card](https://goreportcard.com/badge/github.com/tranotheron/openfaas-bcrypt-hash-verifier)](https://goreportcard.com/report/github.com/tranotheron/openfaas-bcrypt-hash-verifier)
[![CircleCI](https://circleci.com/gh/tranotheron/openfaas-bcrypt-hash-verifier.svg?style=shield)](https://circleci.com/gh/tranotheron/openfaas-bcrypt-hash-verifier)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![OpenFaaS](https://img.shields.io/badge/openfaas-serverless-blue.svg)](https://www.openfaas.com)

### Prerequisites
To build and deploy `openfaas-bcrypt-hash-verifier` function you need to have installed:
- [git](https://git-scm.com/downloads)
- [go](https://golang.org/doc/install)
- [docker](https://docs.docker.com/install/)
- [faas-cli](https://github.com/openfaas/faas-cli#get-started-install-the-cli)
You also need to have an access to installed [OpenFaaS](https://github.com/openfaas/faas#hands-on-labs-detailed-getting-started) (You can do it easly using [Kubernetes](https://github.com/kubernetes/kubernetes#to-start-using-kubernetes) or [Docker Swarm](https://docs.docker.com/engine/swarm/)) know url to [OpenFaaS Gateway](https://github.com/openfaas/faas/tree/master/gateway)
- If you want to build and push a function to your own [docker registry](https://docs.docker.com/registry/) (for example [Docker Hub](https://hub.docker.com/)) you need to be logged in your account ([`docker login`](https://docs.docker.com/engine/reference/commandline/login/))

### Example of deploying function using **faas-cli** using default values from [this repository](https://raw.githubusercontent.com/tranotheron/openfaas-bcrypt-hash-verifier/master/stack.yml)
```bash
$ faas-cli deploy -f https://raw.githubusercontent.com/tranotheron/openfaas-bcrypt-hash-verifier/master/stack.yml
```

### Example of building and using function from own docker registry

You can build, push and deploy your function on many ways, but the easiest (I think) I'll try to describe below:

1. Cloning repository
    ```bash
    $ git clone https://github.com/tranotheron/openfaas-bcrypt-hash-verifier
    ```

2. Going into cloned directory
    ```bash
    $ cd openfaas-bcrypt-hash-verifier
    ```

3. [stack.yml](stack.yml) file adjustments

    You need to open file [stack.yml](stack.yml) and edit gateway url to match your gateway (for example: `http://127.0.0.1:8080`) and set image name to matching your docker repository (for example: `NICKNAME/openfaas-bcrypt-hash-verifier`)

4. Building function image
    ```bash
    $ faas-cli build -f stack.yml
    ```

5. Pushing built image to docker registry
    ```bash
    $ faas-cli push -f stack.yml
    ```

6. Deploying function to **OpenFaaS**
    ```bash
    $ faas-cli deploy -f stack.yml
    ```

### Usage with [curl](https://curl.haxx.se/)
```bash
$ curl -d '{"hash":"$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K","password": "foo"}' -X POST http://localhost:8080/function/bcrypt
{"match":true}

$ curl -d '{"hash":"$2a$12$Y/98WmHkm3k38/suzvvEUuJ.QVA3oUeks74uTDDGt6JGhTqL/RP0K","password": "bar"}' -X POST http://localhost:8080/function/bcrypt
{"match":false,"error":"crypto/bcrypt: hashedPassword is not the hash of the given password"}
```

### Possible errors
- **crypto/bcrypt: hashedPassword is not the hash of the given password**

    Hash doesn't match passed password

- **crypto/bcrypt: bcrypt hashes must start with '$', but hashedSecret started with ' '**

    In that case there is space at the beginning of a hash

- **crypto/bcrypt: hashedSecret too short to be a bcrypted password**

    Passed hash is incorrect, check if you passed all hash and there is no missing characters at the end of it

- **you didn't pass password**

    You probably forgot to put password in the body of your request:
    ```json
    { ..., "password": "your_password" }
    ```

- **you didn't pass hash**

    You probably forgot to put hash in the body of your request
    ```json
    { ..., "hash": "your_hash" }
    ```
    It can also mean that you didn't pass hash **and** password

- **unexpected end of JSON input**

    You probably forgot to put json with hash and password is body of your request
    ```json
    {
        "password": "your_password",
        "hash": "your_hash"
    }
    ```

- **illegal base64 data at input byte 4** (or other number)

    Passed hash is incorrect. Check if in your hash there is no additional character like space somewhere in the middle of it

### Tests
```bash
$ git clone https://github.com/tranotheron/openfaas-bcrypt-hash-verifier
$ go test -v ./openfaas-bcrypt-hash-verifier/function
```

### Code of Conduct
You can find Code of Conduct [here](CODE_OF_CONDUCT.md)

### License
This project is under the [MIT License](https://github.com/tranotheron/openfaas-bcrypt-hash-verifier/blob/master/LICENSE).
