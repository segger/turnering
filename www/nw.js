$(function() {
    
    $.ajax({
        type: 'GET',
        url: '/api/contests/enabled',
        dataType: 'json',
        success: function(data) {
            data.forEach(element => {
                let contest = '<button id="' + element.id + '" type="button" '+
                'class="btn btn-primary register">' + element.name + '</button>';
                $('#contests').append(contest);
            });
        },
        error: function(error) {
            console.log(error);
        }
    });
    
    $('#contests').on('click', 'button.register', function() {
        contestId = $(this).attr('id');
        Cookies.set('registerId',contestId);
        window.location.href = "/register.html";
        return false;
    });
});