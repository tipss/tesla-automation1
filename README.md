# tesla-automation1
# Author hongal@gmail.com  Date 08/5/2021
This repository is place holder of various files(e.g .go, Docker, script, design) etc
for tesla-automation1 (demo) project.

What is solved using this project?.
1.  Identify a unique micro-service.(e.g currency-exchange, timezone calculator, configure-a-virtual-router)
2.  Write cloud aware code to enable micro-service.(e.g golang)
3.  Package code as container(e.g Docker)
4.  Orchestrate this service in Public Cloud using tools(e.g Kubernetes).
4.  Host entire solution in Public Cloud(e.g Google Cloud)
5.  Enable CI/CD on to this code base, so its open for further development, while this demo is running. 



micro-service1
1. timezone calculator
   This is a web interface, when user wants to know time in a specifc
   timezone, he will entry via web-url.
   In return it computes time and the destination timezone and sends the
   output back.

   This is a micro task which is independant.
   This may get used/called many 100 times per second if it becomes popular service.
   So its best fit for to host on a scalable clould.
   Its a service user can access 24/7, so hosting in clould is best option.
 
1.1 Timezone calculator
     curl -s -X PUT http://34.125.164.201/time

2. Device Config Queries and Configurator

   FRRouting has netconf, restconf, or grpc and vtysh based config inputs.
   URL based approach is good for short and quick way to configure, which does not
   need elaborate json parsing skills.
   I have used URL parsing  method.

   In an org there can be many devices which may run 24/7.
   Accessing these device to 'read' and configure is traditionally via CLI, where
   in a single user can access at a time. This can be a bit complicated if we have large
   number of devices.

   Say we have URL based querier and config creater, it will be easy for
   new user to read config on large number of devices , compare and do health check.
   Same applies for routine configuration.


   This microservice via web URL reads config on a given device( we have choosen an FRR BGP Router)
   These devices are hosted on a Google Cloud via K8 cluster PODs as Deployments.
   This micro service is enabled on each node, so user can query from any web URL and any device.


  Typical flow of change expected:
  ================================
  Developer ->  GIT Hub --> Cloud Build ----> Cloud Function(K8 upgrade version)


  How to use these micro service ?.
  1. Locally for testing.
     curl -s -X PUT http://localhost/frr\?cmd\="-c+show+running"     

  2. Remotely via Web 
     curl -s -X PUT http://34.125.225.186/frr?cmd="-c+show+running"





