$('#myModal').on('shown.bs.modal', function () {
    $('#myInput').trigger('focus')
  })
  

$(document).ready(function() { 
  $("#searchData").on("keyup", function() { 
      var value = $(this).val().toLowerCase(); 
      $("#diagnosisData tr").filter(function() { 
          $(this).toggle($(this).text() 
          .toLowerCase().indexOf(value) > -1) 
      }); 
  }); 
}); 



$.getJSON("js/diagnosis.json", 
    function (data) {
    var info = '';
    $.each(data, function (key, value) {
  
        info += '<tr>';
        
        info += '<td>' + 
            value.Diagnosis_id + '</td>';

        info += '<td>' + 
            value.Diagnosis_name + '</td>';

        info += '<td class="text-center"><button type="button" class="btn btn-primary">edit</button></td>';

        info += '</tr>';

     });
     $('#diagnosis_list_datatable').append(info);
    });

    
   
