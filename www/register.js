$(function() {
    let cookie = Cookies.get('register');
    console.log(cookie);

    $('#save').click(function(){ 
        Cookies.remove('register');
        window.location.href = "/";
        return false;
    });
});

