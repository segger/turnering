$(function() {
    let contestId = Cookies.get('registerId');

    $.ajax({
        type: 'GET',
        url: '/api/contests/'+contestId,
        dataType: 'json',
        success: function(data) {
            $('#contestName').html(" | " + data.name);
        },
        error: function(error) {
            // console.error(error);
            window.location.href = "/";
        },
    });

    $.ajax({
        type: 'GET',
        url: '/api/results/'+contestId,
        dataType: 'json',
        success: handleResult,
        error: function(error) {
            console.log(error);
            let empty = '<div>Något gick fel. Försök igen senare.</div>';
            $('#result').append(empty);
        }
    });

    function handleResult(data) {
        if (data && data.length) {
            let nw1data = data[0];
            let nw2data = data[1];
            
            let nw1 = '<h3>NW1</h3>';
            nw1 += buildTables(nw1data);
            let nw2 = '<h3>NW2</h3>';
            nw2 += buildTables(nw2data);
            $('#result').append(nw1).append(nw2);
        } else {
            let empty = '<div>Inga registrerade</div>';
            $('#result').append(empty);
        }
    }

    function buildTables(nwdata) {
        let tables = '';
        for (const eventName in nwdata) {
            tables += buildTable(eventName, nwdata[eventName]);
        }
        return tables;
    }

    function buildTable(eventName, nwdata) {
        var json = {
            "#": [],
            "Namn": [],
            "Poäng": [],
            "Fel":[],
            "Tid": [],
            "SSE": [],
            "TP": []
        };

        for (i = 0; i < nwdata.length; i++) {
            var currRow = nwdata[i];

            let ticks = currRow.time;
            let mm = Math.floor(ticks / (60*100));
            let ss = Math.floor((ticks / 100) % 60);
            let ms = ticks % 100;
            let fmm = String(mm).padStart(2, '0');
            let fss = String(ss).padStart(2, '0');
            let fms = String(ms).padStart(2, '0');
            let time = fmm + ':' + fss + ',' + fms;

            let sse = currRow.sse ? 'x' : '';

            json["#"].push(currRow.placement);
            json.Namn.push(currRow.firstName + ' & ' + currRow.dogName);
            json.Poäng.push(currRow.points);
            json.Fel.push(currRow.errors);
            json.Tid.push(time);
            json.SSE.push(sse);
            json.TP.push(currRow.tournamentPoints);
        }
            
        var headings = Object.keys(json);
        
        var html = '<table class="table">';
        html += '<thead><tr>';
        html += '<th colspan="6">' + eventName +'</th>'
        html += '</tr></thead>';
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