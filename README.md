## go-cli

## Contents ##
- [Brief](#brief)
- [Components](#comps)
  - [hes-client](#hesclient)
  - [pw-client](#pwclient)
  - [data](#data)
- [Gotchas](#gotchas)
- [References](#references)

### [Brief](#brief){#brief}

CLI clients for accessing prosperworks and hes via public API - GET data only for now!

Build with ./autobuild.sh

### [Components](#comps){#comps} 

#### [hes-client](#hes)

Run with:
`
./dist/linux/amd64/hes-client -login -debug
`
#### [pw-client](#pw)

Run with:
`
/dist/linux/amd64/pw-client
 
Cmd: leads | people | projects
--OR--
Cmd: lead <id> | people <id> | project <id>
`
#### [data](#data)

HES/PW structs for (un)marshaling data from API calls.  

Redis to store 

### Gotchas ###

+ 

### References ### 

+ https://developer.prosperworks.com/
+ https://docs.gurusys.co.uk/static/public-api-docs/

