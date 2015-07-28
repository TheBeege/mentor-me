'use strict';

angular.module('mentorMeApp')
  .controller('FindMentorCtrl', function ($scope, $http, socket) {
    $scope.message = 'Hello';

    $http.get('/api/things').success(function(mentors) {
      $scope.mentors = mentors;
      socket.syncUpdates('thing', $scope.mentors);
    });
  });
