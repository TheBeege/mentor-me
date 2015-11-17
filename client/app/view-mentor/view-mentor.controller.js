'use strict';

angular.module('mentorMeApp')
  .controller('ViewMentorCtrl', function ($scope, $state, $stateParams, $http) {
    console.log("stateParams:");
    console.log($stateParams);
    $http.get('/api/mentors/' + $stateParams.mentorID).success(function(mentor) {
      $scope.mentor = mentor;
      $scope.tagText = mentor.tags.join(", ");
      console.log("mentor:");
      console.log($scope.mentor);
      //socket.syncUpdates('thing', $scope.mentors);
    });

    $scope.sendMessage = function(username) {
      if (kik.send) {
          kik.openConversation(username);
      }
    };
  });
