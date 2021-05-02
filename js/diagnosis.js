$.getJSON("js/diagnosis.json", 

    function (data) {
    var info1 = '';
    var info2 = '';
    var x = data.length;
    var r = 0;
    
    $.each(data, function (key, value) {
      if(r < x/2){
      info1 += '<tr>';
      info1 += '<td>' + 
          value.Diagnosis_id + '</td>';

      info1 += '<td>' + 
          value.Diagnosis_name + '</td>';

      info1 += '<td class="text-center"><button type="button" id="newDiagnosis"class="btn btn-primary">edit</button></td>';  
      info1 += '</tr>';
      r++;
      }
      else{
      info2 += '<tr>';
      info2 += '<td>' + 
           value.Diagnosis_id + '</td>';

      info2 += '<td>' + 
           value.Diagnosis_name + '</td>';
     
      info2 += '<td class="text-center"><button type="button" id="newDiagnosis"class="btn btn-primary">edit</button></td>';  
      info2 += '</tr>';
      }
    });
    $('#diagnosis_datatable_1').append(info1);    
    $('#diagnosis_datatable_2').append(info2);
    });








    

$(document).ready(function() { 
      $("#searchData").on("keyup", function() { 
          var value = $(this).val().toLowerCase(); 
          $("#diagnosisData tr").filter(function() { 
              $(this).toggle($(this).text() 
              .toLowerCase().indexOf(value) > -1) 
          }); 
      }); 
    }); 
    
    $('#myModal').on('shown.bs.modal', function () {
    $('#myInput').trigger('focus')
  })
  



