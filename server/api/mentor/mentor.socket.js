/**
 * Broadcast updates to client when the model changes
 */

'use strict';

var Mentor = require('./mentor.model');

exports.register = function(socket) {
  Mentor.schema.post('save', function (doc) {
    onSave(socket, doc);
  });
  Mentor.schema.post('remove', function (doc) {
    onRemove(socket, doc);
  });
}

function onSave(socket, doc, cb) {
  socket.emit('mentor:save', doc);
}

function onRemove(socket, doc, cb) {
  socket.emit('mentor:remove', doc);
}