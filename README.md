# gcp-api-gateway-tests

This was a project to understand how api gateway works and the
authentication as well, so any request that has been done with a
jwt authenticate by gcp can comunicate with others services like
cloud run or functions apis, using github actions as pipeline
with terraform to upload any new version of apis authentication,
gateway or jwt generator.

# Important Annotations

For each modification on config file needed we need to update api gateway with the new file version.

```
gcloud api-gateway gateways update my-gateway \
  --api-config=NEW_CONFIG_ID --api=API_ID --location=GCP_REGION --project=PROJECT_ID
```

On creating http load balancer in backend configuration → advanced seetings → Custom request headers: header name: host - header value: /

# Commands

List of commands for api gateway management

# Creating Api

```
gcloud api-gateway apis create API_ID --project=PROJECT_ID
```
On successful completion
```
gcloud api-gateway apis describe API_ID --project=PROJECT_ID
```

# Create APi Gateway Config
```
gcloud api-gateway api-configs createCONFIG_ID \
  --api=API_ID --openapi-spec=API_DEFINITION \
  --project=PROJECT_ID --backend-auth-service-account=SERVICE_ACCOUNT_EMAIL
```

# Create gateway

```
gcloud api-gateway gateways create GATEWAY_ID \
  --api=API_ID --api-config=CONFIG_ID \
  --location=GCP_REGION --project=PROJECT_ID
```

# Problems found

Problem to setup load balancer with api gateway:
[stackoverlow](https://stackoverflow.com/questions/65877067/gcp-load-balancing-with-api-gateway-returning-404)