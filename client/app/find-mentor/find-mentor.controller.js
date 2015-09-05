'use strict';

angular.module('mentorMeApp')
  .controller('FindMentorCtrl', function ($scope, $http, socket) {
    $scope.message = 'Hello';

    $http.get('/api/mentors').success(function(mentors) {
      $scope.mentors = mentors;
      console.log("mentors:");
      console.log($scope.mentors);
      //socket.syncUpdates('thing', $scope.mentors);
    });
  });
