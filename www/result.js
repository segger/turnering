$(function() {
    let contestId = Cookies.get('registerId');
    $.ajax({
        type: 'GET',
        url: '/api/results/'+contestId,
        dataType: 'json',
        success: function(data) {
            // console.log(data);
            if (data && data.length) {
                data.forEach(element => {
                    // console.log(element);
                    let result = "<div>" + element.points + ", " + element.errors + ", " + element.time + ", " + element.sse + "</div>";
                    $('#results').append(result);
                });
            } else {
                let empty = '<div>Inga resultat hittades</div>';
                $('#results').append(empty);
            }
        },
        error: function(error) {
            console.log(error);
            let empty = '<div>Något gick fel. Försök igen senare.</div>';
            $('#results').append(empty);
        }
    });
});