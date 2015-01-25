"use strict"

var State = function() {
  this.content = new Content();
};

// makes API calls to get the next content
State.prototype.refresh = function() {
  console.log("Refreshing...");
  this.content.getNext();
}

// wrapper for the content, containing data
// as well as helper functions to render
var Content = function() {
  this.count = 0;
  this.getFirst();
}

Content.prototype.onComplete_ = function(data) {
  console.log("COMPLETE:", data);
  this.update(data);
  this.render();
}

Content.prototype.fetch_ = function(params) {
  var apiUrl = window.location.origin + "/api";
  $.get(apiUrl, params)
    .done(this.onComplete_.bind(this));
};

Content.prototype.update = function(resp) {
  this.id = resp.id;
  this.title = resp.title;
  this.game = resp.game;
  this.upvotes = resp.upvotes;
  this.reddit = resp.reddit;
};

Content.prototype.getNext = function() {
  this.count += 1;
  var params = {
    count: this.count,
    upvotes: this.upvotes,
  }; // TODO: Change this

  $(".checkbox label input").each(function() {
    var name = $(this).attr("name");
    if ($(this).is(":checked")) {
      params[name] = "1";
    } else {
      params[name] = "0";
    }
  });
  this.fetch_(params);
};

Content.prototype.getFirst = function() {
  var params = {
    count: this.count
  };
  this.fetch_(params);
};

Content.prototype.render = function() {
  $(".contentWrapper").empty();

  var img = $("<img/>");
  img.attr("data-id", this.id);
  img.addClass("gfyitem");

  $(".contentWrapper").append(img);
  console.log(this.count);
  gfyCollection.init();

  $(".page-header h2").replaceWith("<h2>" + this.title + "</h2>");
  $("#reddit-comment").attr("href", this.reddit);
};

var spaceHandler = function(e) {
  if (e.keyCode === 32) {
    window.state.refresh();
  }
};

$(document).ready(function() {
  window.state = new State();
  $("body").on("keydown", spaceHandler);
  $('[data-toggle=offcanvas]').click(function() {
    $('.row-offcanvas').toggleClass('active');
  });
});
