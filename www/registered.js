$(function() {
    let contestId = Cookies.get('registerId');

    $.ajax({
        type: 'GET',
        url: '/api/contests/'+contestId,
        dataType: 'json',
        success: function(data) {
            let backLink = '<a href="/">&lt;&lt;</a> ';
            $('#contestName').html(backLink + data.name);
        },
        error: function(error) {
            // console.error(error);
            window.location.href = "/";
        },
    });

    $.ajax({
        type: 'GET',
        url: '/api/registered/'+contestId,
        dataType: 'json',
        success: handleResult,
        error: function(error) {
            console.log(error);
            let empty = '<div>Något gick fel. Försök igen senare.</div>';
            $('#registered').append(empty);
        }
    });

    function handleResult(data) {
        if (data && data.length) {
            let nw1data = [];
            let nw2data = [];
            data.forEach(element => {
                let nwdataobj = {
                    "name": element.firstName + ' & ' + element.dogName,
                    "event" : element.eventName
                };
                if (element.classNbr == 1) {
                    nw1data.push(nwdataobj);
                } else if (element.classNbr == 2) {
                    nw2data.push(nwdataobj);
                } else {
                    console.error("Invalid class nbr of element");
                }
            });

            let nw1 = '<h3>NW1</h3>';
            nw1 += buildTable(nw1data);
            let nw2 = '<h3>NW2</h3>';
            nw2 += buildTable(nw2data);
            $('#registered').append(nw1).append(nw2);
        } else {
            let empty = '<div>Inga registrerade</div>';
            $('#registered').append(empty);
        }
    }

    function buildTable(nwdata) {
        var grouped = [],
        names = {},
        i, j, cur;
        for (i = 0, j = nwdata.length; i < j; i++) {
            cur = nwdata[i];
            if (!(cur.name in names)) {
                names[cur.name] = {name: cur.name, events: []};
                grouped.push(names[cur.name]);
            }
            names[cur.name].events.push(cur.event);
        }

        var json = {
            "Namn": [],
            "Sök": [],
        };

        grouped.forEach((e) => {
            json.Namn.push(e.name);
            json.Sök.push(e.events.reverse().join(', '));
        });
        
        var headings = Object.keys(json);
        
        var html = '<table class="table">';
        html += '<thead><tr>';
        $.each(headings, function () {
          html += '<th>' + this + '</th>';
        });
        html += '</tr></thead>';
        
        html += '<tbody>';
        for (i = 0, len = json[headings[0]].length; i < len; i++) {
          html += '<tr>';
          $.each(headings, function () {
            html += '<td>' + json[this][i] + '</td>';
          });
          html += '</tr>';
        }
        html += '</tbody>';
        html += '</table>';
        
        return html;
    }
});