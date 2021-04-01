    function subform(){
    var data=$("#login_form").serialize();
    $.ajax({
        type:'POST',
        url:'user.go'+'?rand=' +newDate().getTime(),
        async:true,
        dataType:"json",
        data:{
            username: $(`#email`).val(),
            password_hash: $(`#password`).val()
        },
        beforeSend: function(){
            $("#info").fadeOut();
            $("#login").html("Please wait...");
        },
        success: function(jsonData){
            if(jsonData.hasErrors){
                displayErrors(jasonData.errors);
            }
            else{
               alert(jsonData.confirm);
               $("#info").html("<p style='color:blue; font-weight:bold'> Login Successful</p>");
               displayLogin();
            }
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            $('#error').html(XMLHttpRequest.responseText).removeClass('hide').fadeIn('slow');

        }
    });
}
function displayErrors(errors) {
	str_errors = '<p><strong>' + (errors.length > 1 ? more_errors : one_error) + '</strong></p><ol>';
	for (var error in errors) 
		if (error != 'indexOf') str_errors += '<li>' + errors[error] + '</li>';
	$('#error').html(str_errors + '</ol>').removeClass('hide').fadeIn('slow');
}

