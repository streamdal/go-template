# go-template

This is a Golang microservice template.

It uses:

1. `httprouter` for the router
1. `logrus` for logging
1. `envconfig` for env parser
1. `kingpin` for CLI args

----

To use this template:

1. Use Github template to generate new project
1. Replace the strings `go-template` and `GO_TEMPLATE` with preferred `service-name` and `SERVICE_NAME`
    1. `find . -type f -exec sed -i "" 's/go-template/writer/g' {} \;`
    1. `find . -type f -exec sed -i "" 's/GO_TEMPLATE/SERVICE_NAME/g' {} \;`
