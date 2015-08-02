# mentor-me
MentorMe is a mentor matchmaking app for kik. This is a community project, primarily worked on by the #programming and #javascript chats in kik.

Built on the MEAN stack - Mongo, Angular, Express, Node. See http://mean.io/. We _might_ sub out Mongo/Mongoose for PostgreS/Bookshelf. Thoughts appreciated.

This project was scaffolded with the yeoman generator angular-fullstack: https://github.com/DaftMonk/generator-angular-fullstack. The juicy stuff is in client/ and server/.

## Participants
* Github: @TheBeege kik: TehBeege
* Add yourself here whenever you submit a pull request!

## Getting started
1. Download and install Node.js: https://nodejs.org/. This automatically installs NPM, Node package manager. You can use this to install additional libraries and utilities.
2. Fork this repo, then clone your fork. Aren't familiar with Git and Github? See https://try.github.io/
3. In your command line of choice, navigate to the project folder and simply run `npm install`. NPM will look at packages.json to find dependencies and install them.
4. Run `grunt serve`. This starts the server and loads the page in your default browser. Updates you make to code will be automatically reloaded.
5. Open the project folder in your editor of choice. I recommend Atom https://atom.io/, Sublime http://www.sublimetext.com/3, or Notepad++ https://notepad-plus-plus.org/.
6. Start coding! The client folder holds all the views that will go to the user. The server folder holds all of the logic for the REST API. The client controllers are designed to pull all data from the REST API. Basically, this is all MVC:  https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller. Then read up on Angular: https://angularjs.org/. I also strongly recommend http://nodeschool.io/.

## Features!
Here's a swag at what we'd like to see. Feel free to add to this. We can discuss details in kik or create issues on here for comment-driven discussion.
* Users can submit themselves as mentors with tags
* Users can search for mentors based on tags

## Under the Hood Changes
* Switch to Bookshelf/PostgreS over Mongoose/Mongo
