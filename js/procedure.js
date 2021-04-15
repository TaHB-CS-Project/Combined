$.getJSON("js/procedure.json", 

    function (data) {
    var info1 = '';
    var info2 = '';
    var x = data.length;
    var r = 0;
    
    $.each(data, function (key, value) {
      if(r < x/2){
      info1 += '<tr>';
      info1 += '<td>' + 
          value.Procedure_id + '</td>';

      info1 += '<td>' + 
          value.Procedure_name + '</td>';

      info1 += '<td class="text-center"><button type="button" id="newProcedure"class="btn btn-primary">edit</button></td>';  
      info1 += '</tr>';
      r++;
      }
      else{
      info2 += '<tr>';
      info2 += '<td>' + 
           value.Procedure_id + '</td>';

      info2 += '<td>' + 
           value.Procedure_name + '</td>';
     
      info2 += '<td class="text-center"><button type="button" id="newProcedure"class="btn btn-primary">edit</button></td>';  
      info2 += '</tr>';
      }
    });
    $('#procedure_datatable_1').append(info1);    
    $('#procedure_datatable_2').append(info2);
    });

    $(document).ready(function() { 
      $("#searchData").on("keyup", function() { 
          var value = $(this).val().toLowerCase(); 
          $("#procedureData tr").filter(function() { 
              $(this).toggle($(this).text() 
              .toLowerCase().indexOf(value) > -1) 
          }); 
      }); 
    }); 

    
$('#myModal').on('shown.bs.modal', function () {
    $('#myInput').trigger('focus')
  })
  

var procedure1;
document.getElementById("newProcedure").onclick=editProcedure;


function editProcedure(){
  procedure1 = prompt("Edit the procedure: ");
  if(procedure===null){
   console.log(document.getElementById("procedure").innerHTML);
  }
  else{
  updateProcedure();
  }
}

function updateProcedure() {
  document.getElementById("procedure").innerHTML = procedure1;
}

/*$('#myModal').on('shown.bs.modal', function () {
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
    
    */
