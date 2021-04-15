$('#myModal').on('shown.bs.modal', function () {
    $('#myInput').trigger('focus')
  })
  

$(document).ready(function() { 
  $("#searchData").on("keyup", function() { 
      var value = $(this).val().toLowerCase(); 
      $("#procedureData tr").filter(function() { 
          $(this).toggle($(this).text() 
          .toLowerCase().indexOf(value) > -1) 
      }); 
  }); 
}); 

$.getJSON("js/procedure.json", 
    function (data) {
    var info = '';
    $.each(data, function (key, value) {
  
        info += '<tr>';
        
        info += '<td>' + 
            value.Procedure_id+ '</td>';

        info += '<td>' + 
            value.Procedure_name + '</td>';

        info += '<td class="text-center"><button type="button" class="btn btn-primary">edit</button></td>';

        info += '</tr>';

     });
     $('#procedure_list_datatable').append(info);
    });