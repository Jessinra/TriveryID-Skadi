# UDCT todo list

## Step 1: Propose and Scope the Project
- [x] Plan what your pipeline will look like.
- [x] Decide which options you will include in your Continuous Integration phase. Use either Circle CI or Jenkins.
- [x] Pick a deployment type - either rolling deployment or blue/green deployment.
- [x] For the Docker application you can either use an application which you come up with, or use an open-source application pulled from the Internet, or if you have no idea, you can use an Nginx “Hello World, my name is (student name)” application.

## Step 2: Use Jenkins or Circle CI, and implement blue/green or rolling deployment.
- [x] ~If you're using Jenkins, create your Jenkins master box and install the plugins you will need.~
- [x] If you're using Circle CI, set up your circle CI account and connect your git repository.
- [x] Set up your environment to which you will deploy code.

## Step 3: Pick AWS Kubernetes as a Service, or build your own Kubernetes cluster.
- [ ] Use Ansible or CloudFormation to build your “infrastructure”; i.e., the Kubernetes Cluster.
- [ ] It should create the EC2 instances (if you are building your own), set the correct networking settings, and deploy software to these instances.
- [ ] As a final step, the Kubernetes cluster will need to be initialized. The Kubernetes cluster initialization can either be done by hand, or with Ansible/Cloudformation at the student’s discretion.

## Step 4: Build your pipeline
- [x] Construct your pipeline in your GitHub repository.
- [x] Set up all the steps that your pipeline will include.
- [x] Configure a deployment pipeline.
- [x] Include your Dockerfile/source code in the Git repository.
- [x] Include with your Linting step both a failed Linting screenshot and a successful Linting screenshot to show the Linter working properly.

## Step 5: Test your pipeline
- [x] Perform builds on your pipeline.
- [x] Verify that your pipeline works as you designed it.
- [x] Take a screenshot of the Circle CI or Jenkins pipeline showing deployment, and a screenshot of your AWS EC2 page showing the newly created (for blue/green) or modified (for rolling) instances. Make sure you name your instances differently between blue and green deployments.