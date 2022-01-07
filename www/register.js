$(function() {
    let cookie = Cookies.get('register');
    console.log(cookie);

    $('#save').click(function(){ 
        Cookies.remove('register');
        window.location.href = "/";
        return false;
    });

    $('#editStep1').hide();
    $('#accordionStep2').hide();
    $('#step3').hide();

    $('#buttonStep1').click(function() {
        let form = document.querySelector('#step1');
        if (!form.checkValidity()) {
            form.classList.add('was-validated');
        } else {
            $("#step1 :input").attr("disabled", true);
            $('#buttonStep1').hide();
            $('#editStep1').show();
            $('#editStep1').attr("disabled", false);
            $('#accordionStep2').show();
        }
    });

    $('#editStep1').click(function() {
        $("#step1 :input").attr("disabled", false);
        $('#buttonStep1').show();
        $('#editStep1').hide();
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

    function setSummary(form) {
        for (let i = 1; i <= 4; i++) {
            let si = parseInt(i);
            let points = form.querySelector('#point_'+si).value;
            let errors = form.querySelector('#error_'+si).value;
            let mm = form.querySelector('#time_mm_'+si).value;
            let ss = form.querySelector('#time_ss_'+si).value;
            let ms = form.querySelector('#time_ms_'+si).value;
            let sse = form.querySelector('#sse_'+si).checked;
            setSummary2(i, points, errors, mm, ss, ms, sse);
        }
    }

    function setSummary2(nbrStr, points, errors, mm, ss, ms, sse) {
        let si = parseInt(nbrStr);
        $('#search_points_'+si).text(Number(points)+"p");
        $('#search_errors_'+si).text(Number(errors)+"p");
        let fmm = String(mm).padStart(2, '0');
        let fss = String(ss).padStart(2, '0');
        let fms = String(ms).padStart(2, '0');
        $('#search_time_'+si).text(fmm+":"+fss+","+fms);
        if (sse) {
            $('#search_sse_'+si).text("sse");
        }
    }

    $('#buttonStep2').click(function() {
        let form = document.querySelector('#step2');
        if (!isValid(form)) {
            form.classList.add('was-validated');
            allCollapsables(form, true);
        } else {
            $("#step2 :input").attr("disabled", true);
            allCollapsables(form, false);
            setSummary(form);
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

