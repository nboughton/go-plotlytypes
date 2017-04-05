# go-plotlytypes
Types definitions for graph structs to be passed to plotly.js

Normally in Go we just put type definitions where they're used. Normally I'd provide functions to 
supplement this code in order to create objects etc. In this particular instance I'm not sure that
would be of much use as the types defined here are usually instantiated with a fair amount of 
customised details. 

So I'm just going to have the types declared so I don't have to boilerplate them every time I want
to make plotly data. Which I do a lot. I may extend this at a later date to provide "new" functions
for common graph types but at the moment I feel that would be inflexible.
