/*Function to search the data*/
$(document).ready(function() { 
    $("#searchData").on("keyup", function() { 
        var value = $(this).val().toLowerCase(); 
        $("#staffData tr").filter(function() { 
            $(this).toggle($(this).text() 
            .toLowerCase().indexOf(value) > -1) 
        }); 
    }); 
  }); 
  

$.getJSON("staff-list.json", 
    function (data) {
    var info = '';
    $.each(data, function (key, value) {
  
        info += '<tr>';
        info += '<td>' + 
            value.Medicalemployee_id + '</td>';

        info += '<td>' + 
            value.Medicalemployee_firstname + '</td>';

        info += '<td>' + 
        value.Medicalemployee_lastname + '</td>';

        info += '<td>' + 
            value.Medicalemployee_department + '</td>';

        info += '<td>' + 
            value.Medicalemployee_classification + '</td>';
            
        if(value.Medicalemployee_supervisor.Valid == false){
            info += '<td> None </td>';  
        }
        else if(value.Medicalemployee_supervisor.Valid == true){
            info += '<td>' +
            value.Medicalemployee_supervisor.String + '</td>';
        }
        else{
            info += '<td> Error </td>';
        }

        info += '</tr>';
    

     });
     $('#datatables-ajax').append(info);
    });

   
