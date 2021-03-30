$('#myModal').on('shown.bs.modal', function () {
    $('#myInput').trigger('focus')
  })
  
/*Functions for editing the HTML*/
var diagnosis1;
document.getElementById("newDiagnosis").onclick=editDiagnosis;


function editDiagnosis(){
  diagnosis1 = prompt("Edit the diagnosis: ");
  if(diagnosis1===null){
   console.log(document.getElementById("diagnosis").innerHTML);
  }
  else{
  updateDiagnosis();
  }
}

function updateDiagnosis() {
  document.getElementById("diagnosis").innerHTML = diagnosis1;
}

/*Functions for editing the HTML ends*/




$(document).ready(function() { 
  $("#searchData").on("keyup", function() { 
      var value = $(this).val().toLowerCase(); 
      $("#diagnosisData tr").filter(function() { 
          $(this).toggle($(this).text() 
          .toLowerCase().indexOf(value) > -1) 
      }); 
  }); 
}); 
