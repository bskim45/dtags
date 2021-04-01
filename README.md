# dtags

![Github build status](https://github.com/bskim45/dtags/workflows/Test%20and%20Build/badge.svg?branch=master)
![GitHub latest release](https://img.shields.io/github/v/release/bskim45/dtags)
![License MIT](https://img.shields.io/github/license/bskim45/dtags)
[![Docker latest version](https://img.shields.io/docker/v/bskim45/dtags?sort=semver)](https://hub.docker.com/r/bskim45/dtags)

**dtags** is a small binary retrieves a list of Docker repositories and Docker Image tags from various Docker registries.

## Supported repositories

| Name | URL | Tag List | Search |
| ---- | --- | -------- | ------ |
| Docker Hub | https://index.docker.io | Y | Y |
| quay.io | https://quay.io | Y | Y |
| Google Container Registry | https://gcr.io | Y | N |
| Elastic Docker Registry | https://docker.elastic.co | Y | N |
| GitHub Container Registry | https://ghcr.io | Y | N |
| Amazon ECR Public Registry | https://public.ecr.aws | N | N |

.. and many more that supports Docker Registry API v2 (https://docs.docker.com/registry/spec/api/)

## Installation

### Binary files

You can download binary files from [release](https://github.com/bskim45/dtags/releases).

### Homebrew

```bash
brew install bskim45/dtags/dtags
```

### Install directly

```bash
# Installed to $HOME/.dtags/bin/dtags by default
$ curl -sfL https://raw.githubusercontent.com/bskim45/dtags/master/scripts/get.sh | sh -s
```

> for additional options, run `curl -sfL https://raw.githubusercontent.com/bskim45/dtags/master/scripts/get.sh | sh -s -- -h`

Now, add `$HOME/.dtags/bin` to your $PATH.

> Replace `~/.bash_profile` according to your favorite shell. (`~/.zshrc` or `~/.bashrc`)

```bash
echo 'export PATH="$HOME/.dtags/bin:$PATH"' >> ~/.bash_profile
```

### Go

> Please make sure **your `go/bin` is in your `$PATH`**. Mostly, Go bin path is `~/go/bin` (MacOS/Linux), or `%HOME%\go\bin`
(Windows). Also, you can find go binary path by `$(go env GOPATH)/bin`.

```bash
go get github.com/bskim45/dtags
```

## Usage

```bash
$ dtags help
```

For those who are big fan of docker (also available in quay.io/bskim45/dtags):
```bash
$ docker run -it --rm bskim45/dtags tags bskim45/dtags
latest
1.0.0
...
```

### Listing Tags

```bash
# official library images
# this is equivalent to 'dtags tags library/ubuntu'
$ dtags tags ubuntu
...
20.04
19.10
19.04
18.04
16.04
...

# normal repo, only last 5 tags (default 100)
$ dtags tags bskim45/helm-kubectl-jq -n 5
3.0.3
latest
3.0.2
3.0.1
3.0.0
...

# quay.io
$ dtags tags quay.io/bitnami/nginx -n 5
latest
1.16.1
1.16
1.16.1-debian-10-r13
1.16-debian-10

# gcr.io
$ dtags tags gcr.io/google_containers/busybox
latest
1.27.2
1.27
1.24

# elastic docker registry
$ dtags tags docker.elastic.co/elasticsearch/elasticsearch -n 5
master-SNAPSHOT
8.0.0-SNAPSHOT
7.x-SNAPSHOT
7.7.0-SNAPSHOT
7.6.0-SNAPSHOT
```

### Search repository

```bash
# search against docker hub
$ dtags search
python
circleci/python
...

# search against quay.io
$ dtags search nginx --endpoint quay.io
kubernetes-ingress-controller/nginx-ingress-controller
bitnami/nginx
...
```

## Contributing

This is my very first project written in Go.
Please feel free to suggest any improvements.


## License

This project are released under the [MIT License](https://github.com/bskim45/dtags/blob/master/LICENSE)
