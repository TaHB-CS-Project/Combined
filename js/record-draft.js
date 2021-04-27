
$.getJSON("js/record-draft.json", 
function (data) {
var info = '';
$.each(data, function (key, value) {

    info += '<tr>';

    info += '<td>' + 
        value.Start_datetime + '</td>';

    info += '<td>' + 
        value.Gender + '</td>';

    info += '<td>' + 
        value.Diagnosis_name + '</td>';

    info += '<td>' + 
        value.Procedure_name + '</td>';
        

    info += '<td class="text-center"><button type="button" class="btn btn-primary">edit</button></td>';

    info += '</tr>';

 });
 $('#record_list_datatable').append(info);
});
