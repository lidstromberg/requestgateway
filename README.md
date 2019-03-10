# Request Gateway - address restricting add-on for [auth].

A Go IP address restriction backend utility which uses Google Cloud Platform Datastore. Intended to be used in conjunction with [auth].

## What?
This is a fairly rudimentary backend client which persists a list of approved address names (IPs, hostnames, whatever you want to use to differentiate), and will provide a boolean check to indicate if a given address name is on the approved list.

## Why?
This was written to augment a [GCP] Appengine service. Appengine already provides a firewall service which can be used to control incoming traffic, so this address restriction utility is useful where a appengine default service already exists serving a superset of incoming traffic. This can be plugged into middleware to permit access to a non-default service for a subset of traffic.

## How?
The best place to start is probably with the examples and tests. Initialise the approval list entity in Datastore by running the tests.

## Examples
See [examples] for a http/appengine implementations which uses requestgateway and auth. This is written for appengine standard 2nd gen, but also works as a standalone.

## Dependencies and services
This utilises the following fine pieces of work:
* [GCP]'s [Datastore Go client] and [Storage Go client]

## Installation
If you want to run the example code, then install using
```sh
$ go get -u github.com/lidstromberg/examples
```
If you only want the requestgateway utility, then install with
```sh
$ go get -u github.com/lidstromberg/requestgateway
```
#### Environment Variables
You will also need to export (linux/macOS) or create (Windows) some environment variables.
```sh
################################
# GCP DETAILS
################################
export GTWAY_GCP_PROJECT='{{PROJECTNAME}}'

################################
# GCP CREDENTIALS
################################
export GOOGLE_APPLICATION_CREDENTIALS="/PATH/TO/GCPCREDENTIALS.JSON"
```
(See [Google Application Credentials])

Change LB_DEBUGON to true/false if you want verbose logging on/off. The other variables don't need to be changed.

```sh
################################
# REQUEST GATEWAY
################################
export LB_DEBUGON='true'
export GTWAY_NAMESP='global'
export GTWAY_KD='gateway'
export GTWAY_CLIPOOL='5'
```

   [auth]: <https://github.com/lidstromberg/auth>
   [GCP]: <https://cloud.google.com/>
   [Datastore Go client]: <https://cloud.google.com/datastore/docs/reference/libraries#client-libraries-install-go>
   [Storage Go client]: <https://cloud.google.com/storage/docs/reference/libraries#client-libraries-install-go>
   [Google Application Credentials]: <https://cloud.google.com/docs/authentication/production#auth-cloud-implicit-go>
   [examples]: <https://github.com/lidstromberg/examples>