$(function() {
    
    $.ajax({
        type: 'GET',
        url: '/api/contests',
        dataType: 'json',
        success: function(data) {
            if (data && data.length) {
                data.forEach(element => {
                    let contestButtonOpen = '<button id="' + element.id + '" type="button" '+
                    'class="btn btn-primary register" ';
                    let contestButtonClose = element.name + '</button>';
                    let disabled = element.enabled ? "" : "disabled";
                    let contest = contestButtonOpen + disabled + '>' + contestButtonClose;
                    let result = contestButtonOpen + '>' + contestButtonClose;
                    $('#contests').append(contest);
                    $('#registered').append(result);
                });
            } else {
                let empty = '<div>Det finns inga tävlingar</div>';
                $('#contests').append(empty);
                $('#registered').append(empty);
            }
        },
        error: function(error) {
            console.log(error);
            let empty = '<div>Något gick fel. Försök igen senare.</div>';
            $('#contests').append(empty);
            $('#registered').append(empty);
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
});