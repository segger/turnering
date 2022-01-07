$(function() {
    let cookie = Cookies.get('register');
    console.log(cookie);

    $('#save').click(function(){ 
        Cookies.remove('register');
        window.location.href = "/";
        return false;
    });

    $('#step2').hide();
    $('#step3').hide();

    $('#buttonStep1').click(function() {
        let form = document.querySelector('#step1');
        if (!form.checkValidity()) {
            form.classList.add('was-validated');
        } else {
            $("#step1 :input").attr("disabled", true);
            $('#buttonStep1').hide();
            $('#step2').show();
        }
    });
});

