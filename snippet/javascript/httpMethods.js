function _delete(id) {
    var xhr = new XMLHttpRequest();
    //console.log(window.location.href)
    xhr.open("DELETE", window.location.href+"/"+id)
    xhr.onload = function() {
        var users = JSON.parse(xhr.responseText);
        if(xhr.readyState == 4 && xhr.status == "200") {
            console.table(users);
        } else {
            console.error(users);
        }
    }
    xhr.send(null);
}

function test(a) {
    alert(a);
}