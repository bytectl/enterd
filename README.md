# enterd

enterd  is an ent extension that create
 [Mermaid-Js](https://mermaid-js.github.io/mermaid/#/entityRelationshipDiagram) (Entity Relationship Diagrams) ERD diagrams of your ent's schema.

inspired by [mermerd](https://github.com/KarnerTh/mermerd)

inspired by [entviz](https://github.com/hedwigz/entviz)

## install

```go
go get github.com/bytectl/enterd
```

add this extension to ent (see [example](examples/ent/entc.go) code)
run

```go
go generate ./ent
```

your html will be saved at `ent/schema-erd.mmd`

## Use from command line

Install the cmd

```go
go get  github.com/bytectl/enterd/cmd/enterd
```

Then run inside your project:

```bash
enterd ./etc/schema
```

## example

![image (3)](https://user-images.githubusercontent.com/8277210/129726965-d3c89f1a-d66a-46b6-82a2-20f1056d350d.png)
