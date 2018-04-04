# codecommit-package-server

A Golang Package Server for CodeCommit Repositories. Say, if you've got myrepo on us-east-1, you can do this:

```
$ go get codecommit.ingenieux.io/repo/myrepo
```

Or if you've got in a region other than us-east-1, you could do this:

```
$ go get codecommit.ingenieux.io/otherregion/repo/myrepo
```

And thats it! Keep reading if you want to build your own package server.

## Installation

requires golang, [gb](https://github.com/constabulary/gb), git, [upx](http://www.upx.org/), [severless](https://serverless.com), AWS and strip

Steps:

  * fork / checkout
  * Edit src/handlers/repo.go to replace ```codecommit.ingenieux.io``` to your repo (TODO: Fix this)
  * ```bash build.sh && serverless deploy```
  * [on AWS, create path mapping to a domain name and route 53 setup](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-edge-optimized-custom-domain-name.html)

## Usage

Replace codecommit.ingenieux.io with your package server host. In the example below, myrepo is a :

```
$ go get -d -v codecommit.ingenieux.io/repo/myrepo
```

# What about other AWS Regions?

Err, nice question. Currently this public server defaults on us-east-1, but you prefix the region on the URL instead:

```
$ go get codecommit.ingenieux.io/us-west-2/repo/myrepo
```

# What about SSH?

You can add ```?protocol=ssh``` into your URLS. 

Be wary that you must also change that in all your source code. 

When running a custom install, simply set "defaultProto" to "ssh" and it should work.
