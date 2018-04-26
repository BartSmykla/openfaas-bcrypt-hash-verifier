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
You also need to have an access to installed [OpenFaaS](https://github.com/openfaas/faas#hands-on-labs-detailed-getting-started) (You can do it easly using [Kubernetes](https://github.com/kubernetes/kubernetes#to-start-using-kubernetes) or [Docker Swarm](https://docs.docker.com/engine/swarm/)), know url to [OpenFaaS Gateway](https://docs.docker.com/engine/swarm/) and have an access to [docker registry](https://docs.docker.com/registry/) where you could push image with function which will be build using **faas-cli** (the easiest way is to create an account on [Docker Hub](https://hub.docker.com/))


### Example instalation

You can build, push and deploy your function on many ways, but the easiest (I think) way I'll try to describe below:

1. Cloning repository
    ```
    $ git clone https://github.com/tranotheron/openfaas-bcrypt-hash-verifier
    ```

2. Going into cloned directory
    ```
    $ cd openfaas-bcrypt-hash-verifier
    ```

3. [stack.yml](stack.yml) file adjustments

    You need to open file [stack.yml](stack.yml) and edit gateway url to match your gateway (for example: `http://127.0.0.1:8080`) and set image name to matching yours docker repository (for example: `NICKNAME/openfaas-bcrypt-hash-verifier`)

4. **CONDITIONAL** If your gateway is protected using basic auth
    ```
    $ faas-cli login --username=LOGIN --password=PASSWORD
    ```

5. Building function image
    ```
    $ faas-cli build -f stack.yml
    ```

6. Pushing built image to docker registry
    ```
    $ faas-cli push -f stack.yml
    ```

7. Deploying function to **OpenFaaS**
    ```
    $ faas-cli deploy -f stack.yml
    ```

### Code of Conduct
You can find Code of Conduct [here](CODE_OF_CONDUCT.md)

### License
This project is under the [MIT License](https://github.com/tranotheron/openfaas-bcrypt-hash-verifier/blob/master/LICENSE).