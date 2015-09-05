'use strict';

var _ = require('lodash');
var Mentor = require('./mentor.model');

// Find mentor by tag
exports.tag = function(req, res) {
  // If there are no tag values, give back
  // ALL THE MENTORS!
  if (req.query.tags.length == 0) {
    console.log("no tags");
    return exports.index(req,res);
  }

  // GET query parameters will appear as a raw value
  // if there's only one of them, instead of as an array
  if (typeof req.query.tags == "string") {
    console.log("single string. Converting to array");
    req.query.tags = [req.query.tags];
  }

  console.log("tags after cleanup");
  console.log(req.query.tags);
  Mentor.find({tags: {$in: req.query.tags}}, function (err, mentors) {
    if(err) {
      console.log("Error finding mentors:");
      console.log("err");
      return handleError(res, err);
    }
    console.log("mentors:");
    console.log(mentors);
    return res.json(200, mentors);
  });
};

// Get list of mentors
exports.index = function(req, res) {
  Mentor.find(function (err, mentors) {
    if(err) { return handleError(res, err); }
    return res.json(200, mentors);
  });
};

// Get a single mentor
exports.show = function(req, res) {
  Mentor.findById(req.params.id, function (err, mentor) {
    if(err) { return handleError(res, err); }
    if(!mentor) { return res.send(404); }
    return res.json(mentor);
  });
};

// Creates a new mentor in the DB.
exports.create = function(req, res) {
  if (!req.body.username) {
    // Malicious attempt. They were never logged in
  } else {
    req.body.tags.forEach(function(part, index, tagArray) {
      tagArray[index] = part.trim();
    });
    req.body.pic = req.body.pic.substring(2);
    req.body.thumbnail = req.body.thumbnail.substring(2);
    Mentor.create(req.body, function(err, mentor) {
      if(err) { return handleError(res, err); }
      return res.json(201, mentor);
    });
  }
};

// Updates an existing mentor in the DB.
exports.update = function(req, res) {
  if(req.body._id) { delete req.body._id; }
  Mentor.findById(req.params.id, function (err, mentor) {
    if (err) { return handleError(res, err); }
    if(!mentor) { return res.send(404); }
    var updated = _.merge(mentor, req.body);
    updated.save(function (err) {
      if (err) { return handleError(res, err); }
      return res.json(200, mentor);
    });
  });
};

// Deletes a mentor from the DB.
exports.destroy = function(req, res) {
  Mentor.findById(req.params.id, function (err, mentor) {
    if(err) { return handleError(res, err); }
    if(!mentor) { return res.send(404); }
    mentor.remove(function(err) {
      if(err) { return handleError(res, err); }
      return res.send(204);
    });
  });
};

function handleError(res, err) {
  return res.send(500, err);
}
