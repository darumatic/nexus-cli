# Nexus Repository Cleaner


The nexus repository cleaner performs the following tasks:
* scans all pods in the K8 cluster and gathers the deployed images
* Deletes tags in the configured nexus registry, if there are more than in KEEP specified tags and these tags are not deployed.

 
## Configuration

The cleaner is deployed into the sdlc namespace and requires the following:

* nexus-cleaner secret in the sdlc namespace 
* nexus-cleaner service account
* list-pods cluster role
* nexus-cleaner role binding


## Installation Steps
1. Setup kubectl to point at your k8 cluster
1. Edit 04-nexus-cleaner-cronjob.yml.
   1. Adjust the KEEP_LIMIT variable and the cron job expression
   1. Change schedule to match cron job expression
1. copy .sample_credentials to .credentials
1. Replace variables in .credentials with the correct values for your nexus repository
1. Run create_secret.sh   
1. Run kubectl create -f ./k8/

