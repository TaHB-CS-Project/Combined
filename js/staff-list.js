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
  