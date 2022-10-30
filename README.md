# DevFest Kutaisi Worksop

This repository contains examples of what we will be doing in the workshop.

## How to use

There are tags numbered 01, 02 and so on. If you miss something during the workshop, just open the tag mentioned and see the code there.

## Requirements

To participate in the workshop you need to have:

- a laptop;
- code editor (vscode by default);
- docker installed;
- golang installed;
- git istalled;
- github account.

You will use a GCP service account provided to you.

## How to run the project locally

Run in terminal to start Redis database.

```bash
make up
```

Run in a second terminal to build and run the application within the container.

```bash
make start
```

To stop the database after you finish:

```bash
make down
```

## How to do it on your own (hands-on manual)

The workshop material is available for everyone, but it may be difficult to reproduce all of the steps without guidance. To help you to achieve that on your own I'll provide a manual.

### Create GitHub repository

Go to GitHub and create a new public repository. Clone it locally on your host.

### Setup a cloud landing zone

To deploy the application to the cloud environment you have to register a GCP account. Don't worry - Google gives credits for you to try out the platform, and this demo will barely spend some bucks.

#### Create a project

You need to create a project (this will be a default project when you create an account).

Copy project id. Go to your GitHub account, go to _Settings_ > _Secrets_ > _Actions_ and add a new secret by clicking "New repository secret". The name is `PROJECT_ID`` and the Secret is the value you've copied.

#### Create a service account

Go to GCP [IAM > Service Accounts](https://console.cloud.google.com/iam-admin/serviceaccounts) and create a service account. Then go to account and switch to the "keys" tab. Add a key in JSON format.

It will be downloaded as a file. Open the file and copy its contents. Go to your GitHub account, go to _Settings_ > _Secrets_ > _Actions_ and add a new secret by clicking "New repository secret". The name is `SERVICE_ACCOUNT_KEY`` and the Secret is the value you've copied.

Go to GCP [IAM](https://console.cloud.google.com/iam-admin/) and add the following roles to your service account:

- Artifact Registry Reader
- Artifact Registry Reader
- Cloud Run Admin
- Cloud Run Service Agent

#### Create a Serverless VPC access

Go to GCP [Serverless VPC access](https://console.cloud.google.com/networking/connectors) under VPC Network. Create a new connector. Chose the name you want and the region where you will deploy your applications (us-central-1 in my example). Choose subnet (10.8.0.0/28 for example) and scaling settings - a minimum of 2 instances and a maximum of 3 instances will be enough for the demo.

#### Create a firewall rule

Go to GCP [Firewall](https://console.cloud.google.com/networking/firewalls) under VPC Network. Create a new firewall rule. Set

- Target: All instances in the network;
- Source IPv4 ranges: CIDR from VPC Connector created on the previous step (10.8.0.0/28 in my example);
- Protocols and Ports: TCP 6379.

#### Create artifact repository

Go to GCP [Actifact Registry](https://console.cloud.google.com/artifacts) and create a new repository called `devfest-kutaisi` with the format "Docker" and region "us-central-1". Note that the repository name and region are set up in the [CI config file](/.github/workflows/gcp.yaml), so in case you want to pick a different name or region for your demo, change them in the file too.

This repository will be used to push and store your Docker images.

#### Create Memorystore instance

Go to GCP [Memorystore](https://console.cloud.google.com/memorystore) and create a Redis instance. Chose the configuration you want (I suggest picking the minimum configuration for the demo since a cloud database is costly). Don't forget to delete the instance after the demo.

Now go to created instance page > Connections and copy the Primary endpoint value (IP address and port). Go to your GitHub account, go to _Settings_ > _Secrets_ > _Actions_ and add a new secret by clicking "New repository secret". The name is `DB_ADDRESS`` and the Secret is the value you've copied.

### Develop the application

The main goal of the workshop is to learn how to develop a simple interactive cloud-native application in Golang, set up the CI pipeline and access the app's API in the cloud environment.

To do that we need first to build the app locally. To do that, run `make up` to start a local Redis instance in a docker container.

Run `make start` to build a docker container and start it.

You now can access the API via its URL on port `8888`. Give it a try:

```bash
# add a wine review
curl -X POST localhost:8888/wine -d '{"name": "Golden Grapes", "winery":"old wine", "vintage":2020, "review":"simple but pleasant"}'
# add one more
curl -X POST localhost:8888/wine -d '{"name": "My Saperavi", "winery":"good year", "vintage":2012, "review":"pretty good one"}'
# read your reviews
curl localhost:8888/wine
```

### Build and push a docker image to the registry

You can build and push an image to the registry locally, i.e. from your host. But we don't want to do that each time, so we automate it. Automatical check and build of application called Continuous Integration (CI).

For this step, we have job `build-push-gcp` in our pipeline. Commit changes and push them to GitHub, and the pipeline will start. It will build and deploy the image to the GCP Artifact Registry.

### Create a Cloud Run application

Go to GCP [Cloud Run](https://console.cloud.google.com/run) and create a new service. In "Deploy one revision from an existing container image" chose _Artifact Registry_ > _.../devfest-kutaisi_ > _ **your GitHub username**_ > latest.

Set autoscaling minimum and maximum to 1. Select "Allow unauthenticated invocations" under "Authentication".

Go to _Container, Connections, Security_ > _Container_ and add the following environment variables:

- API_ADDRESS - 0.0.0.0:8080
- DB_ADDRESS - **address of Memorystore instance**
- DB_COLLECTION - **your GitHub account name**

Go to _Container, Connections, Security_ > _Connections_ and choose `devfest-kutaisi` as a Network connector.

Create the service. When it will be ready, you can access it via generated URL. Check it.

### Automate application deployment

To avoid errors, it's always better to deploy applications automatically. It is called Continuous Deployment (CD).

For deployment, we have another job called `deploy-cloud-run`. It is disabled so far, and to enable it you have to go to _Settings_ > _Secrets_ > _Actions_ and add a new secret by clicking "New repository secret". The name is `DEPLOY_ENABLED` and the Secret is `true`.

Go to the app and change the greeting text (in `internal/api/server.go`) to something meaningful to you.
Commit changes and push them to GitHub. Check the pipeline.

Now your app will be built again but this time it will be deployed to the cloud automatically.
