'use strict';

var mongoose = require('mongoose'),
    Schema = mongoose.Schema;

var MentorSchema = new Schema({
  name: String,
  info: String,
  active: Boolean,
  image: String,
  tags: { type: [String], index: true }
});

module.exports = mongoose.model('Mentor', MentorSchema);
