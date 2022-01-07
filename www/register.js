$(function() {
    let cookie = Cookies.get('register');
    console.log(cookie);

    $('#save').click(function(){ 
        Cookies.remove('register');
        window.location.href = "/";
        return false;
    });

    $('#accordionStep2').hide();
    $('#step3').hide();

    $('#buttonStep1').click(function() {
        let form = document.querySelector('#step1');
        if (!form.checkValidity()) {
            form.classList.add('was-validated');
        } else {
            $("#step1 :input").attr("disabled", true);
            $('#buttonStep1').hide();
            $('#accordionStep2').show();
        }
    });

    function isValid(form) {
        let valid = true;
        let points = form.querySelectorAll("[id^='point']");
        points.forEach(field => {
            if (field.value && field.value > 25) {
                valid = false;
                field.setCustomValidity("Too high value of points");
            } else {
                field.setCustomValidity("");
            }
        });

        let mm = form.querySelectorAll("[id^='time_mm']");
        mm.forEach(field => {
            if (field.value && field.value > 15) {
                valid = false;
                field.setCustomValidity("Too high value of mm");
            } else {
                field.setCustomValidity("");
            }
        });

        let ss = form.querySelectorAll("[id^='time_ss']");
        ss.forEach(field => {
            if (field.value && field.value > 59) {
                valid = false;
                field.setCustomValidity("Too high value of ss");
            } else {
                field.setCustomValidity("");
            }
        });

        let ms = form.querySelectorAll("[id^='time_ms']");
        ms.forEach(field => {
            if (field.value && field.value > 99) {
                valid = false;
                field.setCustomValidity("Too high value of ms");
            } else {
                field.setCustomValidity("");
            }
        });

        return valid;
    }

    function allCollapsables(form, add) {
        let collapsables = form.querySelectorAll(".collapse");
        collapsables.forEach(c => {
            if (add) {
                c.classList.add('show');
            } else {
                c.classList.remove('show');
            }
        });
    }

    $('#buttonStep2').click(function() {
        let form = document.querySelector('#step2');
        if (!isValid(form)) {
            form.classList.add('was-validated');
            allCollapsables(form, true);
        } else {
            $("#step2 :input").attr("disabled", true);
            allCollapsables(form, false);
            $("#step3").show();
            $('#buttonStep2').hide();
        }
    });

    $('#cancel').click(function() {
        let form = document.querySelector('#step2');
        $("#step2 :input").attr("disabled", false);
        $('#collapseOne').collapse();
        $("#step3").hide();
        $('#buttonStep2').show();
    });
});

