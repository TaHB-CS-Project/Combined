$.getJSON("js/hospitalnamelist.json",
function (data) {
    $.each(data,
    function (key, value) {
    $("#hospital").append("<option value='" + value.Hospital_name + "'>" + value.Hospital_name + "</option>");
    });
});
