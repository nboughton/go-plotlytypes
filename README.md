# go-plotlytypes
Types definitions for graph structs to be passed to plotly.js

Normally in Go we just put type definitions where they're used. Normally I'd provide functions to 
supplement this code in order to create objects etc. In this particular instance I'm not sure that
would be of much use as the types defined here are usually instantiated with a fair amount of 
customised details. 

So I'm just going to have the types declared so I don't have to boilerplate them every time I want
to make plotly data. Which I do a lot. I may extend this at a later date to provide "new" functions
for common graph types but at the moment I feel that would be inflexible.

You might also wonder why all the axes values are []string. I made a point about this in the comments
but it's worth repeating here:
   
   Why are all axes strings? Because it's more flexible. You're passing this data to
	 Plotly.js which interprets it accordingly and where strconv has no value it returns
	 an empty string which plotly ignores but maintains the correct number of elements
	 in the set. This removes useless visual noise from the graph and allows you to make
	 your axes anything. Yes, there's a performance overhead to the conversion but it is
	 on the scale of a fraction of a fraction of a second, which I see as a reasonable
	 compromise