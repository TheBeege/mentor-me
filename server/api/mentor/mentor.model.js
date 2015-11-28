'use strict';

var mongoose = require('mongoose'),
    Schema = mongoose.Schema;

var MentorSchema = new Schema({
  username: {
    type: String,
    index: {
      unique: true,
      dropDups: true,
    }
  },
  fullName: String,
  pic: String,
  thumbnail: String,
  info: String,
  active: Boolean,
  tags: { type: [String], index: true }
});

module.exports = mongoose.model('Mentor', MentorSchema);
