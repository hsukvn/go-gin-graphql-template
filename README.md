# go-graphql-template

Basic file hierarchy of graphql API server in golang

## Example

### Query

curl -g 'http://localhost:9527/graphql?query={songs(album:"ts-fearless"){title,duration,id}}'

### Mutation

curl -g 'http://localhost:9527/graphql?query=mutation+_{createSong(id:"7",album:"ts-fearless",title:"Breathe",duration:"4:23"){title,duration,id}}'
