ANSIBLE
========
Ansible is an automation tool for managing IT infrastructure. In DevOps, we use Jenkins to handle testing and building our projects while Ansible is used to configure the servers they will be deployed on and do the final deploy stages such as pulling from a registry and running installs.

This begs the question, if I have Kubernetes, why do I need Ansible? [Answer](https://www.ansible.com/blog/how-useful-is-ansible-in-a-cloud-native-kubernetes-environment)

You may not always use Ansible with K8s (especially if you are using a service like EKS on AWS). Yu are more likely to use Ansible with K8s if you are setting up an on-premise cluster rather than a cloud based one.