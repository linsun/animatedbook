$(document).ready(function() {
  var headerTitleElement = $("#header h1");
  var entriesElement = $("#guestbook-entries");
  var formElement = $("#guestbook-form");
  var submitElement = $("#guestbook-submit");
  var entryContentElement = $("#guestbook-entry-content");
  var hostAddressElement = $("#guestbook-host-address");

  function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
  
  // handle giphy content
  async function loadDoc() {
    console.log('Taking a break to ensure database is populated');
    await sleep(1000);
    console.log('1 second later');
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
      if (this.readyState == 4 && this.status == 200) {
        document.getElementById("giphy").innerHTML = '<p><iframe src=' + this.responseText + ' width="480" height="270" frameBorder="0" class="giphy-embed" allowFullScreen></iframe></p>';
        console.log(this.responseText);
      }
    };
    xhttp.open("GET", "giphy", true);
    xhttp.send();
  }

  var appendGuestbookEntries = function(data) {
    entriesElement.empty();
    $.each(data, function(key, val) {
      entriesElement.append("<p>" + val + "</p>");
    });
  }

  var handleSubmission = function(e) {
    e.preventDefault();
    var entryValue = entryContentElement.val()
    if (entryValue.length > 0) {
      entriesElement.append("<p>...</p>");
      $.getJSON("rpush/guestbook/" + entryValue, appendGuestbookEntries);
	  entryContentElement.val("")
    }

    loadDoc();
    return false;
  }

  submitElement.click(handleSubmission);
  formElement.submit(handleSubmission);
  hostAddressElement.append(document.URL);
  // Poll every second.
  (function fetchGuestbook() {
    $.getJSON("lrange/guestbook").done(appendGuestbookEntries).always(
      function() {
        setTimeout(fetchGuestbook, 1000);
      });
  })();

});
