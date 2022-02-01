## Word Counter API 
___________

Server is running on `8080` port.

#### Endpoints
________

`/api/v1/wordcount` - Counts the number of words in the text.

CURL 
```
curl -X 'POST' \
  'http://localhost:8080/api/v1/wordcount' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "text": "This is just a screenshot; the actual documentation is clickable and expandable, providing a clear description of request parameters, responses and their JSON schemas, etc.There'\''s more, though. Assuming this API is hosted in a publicly-accessible server, we can interact with it straight from the Swagger Editor - just like hand-crafting curl commands, but in an auto-generated schema-aware way.Imagine you'\''ve developed this task API and now want to publish it for use in various clients (web, mobile, etc.); if your API is specified with OpenAPI/Swagger, you get automatic documentation and an interface for clients to experiment with the API. This is doubly important when you have API consumers who aren'\''t SW engineers - for example, this could be UX designers, technical documentation writers and product managers who need to understand an API but may be less comfortable throwing scripts together. Moreover, OpenAPI standardizes things like authorization, which could also be very useful when compared with an ad-hoc description.There are additional tools available once you have a spec - e.g. Swagger UI and Swagger Inspector. You can even use the spec to help integrate your REST server into your cloud provider'\''s infrastructure; for example, GCP has Cloud Endpoints for OpenAPI for setting up monitoring, analysis and other features for published APIs; the API is described to the tool using OpenAPI."
}'
```

`/api/v1/wordcount/file` - Counts the number of words in the text file.

CURL 
```
curl -X 'POST' \
  'http://localhost:8080/api/v1/wordcount/file' \
  -H 'accept: application/json' \
  -H 'Content-Type: multipart/form-data' \
  -F 'myFile=@sample1.txt;type=text/plain'
```


## Swagger Endpoint

`http://localhost:8080/api/v1/swagger/index.html`# wordfrequencycounter
