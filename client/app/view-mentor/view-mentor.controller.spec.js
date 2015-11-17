'use strict';

describe('Controller: ViewMentorCtrl', function () {

  // load the controller's module
  beforeEach(module('mentorMeApp'));

  var ViewMentorCtrl, scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    ViewMentorCtrl = $controller('ViewMentorCtrl', {
      $scope: scope
    });
  }));

  it('should ...', function () {
  });
});
