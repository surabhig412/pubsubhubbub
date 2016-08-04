/*jslint browser: true*/
/*global $*/
'use strict';
var show_topics_details = function () {
    $('#option_subscribe').show();
};

var hide_topics_details = function () {
    $('#option_subscribe').hide();
};

var show_unsubscribe_topics_details = function () {
  $('#option_unsubscribe').show();
}
var hide_unsubscribe_topics_details = function (){
  $('#option_unsubscribe').hide();
}


$(document).ready(function () {
  hide_topics_details();
  hide_unsubscribe_topics_details();
$("#subscribe_option").off('click').on('click', function () {
    var checked = $('#subscribe_option');
    if (checked.prop('checked')) {
        show_topics_details();
        hide_unsubscribe_topics_details();
        document.getElementById("unsubscribe_option").checked = false;
    } else {
        hide_topics_details();
    }
});

$("#unsubscribe_option").off('click').on('click', function () {
    var checked = $('#unsubscribe_option');
    if (checked.prop('checked')) {
        show_unsubscribe_topics_details();
        document.getElementById("subscribe_option").checked = false;
        hide_topics_details();
    } else {
        hide_unsubscribe_topics_details();
    }
});
})
