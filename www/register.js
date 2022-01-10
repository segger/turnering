$(function() {
    let contestId = Cookies.get('registerId');

    $.ajax({
        type: 'GET',
        url: '/api/contests/'+contestId,
        dataType: 'json',
        success: function(data) {
            $('#contestName').text(data.name);
        },
        error: function(error) {
            // console.error(error);
            window.location.href = "/";
        },
    });

    $('#save').click(function(){ 
        Cookies.remove('registerId');

        $("#step1 :input").attr("disabled", false);
        let participant = $('#step1').serializeArray();
        $("#step1 :input").attr("disabled", true);

        $("#step2 :input").attr("disabled", false);
        let results = $('#step2').serializeArray();
        $("#step2 :input").attr("disabled", true);

        var resultData = {};
        $(results).each(function(index, obj) {
            resultData[obj.name] = obj.value;
        });

        let search = [];
        for (let i = 1; i <= 4; i++) {
            let si = parseInt(i);
            let eventName = $('#search_event_name_'+si).text();
            let points = Number(resultData['point_'+si]);
            let errors = Number(resultData['error_'+si]);
            let mm = resultData['time_mm_'+si];
            let ss = resultData['time_ss_'+si];
            let ms = resultData['time_ms_'+si];
            let time = Number(mm)*60*1000+Number(ss)*1000+Number(ms);
            let sse = resultData['sse_'+si] ? true : false;
            var searchObj = {
                "eventName": eventName,
                "points": points,
                "errors": errors,
                "time": time,
                "sse": sse
            };
            search.push(searchObj);
        }

        var participantObj = {};
        $(participant).each(function(index, obj){
            participantObj[obj.name] = obj.value;
        });

        var data = {};
        data['Participant'] = participantObj;
        data['EventResultList'] = search;

        console.log(JSON.stringify(data));

        $.ajax({
            type: 'POST',
            url: '/api/register',
            data: JSON.stringify(data),
            error: function(e) {
                console.log(e);
            },
            contentType: "application/json"
        });

        // window.location.href = "/";
        return false;
    });

    $('#editStep1').hide();
    $('#accordionStep2').hide();
    $('#step3').hide();

    /*
    function generateStep2() {
        for (let i = 1; i <= 4; i++) {
            let outer = $("div").attr("id", "append_"+parseInt(i));
            $('#step2').append(outer);
        }
    } */

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
            // generateStep2();
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
            // setSummary(form);
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

