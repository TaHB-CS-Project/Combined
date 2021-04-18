
$.getJSON("js/hospitalnamelist.json",
function (data) {
    $.each(data,
    function (key, value) {
    $("#hospital").append("<option value='" + value.Hospital_name + "'>" + value.Hospital_name + "</option>");
    });
});

$.getJSON("js/diagnosis.json",
function (data) {
    $.each(data,
    function (key, value) {
    $("#diagnosis").append("<option value='" + value.Diagnosis_id + "'>" + value.Diagnosis_name + "</option>");
    });
});

$.getJSON("js/procedure.json",
function (data) {
    $.each(data,
    function (key, value) {
    $("#procedure").append("<option value='" + value.Procedure_id + "'>" + value.Procedure_name + "</option>");
    });
});
