#!/usr/bin/env node

var comparator = function(a, b) {
  if(a[0] < b[0]){
      return -1;
  }
  if(a[0] > b[0]){
      return 1;
  }
  if(a[1] < b[1]){
      return -1;
  }
  if(a[1] > b[1]){
      return 1;
  }
  return 0;
}

function solution(A) {
	circlePoints = []

	A.forEach(function(radius, index) {
		circlePoints.push([index - radius, false]);
		circlePoints.push([radius + index, true]);
	});

  circlePoints.sort(comparator);

  var intersections = 0, activeDiscs = 0;

  circlePoints.every(function(point) {
    if(!point[1]) {
      intersections += activeDiscs;
      ++activeDiscs;
    } else {
      --activeDiscs;
    }

    if(intersections > 10E6) {
      intersections = -1;
      return false;
    }

    return true;
  })
  return intersections;
}
