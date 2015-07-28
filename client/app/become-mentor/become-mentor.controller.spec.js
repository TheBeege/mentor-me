'use strict';

describe('Controller: BecomeMentorCtrl', function () {

  // load the controller's module
  beforeEach(module('mentorMeApp'));

  var BecomeMentorCtrl, scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    BecomeMentorCtrl = $controller('BecomeMentorCtrl', {
      $scope: scope
    });
  }));

  it('should ...', function () {
    expect(1).toEqual(1);
  });
});
