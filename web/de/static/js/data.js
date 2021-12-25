var twilightVisible = false;

function handlePermission() {
    navigator.permissions.query({name:'geolocation'}).then(function(result) {
        if (result.state == 'granted') {
            navigator.geolocation.getCurrentPosition(savePosition);
        } else if (result.state == 'prompt') {
            navigator.geolocation.getCurrentPosition(savePosition);
        } else if (result.state == 'denied') {
            guesstimated();
        }
    });
    guesstimated();
}
handlePermission();

function savePosition(position) {
    document.cookie = "lat=" + parseFloat(position.coords.latitude.toFixed(5)) + "; max-age=259200; Secure; domain=sunrisewhen.com;path=/";
    document.cookie = "lon=" + parseFloat(position.coords.longitude.toFixed(5)) + "; max-age=259200; Secure; domain=sunrisewhen.com;path=/";
    document.cookie = "times=" + "" + "; max-age=0; Secure; domain=sunrisewhen.com;path=/";
    getUpdate(position);
}
function getUpdate(position) {
    document.getElementById("loading1").style.display = null
    document.getElementById("loading2").style.display = null;  
    clearTimes();
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            times = JSON.parse(this.response);
            updateDivs(times, position);
        }
    };
    xhttp.withCredentials = true;
    xhttp.open("GET", "https://backend.sunrisewhen.com", true);
    xhttp.send();
}
function updateDivs(times, position) {
    times = times;
    var lat = parseFloat(position.coords.latitude.toFixed(5));
    var lon = parseFloat(position.coords.longitude.toFixed(5));
    document.getElementById("loading1").style.display = "none";
    document.getElementById("loading2").style.display = "none";  
    document.getElementById("sunrise").innerHTML = times["Sunrise"];
    document.getElementById("sunrise_utc").innerHTML = 'UTC ' + times["UTC_Sunrise"];
    document.getElementById("sunset").innerHTML = times["Sunset"];
    document.getElementById("sunset_utc").innerHTML = 'UTC ' + times["UTC_Sunset"];
    document.getElementById("twilight_begin").innerHTML = 'DÄMMERUNG ' + times["CivilTwilightBegin"];
    document.getElementById("twilight_end").innerHTML = 'DÄMMERUNG ' + times["CivilTwilightEnd"];
    
    var el = document.querySelector('#loclink');
    
    var newEl = document.createElement('a');
    newEl.setAttribute("class", 'footer-text');
    newEl.setAttribute("id", 'loclink');
    newEl.setAttribute("href", 'https://wego.here.com/?map=' + times["Lat"] + ',' + times["Lon"] + ',12,normal');
    newEl.setAttribute("target", '_blank');
    newEl.setAttribute("rel", 'noreferrer');
    newEl.innerHTML = '📍 lat:' + times["Lat"] + ' lng:' + times["Lon"];
    
    el.parentNode.replaceChild(newEl, el);
}

function guesstimated() {
    getUpdate();
    function getUpdate() {
        document.getElementById("loading1").style.display = null;
        document.getElementById("loading2").style.display = null;
        clearTimes();
        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                times = JSON.parse(this.response);
                updateDivs(times);
            }
        };
        xhttp.withCredentials = true;
        xhttp.open("GET", "https://backend.sunrisewhen.com", true);
        xhttp.send();
    }
    
    function updateDivs(times) {
        times = times;
        document.getElementById("loading1").style.display = "none";
        document.getElementById("loading2").style.display = "none";  
        document.getElementById("sunrise").innerHTML = times["Sunrise"];
        document.getElementById("sunrise_utc").innerHTML = 'UTC ' + times["UTC_Sunrise"];
        document.getElementById("sunset").innerHTML = times["Sunset"];
        document.getElementById("sunset_utc").innerHTML = 'UTC ' + times["UTC_Sunset"];
        document.getElementById("twilight_begin").innerHTML = 'DÄMMERUNG ' + times["CivilTwilightBegin"];
        document.getElementById("twilight_end").innerHTML = 'DÄMMERUNG ' + times["CivilTwilightEnd"];
        
        var el = document.querySelector('#loclink');
        
        var newEl = document.createElement('a');
        newEl.setAttribute("class", 'footer-text');
        newEl.setAttribute("id", 'loclink');
        newEl.setAttribute("href", 'https://wego.here.com/?map=' + times["Lat"] + ',' + times["Lon"] + ',12,normal');
        newEl.setAttribute("target", '_blank');
        newEl.setAttribute("rel", 'noreferrer');
        newEl.innerHTML = 'guesstimated 📍 lat:' + times["Lat"] + ' lng:' + times["Lon"];
        
        el.parentNode.replaceChild(newEl, el);
    }
}

function clearTimes() {
    document.getElementById("sunrise").innerHTML = "";
    document.getElementById("sunrise_utc").innerHTML = "";
    document.getElementById("sunset").innerHTML = "";
    document.getElementById("sunset_utc").innerHTML = "";
    document.getElementById("twilight_begin").innerHTML = "";
    document.getElementById("twilight_end").innerHTML = "";
}

function toggleTwilight() {
    if(twilightVisible == false) {
        document.getElementById("twilight_begin").style.display = null; 
        document.getElementById("twilight_end").style.display = null; 
        twilightVisible = true;
    } else {
        document.getElementById("twilight_begin").style.display = "none"; 
        document.getElementById("twilight_end").style.display = "none";       
        twilightVisible = false; 
    }
}