if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(savePosition);
}
function savePosition(position) {
    document.cookie = "lat=" + parseFloat(position.coords.latitude.toFixed(5)) + "; max-age=259200; path=/; Secure;";
    document.cookie = "lon=" + parseFloat(position.coords.longitude.toFixed(5)) + "; max-age=259200; path=/; Secure;";
    getUpdate(position);
}
function getUpdate(position) {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            times = JSON.parse(this.response);
            updateDivs(times, position);
        }
    };
    xhttp.open("GET", "https://sunrisewhen.com/update", true);
    xhttp.send();
}
function updateDivs(times, position) {
    times = times;
    var lat = parseFloat(position.coords.latitude.toFixed(5));
    var lon = parseFloat(position.coords.longitude.toFixed(5));
    document.getElementById("sunrise").innerHTML = times[0];
    document.getElementById("sunrise_utc").innerHTML = 'UTC ' + times[2];
    document.getElementById("sunset").innerHTML = times[1];
    document.getElementById("sunset_utc").innerHTML = 'UTC ' + times[3];
    
    var el = document.querySelector('#loclink');
    
    var newEl = document.createElement('a');
    newEl.setAttribute("class", 'footer-text');
    newEl.setAttribute("id", 'loclink');
    newEl.setAttribute("href", 'https://wego.here.com/?map=' + lat + ',' + lon + ',12,normal');
    newEl.setAttribute("target", '_blank');
    newEl.setAttribute("rel", 'noreferrer');
    newEl.innerHTML = '📍 lat:' + lat + ' lng:' + lon;
    
    el.parentNode.replaceChild(newEl, el);
}