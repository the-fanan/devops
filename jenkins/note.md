JENKINS
==========
This is an automation server that allows us to easily automate tasks related to continuous integration and delivery. 

In my opinion the bulk of this tool is about knowing the right plugins to use and how to use them.

You would typically use Jenkins to poll your repo for changes (or accept a webhook notification from GitHub) then perform a set of automated tasks like building and testing the project, building and pushing to Docker Registry. 

You could also use Jenkins to deploy to servers and manage them but this could be cumbersome. So you typically will push your build to a docker registry then call another tool like Ansible to handle the management of your servers and do the final deploys [This is my opinion at the moment and might change once I learn more about Ansible].



Parts of the GUI I have found useful

1. `Manage Jenkins` -> `Manage Plugins`: For installing and updating plugins
2. `Manage Jenkins` -> `Configure System`: For configuring global settings and plugins
3. `Manage Jenkins` -> `Configure Credentials`: For managing credentials that Jenkins can use in projects
4. `New Item`: To create a new Project
5. *After selecting a project from dashboard* -> `Configure`: Setting up build, polling and other settings for a project

Links I have found helpful:
1. [Handling Private Github Repos Using Personal Access Tokens](https://stackoverflow.com/questions/61105368/how-to-use-github-personal-access-token-in-jenkins)
2. [Handling Private Repos using SSH](https://shreyakupadhyay.medium.com/integrate-jenkins-with-github-private-repo-8fb335494f7e)
3. [Plugin For Environment Variables In Build](https://plugins.jenkins.io/text-file-operations/)
4. [Basic Laravel DevOps Setup with Jenkins](https://faun.pub/configure-laravel-8-for-ci-cd-with-jenkins-and-github-part-1-58b9be304292)

Sample Groovy script to build and deploy to private docker registry
```
pipeline {
    agent any
    stages {
        stage("Docker build") {
            steps {
                sh "docker build -t [REGISTRY_HOST]/[BUILD_NAME] ."
            }
        }
        stage("Docker push") {
            environment {
                DOCKER_USERNAME = credentials("docker-user")
                DOCKER_PASSWORD = credentials("docker-password")
            }
            steps {
                sh "docker login [REGISTRY_HOST] --username ${DOCKER_USERNAME} --password ${DOCKER_PASSWORD}"
                sh "docker push [REGISTRY_HOST]/[BUILD_NAME]"
            }
        }
    }
}   
```
5. [Pulling From Private Registries](https://www.baeldung.com/linux/docker-compose-private-repositories)
`docker pull [REGISTRY_HOST]/[BUILD_NAME]`
```
other-registry:
    image: [REGISTRY_HOST]/[REGISTRY_USER]/[BUILD_NAME]
```
`docker build -t [REGISTRY_HOST]/[BUILD_NAME] -f DockerFileCustom .`
Note how the user needs to be specified in the docker-compose but not in CLI (I am not sure if the user must be speicifed in docker-compose file. I will experiment and update this)
6. [Tell K8s to pull and deploy new images](https://stackoverflow.com/a/55914480/7715823)