var Rell=function (){
    location.reload();
};
setTimeout(Rell, 300000)
function logOut() {
    var allcookies = document.cookie.split(";");
    for (var i = 0; i < allcookies.length; i++) {
        var cookie = allcookies[i];
        var eqPos = cookie.indexOf("=");
        var name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;
        document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT";
    }
    location.replace("/auth")
}