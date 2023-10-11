# Notes For Maglev Server
Brainstorming, research, idea, and notes regarding the construction of
`maglev` server and its facilities. 

  * It would be incredibly prudent to have the ability to track changes to files
    during development and restart the server automagically. This is done
    through inode monitoring in linux and so can be done pretty easily. And if
    you see how often this is accomplished via plugins for rails developers you
    will see how much they value it. 
