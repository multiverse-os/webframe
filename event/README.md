# event subpackage
Need to divide the generalized event logic from the fs watch code, and have the
fs watch code juse the event code. 

fs watching will be necessary for making a halfway decent development server; we
need to watch for file changes and restart the server automatically for one
reason

and good generalized event code is critical for foundation for jobs but also
other event driven stuff that willb e important for sse and ws and such
