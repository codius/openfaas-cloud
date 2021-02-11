OpenFaaS Cloud
==============

The Multi-user OpenFaaS Platform

## Introduction

[![Build Status](https://github.com/openfaas/openfaas-cloud/workflows/build/badge.svg?branch=master)](https://github.com/openfaas/openfaas-cloud/actions)

Features:

* Portable - self-host on any cloud
* Multi-user
* Sub-domain per user or organization with HTTPS
* Runtime-logs for your functions

The dashboard page for a user:

![Dashboard](/docs/dashboard.png)

The details page for a function:

![Details page](/docs/details.png)

## Overview

### KubeCon video

[![](http://img.youtube.com/vi/sD7hCwq3Gw0/maxresdefault.jpg)](https://www.youtube.com/watch?v=sD7hCwq3Gw0)

[KubeCon: OpenFaaS Cloud + Linkerd: A Secure, Multi-Tenant Serverless Platform - Charles Pretzer & Alex Ellis](https://www.youtube.com/watch?v=sD7hCwq3Gw0&feature=emb_title)

### Blog posts

* [Build your own OpenFaaS Cloud with AWS EKS](https://www.openfaas.com/blog/eks-openfaas-cloud-build-guide/)
* [Introducing OpenFaaS Cloud](https://blog.alexellis.io/introducing-openfaas-cloud/)
* [Sailing through the Serverless Ocean with Spotinst & OpenFaaS Cloud](https://spotinst.com/blog/sailing-through-the-serverless-ocean-with-openfaas-cloud/)

### Documentation

* [Conceptual architecture](https://docs.openfaas.com/openfaas-cloud/architecture).
* [Authentication](https://docs.openfaas.com/openfaas-cloud/authentication/)
* [Multi-stage environments](https://docs.openfaas.com/openfaas-cloud/multi-stage/)
* [Manage secrets](https://docs.openfaas.com/openfaas-cloud/secrets/)
* [User guide](https://docs.openfaas.com/openfaas-cloud/user-guide/)

### Roadmap & Features

See the [Roadmap & Features](docs/ROADMAP.md)

## Get started

You can set up and host your own *OpenFaaS Cloud* or pay an expert to do that for you. OpenFaaS Ltd also offers custom development, if you should have new requirements.

### Option 1: Expert installation

OpenFaaS Ltd provides expert installation and support for OpenFaaS Cloud. You can bring your own infrastructure, or we can install and configure OpenFaaS Cloud for your accounts on a managed cloud.

[Get started today](https://www.openfaas.com/support/)

### Option 2: Automated deployment (self-hosted)

You can set up your own OpenFaaS Cloud with authentication and wildcard certificates using ofc-bootstrap in around 100 seconds using the ofc-bootstrap tool.

This method assumes that you are using Kubernetes, have a public IP available or [are using the inlets-operator](https://github.com/inlets/inlets-operator), and have a domain name. Some basic knowledge of how to setup a GitHub App and GitHub OAuth App along with a DNS service account on DigitalOcean, Google Cloud DNS, Cloudflare or AWS Route53.

A [developer install is also available via this blog post](https://blog.alexellis.io/openfaas-cloud-for-development/), which disables OAuth and TLS. You will still need an IP address and domain name.

Deploy with: [ofc-bootstrap](https://github.com/openfaas-incubator/ofc-bootstrap)

## Getting help

For help join #openfaas-cloud on the [OpenFaaS Slack workspace](https://docs.openfaas.com/community). If you need commercial support, contact [sales@openfaas.com](mailto:sales@openfaas.com)

