$(function() {
    
    $('#register').click(function(){ 
        Cookies.set('register','blue');
        window.location.href = "/register.html";
        return false;
    });
});