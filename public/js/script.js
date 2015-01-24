"use strict"

var State = function() {
  this.count = 0;
  this.content = new Content(this.count);
  this.content.render();
};

// makes API calls to get the next content
State.prototype.refresh = function() {
  this.count += 1;
  console.log("Refreshing...");
}

// wrapper for the content, containing data
// as well as helper functions to render
var Content = function(count, resp) {
  this.count = count;
  if (!resp) {
    this.getFirst();
  }
}

Content.prototype.getFirst = function() {
  var apiUrl = window.location.origin + "/api"
  $.get(apiUrl, { count: this.count })
    .done(function(data) {
      console.log(data);
    });
};

Content.prototype.render = function() {
  $(".contentWrapper").empty();

  var img = $("<img/>");
  img.attr("data-id", "everyillchick");
  img.addClass("gfyitem");

  $(".contentWrapper").append(img);
  gfyCollection.init();
};

var spaceHandler = function() {
  window.state.refresh();
};

$(document).ready(function() {
  window.state = new State();
  $("body").on("keydown", spaceHandler);
});
