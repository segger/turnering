$(function() {
    
    $.ajax({
        type: 'GET',
        url: '/api/contests',
        dataType: 'json',
        success: function(data) {
            if (data && data.length) {
                data.forEach(element => {
                    // let contestButtonOpen = '<button id="' + element.id + '" type="button" '+
                    // 'class="btn btn-primary register" ';
                    // let contestButtonClose = element.name + '</button>';
                    // let registerDisable = element.enabled ? "" : "disabled";
                    // let register = contestButtonOpen + registerDisable + '>' + contestButtonClose;
                    // let registered = contestButtonOpen + '>' + contestButtonClose;
                    // let resultDisable = element.enabled ? "disabled" : "";
                    // let result = contestButtonOpen + resultDisable + '>' + contestButtonClose;
                    let button = '<button id="' + element.id + '" type="button" '+
                    'class="btn btn-primary register">' + element.name + '</button>';
                    let register = button;
                    let registered = button;
                    let result = button;
                    $('#contests').append(register);
                    $('#registered').append(registered);
                    $('#result').append(result);
                });
            } else {
                let empty = '<div>Det finns inga tävlingar</div>';
                $('#contests').append(empty);
                $('#registered').append(empty);
                $('#result').append(empty);
            }
        },
        error: function(error) {
            console.log(error);
            let empty = '<div>Något gick fel. Försök igen senare.</div>';
            $('#contests').append(empty);
            $('#registered').append(empty);
            $('#result').append(empty);
        }
    });
    
    $('#contests').on('click', 'button.register', function() {
        contestId = $(this).attr('id');
        Cookies.set('registerId',contestId);
        window.location.href = "/register.html";
        return false;
    });

    $('#registered').on('click', 'button.register', function() {
        contestId = $(this).attr('id');
        Cookies.set('registerId',contestId);
        window.location.href = "/registered.html";
        return false;
    });

    $('#result').on('click', 'button.register', function() {
        contestId = $(this).attr('id');
        Cookies.set('registerId', contestId);
        window.location.href = "/result.html";
    });
});